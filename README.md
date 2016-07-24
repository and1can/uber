Uber Backend API
================

Uber backend api endpoint that returns total distance traveled and amount charged for Uber rides in certain time interval (time interval is defined by Unix time stamp) written in GO version: go1.6.3 darwin/amd64
----------------------------------------------------------------------

First, install necessary packages.
Then install database so testing will have same data. Make sure to create database called uber in postgresql before running node database.js.

    npm install

    node database.js


Database contents can be populated with two Uber api requests: 
- GET /v.12/history 
- GET /v1/requests/{request_id}/receipt

To run code do usual:
- export GOPATH=~pwd/ 
- export GOBIN=$GOPATH/bin
- go get
-go run Main.go Handlers.go Routes.go Router.go Todo.go repo.go Logger.go 

To run tests:
    go run test
