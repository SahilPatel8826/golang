package main

import (
	"fmt"
	"net/http"

	controllers "mongo-golang/Controllers"
	db "mongo-golang/connectdb.go"

	"github.com/julienschmidt/httprouter"
)

func main() {

	k := httprouter.New()

	client := db.ConnectDB()
	controllers.InitCollections(client)
	k.GET("/user/:id", controllers.GetUser)
	k.POST("/user", controllers.CreateUser)
	// r.PUT("/user/:id", uc.UpdateUser)
	// r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("mongo db connected:")
	if err := http.ListenAndServe(":8080", k); err != nil {
		fmt.Println("❌ Server failed:", err)
	}

}
