'use strict';

function parseJSON(response) {
    return response.json();
}

function checkStatus(response) {
    // we assume 400 as valid code here because it's the default return code when sth has gone wrong,
    // but then we have an error within the response, no?
    if (response.status >= 200 && response.status <= 400) {
        return response;
    }

    const error = new Error(response.statusText);
    error.response = response;
    throw error;
}
  
function checkError(data) {
    if (data.error) {
        throw new Error(data.error.message);
    }

    return data;
}

function doRPCRequest(endpoint, method, params) {
    var body = {
        jsonrpc: "2.0",
        method: method,
        params: params
    };

    return fetch(endpoint, {method: "POST", body: JSON.stringify(body)})
            .then(res => checkStatus(res))
            .then(res => parseJSON(res))
            .then(res => checkError(res))
            .catch(console.log);
}



function ArithClient(client, endpoint) {
    this.client = client;
    this.endpoint = endpoint;
}


ArithClient.prototype.Add = function(params) {
    return doRPCRequest(this.client.endpoint + this.endpoint, "Add", params);
}

ArithClient.prototype.Pow = function(params) {
    return doRPCRequest(this.client.endpoint + this.endpoint, "Pow", params);
}

ArithClient.prototype.IsNegative = function(params) {
    return doRPCRequest(this.client.endpoint + this.endpoint, "IsNegative", params);
}



function GreeterClient(client, endpoint) {
    this.client = client;
    this.endpoint = endpoint;
}


GreeterClient.prototype.SayHello = function(params) {
    return doRPCRequest(this.client.endpoint + this.endpoint, "SayHello", params);
}




function Client(endpoint) {
    this.endpoint = endpoint
    
    
    this.Arith = new ArithClient(this, "/arith");
    
    this.Greeter = new GreeterClient(this, "/greeter");
    
}