//go:build !debug
// +build !debug

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/profile"
)

const webDebugEnabled = false

func initDebugger() profile.Profile {
	return profile.Profile{}
}

func stopDebugger(profiler profile.Profile) {}

func ProfilerHandler() http.Handler {
	r := chi.NewRouter()
	return r
}
