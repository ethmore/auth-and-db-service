package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneUser(name, surname, email, password string) int {
	coll := client.Database("eCommUsers").Collection("users")
	doc := bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	return 200
}

func UpdateOneUser(name, surname, email, password string) int {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err)
	return 200
}

func DeleteOneUser(id string) int {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	CheckError(err)
	return 200

}

func FindOneUser(email string) (id int, x, y string) {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return 0, "not registered", ""
		} else {
			panic(err)
		}
	}

	return result["_id"].(int), result["email"].(string), result["password"].(string)
}
