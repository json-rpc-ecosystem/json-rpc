using System;
using System.IO;
using System.Net;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace example
{
    class arith : JsonRpc.IArith
    {
        public JsonRpc.ArithAddResult Add(JsonRpc.ArithAddParams arithAddParams)
        {
            return new JsonRpc.ArithAddResult(1.0);
        }
        public JsonRpc.ArithPowResult Pow(JsonRpc.ArithPowParams arithPowParams)
        {
            return new JsonRpc.ArithPowResult(1.0);
        }
        public JsonRpc.ArithIsNegativeResult IsNegative(JsonRpc.ArithIsNegativeParams arithIsNegativeParams)
        {
            return new JsonRpc.ArithIsNegativeResult(false);
        }
    }
    class greeter : JsonRpc.IGreeter
    {
        public JsonRpc.GreeterSayHelloResult SayHello(JsonRpc.GreeterSayHelloParams greeterSayHelloParams)
        {
            JsonRpc.GreeterSayHelloResult result = new("Dear " + greeterSayHelloParams.To + "\nJust saying hello!\n"+greeterSayHelloParams.From);
            
            return result;
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            JsonRpc.Server server = new JsonRpc.Server();
            server.Greeter = new greeter();
            server.Arith = new arith();

            HttpListener listener = new HttpListener();
            listener.Prefixes.Add("http://*:8080/");
            listener.Start();
            while (true)
            {
                try
                {
                    HttpListenerContext context = listener.GetContext();
                    server.Process(context);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.ToString());
                }
            }
        }
    }
}