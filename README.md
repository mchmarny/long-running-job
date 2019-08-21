# long-running-job

While the idea of a serverless platform and long running workloads does seem somewhat "unnatural" at first, smart people are already working on that (looking at you @Knative community). In the meantime, a simple approach is sometimes all you may need.

In this demo I will illustrating how to use Google Compute Engine (GCE) container execution option to run variable duration jobs. This approach supports custom VMs, GPU/TPU accelerators and VPC networks… so may be handy alternative to other compute options on GCP. I'll also demo how to auto-terminate the created VM on container completion, so you won't have to pay for idle VM time. 

Finally, in this example I'll parses small gzip file from Google Cloud Storage (GCS), but since this approach is not limited by client timeouts you can use it to do pretty much anything… transformations on bigger files, lightweight ETL, or media format encoding. You can even combine it with GCP Task Queue for more complex pipelines.

## Pre-requirements

If you don't have one already, start by creating new project and configuring [Google Cloud SDK](https://cloud.google.com/sdk/docs/).

## Setup

To start, clone this repo, and navigate into that directory:

```shell
git clone https://github.com/mchmarny/long-running-job.git
cd long-running-job
```

## Run in GCE

> Note, to keep this readme short, I prepared series of scripts that you can execute rather than listing the complete commands. You should absolutely review each one of these scripts for content before executing it. This will help you understand the individual commands and allow you use them in the future.

### Service Account

To execute this sample you will need a GCP service account. You can do that either in UI or using `gcloud` SDK. To find out more read [creating and managing service accounts](https://cloud.google.com/iam/docs/creating-managing-service-accounts).

You can create a specific service account for this demo using the [bin/account](bin/account) script. This script will also assign the new account all the necessary IAM roles and provision a service account key which will be saved in the `~/.gcp-keys` folder in your home directory. You should protect that key or just delete it after this demo.

```shell
bin/account
```

### Container Image

The unit of code delivery to to the GCE VM will be container image. To create an image from this demo, you can use the [bin/image](bin/image) script

```shell
bin/image
```

### Deploy Container to VM

To create a new GCE VM and configure it to run the above built image, execute the [bin/deploy](bin/deploy) script

```shell
bin/deploy
```

### Container Logs

Once the VM started you can monitor the logs output from the VM to Stackdriver using the [bin/monitor](bin/monitor) script

> Note, this command will print only the logs that are output by the user code in the container. You can see the complete list of log entries by removing the `jsonPayload.message:"[LRJ]"` filter.

```shell
bin/monitor
```

After the container exists, the VM will be shutdown but the logs should be still available in Stackdriver for forensic analyses

## Run Locally

You can run this code also locally by executing the [bin/run](bin/run) script

```shell
bin/run
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.