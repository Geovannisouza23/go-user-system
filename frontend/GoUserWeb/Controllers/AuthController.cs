using GoUserWeb.Services;
using Microsoft.AspNetCore.Mvc;

public class AuthController : Controller
{
    private readonly AuthApiService _authApiService;

    public AuthController(AuthApiService authApiService)
    {
        _authApiService = authApiService;
    }

    [HttpGet]
    public IActionResult Login()
    {
        return View();
    }

    [HttpPost]
    public async Task<IActionResult> Login(string email, string password)
    {
        var token = await _authApiService.LoginAsync(email, password);

        if (string.IsNullOrEmpty(token))
        {
            ViewBag.Error = "Email ou senha inválidos";
            return View();
        }

        // 🔐 SALVA O TOKEN NA SESSION
        HttpContext.Session.SetString("JWT", token);

        // 👉 REDIRECIONA PARA PROFILE
        return RedirectToAction("Index", "Profile");
    }

    public IActionResult Logout()
    {
        HttpContext.Session.Clear();
        return RedirectToAction("Login");
    }
}
