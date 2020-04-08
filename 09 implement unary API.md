## Unary "greet" API

- our message is
  - `Greeting`
- contains the strings
  - `first_name`
  - `last_name`
- takes a `GreetRequest` which contains a
  - `Greeting`
- returns a `GreetResponse` which contains the string
  - `result`

in the `greet.proto` file add the message:

```proto
message Greeting {
  string first_name = 1;
  string last_name = 2;
}
```

the tags, 1 and 2, are showing which field they are. That is, first_name is the first field of our Greeting message.

then add the `request` (it uses a Greeting):

```proto
message GreetRequest {
  Greeting greeting = 1;
}
```

then add the `response` (returning the string 'result'):

```proto
message GreetResponse {
  string result = 1;
}
```

---
update `GreetService`

note that `Greet` takes input and returns output: `Greet(input) returns (output)`

the `{}` are for options

```proto
service GreetService {
  // unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

Convention is when naming your RPC 'Something', the request and response are 'SomethingRequest' and 'SomethingResponse'

---
the file should now look like:

```proto
syntax = "proto3";

package greet;
option go_package="greetpb";

message Greeting {
  string first_name = 1;
  string last_name = 2;
}

message GreetRequest {
  Greeting greeting = 1;
}

message GreetResponse {
  string result = 1;
}

service GreetService {
  // unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

run either
  - `sh generate.sh`

or

- `protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.`

to generate the new pb.go file (the old one is overwritten)

`greet.pb.go` is significantly longer now, with

- `Greeting struct`
- `GreetRequest struct`
- `GreetResponse struct`

and more having been created. There is also

- `GreetServiceClient interface`
- `GreetServiceServer interface`

so we can call them.
