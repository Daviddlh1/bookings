package main

import (
	"fmt"
	"testing"

	"github.com/Daviddlh1/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// test passed
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
