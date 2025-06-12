package ctxlog

import (
	"net/http"
	"runtime/debug"
)

// middleware adding a http requests RemoteAddr to log context
func RequestClient(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := Addf(r.Context(), r.RemoteAddr)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

// middleware adding a http requests Host to log context
func RequestHost(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := Addf(r.Context(), r.Host)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

// middleware adding a http requests URL.Path to log context
func RequestPath(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := Addf(r.Context(), r.URL.Path)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

// middleware catching and logging panics
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					Printf(r.Context(), "PANIC: %s\n%s", err, string(debug.Stack()))
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		},
	)
}
