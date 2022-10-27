package main

import (
	"log"

	"github.com/flaiers/fiber-clean-architecture/internal/app"
)

func main() {
	log.Fatal(app.Create().Listen(":3000"))
}
