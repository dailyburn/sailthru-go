package main

import (
	"flag"
	"log"

	sailthru "github.com/dailyburn/sailthru-go"
	"github.com/dailyburn/sailthru-go/params"
)

func main() {
	key := flag.String("key", "", "Sailthru Key (Required)")
	secret := flag.String("secret", "", "Sailthru Secret (Required)")
	template := flag.String("template", "", "Email Template (Required)")
	email := flag.String("email", "", "Email Address (Required)")
	flag.Parse()

	client := sailthru.NewClient(*key, *secret)

	params := &params.Send{
		Email:    *email,
		Template: *template,
	}

	res, err := client.Send(params)

	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("send_id:", res.GetResponse()["send_id"])
}
