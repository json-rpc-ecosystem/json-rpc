'use strict';

const url = require('url');
const http = require('http');

function doRPCRequest(endpoint, method, params) {
    return new Promise((resolve, reject) => {
        const req = http.request(endpoint, {method: "POST"}, res => {
                const buffers = [];
                res.on('data', chunk => {
                    buffers.push(chunk);
                });

                res.on('end', () => {
                    const jsonRPCResponse = JSON.parse(Buffer.concat(buffers).toString());
                    if(jsonRPCResponse.error) {
                        reject(jsonRPCResponse.error);
                        return
                    }

                    resolve(jsonRPCResponse.result);
                });
            });

        const jsonRPCRequest = {
            jsonrpc: "2.0",
            method: method,
            params: params
        }

        req.write(JSON.stringify(jsonRPCRequest));
        req.end();
    })
}



module.exports.Arith = class Arith {
    constructor() {}

    
    Add(params) {
        throw new Error("Arith does not implment method Add");
    }
    
    Pow(params) {
        throw new Error("Arith does not implment method Pow");
    }
    
    IsNegative(params) {
        throw new Error("Arith does not implment method IsNegative");
    }
    
}


module.exports.Greeter = class Greeter {
    constructor() {}

    
    SayHello(params) {
        throw new Error("Greeter does not implment method SayHello");
    }
    
}


module.exports.Server = class Server {
    constructor() {
        
        this.Arith = null;
        
        this.Greeter = null;
        
    }

    requestListener() {
        return (req, res) => {
            const buffers = [];
            req.on('data', chunk => {
                buffers.push(chunk);
            });

            req.on('end', () => {
                const jsonRPCRequest = JSON.parse(Buffer.concat(buffers).toString());
                const jsonRPCResponse = {jsonrpc: "2.0"};
                
                const path = url.parse(req.url, true).pathname;
                switch(path) {
                    
                    
                    case "/arith":
                        switch(jsonRPCRequest.method) {
                            
                            case "Add":
                                jsonRPCResponse.result = this.Arith.Add(jsonRPCRequest.params);
                                break;
                            
                            case "Pow":
                                jsonRPCResponse.result = this.Arith.Pow(jsonRPCRequest.params);
                                break;
                            
                            case "IsNegative":
                                jsonRPCResponse.result = this.Arith.IsNegative(jsonRPCRequest.params);
                                break;
                            
                        }
                        break;
                    
                    
                    case "/greeter":
                        switch(jsonRPCRequest.method) {
                            
                            case "SayHello":
                                jsonRPCResponse.result = this.Greeter.SayHello(jsonRPCRequest.params);
                                break;
                            
                        }
                        break;
                    
                }

                res.write(JSON.stringify(jsonRPCResponse));
                res.end();
            });
        };
    }
}


class ArithClient extends module.exports.Arith {
    constructor(client, endpoint) {
        super();
        this.client = client;
        this.endpoint = endpoint;
    }

    
    Add(params) {
        return doRPCRequest(this.client.endpoint + this.endpoint, "Add", params);
    }
    
    Pow(params) {
        return doRPCRequest(this.client.endpoint + this.endpoint, "Pow", params);
    }
    
    IsNegative(params) {
        return doRPCRequest(this.client.endpoint + this.endpoint, "IsNegative", params);
    }
    
}

class GreeterClient extends module.exports.Greeter {
    constructor(client, endpoint) {
        super();
        this.client = client;
        this.endpoint = endpoint;
    }

    
    SayHello(params) {
        return doRPCRequest(this.client.endpoint + this.endpoint, "SayHello", params);
    }
    
}


module.exports.Client = class Client {
    constructor(endpoint) {
        this.endpoint = endpoint
        
        this.Arith = new ArithClient(this, "/arith");
        
        this.Greeter = new GreeterClient(this, "/greeter");
        
    }
}