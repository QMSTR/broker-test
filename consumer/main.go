package main

import (
   "context"
   "github.com/kubemq-io/kubemq-go"
   "log"
)

func main() {
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   client, err := kubemq.NewClient(ctx,
      kubemq.WithAddress("kubemq-cluster-grpc.kubemq", 50000),
      kubemq.WithClientId("test-command-client-id"),
      kubemq.WithTransportType(kubemq.TransportTypeGRPC))
   if err != nil {
      log.Fatal(err)
   }
   defer client.Close()
   channel := "hi-world-queue"

   receiveResult, err := client.NewReceiveQueueMessagesRequest().
      SetChannel(channel).
      SetMaxNumberOfMessages(1).
      SetWaitTimeSeconds(5).
      Send(ctx)
   if err != nil {
      log.Fatal(err)
   }
   log.Printf("Received %d Messages:\n", receiveResult.MessagesReceived)
   for _, msg := range receiveResult.Messages {
      log.Printf("MessageID: %s, Body: %s", msg.MessageID, string(msg.Body))
   }
}
