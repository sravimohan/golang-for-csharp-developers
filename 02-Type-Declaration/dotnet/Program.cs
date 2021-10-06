using System;

namespace dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            // type declaration
            int id = 0;
            var name = "hello";
            var age = 20;
            Console.WriteLine($"{id}, {name}, {age}");

            // conversion
            int num1 = 10;
            float num2 = 10.5F;
            var total =  (float)num1 + num2;
            Console.WriteLine($"{num1},{num2},{total}");
        }
    }
}
