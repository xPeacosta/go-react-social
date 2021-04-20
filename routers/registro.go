package routers

import (
	"encoding/json"
	"net/http"

	"github.com/xPeacosta/go-react-social/bd"
	"github.com/xPeacosta/go-react-social/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "Debe especificar una contraseña de añ menos 8 caracteres", 400)
		return
	}

	_, exist, _ := bd.CheckExistUser(t.Email)
	if exist {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el usuario: "+err.Error(), 400)
		return
	}

	if !status {
		if err != nil {
			http.Error(w, "No se logró intentar el registro del usuario", 400)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
