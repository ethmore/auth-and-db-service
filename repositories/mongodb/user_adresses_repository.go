package mongodb

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Address struct {
	Id              string `bson:"_id"`
	Title           string
	Name            string
	Surname         string
	PhoneNumber     string
	Province        string
	County          string
	DetailedAddress string
}

type IUserAddressesRepo interface {
	InsertUserAddress(userRepo IUsersRepo, userMail, addressName, name, surname, phoneNumber, province, county, address string) error
	UpdateUserAdress(addressName, name, surname, phoneNumber, province, county, address string) error
	DeleteUserAddress(userMail, addressName string) error
	FindUserAddress(addressID string) (*Address, error)
	FindAllUserAddresses(userRepo IUsersRepo, userMail string) ([]Address, error)
}

type UserAddressesRepo struct{}

func (uar UserAddressesRepo) InsertUserAddress(userRepo IUsersRepo, userMail, addressName, name, surname, phoneNumber, province, county, address string) error {
	user, err1 := userRepo.FindOneUser(userMail)
	if err1 != nil {
		return err1
	}

	coll := client.Database("eCommUsers").Collection("userAddresses")
	doc := bson.D{{Key: "userId", Value: user.Id}, {Key: "title", Value: addressName}, {Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "detailedAddress", Value: address}}
	_, err := coll.InsertOne(context.TODO(), doc)
	return err
}

func (uar UserAddressesRepo) UpdateUserAdress(addressName, name, surname, phoneNumber, province, county, address string) error {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}, {Key: "surname", Value: surname}, {Key: "phoneNumber", Value: phoneNumber}, {Key: "province", Value: province}, {Key: "county", Value: county}, {Key: "address", Value: address}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (uar UserAddressesRepo) DeleteUserAddress(userMail, addressName string) error {
	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "addressName", Value: addressName}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}

func (uar UserAddressesRepo) FindUserAddress(addressID string) (*Address, error) {
	addressID = strings.TrimSpace(addressID)
	objID, convErr := primitive.ObjectIDFromHex(addressID)
	if convErr != nil {
		return nil, convErr
	}

	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "_id", Value: objID}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	strID := result["_id"].(primitive.ObjectID).Hex()
	var address = Address{
		Id:              strID,
		Title:           result["title"].(string),
		Name:            result["name"].(string),
		Surname:         result["surname"].(string),
		PhoneNumber:     result["phoneNumber"].(string),
		Province:        result["province"].(string),
		County:          result["county"].(string),
		DetailedAddress: result["detailedAddress"].(string),
	}

	return &address, nil
}

func (uar UserAddressesRepo) FindAllUserAddresses(userRepo IUsersRepo, userMail string) ([]Address, error) {
	// var address Address
	user, findErr := userRepo.FindOneUser(userMail)
	if findErr != nil {
		return nil, findErr
	}

	coll := client.Database("eCommUsers").Collection("userAddresses")
	filter := bson.D{{Key: "userId", Value: user.Id}}
	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []Address
	for cur.Next(context.TODO()) {
		var elem Address
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	defer cur.Close(context.TODO())
	return results, nil
}
