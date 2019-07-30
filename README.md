# Assignment 2: feedback

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

