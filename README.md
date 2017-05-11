# golang_speedtest

A simple Go speed test to time an http response from Rails and Go.

It fetches the response body for a demo page of each framework running on Heroku with 1 Dyno. 

This is not a controlled test and is only being done to learn Go tools.

Toby Jones


sample output:

$ go run retrieve.go
http transport error is: <nil>
read error is: <nil>
    
    <title>RailsHku</title>

Retrieve time was:  390  Milliseconds.

rails response time:  390
http transport error is: <nil>
read error is: <nil>

    <title>Toby's Golang Demo</title>

Retrieve time was:  325  Milliseconds.

golang response time:  325
