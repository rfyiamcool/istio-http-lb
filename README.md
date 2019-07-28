# istio-http-lb

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

## test

```
curl 127.0.0.1:3001/hello
```
