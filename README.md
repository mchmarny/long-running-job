# long-running-job

Demo project showcasing how to use GCE's container exec feature to run long running jobs. This example also shows hot terminate (delete) the VM after the user code in container exited.

## Pre-requirements

### GCP Project and gcloud SDK

If you don't have one already, start by creating new project and configuring [Google Cloud SDK](https://cloud.google.com/sdk/docs/). Similarly, if you have not done so already, you will have [set up Cloud Run](https://cloud.google.com/run/docs/setup).

### GCP Service Account

> TODO:

## Setup

To start, clone this repo:

```shell
git clone https://github.com/mchmarny/long-running-job.git
```

And navigate into that directory:

```shell
cd long-running-job
```


## Run Locally

```shell
bin/run
```

## Run in GCE

### Container Image

First, build container image from the included source using the [bin/image](bin/image) script

```shell
bin/image
```

### Deploy Container to VM

Create a new GCE VM and configure it to run above built image using the [bin/deploy](bin/deploy) script

> Make sure you have a valid service account key configured in [bin/config](bin/config) script

```shell
bin/deploy
```

### Tail Container Logs

Once the VM started you can monitor the logs output from the VM to Stackdriver using the [bin/monitor](bin/monitor) script

> Note, this command will output only the logs that are output by the user code in the container. To see complete list remove the `jsonPayload.message:"[LRJ]"` filter

```shell
bin/monitor
```

After the container exists, the VM will be shutdown but the logs should be still available in Stackdriver for forensic analyses

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.