using System;

namespace dotnet
{
    class Program
    {
        static void Main(string[] args)
        {
            var status = "ok";

            // if
            if (status == "ok")
                Console.WriteLine($"If : Status is {status}");

            // if else
            status = "bad";
            if (status == "ok")
            {
                Console.WriteLine($"If : Status is {status}");
            }
            else
            {
                Console.WriteLine($"If Else : Status is {status}");
            }

            // switch
            var id = 0;
            Console.Write("Switch");
            switch (id)
            {
                case 0:
                    Console.Write($" : {id}");
                    break;
                case 1:
                    Console.Write($" : {id}");
                    break;
                case 2:
                case 3:
                    Console.Write($" : {id}");
                    break;
                default:
                    Console.WriteLine(" : Invalid");
                    break;
            }

            // for
            Console.Write("\nFor");
            for (int i = 0; i < 10; i++)
            {
                if (i > 5)
                    break;

                if (i == 2)
                    continue;

                Console.Write($" : {i}");
            }

            // for each
            Console.Write("\nFor Each");
            var ids = new int[] { 1, 2, 3 };
            foreach (var i in ids)
            {
                Console.Write($" : {i}");
            }

            // while
            var y = 0;
            Console.Write("\nWhile");
            while (y < 10)
            {
                if (y > 5)
                    break;

                if (y == 2)
                {
                    y++;
                    continue;
                }

                Console.Write($" : {y}");
                y++;
            }

            // do while
            var z = 0;
            Console.Write("\nDo While");
            do
            {
                if (z > 5)
                    break;

                if (z == 2)
                {
                    z++;
                    continue;
                }

                Console.Write($" : {z}");
                z++;
            } while (z < 10);
        }
    }
}
