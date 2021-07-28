# json-rpc

## Very Experimental and a Work in Progress

Generated clients and servers for the [JSON-RPC 2.0 specification](https://www.jsonrpc.org/specification) similar to gRPC and `protoc`.

## Goals

- Support full JSON-RPC 2.0 specification
- Simple alternative to [gRPC](https://grpc.io) and [protocol buffers](https://developers.google.com/protocol-buffers/)
- Language agnostic for clients and servers
- Avoid all type reflection and casting where possible