# Golang ExampleApp
**Clean Architecture**

## Installation

```bash
    go get "github.com/go-playground/locales" 
    go get "github.com/go-playground/universal-translator" 
    go get "github.com/go-playground/validator/v10" 
    go get "github.com/gorilla/mux" 
    go get "go.uber.org/zap" 
    go get "gorm.io/driver/postgres" 
    go get "gorm.io/gorm" 
```

## Folder Structure

```
+-- Cmd
|   +-- main.go
+-- Internal
|   +--- config
|   |   +--- database.go
|   +--- core
|   |   +--- routers.go
|   |   +--- server.go
|   +--- dtos
|   |   +--- account
|   |   |    +--- login_dto.go
|   |   |    +--- register_dto.go
|   |   |    +--- token.go
|   |   +--- user
|   |   |    +--- user_dto.go
|   |   |    +--- user_list.go
|   |   +--- pagination.go
|   |   +--- search.go
|   +--- filters
|   |   +--- user_filter.go
|   +--- handlers
|   |   +--- account_handler.go
|   |   +--- user_handler.go
|   +--- helpers
|   |   +--- authorization.go
|   |   +--- custom_errors.go
|   |   +--- hashing.go
|   |   +--- json.go
|   |   +--- jwt.go
|   |   +--- paginations.go
|   |   +--- validations.go
|   |   +--- zap_logger.go
|   +--- middleware
|   |   +--- authorization.go
|   +--- models
|   |   +--- user.go
|   +--- repository
|   |   +--- user_respository.go
|   +--- services
|   |   +--- account_service.go
|   |   +--- user_service.go
|   go.mod
|   go.sum
|   README.md 
```



