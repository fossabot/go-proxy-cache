// +build all unit

package balancer_test

//                                                                         __
// .-----.-----.______.-----.----.-----.--.--.--.--.______.----.---.-.----|  |--.-----.
// |  _  |  _  |______|  _  |   _|  _  |_   _|  |  |______|  __|  _  |  __|     |  -__|
// |___  |_____|      |   __|__| |_____|__.__|___  |      |____|___._|____|__|__|_____|
// |_____|            |__|                   |_____|
//
// Copyright (c) 2020 Fabio Cicerchia. https://fabiocicerchia.it. MIT License
// Repo: https://github.com/fabiocicerchia/go-proxy-cache

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/fabiocicerchia/go-proxy-cache/config"
	"github.com/fabiocicerchia/go-proxy-cache/server/balancer"
)

func TestGetLBRoundRobinUndefined(t *testing.T) {
	setUp()

	var endpoints []string
	balancer.InitRoundRobin("testing", endpoints)
	endpoint := balancer.GetLBRoundRobin("testing", "8.8.8.8")

	assert.Equal(t, "8.8.8.8", endpoint)

	tearDown()
}

func TestGetLBRoundRobinDefined(t *testing.T) {
	setUp()

	var endpoints = []string{"1.2.3.4"}
	balancer.InitRoundRobin("testing", endpoints)
	endpoint := balancer.GetLBRoundRobin("testing", "8.8.8.8")

	assert.Equal(t, "1.2.3.4", endpoint)

	tearDown()
}

func setUp() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	config.Config = config.Configuration{}
}

func tearDown() {
	config.Config = config.Configuration{}
}
