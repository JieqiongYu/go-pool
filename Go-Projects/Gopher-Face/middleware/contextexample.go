package middleware

import (
	"context"
	"net/http"
)

/*ContextExampleHandler is a context example handler*/
func ContextExampleHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := newExampleContext(r.Context(), r)
			next.ServeHTTP(w, r.WithContext(ctx))
		}()
	})
}

func newExampleContext(ctx context.Context, r *http.Request) context.Context {

	fooID := r.Header.Get("X-Foo-ID")
	if fooID == "" {
		fooID = "bar"
	}
	return context.WithValue(ctx, "fooID", fooID)
}
