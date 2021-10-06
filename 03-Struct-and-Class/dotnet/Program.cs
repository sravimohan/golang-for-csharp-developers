using System;

namespace dotnet
{
    public class Employee
    {
        public int Id { get; set; }

        public string FirstName { get; set; }

        public string LastName { get; set; }

        public int Age { get; set; }

        public void ProcessPayroll()
        {
            Console.WriteLine("processing payroll");
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var emp =
                new Employee {
                    Id = 1,
                    FirstName = "hello",
                    LastName = "world",
                    Age = 20
                };

            Console.WriteLine(emp.FirstName);
            emp.ProcessPayroll();
        }
    }
}
