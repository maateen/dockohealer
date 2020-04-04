package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/maateen/dockohealer/internal/healer"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err == nil {
		log.Info("Connected to Docker.")
	} else {
		log.Fatal(err)
	}

	args := filters.NewArgs(
		filters.Arg("event", "health_status"),
	)

	eventOptions := types.EventsOptions{
		Filters: args,
	}

	ctx, cancel := context.WithCancel(context.Background())
	events, errs := cli.Events(ctx, eventOptions)
	log.Info("Listening from Docker.")
	defer cancel()

	for {
		select {
		case event := <-events:
			go healer.CheckPoint(ctx, cli, event)
		case err := <-errs:
			if err == io.EOF {
				log.Error(err)
			}
		}
	}
}
