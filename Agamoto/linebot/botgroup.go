package linebot

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
)


func RecvMessageLineBot(w http.ResponseWriter, r *http.Request)  {
	log.Print("On recv")
	client := &http.Client{}

	bot, err := linebot.New("ff048454e8eccccd34ff7506335de252", "F1t6Fuj83dNKbGUu1X9K0oEpgmMSCcAGi5WZnlUGb73ayEznUh+iMLOEY6zObIf+1VD9Wycjz6nwUy6T3n5ullGjxAnN/w66bJi4YsL1glRv+KE8LpC32yXXM20GRc3urFNaEru7+WL8vXaQ+qOcbwdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	if err != nil {
		log.Print("Can't crate line bot ")
	}


	events, err := bot.ParseRequest(r)
	if err != nil {
		// Do something when something bad happened.
		log.Print("Can't Parse Data ")
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// Do Something...
			log.Print(event.Message)
			switch text := event.Message.(type) {
			case *linebot.TextMessage:
				if text.Text  == "สวัสดี" {
					// groupID := event.Source.GroupID
					// var messages []linebot.SendingMessage
					// leftBtn := linebot.NewMessageAction("left", "left clicked")
					// rightBtn := linebot.NewMessageAction("right", "right clicked")
					// // if event.Message
					// template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
					// message := linebot.NewTemplateMessage("Sorry :(, please update your app.",template)
					// messages = append(messages, message)
					// _, err := bot.PushMessage(groupID, messages...).Do()
					// if err != nil {
					// 	// Do something when some bad happened
					// }
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}

		}
	}


}