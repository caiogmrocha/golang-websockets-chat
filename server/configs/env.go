package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
  cwd, error := os.Getwd()

  if error != nil {
    log.Fatal(error)
    os.Exit(1)
  }

  err := godotenv.Load(cwd + "/../../.env")

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
