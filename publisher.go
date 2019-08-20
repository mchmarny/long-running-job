package main

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
)

type publisher struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func (p *publisher) publish(ctx context.Context, data []byte) error {
	msg := &pubsub.Message{Data: data}
	result := p.topic.Publish(ctx, msg)
	_, err := result.Get(ctx)
	return errors.Wrap(err, "Error publishing content")
}

func newPublisher(ctx context.Context, projectID, topicName string) (p *publisher, err error) {

	if projectID == "" {
		return nil, errors.New("projectID not set")
	}

	if topicName == "" {
		return nil, errors.New("topicName not set")
	}

	if ctx == nil {
		return nil, errors.New("context not set")
	}

	c, e := pubsub.NewClient(ctx, projectID)
	if e != nil {
		return nil, e
	}

	t := c.Topic(topicName)
	topicExists, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !topicExists {
		logger.Printf("Topic %s not found, creating...", topicName)
		t, err = c.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to create topic: %s - %v", topicName, err)
		}
	}

	o := &publisher{
		client: c,
		topic:  t,
	}

	return o, nil
}
