package main

import (
	"find-server/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
