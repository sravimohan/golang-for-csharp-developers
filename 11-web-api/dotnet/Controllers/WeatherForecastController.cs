using Microsoft.AspNetCore.Mvc;

namespace dotnet.Controllers;

[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    private readonly ILogger<WeatherForecastController> _logger;
    private static IDictionary<string, WeatherForecast> _weatherForecast = new Dictionary<string, WeatherForecast>();

    static WeatherForecastController()
    {
        #region setup mock data

        var summaries = new[]
        {
            "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
        };

        for (var i = 0; i < summaries.Length; i++)
        {
            var date = DateTime.Now.AddDays(i).Date.ToString("yyyy-MM-dd");

            _weatherForecast[date] = new WeatherForecast
            {
                Date = date,
                TemperatureC = Random.Shared.Next(-20, 55),
                Summary = summaries[Random.Shared.Next(summaries.Length)]
            };
        }

        #endregion
    }

    public WeatherForecastController(ILogger<WeatherForecastController> logger)
    {
        _logger = logger;
    }

    [HttpGet]
    public IEnumerable<WeatherForecast> Get() => _weatherForecast.Select(f => f.Value);

    [HttpGet("{date}")]
    public ActionResult<WeatherForecast> GetByDate(string date)
    {
        var found = _weatherForecast.TryGetValue(date, out var weatherforecast);
        return found ? weatherforecast : NotFound();
    }

    [HttpPost]
    public ActionResult<WeatherForecast> Post(WeatherForecast weatherForecast)
    {
        _weatherForecast[weatherForecast.Date] = weatherForecast;
        return CreatedAtAction(nameof(Get), new { Date = weatherForecast.Date }, weatherForecast);
    }
}
