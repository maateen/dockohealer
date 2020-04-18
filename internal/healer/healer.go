package healer

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

// CheckPoint checks the health status of a container and takes action.
func CheckPoint(ctx context.Context, cli *client.Client, event events.Message) {
	healthStatus := strings.Split(event.Status, " ")[1]
	log.WithFields(log.Fields{
		"containerID": event.ID,
	}).Infof("Container is %s.", healthStatus)

	if event.Type == "container" && healthStatus == "unhealthy" {
		log.WithFields(log.Fields{
			"containerID": event.ID,
		}).Info("Restarting container.")
		restartContainer(ctx, cli, event.ID)
	}
}

// FindGhosts finds the containers which are already unhealthy.
func FindGhosts(ctx context.Context, cli *client.Client) {
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

			restartContainer(ctx, cli, container.ID)
		}
	} else {
		log.Error(err)
	}
}

func restartContainer(ctx context.Context, cli *client.Client, containerID string) {
	var timeout *time.Duration
	err := cli.ContainerRestart(ctx, containerID, timeout)
	if err == nil {
		log.WithFields(log.Fields{
			"containerID": containerID,
		}).Info("Successfully restarted container.")
	} else {
		log.WithFields(log.Fields{
			"containerID": containerID,
		}).Error(err)
	}
}
