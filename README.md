# RabbitMQ Work Queue example

Based on [these tutorials](https://www.rabbitmq.com/tutorials/tutorial-two-go.html),
this repo contains a MWE of a master that sends commands to various modules.

## Demo setup

- A [RabbitMQ instance](https://kubernetes.io/docs/tasks/job/coarse-parallel-processing-work-queue/#starting-a-message-queue-service)
- One [master Job](master/master.yaml) that issues one event in the queue
- A [module ReplicaSet](module/module.yaml) containing Pods waiting for messages 

## Run the demo

1. Start RabbitMQ:
    ```bash
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-service.yaml
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-controller.yaml
    ```
1. Start the module:
    ```bash
    kubectl apply -k module
    ```
1. Start the master:
    ```bash
    kubectl apply -k master
    ```

Exactly one Pod of the [module ReplicaSet](module/module.yaml) will consume the message.
