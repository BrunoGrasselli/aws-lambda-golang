# Golang AWS lambda
An example of an AWS lambda function written in golang

# Setup
```
npm install -g aws-sam-local
docker pull eawsy/aws-lambda-go-shim:latest
go get -u -d github.com/eawsy/aws-lambda-go-core/...
```

# Build
```
make
```

# Test
```
sam local invoke "SlackFunction" -e event.json
```
