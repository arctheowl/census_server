package graph

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

// Development env vars
var (
	host     = getEnvVar("PGHOST")
	port     = getEnvVar("PGPORT")
	user     = getEnvVar("PGUSER")
	password = getEnvVar("PGPASSWORD")
	dbname   = getEnvVar("PGDATABASE")
)

// Production env vars
// var (
// 	host     = os.Getenv("PGHOST")
// 	port     = os.Getenv("PGPORT")
// 	user     = os.Getenv("PGUSER")
// 	password = os.Getenv("PGPASSWORD")
// 	dbname   = os.Getenv("PGDATABASE")
// )
