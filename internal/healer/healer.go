package healer

import (
	"context"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func CheckPoint(cli *client.Client, ctx context.Context, event events.Message) {
	healthStatus := strings.Split(event.Status, " ")[1]
	log.WithFields(log.Fields{
		"containerID":  event.ID,
	}).Infof("Container is %s.", healthStatus)

	if event.Type == "container" && healthStatus == "unhealthy" {
		log.WithFields(log.Fields{
			"containerID": event.ID,
		}).Info("Restarting container.")
		restartContainer(cli, ctx, event.ID)
	}
}

func restartContainer(cli *client.Client, ctx context.Context, containerID string) {
	var timeout *time.Duration
	err := cli.ContainerRestart(ctx, containerID, timeout)
	if err == nil {
		log.WithFields(log.Fields{
			"containerID": containerID,
		}).Info("Successfully restarted container.")
	} else {
		log.WithFields(log.Fields{
			"containerID": containerID,
		}).Fatalf("%s", err)
	}
}
