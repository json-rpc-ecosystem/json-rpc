// Code generated by jsonrpcgen. DO NOT EDIT.

using System;
using System.IO;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace JsonRpc
{  
    internal class Request
    {
        [JsonPropertyName("jsonrpc")]
        public string JsonRpc { get; set; }
        
        [JsonPropertyName("method")]
        public string Method { get; set; }

        [JsonPropertyName("params")]
        public object Params { get; set; }

        [JsonPropertyName("id")]
        public int Id { get; set; }
    } 

    internal class Response
    {
        [JsonPropertyName("jsonrpc")]
        public string JsonRpc { get; set; }

        [JsonPropertyName("result")]
        [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
        public object Result { get; set; }

        [JsonPropertyName("error")]
        [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
        public Error Error { get; set; }
        
        [JsonPropertyName("id")]
        public int Id { get; set; }
    }

    internal class Error
    {
        [JsonPropertyName("code")]
        public int Code { get; set; }

        [JsonPropertyName("message")]
        public string Message { get; set; }

        [JsonPropertyName("data")]
        public object Data { get; set; }
    }

    
    
    
    public record ArithAddParams
    {
        
        public double[] Nums { get; set; }
        
    }

    public record ArithAddResult
    {
        
        public double Sum { get; set; }
        
    }
    
    public record ArithPowParams
    {
        
        public double Base { get; set; }
        
        public double Pow { get; set; }
        
    }

    public record ArithPowResult
    {
        
        public double Num { get; set; }
        
    }
    
    public record ArithIsNegativeParams
    {
        
        public double Num { get; set; }
        
    }

    public record ArithIsNegativeResult
    {
        
        public bool Negative { get; set; }
        
    }
    
    
    public interface IArith
    {
        
        ArithAddResult Add(ArithAddParams arithAddParams);
        
        ArithPowResult Pow(ArithPowParams arithPowParams);
        
        ArithIsNegativeResult IsNegative(ArithIsNegativeParams arithIsNegativeParams);
        
    }
    
    
    
    public record GreeterSayHelloParams
    {
        
        public string From { get; set; }
        
        public string To { get; set; }
        
    }

    public record GreeterSayHelloResult
    {
        
        public string Message { get; set; }
        
    }
    
    
    public interface IGreeter
    {
        
        GreeterSayHelloResult SayHello(GreeterSayHelloParams arithSayHelloParams);
        
    }
    
    
    public class Server
    {
        public const int ErrorParsing = -32700;
        public const int ErrorInvalidRequest = -32600;
        public const int ErrorMethodNotFound = -32601;
        public const int ErrorInvalidParams = -32602;
        public const int ErrorInternal = -32603;

        
        public IArith Arith;
        
        public IGreeter Greeter;
        
        
        public void HandleHttpListenerContext(HttpListenerContext context)
        {
            StreamReader reader = new StreamReader(context.Request.InputStream);
            Request request = JsonSerializer.Deserialize<Request>(reader.ReadToEnd());
            reader.Close();

            string rpcparams = ((JsonElement)request.Params).GetRawText();
            object rpcresult = null;

            switch (context.Request.RawUrl)
            {
                
                
                case "/arith":
                    switch (request.Method)
                    {
                        
                        case "Add":
                            rpcresult = this.Arith.Add(JsonSerializer.Deserialize<ArithAddParams>(rpcparams));
                            break;
                        
                        case "Pow":
                            rpcresult = this.Arith.Pow(JsonSerializer.Deserialize<ArithPowParams>(rpcparams));
                            break;
                        
                        case "IsNegative":
                            rpcresult = this.Arith.IsNegative(JsonSerializer.Deserialize<ArithIsNegativeParams>(rpcparams));
                            break;
                        
                        default:
                            break;
                    }
                    break;
                
                
                case "/greeter":
                    switch (request.Method)
                    {
                        
                        case "SayHello":
                            rpcresult = this.Greeter.SayHello(JsonSerializer.Deserialize<GreeterSayHelloParams>(rpcparams));
                            break;
                        
                        default:
                            break;
                    }
                    break;
                
                default:
                    break;
            }

            Response response = new Response{
                JsonRpc = "2.0",
                Result = rpcresult,
                Id = request.Id,
            };

            context.Response.StatusCode = (int)HttpStatusCode.OK;

            StreamWriter writer = new StreamWriter(context.Response.OutputStream);
            writer.WriteLine(JsonSerializer.Serialize(response));
            writer.Close();
        }
    }

    public class Client
    {
        
        public IArith Arith;
        
        public IGreeter Greeter;
        

        public Client(HttpClient httpClient, string endpoint)
        {
            
            this.Arith = new ArithClient(endpoint+"/arith");
            
            this.Greeter = new GreeterClient(endpoint+"/greeter");
            
        }

        internal static R DoRequest<P, R>(string endpoint, string method,  P reqParams)
        {
            Request rpcRequest = new Request
            {
                JsonRpc = "2.0",
                Method = method,
                Params = reqParams,
                Id = 1,
            };

            string json = JsonSerializer.Serialize<Request>(rpcRequest);
            byte[] byteArray = Encoding.UTF8.GetBytes(json);

            WebRequest request = WebRequest.Create(endpoint);
            request.Method = "POST";
            request.ContentType = "application/json";
            request.ContentLength = byteArray.Length;

            Stream reqStream = request.GetRequestStream();
            reqStream.Write(byteArray, 0, byteArray.Length);

            WebResponse response = request.GetResponse();
            Stream respStream = response.GetResponseStream();
            
            StreamReader reader = new StreamReader(respStream);
            Response rpcResponse = JsonSerializer.Deserialize<Response>(reader.ReadToEnd());

            string result = ((JsonElement)rpcResponse.Result).GetRawText();

            return JsonSerializer.Deserialize<R>(result);
        }
    }

    
    
    internal class ArithClient : IArith
    {
        private string endpoint;

        public ArithClient(string endpoint)
        {
            this.endpoint = endpoint;
        }

        
        public ArithAddResult Add(ArithAddParams reqParams)
        {
            return Client.DoRequest<ArithAddParams, ArithAddResult>(this.endpoint, "Add", reqParams);
        }
        
        public ArithPowResult Pow(ArithPowParams reqParams)
        {
            return Client.DoRequest<ArithPowParams, ArithPowResult>(this.endpoint, "Pow", reqParams);
        }
        
        public ArithIsNegativeResult IsNegative(ArithIsNegativeParams reqParams)
        {
            return Client.DoRequest<ArithIsNegativeParams, ArithIsNegativeResult>(this.endpoint, "IsNegative", reqParams);
        }
        
    }
    
    
    internal class GreeterClient : IGreeter
    {
        private string endpoint;

        public GreeterClient(string endpoint)
        {
            this.endpoint = endpoint;
        }

        
        public GreeterSayHelloResult SayHello(GreeterSayHelloParams reqParams)
        {
            return Client.DoRequest<GreeterSayHelloParams, GreeterSayHelloResult>(this.endpoint, "SayHello", reqParams);
        }
        
    }
    
}