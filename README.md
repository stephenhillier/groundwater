# Testing gRPC / NATS / grpc-gateway services

This repo contains notes and testing for communication between microservices in a way that enables services to react to changes in other systems (e.g., we want the **wells** service to have an idea about which aquifers exist, and we want the **aquifers** service to know when new wells are dug into an aquifer).

We can choreograph the way systems respond to events (like the creation of a new well in the wells service) by ensuring each system knows what to do when something is created, updated, or removed in a service it is interested in.
See: https://microservices.io/patterns/data/saga.html

## How to test with docker-compose:

1. Clone repo
2. run `docker-compose up`

As soon as the services start, a job container will start that will send 5 `create well` requests; you can observe the output from the wells and aquifers service.

```
run-tests_1      | 2019/03/16 02:56:05 Running tests...
run-tests_1      | 2019/03/16 02:56:05 1. Create well
run-tests_1      | 2019/03/16 02:56:05   POST http://wells-gateway:8000/v1/wells
wells_1          | 2019/03/16 02:56:05 created well 8491, publishing to stream...
aquifers_1       | 2019/03/16 02:56:05 received well created notification: {"id":8491,"aquifer":1,"owner":"Johnny","depth":10.1}
```

The aquifer system can now act on the information it received about a new well in aquifer `1`.

## Ensuring data reaches the stream

When a request for a new `Well` arrives, we don't need to place it on the event stream for the Wells service to pick up.  We can make a request directly to the **wells** service to create it. If an error occurs, the user can be notified immediately and no event needs to be placed onto the stream.

After the well create request is successfully processed by the wells service, we should publish the `well-created` event onto the stream.

To ensure that the message makes it to the message server, we can place it in an "outbox" table in the wells service.  If the publish event was succesful, we can delete the message from the outbox.  If it fails, we can periodically retry sending messages in the outbox until we verify it was published succesfully.

See:
* http://gistlabs.com/2014/05/the-outbox/
* [Asyncronous Message-based communication](https://docs.microsoft.com/en-us/dotnet/standard/microservices-architecture/architect-microservice-container-applications/asynchronous-message-based-communication#resiliently-publishing-to-the-event-bus)
