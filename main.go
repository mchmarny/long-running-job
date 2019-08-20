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
	shutdownVM()

}

func failOnErr(err error) {
	if err != nil {
		logger.Println(err)
		shutdownVM()
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
