#!/bin/bash

HOST="localhost"           # 目标主机
PORT=8080                  # 目标端口
PATH="/api/data"           # 请求路径
DATA='{
          "serviceName": "123",
          "serviceOwnerID": 12,
          "providerList": {
              "12345": [
                  {
                      "databaseName": "raw_sg",
                      "tableName": "POB_POINT"
                  },
                  {
                      "databaseName": "raw_sg",
                      "tableName": "POB_CUSTOMER"
                  }
              ],
              "12346": [
                  {
                      "databaseName": "raw_sg",
                      "tableName": "POB_POINT"
                  },
                  {
                      "databaseName": "raw_sg",
                      "tableName": "POB_CUSTOMER"
                  }
              ]
          }
      }'  # POST数据

# 构造完整的URL
URL="http://${HOST}:${PORT}${PATH}"

echo "发送POST请求到: $URL"
echo "数据: $DATA"

# 使用curl发送POST请求
curl -X POST -d "$DATA" "$URL"
