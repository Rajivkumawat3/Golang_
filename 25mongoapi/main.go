package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rajivkumawat3/25mongoapi/router"
)

func main() {
	// go mod init  ....package main
	//    go get -u github.com/gorilla/mux
	//   go get go.mongodb.org/mongo-driver/mongo

	fmt.Println("MongoDB API")
	r := router.Router()

	fmt.Println("server is getting started..........")

	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listning Port at  4000 ...")

}
