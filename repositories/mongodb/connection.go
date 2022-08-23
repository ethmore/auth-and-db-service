package mongodb

import (
	"fmt"

	"e-comm/authService/dotEnv"
	"sync"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

var lock = &sync.Mutex{}
var singleInstance *single

type single struct {
}

func Connect() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			URI := dotEnv.GoDotEnvVariable("MONGODBURI")

			if URI == "" {
				log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
			}
			client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
			if err != nil {
				panic(err)
			}
			singleInstance = &single{}
			fmt.Println("MongoDB Connected!")
		} else {
			fmt.Println("MongoDB connection already created.")
		}
	} else {
		fmt.Println("MongoDB connection already created.")
	}
	return singleInstance
}

func Disconnect() {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Println("MongoDB couldn't disconnect")
		}
		fmt.Println("MongoDB Disconnected")
	}()
}
