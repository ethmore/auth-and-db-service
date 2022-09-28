package mongodb

import (
	"auth-and-db-service/repositories/postgresql"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/exp/slices"
)

type Product struct {
	Id  string
	Qty string
}

type Cart struct {
	Id         string `bson:"_id"`
	UserId     primitive.ObjectID
	Products   []Product
	TotalPrice string
}

type IUserCartRepo interface {
	NewCart(up IUsersRepo, userMail string) error
	AddProductToCart(postgresqlRepo postgresql.IProductRepo, up IUsersRepo, userMail, productId, qty string) error
	FindAllCartProducts(up IUsersRepo, userMail string) (*Cart, error)
	RemoveProductFromCart(up IUsersRepo, userMail, productId string) error
	UpdateCartProducts(userId primitive.ObjectID, newProducts []Product) error
	ChangeProductQty(up IUsersRepo, userMail, productId, productQty string) error
	AddTotalToCart(up IUsersRepo, userMail, totalPrice string) error
	GetTotalPrice(up IUsersRepo, userMail string) (string, error)
	ClearCart(up IUsersRepo, userMail string) error
}

type UserCartRepo struct{}

func (ucr UserCartRepo) NewCart(up IUsersRepo, userMail string) error {
	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		fmt.Println("mongodb (find): ", findErr)
		return findErr
	}

	coll := client.Database("eCommUsers").Collection("carts")
	doc := bson.D{{Key: "userId", Value: user.Id}, {Key: "products"}}

	_, err := coll.InsertOne(context.TODO(), doc)
	return err
}

func (ucr UserCartRepo) AddProductToCart(postgresqlRepo postgresql.IProductRepo, up IUsersRepo, userMail, productId, qty string) error {
	if productId == "" {
		return errors.New("empty product id")
	}

	product, getErr := postgresqlRepo.GetProduct(productId)
	if getErr != nil {
		return getErr
	}

	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		fmt.Println("mongodb (find): ", findErr)
		return findErr
	}

	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: user.Id}}

	var result Cart
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}

	fmt.Println(product)
	var prod = Product{
		Id:  productId,
		Qty: qty,
	}

	products := append(result.Products, prod)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "products", Value: products}}}}

	_, updateErr := coll.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (ucr UserCartRepo) FindAllCartProducts(up IUsersRepo, userMail string) (*Cart, error) {
	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		fmt.Println("mongodb (find):", findErr)
		return nil, findErr
	}
	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: user.Id}}

	var result Cart
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ucr UserCartRepo) RemoveProductFromCart(up IUsersRepo, userMail, productId string) error {
	cart, findErr := ucr.FindAllCartProducts(up, userMail)
	if findErr != nil {
		fmt.Println("mongodb (remove): ", findErr)
		return findErr
	}

	i := slices.IndexFunc(cart.Products, func(c Product) bool { return c.Id == productId })

	copy(cart.Products[i:], cart.Products[i+1:])
	cart.Products[len(cart.Products)-1] = Product{} // or the zero value of T
	cart.Products = cart.Products[:len(cart.Products)-1]
	err := ucr.UpdateCartProducts(cart.UserId, cart.Products)

	return err
}

func (ucr UserCartRepo) UpdateCartProducts(userId primitive.ObjectID, newProducts []Product) error {
	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "products", Value: newProducts}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (ucr UserCartRepo) ChangeProductQty(up IUsersRepo, userMail, productId, productQty string) error {
	cart, findErr := ucr.FindAllCartProducts(up, userMail)
	if findErr != nil {
		return findErr
	}
	fmt.Println(productId, productQty)
	var newProducts []Product
	for i := 0; i < len(cart.Products); i++ {
		if cart.Products[i].Id == productId {
			cart.Products[i].Qty = productQty
		}

		newProducts = append(newProducts, cart.Products[i])
		fmt.Println(newProducts)
	}

	user, usrErr := up.FindOneUser(userMail)
	if usrErr != nil {
		return usrErr
	}
	err := ucr.UpdateCartProducts(user.Id, newProducts)
	if err != nil {
		return err
	}

	return nil
}

func (ucr UserCartRepo) AddTotalToCart(up IUsersRepo, userMail, totalPrice string) error {
	cart, findErr := ucr.FindAllCartProducts(up, userMail)
	if findErr != nil {
		return findErr
	}

	cart.TotalPrice = totalPrice

	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		fmt.Println("mongodb (find):", findErr)
		return findErr
	}
	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: user.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "totalPrice", Value: totalPrice}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (ucr UserCartRepo) GetTotalPrice(up IUsersRepo, userMail string) (string, error) {
	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		fmt.Println("mongodb (find):", findErr)
		return "", findErr
	}
	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: user.Id}}
	var result Cart
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(result.TotalPrice)
	return result.TotalPrice, nil
}

func (ucr UserCartRepo) ClearCart(up IUsersRepo, userMail string) error {
	user, findErr := up.FindOneUser(userMail)
	if findErr != nil {
		return findErr
	}

	coll := client.Database("eCommUsers").Collection("carts")
	filter := bson.D{{Key: "userId", Value: user.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "products", Value: nil}, {Key: "totalPrice", Value: ""}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}
