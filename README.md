# Container Platform App

This is a simple golang application that allows you to connect with various component such as Mongo, Redis and RabbitMQ
server. You can integrate and utilize your existing components with this app. 

## Installation

### Docker

```
docker run -p 8080:8080 -e SERVER_PORT=8080 -e RUN_MODE=PRODUCTION quay.io/klovercloud/app-container-platform:v1.0 
```

## Environment Variables

### Default Environment Variables

| Name          |                            Description                            | Value     | Required |
|:--------------|:-----------------------------------------------------------------:|-----------|:--------:|
| `RUN_MODE`    | Specifies the app running mode (options: `DEVELOP`, `PRODUCTION`) | `DEVELOP` |    ◽     |
| `SERVER_PORT` |                         App running port                          | `8080`    |    ◽     |

### Mongo Environment Variables

| Name                                  |                                                                                 Description                                                                                 | Value   | Required |
|:--------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|---------|:--------:|
| `CONNECT_MONGO`                       |                                                   For enable the mongo connection with the app (options: `true`, `false`)                                                   | `false` |    ✅     |
| `MONGODB_CONNECTION_STRING_FOR_WRITE` | Mongo connection string for write privileges. Required if `CONNECT_MONGO` is `true`. Format `mongodb://<username>:<password>@<mongo server>:<mongo port>/?authSource=admin` | `""`    |    ◽     |
| `MONGODB_CONNECTION_STRING_FOR_READ`  | Mongo connection string for read privileges. Required if `CONNECT_MONGO` is `true`.  Format `mongodb://<username>:<password>@<mongo server>:<mongo port>/?authSource=admin` | `""`    |    ◽     |
| `DATABASE_NAME`                       |                                 Any name in Mongo (It will be auto matically created if not exists). Required if `CONNECT_MONGO` is `true`                                  | `""`    |    ◽     |


### Redis Environment Variables

| Name                         |                                                                         Description                                                                          | Value        | Required |
|:-----------------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------:|--------------|:--------:|
| `CONNECT_REDIS`              |                                           For enable the redis connection with the app (options: `true`, `false`)                                            | `false`      |    ✅     |
| `REDIS_CONNECTION_TYPE`      |                         Redis server type. (options: `MASTER_SLAVE`, `SENTINEL`, `CLUSTER`). Required if `CONNECT_REDIS` is `true`.                          | `""`         |    ◽     |
| `REDIS_SERVER_FOR_WRITE`     | Redis connection string for write privileges. Required if `CONNECT_REDIS` is `true` and `REDIS_CONNECTION_TYPE` is `MASTER_SLAVE`.  Format `<server>:<port>` | `""`         |    ◽     |
| `REDIS_SERVER_FOR_READ`      | Redis connection string for read privileges. Required if `CONNECT_REDIS` is `true` and `REDIS_CONNECTION_TYPE` is `MASTER_SLAVE`.  Format `<server>:<port>`  | `""`         |    ◽     |
| `REDIS_SENTINEL_MASTER_NAME` |                         Redis sentinel master name. Required if `CONNECT_REDIS` is `true` and `REDIS_CONNECTION_TYPE` is `SENTINEL`.                         | `"mymaster"` |    ◽     |
| `REDIS_SENTINEL_SERVER`      |              Redis sentinel server. Required if `CONNECT_REDIS` is `true` and `REDIS_CONNECTION_TYPE` is `SENTINEL`.  Format `<server>:<port>`               | `""`         |    ◽     |
| `REDIS_SERVER_PASSWORD`      |                                             Redis password if necessary. Required if `CONNECT_REDIS` is `true`.                                              | `""`         |    ◽     |

