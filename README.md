# Message Broker test

This repository contains a MWE of a message broker handling:

- Synchronization between _n_ identical consumers and a single producer, a.k.a. "master"
- Optionally, data insertion into a database

## Steps

1. All consumers (e.g., a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)) wait for an event to occur
1. The master inserts such an event in the message queue after some time
1. Exactly one consumer should wake up and continue its execution
1. The consumer node produces dummy data to be inserted in a [DGraph instance](https://github.com/QMSTR/qmstr/tree/e2d0401f804ffeed8c7e2aafdd22b7889504cf15/deploy/dgraph)
    - This insertion could be performed by the message broker itself
