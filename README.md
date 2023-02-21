# Service

Прототип сервиса бронирования номеров в отеле.
В текущий момент имеет функциональность бронирования номера в отеле пользователем на определенные даты и получение
списка бронирований для указанного пользователя.

# Usage

## API

**[TODO: create OpenAPI doc in /api!]**

### GetUserOrders

#### Request:

```
GET /orders?user_id=<string>
```

#### Response:

```
{
  "orders": [
    {
      id: "<str>",
      room: "<str>",
      from: "<2028-05-24>",
      to: "<2028-05-27>"
    }
  ]
}

```

### ReserveRoom

#### Request:

```
POST /reservation
{
  	"user_id": "<str>",
	"room_id": "<str>",
	"from": "<2028-05-24>",
	"to": "<2028-05-26>"
}
```

#### Ok response:

```
Http status: 200
{
  "reservation_id": "<str>"
}
```

#### Response "dates nor available"

```
Http status: 209
```

### CreateOrder

#### Request:

```
POST /order
{
  	"user_id": "<str>",
	"reservation_id": "<str>"
}
```

#### Ok response:

```
{
  "order_id": "<str>"
}

```

# Dev

## App layout

* cmd: application executables (app, cron scripts, cli scripts, workers, etc)
* internal:
    * app: application services, e.g. commands/queries, etc
    * domain: domain entities, domain services
    * infra: implementations (postgres repository, kafla databus, etc)
    * presentation: application UI, e.g. http or gRPC handlers
    * api: OpenAPI/Swagger specs, JSON schema files, .proto files, etc.

## Commands

* Run tests: `make test`
* Run linter: `make lint`
