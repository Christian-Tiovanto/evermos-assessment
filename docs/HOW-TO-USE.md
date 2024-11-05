## Prerequisite

- **Go**: v1.22 or later (https://go.dev/)
- **Mockgen (Uber version)**: v0.4.0 or later (https://github.com/uber-go/mock)
- **Golangci-lint**: v1.56.1 or later (https://golangci-lint.run/usage/install)
- **SQLC**: v1.27.0 or later (https://docs.sqlc.dev/en/stable/overview/install.html)
- **Dbmate**: v2.12.0 or later (https://github.com/amacneil/dbmate)

### Project structure
```
├── app                       # Application source code
│   ├── common                # Common code shared across the application
│   ├── delivery              # Delivery layer
│   │   ├── http
│   │   │   ├── handler
│   │   │   └── middleware
│   │   └── subscriber
│   ├── repository            # Repository layer
│   └── usecase               # Usecase layer
├── bin                       # Helper scripts
├── cmd                       # Application entry points
├── db                        # Database scripts
│   ├── migrations
│   └── queries
├── deployment                # Deployment configurations
├── dockerfiles
├── protobuf                  # Protobuf files
│   ├── api                   # API protobuf files
│   ├── swagger               # Swagger files generated from protobuf
└── test
```

### How to use

These commands below will help you to use this project properly.
For the first, you need to clone this repository then run:

```
make vendor
```
### When working

#### Running the tests

Unit test
```shell
make test.unit
```

Coverage report
```shell
make test.coverage
```

#### Verifying code style

```shell
make lint
```

You can reformat the code by executing:

```shell
make format
```

### When running this project

Run docker command
```shell
docker-compose up
```
Start running
```shell
make run
```

#### Add new API

You can add new API with write the protobuf `<package name>.proto` into `protobuf/api/` directory.

You can generated golang code and swagger by executing:
```shell
make generate
```

#### Manipulating database

Create new migrations
```shell
make migration name="< action >"
```

Migrating the migrations
```shell
make migrate url="< postgres url >"
```

Rollback the migrations
```shell
make rollback url="< postgres url >"
```

#### Add new queries
You can add new queries with write the sql `<package name>.sql` into `db/queries` directory.

Execute this command to generated golang code
```shell
make sqlc
```
