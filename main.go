package main

import (
	"context"
	"log"
	"os"

	ev "github.com/mchmarny/gcputil/env"
	pr "github.com/mchmarny/gcputil/project"
)

var (
	logger = log.New(os.Stdout, "", 0)

	projectID  = pr.GetIDOrFail()
	bucketName = ev.MustGetEnvVar("BUCKET", "")
	objectName = ev.MustGetEnvVar("OBJECT", "")
	topicName  = ev.MustGetEnvVar("TOPIC", "")
)

func main() {

	logger.Println("Starting long running job demo...")
	ctx := context.Background()

	logger.Println("Initializing publisher...")
	pub, err := newPublisher(ctx, projectID, topicName)
	failOnErr(err)

	logger.Println("Starting provider...")
	count, err := provide(ctx, pub, bucketName, objectName)
	failOnErr(err)

	logger.Printf("Processed %d records", count)

}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
