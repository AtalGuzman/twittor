package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(id, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	fmt.Println("entrado en borrto tweet!")
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objectID, _ := primitive.ObjectIDFromHex(id)
	condicion := bson.M{"_id": objectID, "userid": userid}
	dc, err := col.DeleteOne(ctx, condicion)
	fmt.Printf("borrados %d\n", dc.DeletedCount)

	return err
}
