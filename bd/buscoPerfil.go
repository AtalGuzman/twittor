package bd

import (
	"context"
	"time"

	"github.com/AtalGuzman/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil: encuentra usuario por ID*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)
	cond := bson.M{"_id": objID}
	err := col.FindOne(ctx, cond).Decode(&perfil)

	perfil.Password = ""

	return perfil, err
}
