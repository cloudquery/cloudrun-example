# CloudQuery Cloud Run Example

This repository contains a small Docker file that starts a webserver on port 8080 for use with Cloud Run. When the server receives a request on the port, it starts a CloudQuery sync using a config file that should be mounted at `/secrets/config.yaml`.

## Deployment

This guide is still incomplete, but the rough steps are:

1. Build the image locally
   ```
   docker build --platform=linux/amd64 -t gcr.io/my-project/cloudquery-cloudrun:3.4.0 .
   ```

2. Upload the image to GCR:

   ```
   docker push gcr.io/my-project/cloudquery-cloudrun:3.4.0
   ```

3. Create a cloud run job using the newly pushed image. Make sure to mount your CloudQuery config file at `/secrets/config.yaml` (using Secrets). Note that it is possible to combine sources and destinations in a single config file by separating the sections with `---` (see [the docs](https://www.cloudquery.io/docs/core-concepts/configuration))

4. Schedule the job via [Cloud Scheduler](https://cloud.google.com/scheduler).
