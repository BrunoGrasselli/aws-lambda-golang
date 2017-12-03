# AWS Lambda with Golang
An example of an AWS lambda function written in golang. This function sends a slack message using incoming webhooks.

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
