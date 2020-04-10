## getting started

Install Go

Install & configure VSCode

On Mac: `brew install protobuf`

`go get -u google.golang.org/grpc`

`go get -u github.com/golang/protobuf/protoc-gen-go`

---

Go to the dir where you'll be working and initialize your module
  
 - ` go mod init go/gRPC/thomas`

Create `greet` dir

Create `greetpb` subdir

Create `greet.proto`

```Go
syntax = "proto3";

package greet;
option go_package="greetpb";

service GreetService{}
```

execute

```Go
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

to create a `pb.go` file (in our case, greet.pb.go)

---

To make life easier, can put that command into a script:

`generate.sh`
```sh
#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

and run with
```sh
generate.sh
```

Take a few minutes to look through the generated file, `greet.pb.go`; find `RegisterGreetServiceServer` and look at it
