package main

import (
	"fmt"
	"log"
	"net/http"
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

func testingStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey the test is working")
}

func main() {
	fmt.Println("App running")
	port:= os.Getenv("PORT")
	fmt.Println(port)

	mux := http.NewServeMux()
	mux.HandleFunc("/test", testingStuff)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}