package middlew

import (
	"net/http"

	"github.com/xPeacosta/go-react-social/bd"
)

/*ChequeoBD es el middlew que permite conocer el estado de la Base de Datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
