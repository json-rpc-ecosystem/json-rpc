# json-rpc

## Very Experimental and a Work in Progress

Generated clients and servers for the [JSON-RPC 2.0 specification](https://www.jsonrpc.org/specification) similar to gRPC and `protoc`.

## Basic Usage

1. Define a specification as seen in `./example/spec.rpc`
2. Generate Go code with `go run cmd/jsonrpcgen/main.go -spec-file=./example/spec.rpc -go-out-dir=./example/rpc`
3. Implement server as and use client as seen in `./example/main.go`

## Goals

- Support full JSON-RPC 2.0 specification
- Simple alternative to [gRPC](https://grpc.io) and [protocol buffers](https://developers.google.com/protocol-buffers/)
- Language agnostic for clients and servers
- Avoid all type reflection and casting where possible