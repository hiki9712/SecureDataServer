#!/bin/bash

# 定义请求的 URL 和 JSON 数据
URL="http://127.0.0.1:8808/api/v1/system/compute/getResult"

JSON='{
  "taskID": 1849743195534004200,
  "status": "string",
  "message": "string",
  "computeResult": [
    {
      "identifier": {
        "fieldName": "CUSTOMERID",
        "fieldValue": "1"
      },
      "result": "1555"
    }
  ]
}'


# 发送 HTTP POST 请求
curl -X POST -H "Content-Type: application/json" -d "$JSON" "$URL"