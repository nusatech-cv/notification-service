package firebase

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"notification/token"
	"io/ioutil"
)

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Message struct {
	Notification Notification `json:"notification"`
	To           string       `json:"to"`
}

func SendNotification(record *token.Record) {
	msg := Message{
		Notification: Notification{
			Title: record.Title,
			Body:  record.Message,
		},
		To: record.User.TokenDevice,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal message: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", "key="+os.Getenv("SERVER_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response status: %s", resp.Status)
		log.Printf("Response body: %s", string(respBody))
		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to send message, status: %s", resp.Status)
		} else {
			log.Println("Successfully sent message")
		}
	}
}
