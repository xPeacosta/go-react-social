package bd

import (
	"context"
	"time"

	"github.com/xPeacosta/go-react-social/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckExistUser recibe un email y revisa en Base de Datos si ya existe*/
func CheckExistUser(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("go-react-social")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
