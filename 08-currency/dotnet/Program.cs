using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace dotnet
{
    class Program
    {
        static HttpClient httpClient = new HttpClient();

        static void Main(string[] args)
        {
            var apis = new string[]
            {
                "https://management.azure.com",
                "https://dev.azure.com",
                "https://api.github.com",
                "https://outlook.office.com/",
                "https://api.somewhereintheinternet.com/",
                "https://graph.microsoft.com",
            };

            var tasks = new List<Task>();
            foreach (var api in apis)
            {
                tasks.Add(CheckApiAsync(api));
            }

            Task.WaitAll(tasks.ToArray());
        }

        private static async Task CheckApiAsync(string api)
        {
            try
            {
                Console.WriteLine($"Checking::{api}");
                await httpClient.GetAsync(api);
                Console.WriteLine($"OK::{api}");
            }
            catch (HttpRequestException)
            {
                Console.WriteLine($"BAD::{api}");
            }
        }
    }
}
