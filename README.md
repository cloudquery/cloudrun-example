# CloudQuery Cloud Run Example

This repository contains a small Docker file that starts a Web server on port 8080 for use with Cloud Run. When the server receives a request on the port, it starts a CloudQuery sync using a config file that should be mounted at `/secrets/config.yaml`.

You should also generate a CloudQuery API key and set it as an environment variable `CLOUDQUERY_API_KEY` in the [Cloud Run configuration](https://cloud.google.com/run/docs/configuring/services/secrets#access-secrets). See more information on generating an API key [here](https://docs.cloudquery.io/docs/deployment/generate-api-key).

## Deployment

This guide is still incomplete, but the rough steps are:

1. Create a Docker repository in Google Cloud Artifact Registry.

2. Get the latest version of CloudQuery Container from [ghcr.io/cloudquery/cloudquery](https://ghcr.io/cloudquery/cloudquery). Update the version in the Dockerfile

3. Build the image locally. Replace the `<REGION>-docker.pkg.dev/<PROJECT>/<REPOSITORY>/<IMAGE_NAME>:<VERSION>` with the proper values.

   ```bash
   docker build --platform=linux/amd64 -t <REGION>-docker.pkg.dev/<PROJECT>/<REPOSITORY>/<IMAGE_NAME>:<VERSION> .
   ```

   Example:

   ```bash
   docker build --platform=linux/amd64 -t europe-north1-docker.pkg.dev/cloudquery-project/cq-repository/cq-image:6.24.2 .

   ```

4. Upload the image to Artifact Registry:

   ```bash
   docker push europe-north1-docker.pkg.dev/cloudquery-project/cq-repository/cq-image:6.24.2
   ```

5. Create a cloud run job using the newly pushed image. Make sure to mount your CloudQuery config file at `/secrets/config.yaml` (using Secrets). Note that it is possible to combine sources and destinations in a single config file by separating the sections with `---` (see [the docs](https://www.cloudquery.io/docs/core-concepts/configuration))

6. Schedule the job via [Cloud Scheduler](https://cloud.google.com/scheduler).
