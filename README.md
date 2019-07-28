# istio-http-lb

## dep

golang 1.12

kubernests 1.15

istio 1.2.2

## build

```
cd frontend; ./build.sh; cd ..
cd backend; ./build.sh; cd ..
```

## run

deploy k8s deployment and service

```
kubectl apply -f k8s-backend.yaml
kubectl apply -f k8s-frontend.yaml
```

deploy base istio virtualService and DestinationRule

```
kubectl apply -f traffic-weights.yaml
```

fault abort

```
kubectl apply -f vs-fault-abort.yaml
```

fault delay

```
kubectl apply -f vs-fault-delay.yaml
```

timeout

```
kubectl apply -f vs-timeout.yaml
```

retry

```
kubectl apply -f vs-retry.yaml
```

match user-agent

```
kubectl apply -f vs-header-match.yaml
```

rate limiter

```
kubectl apply -f vs-ratelimit.yaml
```

## api

hostname

```
curl 127.0.0.1:3001/info
```

version and hostname

```
curl 127.0.0.1:3001/hello
```

resp status code 555

```
curl 127.0.0.1:3001/abort
```

sleep 30 second

```
curl 127.0.0.1:3001/timeout
```

active raise error

```
curl 127.0.0.1:3001/retry
```
