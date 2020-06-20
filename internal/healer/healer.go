package healer

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

func HealContainer(ctx context.Context, cli *client.Client, containerID string) {
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
