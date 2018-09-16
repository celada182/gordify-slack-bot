package main

import (
	"fmt"
	"os"
	"math/rand"
	"strings"
	"github.com/nlopes/slack"
)

var command = "gordify"
var	start = "start"
var	stop = "stop"
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
		organizeGroups(users, event)
	} else if isConfirmation(event.Text) {
		users = addUser(users, event)
	}
}

func checkCommand(message string, isStart bool) bool {
	var function string
	if function = stop; isStart  {
	    function = start
	}
	commandMessage := command + " " + function
	return message == commandMessage && !active == isStart
}

func isConfirmation(message string) bool {
	info := rtm.GetInfo()
	prefix := fmt.Sprintf("<@%s>", info.User.ID)
	return active && strings.HasPrefix(message, prefix)
}

func addUser(users []string, event *slack.MessageEvent) []string {
	if (!contains(users, event.User)){
		users = append(users, event.User)
		fmt.Printf("Users: %v\n", users)
		rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("<@%s> is in!", event.User), event.Channel))
	}
	return users
}

func organizeGroups(users []string, event *slack.MessageEvent) {
	//users = []string{"a","b","c","d","e","f","g","h","i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w"}
	numberOfGroups, groupSize := getNumberOfGroups(len(users))
	groups:=getGroups(users, numberOfGroups, groupSize)
	for index, group := range groups {
		if (len(group) > 0){
			groupMessage(index, group, event)
		}
	}
}

func groupMessage(index int, group []string, event *slack.MessageEvent) {
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Group %d", index + 1), event.Channel))
	for _, user := range group {
		rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("<@%s>", user), event.Channel))
	}
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Lider <@%s>", getGruopLider(group)), event.Channel))
}

func getNumberOfGroups(users int) (int, int) {
	groupSize:= 7
	groups:=0
	lastGruopSize:=0
	for {
		groups, lastGruopSize = divmod(users, groupSize)
		if (lastGruopSize > 1 || groups == 0) {
			groups++
		}
 		if (lastGruopSize == 0 || lastGruopSize == groupSize - 1 || lastGruopSize == 1) {
			break
 		}
 		groupSize--
	}
 	return groups, groupSize
}

func getGroups(users []string, numberOfGroups int, groupSize int) [][]string {
	var groups [][]string
	for i:=1; i <= numberOfGroups; i++ {
		start := (i - 1) * groupSize
		end := (i * groupSize)
		if (i == numberOfGroups) {
			end = len(users)
		}
		groups = append(groups, users[start:end])
	}
	return groups
}

func getGruopLider(group []string) string {
	if (len(group) == 1) {
		return group[0]
	}
	index := rand.Intn(len(group) - 1)
	return group[index]
}

func divmod(numerator, denominator int) (quotient, remainder int) {
    quotient = numerator / denominator
    remainder = numerator % denominator
    return
}

func contains(slice []string, search string) bool {
    for _, value := range slice {
        if value == search {
            return true
        }
    }
    return false
}