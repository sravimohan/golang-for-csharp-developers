using System;

namespace dotnet
{
    public interface IAnimal
    {
        string Speak();

    }

    public class Dog : IAnimal
    {
        public string Speak()
        {
            return "bark";
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }
    }
}
