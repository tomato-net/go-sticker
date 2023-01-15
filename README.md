# go-sticker

A stock ticker written in Go.

## Running locally

The application can be deployed to a local kubernetes cluster using
the config found in the `deploy/` directory.

**NOTE:** If you are running on MacOS, to get Ingress functioning you should use `kind` and follow the ingress
set up instructions [here](https://kind.sigs.k8s.io/docs/user/ingress) and then add stocks-ticker-app.test to `/etc/hosts`
pointing at 127.0.0.1. If you are on linux, you can alternatively use `minikube` with the ingress and ingress-dns plugins.

### Option A: Run published image

The application is published on 
[Docker Hub](https://hub.docker.com/repository/docker/tomatod4r/stock-ticker/general)
and can be deployed to a local Kubernetes cluster (minikube, kind, etc).

```
make deploy-published
```

### Option B: Run local image

The image can be built locally and deployed to a local Kubernetes cluster.
```
IMG=stock-ticker make deploy-local
```

### Option C: No kubectl apply -k or kustomize support
If the above options have failed due to `kubectl` not supporting `apply -k`, then the `kustomize` config has been 
compiled in `deployments/compiled/k8s.yaml` and can be deployed via `kubectl apply -f deployments/compiled/k8s.yaml`.

### Accessing the app

The deployment config creates a deployment, service, and ingress in
Kubernetes to allow access to the application. If you are using `minikube` then
you can enable support for Ingress locally by following 
[this guide](https://minikube.sigs.k8s.io/docs/handbook/addons/ingress-dns/#installation).

Once Ingress is enabled on the cluster, you can access the app via http://stocks-ticker-app.test/v1/stock
```
curl -XGET http://stocks-ticker-app.test/v1/stock | jq .
```

**NOTE:** The `minikube` ingress plugins do not function on MacOS as noted 
[here](https://minikube.sigs.k8s.io/docs/drivers/docker/#known-issues).
In this case you can use `kind` and install an [ingress controller to the cluster](https://kind.sigs.k8s.io/docs/user/ingress).
You will also need to add stocks-ticker-app.test to `/etc/hosts` pointing at 127.0.0.1.
Alternatively, port forward to the service directly, skipping the ingress. 
If port forwarding is required you can run `kubectl port-forward -n stocks-ticker service/stocks-server 8080:8080` and 
access the app on http://localhost:8080/v1/stock


## Notes
* The code is barely tested due to time constraints, most tests are in the `pkg/alphavantage` package.
* Was unable to get the `minikube` ingress plugin to work on MacOS as this is a known error, instead tested ingress
using `kind` and installing the contour ingress controller into the `kind` cluster. See 
[here](https://kind.sigs.k8s.io/docs/user/ingress) for setup instructions.
* The path for the GET request is on `/v1/stock` not `/`.
* The `TIME_SERIES_DAILY` endpoint is a premium endpoint, requiring a premium API key, therefore
this code uses the `TIME_SERIES_DAILY_ADJUSTED` endpoint as this it not premium and returns
mostly identical data.
* Dockerfile is at `build/server/Dockerfile`

### Resilience
Additional resilience steps if not restrained by time:
* Would add metrics to the endpoints to track connection times to AlphaVantage and various other metrics through the application, and
then expose them with a metrics endpoint on the web app, scraped by something such as prometheus.
* Alerting rules for the application uptime, and various error metrics in the app.
* Structured logging throughout the application to ensure error cases and important logs are visible and searchable.
* Add pprof profiling endpoints to the web app
* Add additional unit tests to all areas of the codebase to ensure functionality.
* Add integration tests for the end-to-end web app process.
* Better handling of the AlphaVantage error cases, though the AlphaVantage API is fairly inconsistent with it's errors, with them being
in different formats and still returning 200 HTTP codes on error states, making it trickier to handle.
* Automated pipelines pre-merge to run tests, and post-merge pipelines to deploy the code.
* Backoff loop handling of API quota limit errors when communicating with AlphaVantage API.
* Better documentation in code for the various structs and functions.
* Swagger API documentation
* Better handling of environment variables, fail fast if not set, and validate them.
* Better error wrapping
* NetworkPolicies restricting access to the app
