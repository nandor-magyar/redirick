package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const AppName = "redirick"

func TestLoadConfigDefaults(t *testing.T) {
	// Test default values
	conf := LoadConfig(AppName, nil)
	assert.Equalf(t, DefaultRedirect, conf.Target, "Expected target to be %s, got %s")
	assert.Equalf(t, DefaultPort, conf.Port, "Expected port to be %d, got %d", DefaultPort, conf.Port)
	assert.Equalf(t, DefaultStatusCode, conf.StatusCode, "Expected status code to be %d, got %d")
}

func TestLoadConfigFlags(t *testing.T) {
	// Test target flag
	args := []string{
		"-target", "https://dyrectorio.com",
		"-code", "302",
		"-port", "8090",
	}
	conf := LoadConfig(AppName, args)
	assert.Equalf(t, "https://dyrectorio.com", conf.Target, "Expected target to be https://dyrectorio.com, got %s", conf.Target)
	assert.Equalf(t, 302, conf.StatusCode, "Expected status code to be 302, got %d")
	assert.Equalf(t, 8090, conf.Port, "Expected port to be 8090, got %d", conf.Port)
}

func TestLoadConfigURLArg(t *testing.T) {
	// Test target flag
	// Test first argument
	args := []string{"https://dyrector.io"}
	conf := LoadConfig(AppName, args)
	if conf.Target != "https://dyrector.io" {
		t.Errorf("Expected target to be https://dyrector.io, got %s", conf.Target)
	}
}

func TestLoadConfigURLOverload(t *testing.T) {
	// Test target flag
	// Test first argument
	args := []string{"-target", "https://dyrectorio.com", "https://dyrector.io"}
	conf := LoadConfig(AppName, args)
	if conf.Target != "https://dyrector.io" {
		t.Errorf("Expected target to be https://dyrector.io, got %s", conf.Target)
	}
}

func TestServer(t *testing.T) {
	conf := &AppConfig{
		Target:     "https://dyrectorio.com",
		Port:       8090,
		StatusCode: http.StatusFound,
	}

	go func() {
		err := Server(conf)
		assert.Nilf(t, err, "Expected server to start without errors, got %v", err)
	}()

	// todo: more elegant options without messing with the code. wait for healtheck?
	time.Sleep(100 * time.Millisecond)

	t.Run("Redirect", func(t *testing.T) {
		// Give the server some time to start

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8090/", nil)
		assert.Nilf(t, err, "Unexpected error on request creation, got %v", err)

		client := new(http.Client)
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		res, err := client.Do(req)
		assert.Nilf(t, err, "Unexecpted response error, got %v", err)
		assert.Equalf(t, http.StatusFound, res.StatusCode, "Expected status code to be %d, got %d", http.StatusFound, res.StatusCode)

		loc, err := res.Location()
		assert.Nilf(t, err, "Unexpected error while getting http location", err)
		assert.Equalf(t, "https://dyrectorio.com", loc.String(), "Expected value should be https://dyrectorio.com got %v", loc.String())
	})

	t.Run("Healthcheck", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8090/healthz")
		assert.Nilf(t, err, "Unexpected error on healthcheck, got %v", err)
		assert.Equalf(t, http.StatusOK, resp.StatusCode, "Healthcheck should have status code (200), got: %d", resp.StatusCode)
	})
}
