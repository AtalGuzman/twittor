package middleW

import (
	"fmt"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/gogearbox/gearbox"
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

func ChequeoBd2(ctx gearbox.Context) {
	fmt.Print("*")
	if bd.CheckConnection() == 0 {
		ctx.Status(gearbox.StatusInternalServerError)
		ctx.SendString("Conexión pérdida con la bd")
		return
	}
	ctx.Next()

}
