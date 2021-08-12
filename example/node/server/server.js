const rpc = require("../../rpc/spec.rpc.node");
const http = require('http');

class arith extends rpc.Arith {
    constructor() {
        super();
    }

    Add() {
        console.log("hello");
    }

    Pow() {
        console.log("hello");
    }

    IsNegative() {
        console.log("hello");
    }
}

class greeter extends rpc.Greeter {
    constructor() {
        super();
    }

    SayHello(params) {
        return {Message: "HELLO " + params.To + " from " + params.From};
    }
}

const server = new rpc.Server();
server.Arith = new arith();
server.Greeter = new greeter();

const httpServer = http.createServer(server.requestListener());

httpServer.listen("8080");