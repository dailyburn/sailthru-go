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
	url := flag.String("url", "", "URL of update .csv file (Required)")
	postback := flag.String("postback", "", "URL to postback results to (Required)")
	email := flag.String("email", "", "Email Address (Required)")
	flag.Parse()

	client := sailthru.NewClient(*key, *secret)

	params := &params.UpdateJob{
		PostbackURL:       *postback,
		ReportEmail:       *email,
		IncludeSignupDate: true,
		URL:               *url,
	}

	res, err := client.ProcessUpdateJob(params)

	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("job_id:", res.GetResponse()["job_id"])
}
