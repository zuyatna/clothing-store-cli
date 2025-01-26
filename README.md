## Note!
I start to refactor this project to advanced learning.

## Entity Relationship Diagram
![alt text](https://github.com/zuyatna/clothing-pair-project/blob/main/erd.png?raw=true)

## Project Structures
```
└── project/
    ├── cmd/
    │   └── cli/
    │       └── main.go
    ├── internal/
    │   ├── config/
    │   │   └── config.go
    │   ├── database/
    │   │   ├── user.go
    │   │   ├── product.go
    │   │   └── ..
    │   ├── helper/
    │   │   └── ..
    │   ├── models/
    │   │   ├── user.go
    │   │   ├── product.go
    │   │   └── ..
    │   ├── repository/
    │   │   │   ├── user.go
    │   │   │   ├── product.go
    │   │   │   └── ..
    │   ├── services/
    │   │   ├── user.go
    │   │   ├── product.go
    │   │   └── ..
    │   └── util/
    │       └── menu
    │           ├── user.go
    │           ├── product.go
    │           └── ..
    ├── tests/
    ├── .env
    ├── .env.example
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    └── README.md
```

## Folder Explanation
- `/cmd` Main applications for this project.
- `/internal` Contains application core code that is not intended to be used by external packages.
    - `/config` Manage application configurations (e.g., databases, servers).
    - `/database` Contains interactions with the database.
    - `/helper` Cross-cutting concerns and common utilities.
    - `/models` Defines data structures.
    - `/repository` Manage interactions with databases or other data sources.
    - `/services` Implementing business logic.
    - `/utils` Contains general utility functions.
- `/tests` Contains unit tests and integration tests. The test is designed to check the behavior method, not the actual content.

## Package Requirements:
- `go get github.com/spf13/viper`
- `go get github.com/jmoiron/sqlx`
- `go get github.com/lib/pq`
- `go get github.com/olekukonko/tablewriter`
- `go get github.com/stretchr/testify/mock`
- `go get golang.org/x/term`
- `go get golang.org/x/crypto/bcrypt`