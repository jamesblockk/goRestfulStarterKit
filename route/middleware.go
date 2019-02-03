package route

import (
	"fmt"
	"net/http"
)

const validateToken = "this_is_my_token"

// middleware :
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		var err error
		defer fmt.Println(err)

		if len(tokenStr) == 0 || tokenStr != validateToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("id", "jamesblock")
		next.ServeHTTP(w, r)
	})
}
