package main

import (
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
	// Simple Triangular distribution
	bestToMostLikely := rand.Intn(task.MostLikely-task.BestCase+1) + task.BestCase
	mostLikelyToWorst := rand.Intn(task.WorstCase-task.MostLikely+1) + task.MostLikely
	return (bestToMostLikely + mostLikelyToWorst) / 2
}

func calculatePercentiles(results []int) []Point {
	sortedResults := append([]int(nil), results...)
	sort.Ints(sortedResults)

	points := make([]Point, 101)
	for i := 0; i <= 100; i++ {
		index := int(float64(i) / 100.0 * float64(len(sortedResults)-1))
		points[i] = Point{
			Days:       sortedResults[index],
			Percentage: float64(i),
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
