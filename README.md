# RabbitMQ Work Queue example

Based on [these tutorials](https://www.rabbitmq.com/tutorials/tutorial-two-go.html),
this repo contains a MWE of a master that sends commands to various modules.

## Run the demo

1. Start RabbitMQ:
    ```bash
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-service.yaml
    kubectl create -f https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.3/examples/celery-rabbitmq/rabbitmq-controller.yaml
    ```
1. Start the dummy builder module:
    ```bash
    kubectl apply -f builder/builder.yaml
    ```
1. Start the master:
    ```bash
    kubectl apply -f master/master.yaml
    ```

Exactly one Pod of the [builder ReplicaSet](builder/builder.yaml) will consume the message.
