# RabbitMQ Work Queue example

Based on [these tutorials](https://www.rabbitmq.com/tutorials/tutorial-two-go.html),
this repo contains a MWE of a master that sends messages to various modules:
a [builder](modules/builder),
an [analyzer](modules/analyzer),
and a [reporter](modules/reporter).

Those modules then send a response back to the master.

## Demo setup

- A [RabbitMQ instance](https://kubernetes.io/docs/tasks/job/coarse-parallel-processing-work-queue/#starting-a-message-queue-service)
- One [master Job](master/master.yaml) issuing messages in all queues
- Three dummy [module ReplicaSets](modules) containing Pods waiting for different kind of messages 

## Run the demo

1. Start RabbitMQ:
    ```bash
    kubectl apply -k rabbitmq
    ```
1. Start any subset of dummy modules:
    ```bash
    kubectl apply -k modules/builder
    kubectl apply -k modules/analyzer
    kubectl apply -k modules/reporter
    ```
   or, to start them all:
    ```bash
    kubectl apply -k modules
    ```
1. Start the master:
    ```bash
    kubectl apply -k master
    ```

Exactly one Pod of each [module ReplicaSet](modules) will consume its corresponding message.\
All the modules will then reply after having performed some dummy computation.

## RabbitMQ dashboard

1. Port-forward RabbitMQ's management port:
    ```bash
    kubectl port-forward svc/rabbitmq-service 15672:15672
    ```
1. Visit [localhost:15672](http://localhost:15672) and use the following credentials to log in:
    - Username: `guest`
    - Password: `guest`
