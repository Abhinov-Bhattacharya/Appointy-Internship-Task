package main

import ("net/http"
	)

//UserHandler : Function to handle "/users" and call CreateUser or GetUser accoording to request
func Userhandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Method)
	if request.Method == "POST" {
		Createuser(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
//PostHandler : Function to handle "/posts" and call Createposts or Getposts accoording to request

func Posthandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		Createpost(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
func GetUserhandler(response http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		GetuserwithID(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
func GetPosthandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		GetpostwithID(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
func Getusersposthandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		Getuserspostlist(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}







