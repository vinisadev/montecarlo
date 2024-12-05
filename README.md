# Monte Carlo Simulation API

This API allows you to run Monte Carlo simulations for project management using the Fiber web framework in Go. The simulation takes input for task estimates and returns a curve of days on the x-axis and percentage on the y-axis.

## Prerequisites

    Go installed on your machine
    Fiber web framework

## Installation

1. Clone the repository:

```
git clone https://github.com/vinisadev/montecarlo.git
cd montecarlo
```

2. Install dependencies:

```
go mod tidy
```

## Running the Application


1. Start the application:

```
go run main.go
```

2. The API will be available at http://localhost:3000/simulate.

## Using the API with curl

### Endpoint

    URL: http://localhost:3000/simulate
    Method: POST
    Content-Type: application/json

### Request Body

The request body should be a JSON object with the following structure:

```
{
  "iterations": 1000,
  "task_estimates": [
    {"best_case": 2, "most_likely": 4, "worst_case": 8},
    {"best_case": 3, "most_likely": 5, "worst_case": 7}
  ]
}
```

- iterations: The number of simulations to run.
- task_estimates: An array of task estimates, where each task has best_case, most_likely, and worst_case durations.

### Example curl Command

```
curl -X POST http://localhost:3000/simulate -H "Content-Type: application/json" -d '{
  "iterations": 1000,
  "task_estimates": [
    {"best_case": 2, "most_likely": 4, "worst_case": 8},
    {"best_case": 3, "most_likely": 5, "worst_case": 7}
  ]
}'
```

### Response

The API will return a JSON array of points, where each point represents a day and the percentage of simulations that completed by that day. Example response:

```
[
  {"days": 0, "percentage": 0},
  {"days": 1, "percentage": 0},
  {"days": 2, "percentage": 5.3},
  {"days": 3, "percentage": 12.7},
  ...
]
```

## Explanation

    Triangular Distribution: The simulation uses a triangular distribution to model task durations based on best-case, most-likely, and worst-case estimates.
    Cumulative Distribution Function (CDF): The response provides a CDF, showing the cumulative probability of completing the project by each day.

## Notes

    Ensure that the task estimates are reasonable and that the best-case duration is less than or equal to the most-likely duration, which is less than or equal to the worst-case duration.