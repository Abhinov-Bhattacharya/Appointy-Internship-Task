package main

import (
"context"
"encoding/json"
"errors"
"fmt"
"net/http"
"path"
"time"

"go.mongodb.org/mongo-driver/bson/primitive"
)


func postwithId(id primitive.ObjectID) (post, error) {
var Posts post
collection := client.Database("Appointy").Collection("Instagramposts")
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
err := collection.FindOne(ctx, post{Id: id}).Decode(&Posts)
if Posts.Id != id {
err = errors.New("Error 400: ID not present")
}
return Posts, err
}

//Getpost : Searches a post in the database
func GetpostwithID(response http.ResponseWriter, request *http.Request) {
if request.Method != "GET" {
response.WriteHeader(http.StatusMethodNotAllowed)
response.Write([]byte(`{ "message": "Incorrect Method" }`))
return
}
response.Header().Set("content-type", "application/json")
fmt.Println(path.Base(request.URL.Path))
id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
getpostwithId, err := postwithId(id)
if err != nil {
response.WriteHeader(http.StatusBadRequest)
response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
return
}
json.NewEncoder(response).Encode(getpostwithId)

}

