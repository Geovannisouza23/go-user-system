using System.Net.Http.Headers;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace GoUserWeb.Services
{
    public class UserApiService
    {
        private readonly HttpClient _httpClient;

        public UserApiService()
        {
            _httpClient = new HttpClient
            {
                BaseAddress = new Uri("http://localhost:8080")
            };
        }

        public async Task<UserProfile?> GetProfileAsync(string token)
{
    _httpClient.DefaultRequestHeaders.Authorization =
        new AuthenticationHeaderValue("Bearer", token);

    var response = await _httpClient.GetAsync("/me");

    if (!response.IsSuccessStatusCode)
        return null;

    var json = await response.Content.ReadAsStringAsync();
Console.WriteLine("JSON /me => " + json);

return JsonSerializer.Deserialize<UserProfile>(
    json,
    new JsonSerializerOptions
    {
        PropertyNameCaseInsensitive = true
    });
        }
    }

    public class UserProfile
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("name")]
        public string Name { get; set; } = "";

        [JsonPropertyName("email")]
        public string Email { get; set; } = "";
    }
}
