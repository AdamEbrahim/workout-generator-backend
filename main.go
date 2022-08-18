package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"
	"github.com/joho/godotenv"

)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("App running")
	port:= os.Getenv("PORT")
	fmt.Println(port)

	//log.fatal(http.ListenAndServe())
}