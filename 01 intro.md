## API

fundamentally

- send a REQUEST (client)
- send a RESPONSE (server)

## gRPC

- free and open-source framework developed by Google
- part of Cloud Native Computational Foundation (CNCF) - as is Docker, Kubernetes, etc

You define REQUEST and RESPONSE for RPC (Remote Procedure Calls), the rest is handled for you

Built on top of HTTP/2, has low latency, supports streaming, is language independent

- you can plug in
  - authentication
  - load balancing
  - logging
  - monitoring

---

## RPC

RPC is a Remote Procedure Call

In the client, just calling a function on the server

Client

```Go
server.CreateUser(user) ...
```

Server

  ```Go
  // function to create users
  def CreateUser(User user) { ...}
```
