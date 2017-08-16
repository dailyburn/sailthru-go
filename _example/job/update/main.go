package main

import (
	"flag"
	"log"

	"github.com/dailyburn/sailthru-go/client"
	"github.com/dailyburn/sailthru-go/job"
)

func main() {
	key := flag.String("key", "", "Sailthru Key (Required)")
	secret := flag.String("secret", "", "Sailthru Secret (Required)")
	url := flag.String("url", "", "URL of update .csv file (Required)")
	postback := flag.String("postback", "", "URL to postback results to (Required)")
	email := flag.String("email", "", "Email Address (Required)")
	flag.Parse()

	client := client.NewClient(*key, *secret)

	updater := job.NewUpdate(client)
	updater.Params.PostbackURL = *postback
	updater.Params.ReportEmail = *email
	updater.Params.IncludeSignupDate = true

	res, err := updater.ProcessURL(*url)

	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("job_id:", res.GetResponse()["job_id"])
}
