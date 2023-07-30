package main

import (
	"net/http"
)

func main() {
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
	<body class="flex flex-col items-center justify-center place-content-evenly w-min-full h-full m-auto bg-violet-700">
	<main class="flex flex-col items-center justify-center place-content-evenly w-7/10 h-8/10 bg-violet-950/[.45] rounded-md mt-4 pb-4">
	<h1 class="bg-gradient-to-tl from-amber-300 via-fuchsia-700 to-amber-200 bg-clip-text text-transparent text-5xl mt-4 mb-6 font-mono font-extrabold">Awesome Frontend GO Test</h1>
	<section class="text-zinc-50 w-3/4 self-center text-center mb-6">
	<p>Explore this serious Golang-powered site, exclusively dedicated to testing frontend development performance.</p>
	<p>Witness the remarkable capabilities of Go in crafting robust and efficient user interfaces.</p>
	<p>Join us on this exciting journey today!</p>
	</section>
	
	<button class="p-3 rounded-md font-semibold text-slate-50 bg-gradient-to-bl from-slate-800 via-purple-600 to-slate-800" onclick="redirectToVideo()">Let's GO!</button>
	<script>
	function redirectToVideo() {
		window.location.href = 'https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&controls=0';
	}
	</script>
	</main>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
