package database_mongo

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type singletonClient struct {
	Client *mongo.Client
}

var (
	instanceClient *singletonClient
	onceDatabase   sync.Once
)

func GetClient(path string) *mongo.Client {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return getClientInstance().Client
}

func getClientInstance() *singletonClient {
	onceDatabase.Do(func() {
		instanceClient = &singletonClient{}
		mongoURL := os.Getenv("MONGO_URL")
		client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
		}

		initClient(client)
	})
	return instanceClient
}

func initClient(client *mongo.Client) {
	instanceClient.Client = client
}

type mongoCollection struct {
	coll *mongo.Collection
}
