package main

import (
"context"
"encoding/json"
"fmt"
"net/http"
"time"

"go.mongodb.org/mongo-driver/bson/primitive"
)
//CreateMeeting : Adds another post to the database
func Createpost(response http.ResponseWriter, request *http.Request) {

response.Header().Set("content-type", "application/json")
var Post post
_ = json.NewDecoder(request.Body).Decode(&Post)
Post.def()
lock.Lock()
defer lock.Unlock()
collection := client.Database("Appointy").Collection("Instagramposts")
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
result, _ := collection.InsertOne(ctx, Post)
Post.Id = result.InsertedID.(primitive.ObjectID)
json.NewEncoder(response).Encode(Post)
fmt.Println(Post)

}