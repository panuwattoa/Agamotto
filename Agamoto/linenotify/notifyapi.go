package linenotify

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var lineToken = "Bearer KWvSoPKjsebCYoLBulS07CBXFvd0XevBQDAdriNH7lr"

func SendLineNotifyMessage(message string,stckerPacketID string, stickerID string) {
    log.Print("SendLineNotifyMessage ..")
	url := "https://notify-api.line.me/api/notify"

	var r http.Request
	r.ParseForm()
	r.Form.Add("message", message)
	r.Form.Add("stickerPackageId", stckerPacketID)
	r.Form.Add("stickerId", stickerID)
	body := strings.NewReader(r.Form.Encode())

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authorization", lineToken)

	client := &http.Client{Timeout: time.Second * 10}
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}
