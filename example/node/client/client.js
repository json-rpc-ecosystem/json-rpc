const rpc = require("../../rpc/spec.rpc.node");

const c = new rpc.Client("http://localhost:8080");
c.Greeter.SayHello({From: "Blain", To: "Austin"}).then(console.log);