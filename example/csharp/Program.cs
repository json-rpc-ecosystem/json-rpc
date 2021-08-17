using System;
using System.IO;
using System.Net;

namespace example
{
    class arith : Rpc.IArith
    {

    }
    class greeter : Rpc.IGreeter
    {
        public Rpc.GreeterSayHelloResult SayHello(Rpc.GreeterSayHelloParams greeterSayHelloParams)
        {
            Rpc.GreeterSayHelloResult result = new("Just saying hello!");
            
            return result;
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Rpc.Server server = new Rpc.Server();
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
    
                }
            }
        }
    }
}

namespace Rpc
{
    public class Server
    {
        public IArith Arith;
        public IGreeter Greeter;
        public void Process(HttpListenerContext context)
        {
            Console.WriteLine(context.Request.RawUrl);

            context.Response.StatusCode = (int)HttpStatusCode.OK;

            StreamWriter writer = new StreamWriter(context.Response.OutputStream);
            writer.WriteLine("OK");
            writer.Close();

            context.Response.OutputStream.Close();
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