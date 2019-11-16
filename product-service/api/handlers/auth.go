package handlers

import (
	"net/http"
	"strconv"

	"github.com/klferreira/hash-challenge/product-service/pkg/util/authctx"
)

// AuthenticationMiddleware retrieves user id from the headers and places it on the request's Context
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sid := r.Header.Get("X-USER-ID")
		if len(sid) > 0 {
			uid, err := strconv.Atoi(sid)
			if err != nil {
				http.Error(w, "Invalid User ID", http.StatusBadRequest)
				return
			}
			ctx = authctx.NewContext(ctx, int64(uid))
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
