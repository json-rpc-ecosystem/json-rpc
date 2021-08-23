import os, sys
currentdir = os.path.dirname(os.path.realpath(__file__))
parentdir = os.path.dirname(currentdir)
sys.path.append(parentdir)

from http.server import HTTPServer
import jsonrpc

class arith(jsonrpc.Arith):
    def Add(self, params: jsonrpc.ArithAddParams):
        result = jsonrpc.ArithAddResult()

        for num in params.Nums:
            result.Sum += int(num)

        return result

    def Pow(self, params: jsonrpc.ArithPowParams):
        result = jsonrpc.ArithPowResult()

        return result

    def IsNegative(self, params: jsonrpc.ArithIsNegativeParams):
        result = jsonrpc.ArithIsNegativeResult()

        return result

class greeter(jsonrpc.Greeter):
    def SayHello(self, params: jsonrpc.GreeterSayHelloParams):
        result = jsonrpc.GreeterSayHelloResult()

        result.Message = "Dear " + params.To + ",\nHello!\nSincerely," + params.From

        return result

server = jsonrpc.Server()
server.Arith = arith()
server.Greeter = greeter()

httpd = HTTPServer(('localhost', 8080), server.HTTPRequestHandler())
httpd.serve_forever()