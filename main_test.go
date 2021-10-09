package main
	
	import (
		"bytes"
		//"context"
		"encoding/json"
		
		"net/http"
		"testing"

	)
	
	func TestMainCreateuser(t *testing.T) {
		var reguser user
		//var part participant
		reguser.Email = "Abhinovbhattacharya@gmail.com"
		reguser.Name = "Abhinov Bhattacharya"
		reguser.Password = "123456"
	
		bytesRepresentation, _ := json.Marshal(reguser)
		resp, err := http.Post("http://localhost:12345/users", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Error("Fail")
		}
		if resp == nil {
			t.Error("NO response")

		}
	}