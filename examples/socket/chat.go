package main

import (
	"fmt"

	wrapper "github.com/post04/rustchance-api-wrapper"
)

func getRooms(session *wrapper.Session, rooms *wrapper.ChatRooms) {
	session.SwitchRoom("en")
}

func onMessage(session *wrapper.Session, message *wrapper.ChatMessage) {
	fmt.Printf("User %s sent message %s\n", message.Data.Profile.Username, message.Data.Content)
}

func main() {
	session, err := wrapper.New("token" /*this can be blank, if it is blank it will not be able to use auth restricted http*/, []string{ /*I input nothing because I wanted it to default to listen for all*/ }, "" /*again I input nothing so it will default to EN*/)
	if err != nil {
		panic(err)
	}
	session.AddHandler(getRooms)
	session.AddHandler(onMessage)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	select {} // this keeps the program running
}
