1. init project
    go mod init github.com/dumpplane/test/test_rest_api

2. generate docs
   swag init   

3. run:
    go run main.go

4. swagger ui
    http://localhost:8080/swagger/index.html

5. build
    go build -o a.out main.go
    ./a.out

6. test via curl
    curl -X 'POST' 'http://localhost:8080/api/v1/accounts' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"name": "Kylin Soong"}'
    curl -X 'GET' 'http://localhost:8080/api/v1/accounts' -H 'accept: application/json'

    http://localhost:8080/swagger/index.html
