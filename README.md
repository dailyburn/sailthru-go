# Sailthru-Go

An unauthorized, and not guaranteed to work for you, implementation of the [Sailthru API](https://getstarted.sailthru.com/developers/overview/) in Go. Not all methods are currently implemented; they will be added as we need them. Pull requests welcomed.

## Examples

### Send
Sailthru [Template](https://getstarted.sailthru.com/developers/api/send/#POST_to_Send_Schedule_or_Update) docs

```go
client := sailthru.NewClient(key, secret)

vars := map[string]string{"name": "Bob", "welcome": "Hello!"}

params := &params.Send{
  Email:    "you@example.org",
  Template: "A-Template",
  Vars:     vars,
}

res, err := client.Send(params)
```

You can also set a schedule time;

```go
params.ScheduleTime = &ScheduleTime{
  StartTime: "Tue, 27 Jul 2010 12:10:00 -0400",
  EndTime: "Tue, 28 Jul 2010 12:10:00 -0400", //Optional
}
```

### Job
Sailthru [Update Job](https://getstarted.sailthru.com/developers/api/job/#update) docs

```go
client := sailthru.NewClient(key, secret)

params := &params.UpdateJob{
  PostbackURL:       "http://example.org/webhooks",
  ReportEmail:       "you@example.org",
  IncludeSignupDate: true,
  URL:               "https://some-bucket.s3.amazonaws.com/path/to/file1.json",
}

res, err := client.ProcessUpdateJob(params)
```
