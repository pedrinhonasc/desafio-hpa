# Desafio HPA

In this challenge, it is expected that we do the following:
 - Create a Go application that executes a loop that sums the square of a number and display the message `Code.education Rocks!`. Moreover, this application needs to start a web server on port 8000;
 - Create a unit test for the Go application;
 - Build an image using this Go application, name it as `go-hpa` and push it to DockerHub;
 - Create a CI process that executes the unit test when triggered;
 - Create a deployment and a service named `go-hpa`. Each replica must consume at leat 50m and at most 100m;
 - Implement a HPA that has the following features:
    - The scaling process starts when the CPU utilization is higher than 15%
    - Minimum number of pods: 1
    - Maximum number of pods: 6
 - Create a POD that ueses an endless loop to send requests to the web server and verify if the autoscaler is working as expected: increasing the number of replicas when the CPU utilization is higher than 15% and decreaing it when the CPU utilization is lower than 15%.

**Note**: The Go application image can be pulled from [here](https://hub.docker.com/r/jpedronascimentofilho/go-hpa).

## Goal

The goal of this challenge is to learn about `Horizontal Pod Autoscaler` in practice.

## Usage

To run this challenge, it is expected that `minikube` and `kubectl` was installed previously.

First, enable the `metrics-server` on minikube:
```bash
$ minikube addons enable metrics-server
```

Start minikube:
```bash
$ minikube start --extra-config=controller-manager.horizontal-pod-autoscaler-upscale-delay=1m --extra-config=controller-manager.horizontal-pod-autoscaler-downscale-delay=1m --extra-config=controller-manager.horizontal-pod-autoscaler-sync-period=10s --extra-config=controller-manager.horizontal-pod-autoscaler-downscale-stabilization=1m
```

Then start the deployment, the service and the hpa:
```bash
$ kubectl apply -f k8s/
```

Watch the CPU utilization and the number of replicas:
```bash
$ watch kubectl get hpa
```

Crete the POD:
```bash
$ kubectl run -it loader --image=busybox /bin/sh
```

Inside the POD, run:
```bash
$ while true; do wget -q -O- http://go-hpa.default.svc.cluster.local:8000: done;
```

**Note**: While the endless loop is running, it is possible to notice that as the CPU utilization increases, the number of replicas increases too. When the endless loop is stopped the CPU utilization decreases and the number of replicas does the same as well.