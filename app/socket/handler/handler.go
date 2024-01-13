package handler

import (
	"net/http"

	"github.com/centrifugal/centrifuge"
)

func ApiChannel(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("uuid")
		ctx := r.Context()
		cred := &centrifuge.Credentials{
			UserID: uuid,
		}
		newCtx := centrifuge.SetCredentials(ctx, cred)
		r = r.WithContext(newCtx)
		h.ServeHTTP(w, r)
	})
}
