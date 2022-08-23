package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

func InsertOneUser(name, surname, email, password string) error {
	coll := client.Database("eCommUsers").Collection("users")
	doc := bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}
	_, err := coll.InsertOne(context.TODO(), doc)
	return err
}

func UpdateOneUser(name, surname, email, password string) error {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "email", Value: email}, {Key: "password", Value: password}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteOneUser(id string) error {
	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}

type User struct {
	Id       primitive.ObjectID
	Name     string
	Surname  string
	Email    string
	Password string
}

func FindOneUser(email string) (*User, error) {
	var user User

	coll := client.Database("eCommUsers").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	user.Id = result["_id"].(primitive.ObjectID)
	user.Name = result["name"].(string)
	user.Surname = result["surname"].(string)
	user.Email = result["email"].(string)
	user.Password = result["password"].(string)

	return &user, nil
}
