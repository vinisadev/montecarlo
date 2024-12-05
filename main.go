package main

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TaskEstimate struct {
	BestCase   int `json:"best_case"`
	MostLikely int `json:"most_likely"`
	WorstCase  int `json:"worst_case"`
}

type Point struct {
	Days       int     `json:"days"`
	Percentage float64 `json:"percentage"`
}

func randomDuration(task TaskEstimate) int {
	// Use a triangular distribution
	bestCase := float64(task.BestCase)
	mostLikely := float64(task.MostLikely)
	worstCase := float64(task.WorstCase)

	u := rand.Float64()
	if u < (mostLikely-bestCase)/(worstCase-bestCase) {
		return int(bestCase + math.Sqrt(u*(worstCase-bestCase)*(mostLikely-bestCase)))
	} else {
		return int(worstCase - math.Sqrt((1-u)*(worstCase-bestCase)*(worstCase-mostLikely)))
	}
}

func calculatePercentiles(results []int) []Point {
	sortedResults := append([]int(nil), results...)
	sort.Ints(sortedResults)

	maxDays := sortedResults[len(sortedResults)-1]
	points := make([]Point, maxDays+1)

	for day := 0; day <= maxDays; day++ {
		count := 0
		for _, result := range sortedResults {
			if result <= day {
				count++
			}
		}
		points[day] = Point{
			Days:       day,
			Percentage: float64(count) / float64(len(sortedResults)) * 100,
		}
	}
	return points
}

func runMonteCarloSimulation(iterations int, taskEstimates []TaskEstimate) []Point {
	rand.Seed(time.Now().UnixNano())
	results := make([]int, iterations)

	for i := 0; i < iterations; i++ {
		totalDays := 0
		for _, task := range taskEstimates {
			totalDays += randomDuration(task)
		}
		results[i] = totalDays
	}

	return calculatePercentiles(results)
}

func main() {
	app := fiber.New()

	app.Post("/simulate", func(c *fiber.Ctx) error {
		var input struct {
			Iterations    int            `json:"iterations"`
			TaskEstimates []TaskEstimate `json:"task_estimates"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		results := runMonteCarloSimulation(input.Iterations, input.TaskEstimates)
		return c.JSON(results)
	})

	app.Listen(":3000")
}
