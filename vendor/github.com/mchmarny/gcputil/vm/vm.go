package vm

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

	mt "github.com/mchmarny/gcputil/meta"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

func failOnErr(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

// ShutdownHostVM looks up the current project ID as well as
// host VM name and zone in which it executes and calls the
// compute API to shut it down. Helpful when container
// finished running in VM and wants to shutdown the host VM
// to avoid paying for idle time. Since there isn't probably
// anyone waiting for response to this, any critical errors
// in this method will be fatal
func ShutdownHostVM(ctx context.Context, agent string) {

	mc := mt.GetClient(agent)

	projectID, err := mc.ProjectID()
	failOnErr(err)

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
