package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUserAddress(userMail, addressName, name, surname, phoneNumber, province, county, address string) int {
	userId, _, _ := FindOneUser(userMail)

	coll := client.Database("eCommUsers").Collection("userAddresses")
	doc := bson.D{{Key: "userId", Value: userId}, {Key: "addressName", Value: addressName}, {Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "address", Value: address}}
	_, err := coll.InsertOne(context.TODO(), doc)
	CheckError(err)
	return 200
}

func UpdateUserAdress(addressName, name, surname, phoneNumber, province, county, address string) int {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "address", Value: address}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err)
	return 200
}

func DeleteUserAddress(userMail, addressName string) int {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	CheckError(err)
	return 200
}
