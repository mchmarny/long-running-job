package main

import (
	"context"
	"log"
	"os"

	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iam/v1"

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

	logger.Println("Starting long running job demo...")
	ctx := context.Background()

	defer shutdownVM(ctx)

	// Sample Code
	// Replace this with your code
	logger.Println("Initializing publisher...")
	pub, err := newPublisher(ctx, projectID, topicName)
	failOnErr(err)

	logger.Println("Starting provider...")
	count, err := provide(ctx, pub, bucketName, objectName)
	failOnErr(err)

	logger.Printf("Processed %d records", count)
	// End of sample code

}

func failOnErr(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

func shutdownVM(ctx context.Context) {

	mc := mt.GetClient("long-running-job-demo")

	vmName, err := mc.InstanceName()
	failOnErr(err)

	vmZone, err := mc.Zone()
	failOnErr(err)

	client, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	failOnErr(errors.Wrap(err, "Error on client create"))

	u := fmt.Sprintf(
		"https://www.googleapis.com/compute/v1/projects/%s/zones/%s/instances/%s",
		projectID, vmZone, vmName)
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	failOnErr(errors.Wrap(err, "Error on client request create"))

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	failOnErr(errors.Wrap(err, "Error while executing request"))
	defer resp.Body.Close()

	if err := googleapi.CheckResponse(resp); err != nil {
		failOnErr(errors.Wrap(err, "Invalid shutdown command response"))
	}

}
