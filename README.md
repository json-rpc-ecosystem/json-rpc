# json-rpc

**Warning: Very Experimental and a Work in Progress**

## Goals

- Support full [JSON-RPC 2.0 specification](https://www.jsonrpc.org/specification)
- Simple alternative to [gRPC](https://grpc.io) and [protocol buffers](https://developers.google.com/protocol-buffers/)
- Language agnostic for clients and servers
- Avoid all type reflection and casting where possible
- Use only standard library where possible

## Basic Usage

### 1. Define a specification as seen in `./example/spec.rpc`

```
version = "2.0"

namespace {
    go = "rpc"
    java = "io.json-rpc"
    kotlin = "io.json-rpc"
    rust = "jsonrpc"
    cpp = "JsonRpc"
    csharp = "JsonRpc"
    c = "jsonrpc"
}

service "Arith" {
    description = "exposes basic arithmatic functions"

    endpoint = "/arith"

    method "Add" {
        description = "adds all the numbers together"
        params = { Nums: "[]Number" }
        result = { Sum: "Number" }
    }

    method "Pow" {
        description = "raises the base number to the pow number"
        params = { Base: "Number", Pow: "Number" }
        result = { Num: "Number" }
    }

    method "IsNegative" {
        description = "checks if the number is negative or not"
        params = { Num: "Number" }
        result = { Negative: "Boolean" }
    }
}

service "Greeter" {
    description = "does lots of greetings"

    endpoint = "/greeter"
    
    method "SayHello" {
        description = "composes a message with the from and to parties"
        params = { From: "String", To: "String" }
        result = { Message: "String" }
    }
}
```

### 2. Generate client and server code.

```
> go run cmd/jsonrpcgen/main.go \
    -spec-file=./example/spec.rpc \
    -browser-out-file=./example/rpc/browser-jsonrpc.js \
    -csharp-out-file=./example/csharp/jsronrpc.cs \
    -go-out-file=./example/rpc/jsonrpc.go \
    -node-out-file=./example/rpc/node-jsonrpc.js \
    -python-out-file=./example/python/jsonrpc.py

2021/08/23 14:28:50 Generated:  ./example/rpc/browser-jsonrpc.js
2021/08/23 14:28:50 Generated:  ./example/csharp/jsronrpc.cs
2021/08/23 14:28:50 Generated:  ./example/rpc/jsonrpc.go
2021/08/23 14:28:50 Generated:  ./example/rpc/node-jsonrpc.js
2021/08/23 14:28:50 Generated:  ./example/python/jsonrpc.py
```
### 3. Implement the server in any language.

```go
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

type arith struct{}

func (s *arith) Add(params *rpc.ArithAddParams) (*rpc.ArithAddResult, error) {
	return nil, nil
}

func (s *arith) Pow(params *rpc.ArithPowParams) (*rpc.ArithPowResult, error) {
	result := &rpc.ArithPowResult{}
	result.Num = math.Pow(params.Base, params.Pow)
	return result, nil
}

func (s *arith) IsNegative(params *rpc.ArithIsNegativeParams) (*rpc.ArithIsNegativeResult, error) {
	return nil, nil
}

type greeter struct{}

func (s *greeter) SayHello(params *rpc.GreeterSayHelloParams) (*rpc.GreeterSayHelloResult, error) {
	result := &rpc.GreeterSayHelloResult{}
	result.Message = fmt.Sprintf("Dear %s, Someone named %s says 'hello!'", params.To, params.From)
	return result, nil
}

func main() {
	server := rpc.Server{
		Arith:   &arith{},
		Greeter: &greeter{},
	}

	err := http.ListenAndServe(":8080", server.Mux())
	if err != nil {
		log.Fatal(err)
	}
}

```

#### 4. Use any of the generated clients.

**Browser**

```html
<html>
    <head>
        <title>JSON-RPC Client</title>
    </head>
    <body>
        See the debug console.
    </body>
    <script src="../../rpc/spec.rpc.browser.js"></script>
    <script>
        const c = new window.Client("http://localhost:8080");
        c.Greeter.SayHello({From: "Blain", To: "Austin"}).then(console.log);
    </script>
</html>
```

**C#**

```csharp
using System;
using System.Net;
using System.Net.Http;

namespace example
{
    class Program
    {
        static void Main(string[] args)
        {
            JsonRpc.Client client = new JsonRpc.Client(new HttpClient(), "http://localhost:8080");
            JsonRpc.GreeterSayHelloResult result = client.Greeter.SayHello(new JsonRpc.GreeterSayHelloParams{From = "Blain", To = "Austin"});
            Console.Write(result.Message);
        }
    }
}
```

**Go**

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/json-rpc-ecosystem/json-rpc/example/rpc"
)

func main() {
	c := rpc.NewClient(http.DefaultClient, "http://localhost:8080")

	fmt.Println(c.Greeter.SayHello(&rpc.GreeterSayHelloParams{From: "Blain", To: "Austin"}))
}
```

**Node**

```js
const rpc = require("../../rpc/spec.rpc.node");

const c = new rpc.Client("http://localhost:8080");
c.Greeter.SayHello({From: "Blain", To: "Austin"}).then(console.log);
```

**Python**

```python
import jsonrpc

client = jsonrpc.Client("http://localhost:8080")

params = jsonrpc.GreeterSayHelloParams()
params.From = "Blain"
params.To = "Austin"

result = client.Greeter.SayHello(params)
print(result.Message)
```

## Open-source, not open-contribution

For now, this JSON-RPC tooling is open source but closed to code contributions. This keeps the code base free of proprietary or licensed code but it also helps me continue to maintain and build it.

I am grateful for community involvement, bug reports, & feature requests. I do not wish to come off as anything but welcoming, however, I've made the decision to keep this project closed to contributions for my own mental health and long term viability of the project.