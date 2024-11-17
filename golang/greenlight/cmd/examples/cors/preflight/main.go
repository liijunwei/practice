package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":9000", "server address")
	flag.Parse()

	log.Println("server started", *addr)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	err := http.ListenAndServe(*addr, handler)
	log.Fatal(err)
}

const html = `
<!doctype html>
<html>

<head>
  <meta charset="UTF-8" />
</head>

<body>
	<h1>Preflight CORS</h1>
	<div id="output"></div>
	<script>
		document.addEventListener('DOMContentLoaded', function(){
			fetch("http://localhost:4000/v1/tokens/authentication", {
				method: "POST",
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					email: 'Emiliano_Cruickshank@example.com',
					password: 'pa55word'
				})
			}).then(
				function(response){
					response.text().then(function(text){
						document.getElementById("output").innerHTML = text;
					});
				},
				function(err){
					document.getElementById("output").innerHTML = err;
				}
			);
		});
	</script>
</body>
</html>
`
