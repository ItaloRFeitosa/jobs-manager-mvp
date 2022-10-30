package api

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/jobs-manager-mvp/manager/di"
)

func Start(deps *di.Container) {
	app := fiber.New()

	app.Get("/v1/health", func(c *fiber.Ctx) error {
		stats := new(runtime.MemStats)
		runtime.ReadMemStats(stats)
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"NumCPU":       runtime.NumCPU(),
			"NumGoroutine": runtime.NumGoroutine(),
			"MemStats":     stats,
		})
	})

	app.Get("/v1/jobs", func(c *fiber.Ctx) error {
		var jobs []fiber.Map

		jobsOnStore := deps.JobsStore.GetAll()

		for _, jobOnStore := range jobsOnStore {
			jobs = append(jobs, fiber.Map{
				"name":      jobOnStore.Name(),
				"nextRun":   jobOnStore.NextRun(),
				"lastRun":   jobOnStore.LastRun(),
				"isRunning": jobOnStore.IsRunning(),
				"runCount":  jobOnStore.RunCount(),
				"isActive":  jobOnStore.IsActive(),
			})
		}

		return c.Status(http.StatusOK).JSON(jobs)
	})

	app.Patch("/v1/jobs/:name/run", func(c *fiber.Ctx) error {
		job, err := deps.JobsStore.Get(c.Params("name"))
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Errorf("error on get job: %w", err).Error(),
			})
		}

		err = job.Run()

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Errorf("error on run job: %w", err).Error(),
			})
		}

		return c.SendStatus(http.StatusNoContent)
	})

	app.Patch("/v1/jobs/:name/stop", func(c *fiber.Ctx) error {
		job, err := deps.JobsStore.Get(c.Params("name"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Errorf("error on get job: %w", err).Error(),
			})
		}

		job.Stop()
		return c.SendStatus(http.StatusNoContent)
	})

	app.Patch("/v1/jobs/:name/start", func(c *fiber.Ctx) error {
		job, err := deps.JobsStore.Get(c.Params("name"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Errorf("error on get job: %w", err).Error(),
			})
		}

		job.Start()
		return c.SendStatus(http.StatusNoContent)
	})

	app.Listen(":8080")
}
