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
    
    export GOPATH=~pwd/ 
    
    export GOBIN=$GOPATH/bin
    
    go get
    
    go run Main.go Cost.go Logger.go Routes.go Uber.go Handlers.go Router.go 

There are three rides in the database:
    the database looks like:
    
    id | start         | distance | reqid   | totalcharged
    
    1  | 1468972800    |       10 | someid1 |           15
    
    2  | 1469059200    |        3 | someid2 |            5
    
    3  | 1469145600    |        2 | someid3 |            3
    
    The time interval that client requests uses the number in the start column.


Request: 
    
    GET http://localhost:8080/intervalhistory/1468972800/1469059200

Response: 
    
    {"distance":13,"totalcharged":20}
    
    This response includes interval of entry with id 1 and entry with id 2

Request: 
    
    GET http://localhost:8080/intervalhistory/1469145600/1469145600

Response: 
    
    {"distance":2,"totalcharged":3}
    
    This response only includes entry with id 3

Request:
    
    GET http://localhost:8080/intervalhistory/1469059200/1469145600

Response:
    
    {"distance":5,"totalcharged":8}
    
    This response includes entry with id 1 and id 2

Request:
    
    GET http://localhost:8080/intervalhistory/1468972800/1469145600

Response:
    
    {"distance":15,"totalcharged":23}
    
    This response includes entry with id 1, id 2, and id 3

Next steps:

     1) Make tests with GO

     2) Bug when no entries are retrieved from query, server crashes

     3) Use OAUTH to sign in and populate database with UBER API calls

     
