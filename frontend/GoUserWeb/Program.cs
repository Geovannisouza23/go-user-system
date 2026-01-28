using GoUserWeb.Services;
using GoUserWeb.Filters;

namespace GoUserWeb
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            // MVC
            builder.Services.AddControllersWithViews();

            // 🔴 Necessário para Session
            builder.Services.AddDistributedMemoryCache();

            // Session
            builder.Services.AddSession(options =>
            {
                options.IdleTimeout = TimeSpan.FromHours(1);
                options.Cookie.HttpOnly = true;
                options.Cookie.IsEssential = true;
                options.Cookie.SameSite = SameSiteMode.Lax;
            });

            builder.Services.AddHttpContextAccessor();

            // Services
            builder.Services.AddScoped<AuthApiService>();
            builder.Services.AddScoped<UserApiService>();

            // 🔐 Auth Filter (substitui o middleware)
            builder.Services.AddScoped<AuthFilter>();

            var app = builder.Build();

            if (!app.Environment.IsDevelopment())
            {
                app.UseExceptionHandler("/Profile/Error");
                app.UseHsts();
            }

            app.UseHttpsRedirection();
            app.UseStaticFiles();

            app.UseRouting();

            // 🔑 Session (SEM middleware de auth)
            app.UseSession();

            app.UseAuthorization();

            app.MapControllerRoute(
                name: "default",
                pattern: "{controller=Profile}/{action=Index}/{id?}");

            app.Run();
        }
    }
}
