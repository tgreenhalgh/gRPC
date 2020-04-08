## Getting Started

- at the core of gRPC, you need to define the messages and services using **Protocol Buffers**
- the rest of the gRPC code will be generated for you
  - you'll have to provide the implementation for it
- One **.proto** file works for ~12 programming languages, server & client, and allows you to use a framework that scales to millions of RPC per second
  
example.proto

```Go
syntax = "proto3";

message Greeting {
  string first_name = 1;
}

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
