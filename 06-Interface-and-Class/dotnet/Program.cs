using System;

namespace dotnet
{
    public interface IAnimal
    {
        string Name { get; }
        string Speak();

    }

    public class Dog : IAnimal
    {
        public string Name { get; }

        public Dog(string name)
        {
            Name = name;
        }

        public string Speak()
        {
            return "bark";
        }
    }

    public class Cat : IAnimal
    {
        public string Name { get; }
        public Cat(string name)
        {
            Name = name;
        }

        public string Speak()
        {
            return "meow";
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var rover = new Dog("rover");
            var tubby = new Cat("tubby");

            var animals = new IAnimal[] { rover, tubby };

            foreach (var animal in animals)
            {
                Console.WriteLine($"{animal.Name} the {animal.GetType()} says {animal.Speak()}");
            }
        }
    }
}
