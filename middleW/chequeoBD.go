package middleW

import (
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
)

/*ChequeoBd: valida conexión a la BD*/
func ChequeoBd(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(rw, "Conexión pérdida con la bd", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
