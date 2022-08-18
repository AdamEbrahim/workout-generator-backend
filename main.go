package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"
	"github.com/joho/godotenv"
	"errors"

)

func init() {
	if _, err := os.Stat("./.env"); errors.Is(err, os.ErrNotExist) {
		// .env file does not exist (production)
		fmt.Println("cant find env")
	} else {
		// .env file exists (development)
		err2 := godotenv.Load(".env")

		if err2 != nil {
			log.Fatal("Error loading .env file")
		}
	}

	

}

func main() {
	fmt.Println("App running")
	port:= os.Getenv("PORT")
	fmt.Println(port)

	//log.fatal(http.ListenAndServe())
}