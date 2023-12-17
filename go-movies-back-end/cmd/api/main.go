package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = 8080
)

type application struct {
	Domain string
}

func main() {

	// установить конфигураци приложения
	app := application{}
	app.Domain = "example.com"
	//db := os.Getenv("MOVIES")
	// лог о запуске сервера
	log.Println("Starting application in port:", port)
	// слушаем указанный порт
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
