import abc
import json
import requests

from http.server import BaseHTTPRequestHandler
from io import BytesIO




class ArithAddParams:
    
    Nums = []
    

    def to_dict(self):
        return {
            
            "Nums": self.Nums,
            
        }

class ArithAddResult:
    
    Sum = 0
    

    def to_dict(self):
        return {
            
            "Sum": self.Sum,
            
        }

class ArithPowParams:
    
    Base = 0
    
    Pow = 0
    

    def to_dict(self):
        return {
            
            "Base": self.Base,
            
            "Pow": self.Pow,
            
        }

class ArithPowResult:
    
    Num = 0
    

    def to_dict(self):
        return {
            
            "Num": self.Num,
            
        }

class ArithIsNegativeParams:
    
    Num = 0
    

    def to_dict(self):
        return {
            
            "Num": self.Num,
            
        }

class ArithIsNegativeResult:
    
    Negative = False
    

    def to_dict(self):
        return {
            
            "Negative": self.Negative,
            
        }


class Arith(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (
            
            hasattr(subclass, 'Add') and callable(subclass.Add) or
            
            hasattr(subclass, 'Pow') and callable(subclass.Pow) or
            
            hasattr(subclass, 'IsNegative') and callable(subclass.IsNegative) or
            
            NotImplemented
        )

    
    @abc.abstractmethod
    def Add(self, params: ArithAddParams):
        raise NotImplementedError
    
    @abc.abstractmethod
    def Pow(self, params: ArithPowParams):
        raise NotImplementedError
    
    @abc.abstractmethod
    def IsNegative(self, params: ArithIsNegativeParams):
        raise NotImplementedError
    



class GreeterSayHelloParams:
    
    From = ""
    
    To = ""
    

    def to_dict(self):
        return {
            
            "From": self.From,
            
            "To": self.To,
            
        }

class GreeterSayHelloResult:
    
    Message = ""
    

    def to_dict(self):
        return {
            
            "Message": self.Message,
            
        }


class Greeter(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (
            
            hasattr(subclass, 'SayHello') and callable(subclass.SayHello) or
            
            NotImplemented
        )

    
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
                    
                    elif rpc_request["method"] == "Pow":
                        params = ArithPowParams()
                        
                        params.Base = rpc_request["params"]["Base"]
                        
                        params.Pow = rpc_request["params"]["Pow"]
                        

                        result = server.Arith.Pow(params)
                    
                    elif rpc_request["method"] == "IsNegative":
                        params = ArithIsNegativeParams()
                        
                        params.Num = rpc_request["params"]["Num"]
                        

                        result = server.Arith.IsNegative(params)
                    
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



class ArithClient(Arith):
    endpoint = ""

    def __init__(self, endpoint: ""):
        self.endpoint = endpoint
        return

    
    def Add(self, params: ArithAddParams):
        rpc_request = {
            "jsonrpc": "2.0",
            "method": "Add",
            "params": params.to_dict()
        }

        response = requests.post(self.endpoint, json=rpc_request).json()
        result = ArithAddResult()
        result.Message = response["result"]
        
        return result
    
    def Pow(self, params: ArithPowParams):
        rpc_request = {
            "jsonrpc": "2.0",
            "method": "Pow",
            "params": params.to_dict()
        }

        response = requests.post(self.endpoint, json=rpc_request).json()
        result = ArithPowResult()
        result.Message = response["result"]
        
        return result
    
    def IsNegative(self, params: ArithIsNegativeParams):
        rpc_request = {
            "jsonrpc": "2.0",
            "method": "IsNegative",
            "params": params.to_dict()
        }

        response = requests.post(self.endpoint, json=rpc_request).json()
        result = ArithIsNegativeResult()
        result.Message = response["result"]
        
        return result
    


class GreeterClient(Greeter):
    endpoint = ""

    def __init__(self, endpoint: ""):
        self.endpoint = endpoint
        return

    
    def SayHello(self, params: GreeterSayHelloParams):
        rpc_request = {
            "jsonrpc": "2.0",
            "method": "SayHello",
            "params": params.to_dict()
        }

        response = requests.post(self.endpoint, json=rpc_request).json()
        result = GreeterSayHelloResult()
        result.Message = response["result"]
        
        return result
    


class Client:
    
    Arith = None
    
    Greeter = None
    

    def __init__(self, endpoint: ""):
        
        self.Arith = ArithClient(endpoint + "/arith")
        
        self.Greeter = GreeterClient(endpoint + "/greeter")
        
        return