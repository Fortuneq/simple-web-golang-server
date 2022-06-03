package main

import (
	"flag"
	"log"
)

func main() {
	t:= MustToken()
	//tgClient = telegram.New(token)
	//fetcher = fetcher.New(tgClient)
	//processor = processor.New(tgClient)
	//consumer.Start(fetcher, processor)
}
func MustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == ""{
		log.Fatal("Not acces to tg token")
	}
	return *token
}
