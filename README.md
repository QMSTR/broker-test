# RabbitMQ Work Queue example

Based on [these tutorials](https://www.rabbitmq.com/tutorials/tutorial-two-go.html),
this repo contains a MWE of a master that sends messages to various modules:
a [builder](modules/builder),
an [analyzer](modules/analyzer),
and a [reporter](modules/reporter).

## Demo setup

- A [RabbitMQ instance](https://kubernetes.io/docs/tasks/job/coarse-parallel-processing-work-queue/#starting-a-message-queue-service)
- One [master Job](master/master.yaml) issuing messages in all queues
- Three dummy [module ReplicaSets](modules) containing Pods waiting for different kind of messages 

## Run the demo

1. Start RabbitMQ:
    ```bash
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-service.yaml
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-controller.yaml
    ```
1. Start any subset of dummy modules:
    ```bash
    kubectl apply -k modules/builder
    kubectl apply -k modules/analyzer
    kubectl apply -k modules/reporter
    ```
1. Start the master:
    ```bash
    kubectl apply -k master
    ```

Exactly one Pod of each [module ReplicaSet](modules) will consume its corresponding message.
