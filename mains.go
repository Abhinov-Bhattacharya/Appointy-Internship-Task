package main

import (
	
    "context"
    "fmt"
    "net/http"
    "time"
    
    "go.mongodb.org/mongo-driver/mongo"
   
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client





func main() {
    fmt.Println("Application is Running")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
    client, _ = mongo.Connect(ctx, clientOptions)
    http.HandleFunc("/users", Userhandler)
    http.HandleFunc("/posts", Posthandler)
    http.HandleFunc("/users/", GetUserhandler)
    http.HandleFunc("/posts/", GetPosthandler)
    http.HandleFunc("/posts/users/", Getusersposthandler)
    http.ListenAndServe(":12345", nil)
}

	
	
	
