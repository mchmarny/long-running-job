# cloud-run-gke-go-starterkit

Simple [Cloud Run on GKE](https://cloud.google.com/run/) starter project for `go` meant to be used as a GitHub repository template to speed up new Cloud Run service development by removing the need to re-create the common "boilerplate" code and simplifying common steps.

> The `go` template optimized for the managed version of Cloud Run is located [here](https://github.com/mchmarny/gorunstarter)

## Audience

For developers wanting to quickly start their `go` service development on [Cloud Run on GKE](https://cloud.google.com/run/) from a well-structured template with minimal amount of external dependencies. Note, this starter assumes you already have a Cloud Run on GKE cluster. For help on setting up you cluster on GKE see [here](https://cloud.google.com/run/docs/gke/setup).

> This is not an official or standard Cloud Run project layout, just a set of common initial bits that are helpful to me.

## Usage

To use this template when creating a new Cloud Run on GKE service, just click on the "use this template" button and follow the prompts.

![](https://help.github.com/assets/images/help/repository/use-this-template-button.png)

Your newly created project based on the `cloud-run-gke-go-starterkit` template will also include the two basic steps of Cloud Run developer workflow. You can click on the below links to see the content of these commands.

> For complete build and deploy walk-through see the [Cloud Run Quickstart](https://cloud.google.com/run/docs/quickstarts/build-and-deploy)

[Building a container image](bin/build) which submits job to Cloud Build using the included [Dockerfile](./Dockerfile) and results in versioned, non-root container image URI which will be used to deploy your service to Cloud Run.

```shell
bin/build
```

[Deploying Cloud Run service](bin/deploy) which deploys public Cloud Run service configured with environment variable using the previously built container image.

```shell
bin/deploy
```

## Cleanup

To cleanup all resources created by this sample execute

```shell
bin/delete
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.