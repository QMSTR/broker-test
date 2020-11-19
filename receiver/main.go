package main

import (
   "context"
   "github.com/kubemq-io/kubemq-go"
   "log"
   "time"
)

func main() {
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   client, err := kubemq.NewClient(ctx,
      kubemq.WithAddress("localhost", 50000),
      kubemq.WithClientId("test-command-client-id"),
      kubemq.WithTransportType(kubemq.TransportTypeGRPC))
   if err != nil {
      log.Fatal(err)
   }
   defer client.Close()
   channel := "hello-world-queue"

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
      log.Printf("MessageID: %s, Body: %s", msg.Id, string(msg.Body))
   }
}
