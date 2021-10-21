using Microsoft.Extensions.Logging;

namespace dotnet
{
    class Program
    {
        public static void Main(string[] args = null)
        {
            using (var loggerFactory = LoggerFactory.Create(
                builder => builder.AddConsole()
                    .AddFilter("dotnet.Program", LogLevel.Trace)))
            {
                ILogger logger = loggerFactory.CreateLogger<Program>();
                logger.LogTrace("Hello World!");
                logger.LogDebug("Hello World!");
                logger.LogInformation("Hello World!");
                logger.LogWarning("Hello World!");
                logger.LogError("Hello World!");
                logger.LogCritical("Hello World!");
            }
        }
    }
}
