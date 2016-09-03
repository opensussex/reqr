#Reqr
====

Simple HTTP request inspector

##Send Json

To send some json as the body as a test:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST --data @example.json http://127.0.0.1:8666
