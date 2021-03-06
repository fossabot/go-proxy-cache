package storage

//                                                                         __
// .-----.-----.______.-----.----.-----.--.--.--.--.______.----.---.-.----|  |--.-----.
// |  _  |  _  |______|  _  |   _|  _  |_   _|  |  |______|  __|  _  |  __|     |  -__|
// |___  |_____|      |   __|__| |_____|__.__|___  |      |____|___._|____|__|__|_____|
// |_____|            |__|                   |_____|
//
// Copyright (c) 2020 Fabio Cicerchia. https://fabiocicerchia.it. MIT License
// Repo: https://github.com/fabiocicerchia/go-proxy-cache

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/fabiocicerchia/go-proxy-cache/cache"
	"github.com/fabiocicerchia/go-proxy-cache/config"
	"github.com/fabiocicerchia/go-proxy-cache/server/response"
	"github.com/fabiocicerchia/go-proxy-cache/utils"
	"github.com/fabiocicerchia/go-proxy-cache/utils/ttl"
)

// RequestCallDTO - DTO object containing request and response.
type RequestCallDTO struct {
	Response response.LoggedResponseWriter
	Request  http.Request
	Scheme   string
	CacheObj cache.CacheObj
}

// RetrieveCachedContent - Retrives the cached response.
func RetrieveCachedContent(rc RequestCallDTO) (cache.URIObj, error) {
	method := rc.Request.Method
	reqHeaders := rc.Request.Header

	url := *rc.Request.URL
	url.Host = rc.Request.Host

	err := rc.CacheObj.RetrieveFullPage(method, url, reqHeaders)
	if err != nil {
		log.Warnf("Cannot retrieve page %s: %s\n", url.String(), err)
	}

	ok, err := rc.CacheObj.IsValid()
	if !ok || err != nil {
		return cache.URIObj{}, err
	}

	return rc.CacheObj.CurrentObj, nil
}

// StoreGeneratedPage - Stores a response in the cache.
func StoreGeneratedPage(rc RequestCallDTO, domainConfigCache config.Cache) (bool, error) {
	ttl := ttl.GetTTL(rc.Response.Header(), domainConfigCache.TTL)

	rc.CacheObj.CurrentObj = cache.URIObj{
		URL:             *rc.Request.URL,
		Host:            rc.Request.Host,
		Method:          rc.Request.Method,
		StatusCode:      rc.Response.StatusCode,
		RequestHeaders:  rc.Request.Header,
		ResponseHeaders: rc.Response.Header(),
		Content:         rc.Response.Content,
	}

	done, err := rc.CacheObj.StoreFullPage(ttl)

	return done, err
}

// PurgeCachedContent - Purges a content in the cache.
func PurgeCachedContent(upstream config.Upstream, rc RequestCallDTO) (bool, error) {
	scheme := utils.IfEmpty(upstream.Scheme, rc.Scheme)

	proxyURL := *rc.Request.URL
	proxyURL.Scheme = scheme
	proxyURL.Host = upstream.Host

	return rc.CacheObj.PurgeFullPage(rc.Request.Method, proxyURL)
}
