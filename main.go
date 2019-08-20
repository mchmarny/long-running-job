package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	ev "github.com/mchmarny/gcputil/env"
	mt "github.com/mchmarny/gcputil/meta"
	pr "github.com/mchmarny/gcputil/project"
)

var (
	logger = log.New(os.Stdout, "[LRJ]", 0)

	projectID  = pr.GetIDOrFail()
	bucketName = ev.MustGetEnvVar("BUCKET", "")
	objectName = ev.MustGetEnvVar("OBJECT", "")
	topicName  = ev.MustGetEnvVar("TOPIC", "")
)

func main() {

	defer shutdownVM()

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
		logger.Println(err)
	}
}

func shutdownVM() {

	mc := mt.GetClient("long-running-job-demo")

	vmName, err := mc.InstanceName()
	failOnErr(err)

	vmZone, err := mc.Zone()
	failOnErr(err)

	cmd := exec.Command("gcloud", "compute", "instances", "delete",
		vmName, "--zone", vmZone)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		logger.Fatalf("Error on VM shutdown %v", err)
	}
}
