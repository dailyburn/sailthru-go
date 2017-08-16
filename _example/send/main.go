package main

import (
	"flag"
	"log"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/dailyburn/sailthru-go/send"
)

func main() {
	key := flag.String("key", "", "Sailthru Key (Required)")
	secret := flag.String("secret", "", "Sailthru Secret (Required)")
	template := flag.String("template", "", "Email Template (Required)")
	email := flag.String("email", "", "Email Address (Required)")
	flag.Parse()

	client := client.NewClient(*key, *secret)
	send := send.NewSingle(client)
	vars := map[string]string{}

	res, err := send.Deliver(*email, *template, vars)

	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("send_id:", res.GetResponse()["send_id"])
}
