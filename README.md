# Sailthru-Go

An unauthorized, and not guaranteed to work for you, implementation of the [Sailthru API](https://getstarted.sailthru.com/developers/overview/) in Go. Not all methods are currently implemented; they will be added as we need them. Pull requests welcomed.

## Examples

### Send
[Template](https://getstarted.sailthru.com/developers/api/send/#POST_to_Send_Schedule_or_Update)

Note: If either `StartTime` or `EndTime` is set, then `ScheduleTime` is ignored (even if also set).

```go
client := sailthru_client.NewClient(key, secret)

send := sailthru_send.NewSingle(client)

// send.Params.CC = "me@example.org"
// send.Params.BCC = "me2@example.org"
// send.Params.ReplyTo = "reply-to@example.org"
// send.Params.Test = true
// send.Params.DataFeedURL = "https://some-bucket.s3.amazonaws.com/path/to/feed"
// send.Params.BehalfOf = "admin@example.rog"
// send.Params.ScheduleTime = "2006-01-02 15:04:05 UTC"
// send.Params.StartTime = "2006-01-02 15:04:05 UTC"
// send.Params.EndTime =  = "2006-01-02 15:04:05 UTC"

vars := map[string]string{"name": "Bob", "welcome": "Hello!"}
send.Deliver("you@example.org", "A-Template", vars)
```


### Job

[Update](https://getstarted.sailthru.com/developers/api/job/#update)

```go
client := sailthru_client.NewClient(key, secret)

updater := sailthru_job.NewUpdate(client)
updater.Params.PostbackURL = "http://example.org/webhooks"
updater.Params.ReportEmail = 'me@example.org'
updater.Params.IncludeSignupDate = true

updater.ProcessURL("https://some-bucket.s3.amazonaws.com/path/to/file1.csv")
updater.ProcessURL("https://some-bucket.s3.amazonaws.com/path/to/file2.csv")
```
