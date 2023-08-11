package main

import (
	"html/template"
)

func main() {
	tmp := template.New("").Parse( `<!DOCTYPE html>
		<html>
		<head>
			<title>Extreme Front in Go</title>
			<script src="https://cdn.tailwindcss.com"></script>
		</head>
		<body class="flex flex-col items-center justify-center place-content-evenly w-min-full h-full m-auto bg-gray-700">
		<main class="container mx-auto flex flex-col items-center px-4 py-16 text-center md:py-32 md:px-10 lg:px-32 xl:max-w-3xl">
		<h1 class="mb-4 bg-gradient-to-tl from-amber-300 via-fuchsia-700 to-amber-200 bg-clip-text text-transparent md:text-6xl font-bold sm:text-9xl">Awesome Frontend GO Test</h1>
		<section class="self-center text-center mb-6 text-zinc-50 sm:mx-1 sm:h-full sm:mt-24 sm:text-6xl sm:w-full md:my-2 md:text-lg md:max-h-1/3">
		<p>Explore this serious Golang-powered site, exclusively dedicated to testing frontend development performance.</p>
		<p>Witness the remarkable capabilities of Go in crafting robust and efficient user interfaces.</p>
		<p>Join us on this exciting journey today!</p>
		</section>
		
		<button class="px-8 py-3 font-semibold rounded sm:text-4xl sm:m-4 md:text-lg md:m-2 dark:bg-purple-400 dark:text-gray-900" onclick="redirectToVideo()">Let's GO!</button>
		<script>
		function redirectToVideo() {
			window.location.href = 'https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0';
		}
		</script>
		</main>
		<footer class="w-full fixed bottom-0 px-4 py-4 dark:bg-gray-800 dark:text-gray-400 sm:h-48 md:h-auto">
		<div class="container flex flex-wrap items-center justify-center mx-auto space-y-4 sm:justify-between sm:h-full sm:space-y-0">
			<div class="flex flex-row pr-3 space-x-4 sm:space-x-8">
				<div class="animate-bounce flex items-center justify-center flex-shrink-0 rounded-full sm:w-24 sm:h-24 md:w-12 md:h-12 dark:bg-purple-400">
					<img src="https://i.imgur.com/j1cwHuu.png" class="rounded-full sm:w-24 sm:h-24 md:w-12 md:h-12 " alt="angel-v-dev-logo"/>
				</div>
				<ul class="flex flex-wrap items-center space-x-4 sm:text-xl md:text-base sm:space-x-8">
					<li>
						<a href="https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0">Terms of Use</a>
					</li>
					<li>
						<a href="https://go.dev/">About GO</a>
					</li>
				</ul>
			</div>
			<ul class="flex flex-wrap pl-3 space-x-4 sm:text-xl md:text-base sm:space-x-8">
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
        `)
		tmp.Execute()
}
