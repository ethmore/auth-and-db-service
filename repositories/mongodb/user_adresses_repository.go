package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUserAddress(userMail, addressName, name, surname, phoneNumber, province, county, address string) error {
	user, err1 := FindOneUser(userMail)
	if err1 != nil {
		return err1
	}

	coll := client.Database("eCommUsers").Collection("userAddresses")
	doc := bson.D{{Key: "userId", Value: user.Id}, {Key: "addressName", Value: addressName}, {Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "address", Value: address}}
	_, err := coll.InsertOne(context.TODO(), doc)
	return err
}

func UpdateUserAdress(addressName, name, surname, phoneNumber, province, county, address string) error {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "address", Value: address}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteUserAddress(userMail, addressName string) error {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}
