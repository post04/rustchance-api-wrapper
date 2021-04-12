package main

import (
	"fmt"

	wrapper "github.com/post04/rustchance-api-wrapper"
)

func main() {
	session, err := wrapper.New("token" /*this can be blank, if it is blank it will not be able to use auth restricted http*/, []string{ /*I input nothing because I wanted it to default to listen for all*/ }, "" /*again I input nothing so it will default to EN*/)
	if err != nil {
		panic(err)
	}
	earnings, err := session.AccountEarnings()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf("%+v\n", earnings)
}
