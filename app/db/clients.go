package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
	"main/db/connect"
)

type StoreClients struct {
	Mongo    *mongo.Client
	Postgres *gorm.DB
}

func InitConnections() *StoreClients {
	store := &StoreClients{}

	postgresClient := connect.PostgresConnect()
	store.Postgres = postgresClient

	mongodbClient := connect.MongoConnect()
	store.Mongo = mongodbClient

	return store
}
