## Unary API

- the basic request/response you're familiar with
  - client sends one message to the server
  - client receives one response from server
- Unary RPC calls are the most common for your APIs
  - well suited when the data is small
  - start with Unary, switch to steaming API if performance becomes an issue

---
In gRPC, Unary calls are defined using Protocol Buffers

e.g.

```go
message GreetRequest {
  Greeting greeting = 1;
}

message GreetResponse {
  string result = 1;
}

service GreetService {
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

For each gRPC call, need to define a "request" message and a "response" message
