# Assignment 3: feedback

## Feedback
Implement a simple passenger feedback service, with basic functions:

* Add passenger feedback
* Get by passenger id
* Get by booking code
* Delete by passenger id

Requirements:

* Implement gRPC server/client
* Simply use local variable as storage
* 1 booking has only 1 feedback
* 1 passenger can add multiple feedbacks

```proto
message PassengerFeedback {
    string bookingCode = 1;
    int32 passengerID = 2;
    string feedback = 3;
}
```
## Additional

* [Secure gRPC](https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html)

* 

## Setup

### generate proto

```shell script
protoc -I api/ api/feedback.proto --go_out=plugins=grpc:internal/bootcamp/feedback
```

### Create cert

```shell script
bash scripts/cert.sh
```

### Run postgres
```shell script
cd deployments
docker-compose up
```