# Notification Project

This project focuses on implementation of a notification system with Golang, Firebase, and RabbitMQ. It processes message queue from RabbitMQ and sends notifications through Firebase Cloud Messaging (FCM).

## Project Structure

The project's directory structure is organized as follows:

```
.
├── Dockerfile
├── firebase
│   └── firebase.go
├── go1.17.5.linux-amd64.tar.gz
├── go.mod
├── go.sum
├── key
│   └── serviceAccountKey.json
├── main.go
├── rabbitmq
│   └── rabbitmq.go
└── token
    └── token.go
```

## Usage

### Build & Run

First, ensure that Golang (version 1.17.5 or above) is installed on your system. Then, build the project using Go:

```bash
$ go build
```

Run the application:

```bash
$ ./main
```

### Docker Build

The Dockerfile allows for building Docker image:

```bash
$ docker build -t your-docker-username/notification-project .
```

Then run the image:

```bash
$ docker run -p 8080:8080 -d your-docker-username/notification-project
```

Remember to replace "your-docker-username" with your actual Docker username.

## Key File

The application uses Firebase for sending notifications. For this, you need a Firebase service account key `serviceAccountKey.json`, which is stored in the `key/` directory. Please replace this file with your own key.

## RabbitMQ

The `main.go` file starts a server that listens to a RabbitMQ service. The `rabbitmq/rabbitmq.go` file contains logic for connecting to the RabbitMQ service and listening to the message queue. It processes received messages, and sends notification.

## Firebase

This project leverages Firebase for sending notifications. The corresponding logic can be found within `firebase/firebase.go`. 

## Token

The `token/token.go` file is used to validate the token associated with each request. Write your own logic according to your use case.

Remember to set the RabbitMQ server address and Firebase key in their respective files or through environment variables for a production setting.

## Tests

No automated tests are written for this project as of now.

## Versioning

Check the `go.mod` and `go.sum` files for dependency versioning. 

For the version of Go, refer to `go1.17.5.linux-amd64.tar.gz`. The Go version used for this project is `1.17.5`.

Note: Please make sure to replace the service key file and set up RabbitMQ properly for successful operation.