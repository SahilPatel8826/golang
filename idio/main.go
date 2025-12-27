package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	//////File DOWNLOAD
	/*resp, err := http.Get("https://example.com/file.jpg")
	if err != nil {
		fmt.Println("Network error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Download failed:", resp.StatusCode)
		return
	}

	file, err := os.Create("file.jpg")
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	fmt.Println("File downloaded successfully 👍")*/
	// http.HandleFunc("/foo", fooHandler)

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	// Serve files from multiple directories
	// Custom config
	app.Static("/", "./", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Index:     "go.sum",
	})

	// Serve files from "./files" directory:

	log.Fatal(app.Listen(":8080"))
}

// func fooHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// }
