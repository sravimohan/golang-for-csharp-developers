using Microsoft.AspNetCore.Mvc;

namespace dotnet.Controllers;

[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    private readonly ILogger<WeatherForecastController> _logger;
    private static IList<WeatherForecast> _weatherForecast = new List<WeatherForecast>();

    static WeatherForecastController()
    {
        #region setup mock data

        var summaries = new[]
        {
            "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
        };

        _weatherForecast = Enumerable.Range(1, 5).Select(index => new WeatherForecast
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
    public IEnumerable<WeatherForecast> Get() => _weatherForecast.AsEnumerable();

    [HttpGet("{date}")]
    public WeatherForecast? GetByDate(DateTime date)
        => _weatherForecast.FirstOrDefault(f => f.Date == date);

    [HttpPost]
    public ActionResult<WeatherForecast> Post(WeatherForecast weatherForecast)
    {
        _weatherForecast.Add(weatherForecast);
        return CreatedAtAction(nameof(Get), new { Date = weatherForecast.Date }, weatherForecast);
    }
}
