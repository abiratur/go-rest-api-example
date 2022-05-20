# go-rest-api-example
A simple rest API service in go  

## Notes
- passwords were taken from: https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt
- in production this service must be run behind an nginx with https only enabled
- GET methods should not have bodies so i changed the http method to POST

## Building
run this command from root repository folder:  
```go build```  
or if using docker:  
```docker build -t go-rest-api .```

## Running
run this command from root  
```go run main.go```  
or if using docker:  
```docker run --rm -it -p 8080:8080 go-rest-api```

## Testing

### Manual
1. test a password is valid  
```curl -d "{\"password\": \"RqGVydeNAU;,p7f*\"}" -H "Content-Type: application/json" -X POST localhost:8080/v1/validatePassword```

2. test a password is not valid  
```curl -d "{\"password\": \"123456\"}" -H "Content-Type: application/json" -X POST localhost:8080/v1/validatePassword```

3. test a request is not valid  
```curl -d "{ }" -H "Content-Type: application/json" -X POST localhost:8080/v1/validatePassword```

## TODO
- move REST api endpoint to a different file
- use redis as a cache for common passwords
- check idiomatic way of handling global state in golang / gin for the 'passwords' map
- dynamically download passwords during build time (in docker build) 
- use multi stage build to reduce image size
