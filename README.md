# Event Service


### Application Architecture
```shell
.
├── Dockerfile
├── README.md
├── api
│   └── v1
│       ├── controllers
│       └── routes.go
├── app
│   ├── app.go
│   └── routes.go
├── config
│   └── date_format.go
├── db
│   └── postgres.go
├── dto
│   ├── db.go
│   ├── errors.go
│   └── event.go
├── filters
│   └── event_filter_impl.go
├── go.mod
├── go.sum
├── interfaces
│   ├── controllers
│   │   └── event_controller_interface.go
│   ├── db
│   │   └── db_interface.go
│   ├── filters
│   │   └── event.go
│   ├── models
│   │   └── models.go
│   ├── repository
│   │   └── event_repository_interface.go
│   └── services
│       └── event_service_interface.go
├── main.go
├── models
│   └── event.go
├── repository
│   └── event_repository_impl.go
├── services
│   └── event_service_impl.go
└── utils
    ├── hash_utils.go
    └── os_utils.go
```