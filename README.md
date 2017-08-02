# Sailthru-Go

An unauthorized, and not guaranteed to work for you, implementation of the [Sailthru API](https://getstarted.sailthru.com/developers/overview/) in Go. Not all methods are currently implemented; they will be added as we need them. Pull requests welcomed.

## Examples

### Job

[Update](https://getstarted.sailthru.com/developers/api/job/#update)

```go
client := sailthru_client.NewClient(key, secret)

updater := sailthru_job.NewUpdate(client)
updater.Params.PostbackURL = "http://example.org/webhooks"
updater.Params.IncludeSignupDate = true

updater.ProcessURL("https://some-bucket.s3.amazonaws.com/path/to/file1.csv")
updater.ProcessURL("https://some-bucket.s3.amazonaws.com/path/to/file2.csv")
```
