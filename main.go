package main

import (
	"context"
	"log"
	"os"

	"github.com/mchmarny/gcputil/env"
	"github.com/mchmarny/gcputil/project"
	"github.com/mchmarny/gcputil/vm"
)

const (
	defaultShutdownFlag = "yes"
)

var (
	logger = log.New(os.Stdout, "[LRJ]", 0)

	projectID  = project.GetIDOrFail()
	bucketName = env.MustGetEnvVar("BUCKET", "")
	objectName = env.MustGetEnvVar("OBJECT", "")
	topicName  = env.MustGetEnvVar("TOPIC", "")
	doShutdown = env.MustGetEnvVar("SHUTDOWN", defaultShutdownFlag)
)

func main() {

	logger.Println("Starting long running job...")
	ctx := context.Background()

	if doShutdown == defaultShutdownFlag {
		defer vm.ShutdownHostVM(ctx, "long-runing-job-demo")
	}

	// Sample Code
	// Replace this with your code
	logger.Println("Initializing publisher...")
	pub, err := newPublisher(ctx, projectID, topicName)
	failOnErr(err)

	logger.Println("Starting provider...")
	count, err := provide(ctx, pub, bucketName, objectName)
	failOnErr(err)

	logger.Printf("Processed %d records", count)
	// End Sample Code

}

func failOnErr(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
