package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	cache_redis "github.com/fabiocicerchia/go-proxy-cache/cache/redis"
	"github.com/fabiocicerchia/go-proxy-cache/utils"
)

// --- LOG

// Log the redirect url
func logRequest(proxyUrl string) {
	log.Printf("proxy_url: %s\n", proxyUrl)
}

// Log the env variables required for a reverse proxy
func logSetup() {
	forward_to := utils.GetEnv("FORWARD_TO", "")

	log.Printf("Server will run on: %s\n", GetListenAddress())
	log.Printf("Redirecting to url: %s\n", forward_to)
}

// --- LOGIC

// Get the port to listen on
func GetListenAddress() string {
	port := utils.GetEnv("PORT", "8080")
	return ":" + port
}

func Start() {
	// Log setup values
	logSetup()

	// redis connect
	cache_redis.Connect()

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)

	if err := http.ListenAndServe(GetListenAddress(), nil); err != nil {
		panic(err)
	}
}

func storeGeneratedPage(url string, lrw loggedResponseWriter) {
	status := lrw.StatusCode
	headers := make(map[string]string)
	for k, values := range lrw.Header() {
		headers[k] = strings.Join(values, "")
	}
	content := string(lrw.Content)
	cache_redis.StoreFullPage(url, status, headers, content, 0)

}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res *loggedResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)

	storeGeneratedPage(url, res)
}

func serveCachedContent(res http.ResponseWriter, fullUrl string) bool {
	code, headers, page, _ := cache_redis.RetrieveFullPage(fullUrl)

	if code == 200 && page != "" {
		res.WriteHeader(code)
		for k, v := range headers {
			res.Header().Add(k, v)
		}
		res.Write(page)

		return true
	}

	return false
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := utils.GetProxyUrl()

	logRequest(url)

	fullUrl := req.URL.String()
	if !serveCachedContent(res, fullUrl) {
		lrw := newLoggedResponseWriter(res)
		serveReverseProxy(url, lrw, req)
	}
}
