#Reqr
====

Simple HTTP request inspector

This was made for a very specific task, to help debug calls that were sent from a course to an LRS using Tin Can (xAPI).

###To run

    go run reqr.go
    
or if you have go you can build reqr 

    go build reqr.go
    
when it runs it will get the ip address and this can be used as the end point


###Send Json

To send some json as the body as a test:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --data @example.json http://192.168.1.118:8666
