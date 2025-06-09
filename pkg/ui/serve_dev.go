//go:build !prod

package ui

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func isFrontendRunning() bool {
	client := http.Client{
		Timeout: 1 * time.Second, // short timeout for responsiveness
	}

	resp, err := client.Get("http://localhost:3000")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 500
}

var frontendReachableOnce = false

func Handler() func(writer http.ResponseWriter, request *http.Request) {
	slog.Warn("Enabling frontend handler for development")

	frontendReachableOnce = isFrontendRunning()
	target, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(w http.ResponseWriter, r *http.Request) {
		if handleAPIRoute(w, r) {
			return
		}

		// Try reach frontend once if not already running since start
		if !frontendReachableOnce {
			slog.Info("Checking if frontend is reachable now")
			running := isFrontendRunning()
			frontendReachableOnce = running
		}

		if frontendReachableOnce {
			proxy.ServeHTTP(w, r)
		} else {
			_, _ = w.Write([]byte(`Frontend development server is not running.`))
		}
	}
}
