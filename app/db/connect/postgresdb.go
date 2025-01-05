package connect

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"main/utils"
	"os"
	"time"
)

func PostgresConnect() *gorm.DB {

	err := godotenv.Load("../_.env")

	portEnv := "STORE_PG_DB_PORT"
	if err == nil {
		portEnv = "STORE_PG_DB_EXTERNAL_PORT"
	}

	ctxPing, ctxPingFn := context.WithTimeout(context.Background(), 2*time.Second)
	defer ctxPingFn()

	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("STORE_PG_DB_USER"),
		os.Getenv("STORE_PG_DB_PASSWORD"),
		"localhost",
		os.Getenv(portEnv),
		os.Getenv("STORE_PG_DB"),
	)

	client, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	
	if err != nil {
		log.Fatal(conn)
		panic(err)
	}

	db, err := client.DB()
	if err != nil {
		log.Fatal(conn)
		panic(err)
	}

	if err := db.PingContext(ctxPing); err != nil {
		panic(err)
	}

	pgCheckDb(db)

	return client
}

func pgCheckDb(dbConnect *sql.DB) {
	ctx, ctxFn := utils.GetCtx()
	defer ctxFn()

	query := `
	SELECT table_name as "name"
	FROM information_schema.tables 
	WHERE table_schema='public'
	`

	var records []string
	rows, _ := dbConnect.QueryContext(ctx, query)
	for rows.Next() {
		var r string
		rows.Scan(&r)
		records = append(records, r)
	}
	rows.Close()

	fmt.Println("pg tables: ", records)
}
