package main

import (
	"fmt"
	"math"
	"net/http"
)

func greeting(text string) string {
	return "<b>" + text + "</b>"
}

func sqrtDelayLoop(x float64) {
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	sqrtDelayLoop(0.0001)
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}
func main() {
	fmt.Printf("Starting server at port 8000\n")
	http.HandleFunc("/", handleGreeting)
	http.ListenAndServe(":8000", nil)

}