/api
contains api endpoints

/internal
application code

/internal/entities
contains fine business rules and repository interface contracts

/internal/usecase
user actions and application-level business rules

/internal/infra
contains data layer and http handlers to interact with usecases

/internal/handler
acts as controller, contains HTTP routes and methods that call usecases

## TODO

- [ ] Better error handling
- [ ] Logging
