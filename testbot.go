package main

import (
	"fmt"
	"os"
	// "strings"

	"github.com/nlopes/slack"
)

var command = "gordify"
var	start = "start"
var	stop = "stop"
var	confirmation = "count me in"
var	active = false
var	users = []string{}
var token = os.Getenv("SLACK_TOKEN")
var	api = slack.New(token)
var	rtm = api.NewRTM()

func main() {
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)
			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev.Text)
				handleMessageEvent(ev)
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
				//Take no action
			}
		}
	}
}

func handleMessageEvent(event *slack.MessageEvent) {
	if checkCommand(event.Text, true) {
		active=true
		rtm.SendMessage(rtm.NewOutgoingMessage("Ey! Who is going to have lunch out today?", event.Channel))
	} else if checkCommand(event.Text, false) {
		active=false
		rtm.SendMessage(rtm.NewOutgoingMessage("Ok, lets have lunch!", event.Channel))
	} else if isConfirmation(event.Text) {
		users = addUser(users, event)
	}
}

func checkCommand(message string, isStart bool) bool {
	var function string
	if function = stop; isStart  {
	    function = start
	}
	var commandMessage = command + " " + function
	return message == commandMessage && !active == isStart
}

func isConfirmation(message string) bool {
	return active && message == confirmation
}

func addUser(users []string, event *slack.MessageEvent) []string {
	if (!contains(users, event.User)){
		users = append(users, event.User)
		fmt.Printf("Users: %v\n", users)
		rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("<@%s> is in!", event.User), event.Channel))
	}
	return users
}

func contains(slice []string, search string) bool {
    for _, value := range slice {
        if value == search {
            return true
        }
    }
    return false
}