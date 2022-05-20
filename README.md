# go-rest-api-example
A simple rest API service in go

## Building
run this command from root repository folder  
```go build```

## Running
run this command from root  
```go run main.go```

## Testing

### Manual
1. test a password is valid

curl -d "{\"password\": \"RqGVydeNAU;,p7f*\"}" -H "Content-Type: application/json" -X POST localhost:8080/v1/validatePassword

2. test a password is not valid

curl -d "{\"password\": \"\"}" -H "Content-Type: application/json" -X POST localhost:8080/v1/validatePassword