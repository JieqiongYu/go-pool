# Sunrise Sunset Demo

It's a more "realistic" example of how one might use Cloud Functions. It implements a basic version of the [sunrise-sunset.org API](https://sunrise-sunset.org/api) (which provides sunrise/sunset times for a given date and location).

## Declaring Dependencies 
For any non-trivial web service, we'll need to depend on existing Go libraries. Cloud Functions uses [go modules](https://github.com/golang/go/wiki/Modules) to define application dependencies. To declare the dependencies we need for our API, we'll first initialize a `mod.go` file in our project with `go mod init`, and then `go get` our dependencies. 

```bash
go mod init sundemo
go get github.com/nathan-osman/go-sunrise github.com/gorilla/schema
```

## How it works about `dependencies`

Upon deployment, Cloud Functions uploads `application code` (not a compiled binary) to a build server, and fetches dependencies remotely. 

Thus, we should not rely on any local dependencies that we won't be accessible by the Cloud Functions build server. 

## Commands about GCP project

```bash
# Check current project
gcloud config get-value core/project
# Set project
gcloud config set project <PROJECT_ID>
```

## Deploy

```bash
# cd /path/to/Sunrise-Sunset
gcloud functions deploy SunTimes --runtime go111 --trigger-http
```

## Test

```bash
curl "https://us-central1-kouzoh-p-jieqiong-yu.cloudfunctions.net/SunTimes?lat=47.6&lon=122.3&date=2019-05-27"
```

Response

```bash
{"sunset":"2019-05-27T11:34:25Z","sunrise":"2019-05-26T20:01:40Z"}
```

## Reference
https://benjamincongdon.me/blog/2019/01/21/Getting-Started-with-Golang-Google-Cloud-Functions/