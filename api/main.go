package handler

import (
	"net/http"
)

func Main() {
	// Defining the route
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>This is not a button</title>
		<script src="https://cdn.tailwindcss.com"></script>
	</head>
	<body class="flex flex-col items-center justify-center place-content-evenly w-min-full h-full m-auto bg-gray-700">
	<main class="container mx-auto flex flex-col items-center px-4 py-16 text-center md:py-32 md:px-10 lg:px-32 xl:max-w-3xl">
	<h1 class="mb-4 bg-gradient-to-tl from-amber-300 via-fuchsia-700 to-amber-200 bg-clip-text text-transparent text-4xl font-bold leadi sm:text-5xl">Awesome Frontend GO Test</h1>
	<section class="text-zinc-50 w-3/4 self-center text-center mb-6">
	<p>Explore this serious Golang-powered site, exclusively dedicated to testing frontend development performance.</p>
	<p>Witness the remarkable capabilities of Go in crafting robust and efficient user interfaces.</p>
	<p>Join us on this exciting journey today!</p>
	</section>
	
	<button class="px-8 py-3 m-2 text-lg font-semibold rounded dark:bg-purple-400 dark:text-gray-900" onclick="redirectToVideo()">Let's GO!</button>
	<script>
	function redirectToVideo() {
		window.location.href = 'https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0';
	}
	</script>
	</main>
	<footer class="w-full fixed bottom-0 px-4 py-4 dark:bg-gray-800 dark:text-gray-400">
	<div class="container flex flex-wrap items-center justify-center mx-auto space-y-4 sm:justify-between sm:space-y-0">
		<div class="flex flex-row pr-3 space-x-4 sm:space-x-8">
			<div class="flex items-center justify-center flex-shrink-0 w-12 h-12 rounded-full dark:bg-purple-400">
				<img src="https://i.imgur.com/j1cwHuu.png" class="w-12 h-12 rounded-full" alt="angel-v-dev-logo"/>
			</div>
			<ul class="flex flex-wrap items-center space-x-4 sm:space-x-8">
				<li>
					<a href="https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0">Terms of Use</a>
				</li>
				<li>
					<a href="https://go.dev/">About GO</a>
				</li>
			</ul>
		</div>
		<ul class="flex flex-wrap pl-3 space-x-4 sm:space-x-8">
			<li>
				<a href="https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0">Instagram</a>
			</li>
			<li>
				<a href="https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0">Facebook</a>
			</li>
			<li>
				<a href="https://twitter.com/AngelScutariV">Twitter</a>
			</li>
		</ul>
	</div>
</footer>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
