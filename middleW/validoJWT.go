package middleW

import (
	"fmt"
	"net/http"

	"github.com/AtalGuzman/twittor/routers"
	"github.com/gogearbox/gearbox"
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

func ValidoJWT2(ctx gearbox.Context) {
	fmt.Print("*")
	_, _, _, err := routers.ProcesoToken(ctx.Get("authorization"))
	if err != nil {
		ctx.Status(gearbox.StatusBadRequest)
		ctx.SendString("Error en el token! " + err.Error())
		return
	}
	ctx.Next()
}
