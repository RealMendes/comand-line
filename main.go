package main

import (
	"comandLine/app"
	"log"
	"os"
)

func main() {
	app := app.Gerar()
	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
