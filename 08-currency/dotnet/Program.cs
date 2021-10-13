using System;
using System.Threading;
using System.Threading.Tasks;

namespace dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            var task1 = Task.Run(() =>
            {
                Work("Task 1", 2);
            });

            var task2 = Task.Run(() =>
            {
                Work("Task 2", 1);
            });

            Task.WaitAll(task1, task2);
        }

        private static void Work(string name, int sleep)
        {
            Thread.Sleep(sleep);
            Console.WriteLine($"{name} completed");
        }
    }
}
