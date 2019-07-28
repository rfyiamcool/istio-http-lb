# istio-http-lb

## dep

golang

## build

```
cd frontend; ./build.sh; cd ..
cd backend; ./build.sh; cd ..
```

## run

```
kubectl apply -f backend.yaml
kubectl apply -f frontend.yaml
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

## test

```
curl 127.0.0.1:3001/hello
curl 127.0.0.1:3001/abort
curl 127.0.0.1:3001/timeout
curl 127.0.0.1:3001/retry
```
