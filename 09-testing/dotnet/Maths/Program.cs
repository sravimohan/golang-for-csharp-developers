using System;

namespace dotnet.Maths
{
    class Program
    {
        static void Main(string[] args)
        {
            var x = int.Parse(args[0]);
            var y = int.Parse(args[1]);

            var calculator = new Calculator();
            Console.WriteLine($"{x} + {y} = {calculator.Add(x, y)}");
            Console.WriteLine($"{x} - {y} = {calculator.Subtract(x, y)}");
            Console.WriteLine($"{x} * {y} = {calculator.Multiply(x, y)}");
            Console.WriteLine($"{x} / {y} = {calculator.Divide(x, y)}");
        }
    }
}
