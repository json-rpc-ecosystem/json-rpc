using System;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;

namespace example
{
    class arith : JsonRpc.IArith
    {
        public JsonRpc.ArithAddResult Add(JsonRpc.ArithAddParams arithAddParams)
        {
            return new JsonRpc.ArithAddResult{Sum = 1.0};
        }
        public JsonRpc.ArithPowResult Pow(JsonRpc.ArithPowParams arithPowParams)
        {
            return new JsonRpc.ArithPowResult{Num = Math.Pow(arithPowParams.Base, arithPowParams.Pow)};
        }
        public JsonRpc.ArithIsNegativeResult IsNegative(JsonRpc.ArithIsNegativeParams arithIsNegativeParams)
        {
            if(arithIsNegativeParams.Num < 0)
            {
                return new JsonRpc.ArithIsNegativeResult{Negative = true};
            }
            
            return new JsonRpc.ArithIsNegativeResult{Negative = false};
        }
    }
    class greeter : JsonRpc.IGreeter
    {
        public JsonRpc.GreeterSayHelloResult SayHello(JsonRpc.GreeterSayHelloParams greeterSayHelloParams)
        {
            return new JsonRpc.GreeterSayHelloResult{Message = "Dear " + greeterSayHelloParams.To + "\nJust saying hello!\n"+greeterSayHelloParams.From};
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            // Create a server
            JsonRpc.Server server = new JsonRpc.Server();

            // Set your service implementations for JsonRpc.IArith and JsonRpc.IGreeter
            server.Greeter = new greeter();
            server.Arith = new arith();

            // Create an HTTP Listener
            HttpListener listener = new HttpListener();
            listener.Prefixes.Add("http://*:8080/");
            listener.Start();

            // Handle each HttpListenerContext from the listener
            Task.Run(() => {
                while (true)
                {
                    HttpListenerContext context = listener.GetContext();
                    server.HandleHttpListenerContext(context);
                }
            });

            // Create a client and call a method
            JsonRpc.Client client = new JsonRpc.Client(new HttpClient(), "http://localhost:8080");
            JsonRpc.GreeterSayHelloResult result = client.Greeter.SayHello(new JsonRpc.GreeterSayHelloParams{From = "Blain", To = "Austin"});
            Console.Write(result.Message);
        }
    }
}