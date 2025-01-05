package connect

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"main/utils"
	"os"
	"time"
)

func MongoConnect() *mongo.Client {
	err := godotenv.Load("../_.env")

	portEnv := "STORE_MDB_DB_PORT"
	if err == nil {
		portEnv = "STORE_MDB_DB_EXTERNAL_PORT"
	}

	ctxPing, ctxPingFn := context.WithTimeout(context.Background(), 2*time.Second)
	defer ctxPingFn()

	url := fmt.Sprintf(
		"mongodb://%s:%s",
		"localhost",
		os.Getenv(portEnv),
	)

	auth := options.Credential{
		Username: os.Getenv("STORE_MDB_DB_USER"),
		Password: os.Getenv("STORE_MDB_DB_PASSWORD"),
	}

	clientOpts := options.Client().ApplyURI(url).SetAuth(auth)
	client, err := mongo.Connect(clientOpts)

	if err := client.Ping(ctxPing, nil); err != nil {
		panic(err)
	}

	mongoCheckDb(client)

	return client
}

func mongoCheckDb(dbConnect *mongo.Client) {
	ctx, ctxFn := utils.GetCtx()
	defer ctxFn()

	list, _ := dbConnect.ListDatabaseNames(ctx, bson.D{{}})

	fmt.Println("mongo tables: ", list)
}

func getMongoCollection(
	client *mongo.Client,
	database string,
	collection string,
) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
