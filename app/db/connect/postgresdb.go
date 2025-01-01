package connect

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"main/utils"
	"os"
	"time"
)

func PostgresConnect() *sql.DB {

	err := godotenv.Load("../_.env")

	portEnv := "STORE_PG_DB_PORT"
	if err == nil {
		portEnv = "STORE_PG_DB_EXTERNAL_PORT"
	}

	ctxPing, ctxPingFn := context.WithTimeout(context.Background(), 2*time.Second)
	defer ctxPingFn()

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		os.Getenv(portEnv),
		os.Getenv("STORE_PG_DB_USER"),
		os.Getenv("STORE_PG_DB_PASSWORD"),
		os.Getenv("STORE_PG_DB"),
	)

	client, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	if err := client.PingContext(ctxPing); err != nil {
		panic(err)
	}

	pgCheckDb(client)

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
