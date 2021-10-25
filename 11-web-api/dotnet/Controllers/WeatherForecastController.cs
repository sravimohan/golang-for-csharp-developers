using Microsoft.AspNetCore.Mvc;

namespace dotnet.Controllers;

[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    private readonly ILogger<WeatherForecastController> _logger;
    private static IList<WeatherForecast> WeatherForecast = new List<WeatherForecast>();

    static WeatherForecastController()
    {
        #region setup mock data

        var summaries = new[]
        {
            "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
        };

        WeatherForecast = Enumerable.Range(1, 5).Select(index => new WeatherForecast
        {
            Date = DateTime.Now.AddDays(index).Date,
            TemperatureC = Random.Shared.Next(-20, 55),
            Summary = summaries[Random.Shared.Next(summaries.Length)]
        }).ToList();

        #endregion
    }

    public WeatherForecastController(ILogger<WeatherForecastController> logger)
    {
        _logger = logger;
    }

    [HttpGet]
    public IEnumerable<WeatherForecast> Get() => WeatherForecast.AsEnumerable();

    [HttpPost]
    public ActionResult<WeatherForecast> Post(WeatherForecast weatherForecast)
    {
        WeatherForecast.Add(weatherForecast);
        return CreatedAtAction(nameof(Get), new { Date = weatherForecast.Date }, weatherForecast);
    }
}
