package ctxlog

import "net/http"

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
