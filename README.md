# go-pubsub

This repository is for my article on Medium about Messaging with PubSub and Golang.

# How to run the project

Replace the [service_account.json](https://github.com/moemoe89/go-pubsub/blob/main/service_account.json) with your key.

```json
{
  "type": "service_account",
  "project_id": "",
  "private_key_id": "",
  "private_key": "",
  "client_email": "",
  "client_id": "",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": ""
}

```

Set environment variable for the `service_account.json` path.
```shell
export GOOGLE_APPLICATION_CREDENTIALS="path/service_account.json"
```

Run the publisher server
```shell
go run ./cmd/publisher/main.go
```

Run the subscriber server
```shell
go run ./cmd/subscriber/main.go
```

Run the publisher server with docker-compose
```shell
docker-compose up publisher
```

Run the subscriber server with docker-compose
```shell
docker-compose up subscriber
```

Publish a message by access the endpoint

[http://localhost:8080/publish](http://localhost:8080/v1/publish)
