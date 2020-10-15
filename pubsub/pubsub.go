package dc_pubsub

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

var (
	pubsubClient *PubSub
)

type PubSub struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func New(projectID string, topicName string) *PubSub {
	if pubsubClient != nil {
		return pubsubClient
	}

	client, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		logrus.WithError(err).Error("pubsub.NewClient")
		os.Exit(1)
		return nil
	}

	topic := client.Topic(topicName)
	pubsubClient = &PubSub{
		client: client,
		topic:  topic,
	}

	return pubsubClient
}

func (p *PubSub) Publish(encodedMessage []byte) (string, error) {
	// create pubsub message data
	pubsubMessage := &pubsub.Message{
		Data: encodedMessage,
	}

	// publish to pubsub
	ctx := context.Background()
	id, err := p.topic.Publish(ctx, pubsubMessage).Get(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to publish message")
	}

	return id, nil
}
