const rpc = require("./rpc/spec.rpc.node");
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
        return {Message: "Hello " + params.To + " from " + params.From};
    }
}

const server = new rpc.Server();
server.Arith = new arith();
server.Greeter = new greeter();

const httpServer = http.createServer(server.requestListener());

httpServer.listen("8080", () => {
    const c = new rpc.Client("http://localhost:8080");

    c.Greeter.SayHello({From: "Blain", To: "Austin"}).then((data) => {
        console.log("caller", data)
    });
});