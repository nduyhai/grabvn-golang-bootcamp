# Assignment 4

## Docs

### Mock test

* [Google moc](https://github.com/golang/mock)
* [Testify](https://github.com/stretchr/testify)

### Integration test

* [Pact](https://docs.pact.io/)


## Requirement

Complete the unit tests and integration tests for all endpoints of the ​to-do​ service.
* A sample unit test is provided for the​ GetToDo ​method.
* A sample integration test is provided for the​ Insert ​method of ToDo repository.
* Use ​pact-go​ to write integration tests for endpoints of the to-do service.
* All logic of the​ service ​package is unit-tested.
* Mock all dependencies in unit tests.
* Write integration tests for all methods inside​ repository ​package to verify that data
is saved into database.
* Write integration tests for all endpoints of the to-do service using ​pact-go