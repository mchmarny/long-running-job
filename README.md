# long-running-job

Demo project showcasing how to use GCE's container exec feature to run long running jobs. This example also shows how to terminate (delete) the VM upon finish of the container execution.

## Pre-requirements

If you don't have one already, start by creating new project and configuring [Google Cloud SDK](https://cloud.google.com/sdk/docs/). Similarly, if you have not done so already, you will have [set up Cloud Run](https://cloud.google.com/run/docs/setup).

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

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.