# golang_speedtest

by Toby Jones

A simple Go speed test to time an http response from Rails and Go web apps. It fetches the response body for a demo page of each framework running on Heroku with 1 Dyno. 

This is not a controlled test and is only being done to learn Go tools.

I've added goroutines to do asyncronous function calls. Now the fastest site returns first like a proper race.




sample output:

$ go run retrieve.go

  RailsHku

Retrieve time was:  390  Milliseconds.

rails response time:  390


 Toby's Golang Demo

Retrieve time was:  325  Milliseconds.

golang response time:  325
