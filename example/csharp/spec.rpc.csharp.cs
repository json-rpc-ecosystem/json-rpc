using System;
using System.IO;
using System.Net;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace JsonRpc
{  
    public class Request
    {
        [JsonPropertyName("jsonrpc")]
        public string JsonRpc { get; set; }
        
        [JsonPropertyName("method")]
        public string Method { get; set; }

        [JsonPropertyName("params")]
        public JsonDocument Params { get; set; }

        [JsonPropertyName("id")]
        public int Id { get; set; }
    } 

    public class Response
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

    public class Error
    {
        [JsonPropertyName("code")]
        public int Code { get; set; }

        [JsonPropertyName("message")]
        public string Message { get; set; }

        [JsonPropertyName("data")]
        public object Data { get; set; }
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
        
        public void Process(HttpListenerContext context)
        {
            StreamReader reader = new StreamReader(context.Request.InputStream);
            Request request = JsonSerializer.Deserialize<Request>(reader.ReadToEnd());
            reader.Close();

            string rpcparams = request.Params.RootElement.GetRawText();
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
        private string endpoint;

        public IArith Arith;
        public IGreeter Greeter;

        public Client(string endpoint)
        {
            this.endpoint = endpoint;
        }
    }
    
    public record ArithAddParams(double[] Nums);
    public record ArithAddResult(double Sum);

    public record ArithPowParams(double Base, double Pow);
    public record ArithPowResult(double Num);

    public record ArithIsNegativeParams(double Num);
    public record ArithIsNegativeResult(bool Negative);

    public interface IArith
    {
        ArithAddResult Add(ArithAddParams arithAddParams);
        ArithPowResult Pow(ArithPowParams arithPowParams);
        ArithIsNegativeResult IsNegative(ArithIsNegativeParams arithIsNegativeParams);
    }

    public record GreeterSayHelloParams(string From, string To);
    public record GreeterSayHelloResult(string Message);

    public interface IGreeter
    {
        GreeterSayHelloResult SayHello(GreeterSayHelloParams greeterSayHelloParams);
    }
}