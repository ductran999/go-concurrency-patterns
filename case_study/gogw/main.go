package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type ReverseProxy struct {
	authGroup singleflight.Group
}

func (a *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimSpace(authHeader)

	if token == "" {
		http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
		return
	}

	v, err, shared := a.authGroup.Do(token, func() (any, error) {
		return a.validateTokenWithAuthService(token)
	})

	if err != nil {
		http.Error(w, "Auth Error", http.StatusInternalServerError)
		return
	}

	isValid, ok := v.(bool)
	if !ok {
		http.Error(w, "invalid data type", http.StatusInternalServerError)
		return
	}
	if shared {
		fmt.Printf("Token [%s]: reuse result!\n", token)
	}

	if !isValid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, _ = w.Write([]byte("OK - Forwarding to Model"))
}

func (a *ReverseProxy) validateTokenWithAuthService(token string) (bool, error) {
	slog.Info("Call Auth Service...")
	time.Sleep(500 * time.Millisecond)
	if !strings.HasPrefix(token, "Bearer") {
		return false, ErrInvalidToken
	}

	return true, nil
}

func main() {
	ReverseProxy := &ReverseProxy{}
	if err := http.ListenAndServe(":8080", ReverseProxy); err != nil && errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
	}
}
