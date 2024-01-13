"# Clean Architecture Go Project

This project follows the principles of Clean Architecture to structure a Go application.

## Project Structure

/api
  - Contains API endpoints

/internal
  - Application code

/internal/entities
  - Contains fine business rules and repository interface contracts

/internal/usecase
  - User actions and application-level business rules

/internal/infra
  - Data layer and HTTP handlers to interact with usecases

/internal/handler
  - Acts as a controller, contains HTTP routes and methods that call usecases

## How to Run the Project

### 1. Clone the Repository

git clone \"https://github.com/PGabriel20/clean-architecture-go.git\"

### 2. Navigate to the Project Directory

cd \"clean-architecture-go\"

### 3. Copy Contents of .env.example to .env

cp .env.example .env

Update the values in the .env file as needed.

### 4. Run Docker Compose to Start Services

docker-compose up -d

Optionally, add the --build flag to rebuild the images:

docker-compose up -d --build

The -d flag runs the containers in the background.

### 5. Access the endpoints

All of the available endpoints are listed /api/api.http file

### 6. Stopping the Application

To stop and remove the containers:

docker-compose down

## TODO

- [ ] Better error handling
- [ ] Logging"
