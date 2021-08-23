import os, sys
currentdir = os.path.dirname(os.path.realpath(__file__))
parentdir = os.path.dirname(currentdir)
sys.path.append(parentdir)

import jsonrpc

client = jsonrpc.Client("http://localhost:8080")

params = jsonrpc.GreeterSayHelloParams()
params.From = "Blain"
params.To = "Austin"

result = client.Greeter.SayHello(params)
print(result.Message)