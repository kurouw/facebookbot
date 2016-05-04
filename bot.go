package main

import (
	"github.com/facebookbot/reqCafe"
	"github.com/facebookbot/fbmessenger"
	"regexp"
	"time"
)


type DistributeMenu struct {
	Judgment []string
	Jf       bool
}

func main() {
	fbmessenger.Listen(handleRecieveMessage)
}

func handleRecieveMessage(event fbmessenger.Messaging) {
	recipient := new(fbmessenger.Recipient)
	recipient.ID = event.Sender.ID
	fbmessenger.SendTextMessage(*recipient, getMessageText(event.Message.Text))
}

func selectMenu(txt string) string {
	foods := new(DistributeMenu)
	foods.Judgment = []string{"kondate","こんだて","献立", "学食","めにゅー", "メニュー","menu"}
	foods.Jf = false

	computers := new(DistributeMenu)
	computers.Judgment = []string{"演習室", "パソコン", "pc"}
	computers.Jf = false

	eves := new(DistributeMenu)
	eves.Judgment = []string{"hoge"}
	eves.Jf = false

	stringnames := []string{"foods","computers","eves"}
	allEvents := []DistributeMenu{*foods,*computers,*eves}
	
	for i := range allEvents { 
		for j := 0; j < len(allEvents[i].Judgment); j++ {
			r := regexp.MustCompile(allEvents[i].Judgment[j])
			if r.MatchString(txt) {
				allEvents[i].Jf = true
			}
		}
	}
	
	for i := range allEvents {
		if allEvents[i].Jf {
			allEvents[i].Jf = false
			return stringnames[i]
		}
	}
	return "notthing"
}

func getMessageText(receivedText string) string {
	if selectMenu(receivedText) == "foods" {
 		return reqCafe.RtCafeInfo(time.Now())
	}
	return receivedText
}
