This is an example on creating a simple webserver to connect to regiboxd using REST calls.

IMPORTANT: Before running
1) you need to define the config directory of your regiboxd instance!
2) you need to define listing interface and port (default localhost:8081)
3) the regiboxctl command must be on PATH

If echo framework is missing, try running 

 go mod tidy

to install any missing dependencies.


In order to do a REST call, you can execute function calls like this:

 curl -X POST http://localhost:8081/post -H "Content-Type: application/json" -d '{"data": "{\"op\": \"someFunction\"}"}'

