# Message Broker test

This repository contains a MWE of a message broker handling:

- Synchronization between _n_ identical consumers and a single producer, a.k.a. "master"
- Optionally, data insertion into a database

## Prerequisite

1. Provisioning a cluster
2. deploying KubeMQ package
3. Port forwarding 
   - kubemqctl set cluster proxy 
       (https://docs.kubemq.io/getting-started/message-patterns/queues)
   or 

   - kubectl -n kubemq port-forward svc/kubemq-cluster-grpc 50000:50000

       


## Steps

1. All consumers (e.g., a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)) wait for an event to occur
1. The master inserts such an event in the message queue after some time
1. Exactly one consumer should wake up and continue its execution
1. The consumer node produces dummy data to be inserted in a [DGraph instance](https://github.com/QMSTR/qmstr/tree/e2d0401f804ffeed8c7e2aafdd22b7889504cf15/deploy/dgraph)
    - This insertion could be performed by the message broker itself

## Resources

- [minikube](https://github.com/kubernetes/minikube), to run Kubernetes locally
- [KubeMQ queues API](https://docs.kubemq.io/getting-started/message-patterns/queues)
- [Golang Docker images based on `scratch`](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)
