using System;

namespace dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            try
            {
                Console.Write("Enter x:");
                float x = float.Parse(Console.ReadLine());
                Console.Write("Enter y:");
                float y = float.Parse(Console.ReadLine());
                var answer = Divide(x, y);
                Console.WriteLine($"{x} / {y} = {answer}");
            }
            catch (DivideByZeroException e)
            {
                Console.WriteLine(e.Message);
            }
            finally
            {
                CleanUp();
            }
        }

        static float Divide(float x, float y)
        {
            if (y == 0)
                throw new DivideByZeroException("Cannot divide by Zero");

            return x / y;
        }

        static void CleanUp()
        {
            Console.WriteLine("Cleaning Up");
        }
    }

    class DivideByZeroException : Exception
    {
        public DivideByZeroException(string message) : base(message)
        {
        }
    }
}
