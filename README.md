# long-running-job

Demo of how to use GCE's `create-with-container` feature for long running (potentially stateless) jobs.

## Pre-requirements

### GCP Project and gcloud SDK

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

```shell
bin/image
```

```shell
bin/deploy
```

```shell
bin/monitor
```


## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.