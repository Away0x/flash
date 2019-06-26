# Flash
Easy flash notifications

> go get -u github.com/Away0x/flash

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/Away0x/flash"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		flash.NewFlash("flash", r).
			Set(map[string][]string{
				"error": {"flash data"},
			}).
			Save(w)

		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusFound)
	} else {
		val := ""
		data := flash.NewFlash("flash", r).Read(w)
		if data["error"] != nil && data["error"][0] != "" {
			val = data["error"][0]
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<form method=POST action=/>
				<button>Submit</button>
				<button type=button><a href=/>Refresh</a></button>
			</form>
		`+"<hr><p>"+val+"</p>")
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
```