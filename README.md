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

var count = 0
 
func handler(w http.ResponseWriter, r *http.Request) {
	if count % 2 == 0 {
		flash.NewFlash("flash", r.Request).
			Set(map[string][]string{
				"count": count,
			}).
			Save(w)
	
		count++
	}

	
	fmt.Fprintf(w, "count: %d", flash.NewFlash("flash", r.Request).Read(w)["count"][0])
}
 
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
```