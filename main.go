package main

import (
  "encoding/json"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
  "net/http"
  "bytes"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  payload := "{\"text\":\"Testing AWS lambda.\"}"
  response, err := http.Post("https://hooks.slack.com/services/x/y/z","application/json",bytes.NewBuffer([]byte(payload)))

  return response.Body, err
}
