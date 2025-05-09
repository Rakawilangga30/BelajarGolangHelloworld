package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Menyajikan file statis dari folder public
	app.Static("/", "./public")

	// Route utama yang menampilkan HTML dengan link ke CSS
	app.Get("/", func(c *fiber.Ctx) error {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Your Brand - Welcome</title>
			<link rel="stylesheet" href="/css/style.css">
		</head>
		<body>
			<header>
				<div class="logo">
					<img src="your-logo.png" alt="Your Brand Logo">
				</div>
				<nav>
					<ul>
						<li><a href="#">Home</a></li>
						<li><a href="#about">About</a></li>
						<li><a href="#services">Services</a></li>
						<li><a href="#contact">Contact</a></li>
					</ul>
				</nav>
			</header>

			<section class="hero">
				<div class="hero-content">
					<h1>Welcome to Your Brand</h1>
					<p>Empowering your business with innovative solutions.</p>
					<a href="#services" class="cta-button">Learn More</a>
				</div>
			</section>

			<section id="about" class="about">
				<h2>About Us</h2>
				<p>Your brand story goes here...</p>
			</section>

			<section id="services" class="services">
				<h2>Our Services</h2>
				<div class="service-item">
					<h3>Service 1</h3>
					<p>Details of the service.</p>
				</div>
				<div class="service-item">
					<h3>Service 2</h3>
					<p>Details of the service.</p>
				</div>
				<div class="service-item">
					<h3>Service 3</h3>
					<p>Details of the service.</p>
				</div>
			</section>
  
			<footer>
				<p>&copy; 2025 Your Brand. All Rights Reserved.</p>
			</footer>

			<script src="/js/script.js"></script>
		</body>
		</html>
		`
		return c.SendString(html)
	})

	// Menjalankan server di port 3000
	app.Listen(":3000")
}
