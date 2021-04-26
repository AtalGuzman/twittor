package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/AtalGuzman/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuarioTodos(Id string, pagina int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	fmt.Print("leyendo usuarios!")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println("No se encontro ningun usuario")
		fmt.Println(err)
		return results, false
	}

	var encontrado, incluir bool
	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = Id
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		encontrado, _ = ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == Id {
			incluir = false
		}

		fmt.Print(tipo, encontrado, incluir)
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.Email = ""
			s.SitioWeb = ""
			s.Ubicacion = ""

			results = append(results, &s)
			fmt.Println(s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)

	return results, true
}
