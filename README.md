#Reqr
====

Simple HTTP request inspector - this is a tiny web server that is waiting for post / get requests and it will output the GET / POST parameters and the JSON body.

This was made for a very specific task, to help debug calls that were sent from a course to an LRS using Tin Can (xAPI).

Following the YAGNI (You Arn't Gonna Need It) approach, I've only developed this for a very specific use case - and it might not meet general purpose needs.  Over time when I need to adapt it I'll make changes.

You'll need go to run this.... 

###To run

    go run reqr.go
    
or to build

    go build reqr.go
    
Once build, then copy to a location in your path so you can run - for example /usr/local/bin    
    
when it runs it will get the ip address and this can be used as the end point


###Send Json

To send some json as the body as a test:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --data @example.json http://192.168.1.118:8666
