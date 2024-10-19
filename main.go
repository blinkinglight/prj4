package main

import (
	"log"
	"net/http"
	"time"

	"github.com/delaneyj/datastar"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
    <html>
        <head>
            <script type="module" defer src="https://cdn.jsdelivr.net/npm/@sudodevnull/datastar"></script>
        </head>
        <body>
            <div>
               <h1 data-on-click="$$post('/clock')">Clock</h1>
                <h2>Time: <span id="time"></span></h2>
            </div>
        </body>
    </html>`))
	})

	http.HandleFunc("/clock", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		datastar.RenderFragmentString(sse, `<span id="time">`+time.Now().Format(time.DateTime)+`</span>`)
	})

	log.Printf("Server started on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
