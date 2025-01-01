package db

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"main/db/connect"
)

type StoreClients struct {
	Mongo    *mongo.Client
	Postgres *sql.DB
}

func InitConnections() *StoreClients {
	store := &StoreClients{}

	postgresClient := connect.PostgresConnect()
	//defer postgresClient.Close()
	store.Postgres = postgresClient

	mongodbClient := connect.MongoConnect()
	store.Mongo = mongodbClient

	return store
}
