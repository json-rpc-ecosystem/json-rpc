import abc
import json

from http.server import BaseHTTPRequestHandler
from io import BytesIO

class ArithAddParams:
    Nums = []

class ArithAddResult:
    Sum = 0

    def to_dict(self):
        return {
            "Sum": self.Sum,
        }

class Arith(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (hasattr(subclass, 'Add') and callable(subclass.Add) or
            NotImplemented)

    @abc.abstractmethod
    def Add(self, params: ArithAddParams):
        raise NotImplementedError

class GreeterSayHelloParams:
    From = ""
    To = ""

class GreeterSayHelloResult:
    Message = ""

    def to_dict(self):
        return {
            "Message": self.Message,
        }

class Greeter(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (hasattr(subclass, 'SayHello') and callable(subclass.SayHello) or
            NotImplemented)

    @abc.abstractmethod
    def SayHello(self, params: GreeterSayHelloParams):
        raise NotImplementedError

class Server:
    Arith = None
    Greeter = None

    def HTTPRequestHandler(server):
        class HTTPRequestHandler(BaseHTTPRequestHandler):

            def do_POST(handler):
                content_length = int(handler.headers['Content-Length'])
                body = handler.rfile.read(content_length)
                rpc_request = json.loads(body)

                result = None
                if handler.path == "/":
                    return
                elif handler.path == "/arith":
                    if rpc_request["method"] == "":
                        return
                    elif rpc_request["method"] == "Add":
                        params = ArithAddParams()
                        params.Nums = rpc_request["params"]["Nums"]

                        result = server.Arith.Add(params)
                    else:
                        return
                elif handler.path == "/greeter":
                    if rpc_request["method"] == "":
                        return
                    elif rpc_request["method"] == "SayHello":
                        params = GreeterSayHelloParams()
                        params.From = rpc_request["params"]["From"]
                        params.To = rpc_request["params"]["To"]

                        result = server.Greeter.SayHello(params)
                    else:
                        return
                else:
                    return

                rpc_response = {
                    "jsonrpc": "2.0",
                    "result": result.to_dict()
                }

                handler.send_response(200)
                handler.end_headers()
                handler.wfile.write(bytes(json.dumps(rpc_response), 'utf-8'))

        return HTTPRequestHandler