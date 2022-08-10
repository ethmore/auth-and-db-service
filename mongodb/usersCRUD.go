package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneUser(name, surname, email, password string) {
	coll := client.Database("eCommUsers").Collection("users")
	doc := bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func UpdateOneUser(name, surname, email, password string) {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err)
	fmt.Println(result)
}

func DeleteOneUser(id string) {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}

	result, err := coll.DeleteOne(context.TODO(), filter)
	CheckError(err)
	fmt.Println(result)

}

func FindOneUser(email string) string {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return "not registered"
		} else {
			panic(err)
		}
	}

	return result["email"].(string)
}
