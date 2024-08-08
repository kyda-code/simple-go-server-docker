package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello from Dockerized Go App!")
		if err != nil {
			return
		}
	})

	fmt.Println("Server is listening on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
