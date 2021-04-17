package middleW

import (
	"net/http"

	"github.com/AtalGuzman/twittor/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("authorization"))
		if err != nil {
			http.Error(rw, "Error en el token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
