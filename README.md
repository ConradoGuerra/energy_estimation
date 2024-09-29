# Energy Estimation Service

The **Energy Estimation Service** is a Go-based project designed using the **Domain-Driven Design (DDD)** approach. This service is responsible for estimating energy consumption based on historical data and applying specific estimation rules. The project, at the moment, uses **in-memory repositories** for tariff rules. The **Continuous Integration (CI) and Continuous Deployment (CD)** are managed through **GitHub Actions** to deploy the project at AWS.

## Project Structure

```bash
.
│   Dockerfile
│   go.mod
│   go.sum
│   main.go
│   README.md
├───.github
│   └───workflows
│           energy_estimation_pipeline.yaml
├───domain
│   ├───estimation
│   │   │   estimation.go
│   │   │   service.go
│   │   │   service_test.go
│   │   └───factory
│   │           factory.go
│   ├───historic_consumption
│   │       historic.go
│   └───tariff
│           service.go
│           service_test.go
│           tariff.go
└───infrastructure
    ├───dtos
    │       historic_consumption.go
    ├───http_handlers
    │       estimation_handler.go
    └───repositories
        ├───in_memory
        │       tariffs_repository.go
        └───tariffs
```

## Key Folders

- **App**: Contains services and controllers to handle business logic and API interaction.
- **Domain**: Contains core domain models and business rules, following DDD practices.
- **Infrastructure**: Contains in-memory repositories for tariff rules and data persistence, and contains the controllers.
- **Test**: Contains unit tests for ensuring the correctness of the application.

## Installation

To install and run the project locally, follow these steps:

#### 1. Clone the repository:

```bash
git clone https://github.com/ConradoGuerra/energy_estimation.git
cd energy_estimation
```

#### 2. Install project dependencies:

```bash
bash
go mod tidy
```

#### 3. Run the project:

```bash
bash
go run ./main.go
```

#### 4. Run the tests:

```bash
go test ./...
```

#### 5. Build the project:

```bash
bash
go build ./main.go
```

## Usage

You can interact with the Energy Estimation Service through an **API** once the project is running. Example:

#### API Request:

`POST http://localhost:8080/api/estimation`

```json
{
  "client_id": "1",
  "measures": [
    {
      "consumption": 23,
      "begin": "2024/09/01",
      "end": "2024/09/30"
    },
    {
      "consumption": 4,
      "begin": "2024/08/01",
      "end": "2024/08/31"
    },
    {
      "consumption": 54,
      "begin": "2024/07/01",
      "end": "2024/07/31"
    }
  ]
}
```

## Estimation Logic

- The service uses `EstimationRules` to apply specific ratios for calculating energy consumption.
- Historical consumption data is evaluated to generate an estimate of energy usage for each measure.
- The final estimation combines historical consumption with applicable tariff rules.

## Testing

Unit tests ensure the correctness of the core estimation logic:

- `domain/estimation/service_test.go`
- `domain/tariff/service_test.go`

To run the unit tests:

```bash
go test ./domain/estimation/...
go test ./domain/tariff/...
```

## CI/CD Pipeline

This project uses **GitHub Actions** for CI/CD to build, test, and deploy the application.

### Key Jobs:

**Build Job**:

- Runs when changes are pushed or a pull request is made to the `main` branch.
- Steps:
  - Set up Go, install dependencies, run unit tests, and build the project.
  - Upload the built binary as an artifact for deployment.

**Deploy Job**:

- Runs after the build job.
- Steps:
  - Log in to AWS ECR, build and push a Docker image.
  - Update ECS task definition and deploy the new image to Amazon ECS.
