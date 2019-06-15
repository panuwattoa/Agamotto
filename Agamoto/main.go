package main

import (
	"Agamoto/linebot"
	"Agamoto/monitoring"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("----- Start server at port 80 -----")
 	// go monitoring.CheckServerAlive()
	http.HandleFunc("/updateCCU", monitoring.RecvUpdateCCU)
	http.HandleFunc("/linebot",linebot.RecvMessageLineBot)
	log.Fatal(http.ListenAndServe(":80", nil))
}



