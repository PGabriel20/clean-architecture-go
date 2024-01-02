/api
contains api endpoints

/internal
application code

/internal/entities
contains fine business rules

/internal/usecase
user actions and application business rules

/internal/infra
contains adapters to interact with usecases

/internal/handler
acts as controller, contains HTTP routes and methods that call usecases

## TODO

- [ ] Better error handling
- [ ] Better usecase naming
- [ ] Logging
