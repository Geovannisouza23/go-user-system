using GoUserWeb.Services;
using GoUserWeb.Models;
using Microsoft.AspNetCore.Mvc;

namespace GoUserWeb.Controllers
{
    public class ProfileController : Controller
    {
        private readonly UserApiService _userApiService;

        public ProfileController(UserApiService userApiService)
        {
            _userApiService = userApiService;
        }

        public async Task<IActionResult> Index()
        {
            var token = HttpContext.Session.GetString("JWT");

            if (string.IsNullOrEmpty(token))
                return RedirectToAction("Login", "Auth");

            var profile = await _userApiService.GetProfileAsync(token);

            if (profile == null)
                return RedirectToAction("Login", "Auth");
        
         Console.WriteLine($"DEBUG API -> Id={profile.Id}, Name={profile.Name}, Email={profile.Email}");

            // 🔴 ESTE MAPEAMENTO É O QUE FALTAVA
            var viewModel = new UserProfileViewModel
            {
                Id = profile.Id,
                Name = profile.Name,
                Email = profile.Email
            };

            return View(viewModel);
        }
    }
}
