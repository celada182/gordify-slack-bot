package main

import (
	"fmt"
	"os"
	// "strings"

	"github.com/nlopes/slack"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	command := "gordify"
	start := "start"
	stop := "stop"
	confirmation := "count me in"
	active := false
	users := []string{}
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev.Text)
				if ev.Text == command + " " + start {
					active=true
					rtm.SendMessage(rtm.NewOutgoingMessage("Ey! Who is going to have lunch out today?", ev.Channel))
				} else if ev.Text == command + " " + stop {
					active=false
					rtm.SendMessage(rtm.NewOutgoingMessage("Ok, lets have lunch!", ev.Channel))
				} else if active && !contains(users, ev.User) && ev.Text == confirmation {
					users = append(users, ev.User)
					fmt.Printf("Users: %v\n", users)
					rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("<@%s> is in!", ev.User), ev.Channel))
				}

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

func contains(slice []string, search string) bool {
    for _, value := range slice {
        if value == search {
            return true
        }
    }
    return false
}