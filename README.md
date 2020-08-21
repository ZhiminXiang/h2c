This doc explans how to set up HTTP2 service with Knative.

The HTTP2 service in this repo simply accepts HTTP2 request, and display

```
"Hello word: path {Your-request-path}
```

## Prerequesties

This doc requires to have a k8s cluster with Knative installed.


## Configure the environment variables

Run the following command to configure the needed environment variables.

```shell
# set REPO to the URL of your container repository
# e.g. export REPO="gcr.io/h2c-test"
export REPO={Your-container-repository}
```

## Build the container image

Run the following command to build the container of the HTTP2 service

```shell
docker build   --tag "${REPO}/h2c"   --file=Dockerfile .
```

## Push the container image to container repository

Run the following command to push the container image to your container repository

```shell
docker push "${REPO}/h2c"
```

## Deploy Knative Service

Run the following command to replace the image of the Knative service with your container repository

```shell
sed -i -e "s@github.com/ZhiminXiang@${REPO}@" h2c.yaml
```

Then run the following command to deploy the Knative Service

```shell
kubectl apply -f h2c.yaml
```

## Verify HTTP2 request

You should be able to see the the Knative service becomes ready by running the command

```shell
kubectl get ksvc h2c
```

And then you can send HTTP2 request to your ingress by running the following command

```shell
curl http://{YOUR-INGRESS-IP} --header "Host:h2c.default.example.com" --http2
```

## Verify the Knative service only accepts HTTP2 request

We require to set port name to `h2c` within the Knative service in order to be able to serve HTTP2 traffic.
If you remove the `h2c` port from your Knative Service, the Knative service should not be able to servce traffic.
