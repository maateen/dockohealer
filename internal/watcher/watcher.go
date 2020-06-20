package watcher

import (
	"context"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/maateen/dockohealer/internal/healer"
	log "github.com/sirupsen/logrus"
)

// checkPoint checks the health status of a container and takes action.
func checkPoint(ctx context.Context, cli *client.Client, event events.Message) {
	healthStatus := strings.Split(event.Status, " ")[1]
	log.WithFields(log.Fields{
		"containerID": event.ID,
	}).Infof("Container is %s.", healthStatus)

	if event.Type == "container" && healthStatus == "unhealthy" {
		log.WithFields(log.Fields{
			"containerID": event.ID,
		}).Info("Restarting container.")
		go healer.HealContainer(ctx, cli, event.ID)
	}
}

// findGhosts finds the containers which are already unhealthy.
func findGhosts(ctx context.Context, cli *client.Client) {
	args := filters.NewArgs(
		filters.Arg("health", "unhealthy"),
	)
	containerListOptions := types.ContainerListOptions{
		Filters: args,
	}

	containerList, err := cli.ContainerList(ctx, containerListOptions)

	if err == nil {
		for _, container := range containerList {
			log.WithFields(log.Fields{
				"containerID": container.ID,
			}).Infof("Container is unhealthy.")

			log.WithFields(log.Fields{
				"containerID": container.ID,
			}).Info("Restarting container.")

			go healer.HealContainer(ctx, cli, container.ID)
		}
	} else {
		log.Error(err)
	}
}

// Watch function watches the containers and listens to Docker events
func Watch() {
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
	defer cancel()

	dockerEvents, errs := cli.Events(ctx, eventOptions)
	if err == nil {
		log.Info("Listening from Docker.")
	} else {
		log.Fatal(err)
	}

	go findGhosts(ctx, cli)

	for {
		select {
		case event := <-dockerEvents:
			go checkPoint(ctx, cli, event)
		case err := <-errs:
			if err == io.EOF {
				log.Error(err)
			}
		}
	}
}
