## client

Create `greet_client` dir

Create `client.go` file in that dir

```go
package main

import "fmt"

func main() {
  fmt.Println("gRPC client is running")
}
```

You can test with

```go
go run greet/greet_client/client.go
```

if you want

---
## Create the connection

- Dial takes an address and options. gRPC uses SSL by default. `WithInsecure` turns that off.

  ```go
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  ```

(again, make sure "google.golang.org/grpc" appears in import)

Use defer to close the connection when things are done executing
  
```go
// close the connection when done
defer cc.Close()
```

Actually create the client (the `\n` is to make it easier to read)

```go
c := greetpb.NewGreetServiceClient(cc)
fmt.Printf("\nCreated client: %f\n\n", c)
```

(make sure "go/gRPC/thomas/greet/greetpb" is imported)

---
`client.go` should now look like:

```go
package main

import (
  "fmt"
  "go/gRPC/thomas/greet/greetpb"
  "log"

  "google.golang.org/grpc"
)

func main() {
  fmt.Println("gRPC client is running")
  // create the connection
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %v", err)
  }

  // close the connection when done
  defer cc.Close()

  // create the client
  c := greetpb.NewGreetServiceClient(cc)
  fmt.Printf("\nCreated client: %f\n\n", c)
}
```

And you can test with

```go
go run greet/greet_client/client.go
```

you should get a message saying "Created client: &{%!f(*grpc." with a bunch more stuff after
