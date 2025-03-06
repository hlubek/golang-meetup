package main

import (
	"asynq-jobqueue/tasks"
	"fmt"
	"log/slog"
	"os"

	"github.com/hibiken/asynq"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "asynq-jobqueue"
	app.Usage = "A job queue example for image resizing with asynq"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "redis-addr",
			Value: "127.0.0.1:6379",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "worker",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				redisAddr := c.String("redis-addr")
				srv := asynq.NewServer(
					asynq.RedisClientOpt{Addr: redisAddr},
					asynq.Config{
						// Specify how many concurrent workers to use
						Concurrency: 2,
						// Optionally specify multiple queues with different priority.
						Queues: map[string]int{
							"medium": 5,
						},
					},
				)

				imgProcessor := tasks.NewImageProcessor()
				err := imgProcessor.Startup(c.Context)
				if err != nil {
					return fmt.Errorf("starting image processor: %v", err)
				}
				defer imgProcessor.Shutdown(c.Context)

				// mux maps a type to a handler
				mux := asynq.NewServeMux()
				mux.Handle(tasks.TypeImageResize, imgProcessor)

				if err := srv.Run(mux); err != nil {
					return fmt.Errorf("running asynq server: %v", err)
				}

				return nil
			},
		},
		{
			Name: "queue-resize",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "width",
					Usage: "width of the resized image",
					Value: 500,
				},
				&cli.IntFlag{
					Name:  "height",
					Usage: "height of the resized image",
					Value: 500,
				},
			},
			ArgsUsage: "<source-url> <target-path>",
			Before: func(c *cli.Context) error {
				if c.NArg() != 2 {
					return fmt.Errorf("source-url and target-path are required")
				}
				return nil
			},
			Action: func(c *cli.Context) error {
				redisAddr := c.String("redis-addr")
				client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
				defer client.Close()

				sourceURL := c.Args().Get(0)
				targetPath := c.Args().Get(1)

				task, err := tasks.NewImageResizeTask(sourceURL, targetPath, c.Int("width"), c.Int("height"))
				if err != nil {
					return fmt.Errorf("creating task: %v", err)
				}

				_, err = client.Enqueue(task, asynq.Queue("medium"))
				if err != nil {
					return fmt.Errorf("enqueueing task: %v", err)
				}

				slog.Info("Enqueued task", "type", task.Type, "source", sourceURL, "target", targetPath)

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
