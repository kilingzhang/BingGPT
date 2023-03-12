# BingGPT

BingGPT

# Updates

注意:

当前为开发不稳定版本！！！请酌情使用。

当前为开发不稳定版本！！！请酌情使用。

当前为开发不稳定版本！！！请酌情使用。

# BingGPT API

Bing AI 的客户端实现。作为 Go 模块、REST API 服务器和 CLI 应用程序提供。

# Getting Started

```shell
make run
```

```shell
curl --location 'http://127.0.0.1:12527/conversation/create' \
--header 'Content-Type: application/json' \
--data '{
    "cookies": ""
}'

----
{
    "conversationId": "",
    "clientId": "",
    "conversationSignature": "",
    "result": {
        "value": "Success",
        "message": null
    }
}
```

```shell
curl --location 'http://127.0.0.1:12527/conversation' \
--header 'Content-Type: application/json' \
--data '{
    "message": "Bing Ai",
    "conversationId": "",
    "clientId": "",
    "conversationSignature": "",
    "invocationId": 0
}'
```