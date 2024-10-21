package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
  err := godotenv.Load("configs/.env")

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
