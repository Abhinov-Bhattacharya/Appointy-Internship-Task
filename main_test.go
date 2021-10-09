package main
	
	import (
		"bytes"
		//"context"
		"encoding/json"
		
		"net/http"
		"testing"
		//"time"
	
		//"go.mongodb.org/mongo-driver/mongo"
		//"go.mongodb.org/mongo-driver/mongo/options"
	)
	
	func TestMainCreateuser(t *testing.T) {
		var reguser user
		//var part participant
		reguser.Email = "Abhinovbhattacharya@gmail.com"
		reguser.Name = "Abhinov Bhattacharya"
		reguser.Password = "123456"
		// message.Title = "Title"
		// message.Participants = append(message.Participants, part)
		// message.Starttime = "2021-09-01T09:52:12+05:30"
		// message.Endtime = "2021-09-01T10:52:12+05:30"
		bytesRepresentation, _ := json.Marshal(reguser)
		resp, err := http.Post("http://localhost:12345/users", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Error("Fail")
		}
		if resp == nil {
			t.Error("NO response")

		}
	}