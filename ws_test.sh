#!/bin/bash

# 定义请求的 URL 和 JSON 数据
URL="http://127.0.0.1:8808/ws"

JSON='{
  "taskID": 1843272151252078600,
  "status": "string",
  "message": "string",
  "computeResult": [
    {
      "identifier": {
        "fieldName": "CUSTOMERID",
        "fieldValue": "17"
      },
      "result": "1269"
    }
  ]
}'


# 发送 HTTP POST 请求
curl -X POST -H "Content-Type: application/json" -d "$JSON" "$URL"