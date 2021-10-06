using System;

namespace dotnet
{
    struct PersonStruct
    {
        public string Name { get; set; }
    }

    class PersonClass
    {
        public string Name { get; set; }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var p = new Program();

            // void return
            p.FunctionWithVoidReturn();

            // string return
            var name = p.FunctionWithStringReturn();
            Console.WriteLine($"FunctionWithStringReturn : {name}");

            // tuple return
            var (result, value) = p.FunctionWithTupleReturn();
            Console.WriteLine($"FunctionWithTupleReturn : {result}, {value}");

            // by value
            name = "Hello";
            p.FunctionByValue(name);
            Console.WriteLine($"FunctionByValue : {name}");

            // by ref
            name = "Hello";
            p.FunctionByReference(ref name);
            Console.WriteLine($"FunctionByReference : {name}");

            // struct
            var personStruct = new PersonStruct { Name = "Hello" };
            p.FunctionWithStruct(personStruct);
            Console.WriteLine($"FunctionWithStruct : {personStruct.Name}");

            // class
            var personClass = new PersonClass { Name = "Hello" };
            p.FunctionWithClass(personClass);
            Console.WriteLine($"FunctionWithClass : {personClass.Name}");
        }

        private void FunctionWithVoidReturn()
        {
            // nothing
        }

        private string FunctionWithStringReturn()
        {
            return "Hello";
        }

        private (bool, string) FunctionWithTupleReturn()
        {
            return (true, "Hello");
        }

        private void FunctionByValue(string name)
        {
            name = "Hello World";
        }

        private void FunctionByReference(ref string name)
        {
            name = "Hello World";
        }

        private void FunctionWithStruct(PersonStruct person)
        {
            person.Name = "Hello World";
        }

        private void FunctionWithClass(PersonClass person)
        {
            person.Name = "Hello World";
        }
    }
}
