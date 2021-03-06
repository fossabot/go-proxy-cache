package handler

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
)

// RedirectToHTTPS - Redirects from HTTP to HTTPS.
func (rc RequestCall) RedirectToHTTPS(redirectStatusCode int) {
	targetURL := rc.Request.URL
	targetURL.Scheme = "https"
	targetURL.Host = rc.Request.Host

	target := targetURL.String()

	log.Infof("Redirect to: %s", target)

	http.Redirect(rc.Response, rc.Request, target, redirectStatusCode)
}
