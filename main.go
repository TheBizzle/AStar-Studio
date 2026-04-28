package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	runner "github.com/TheBizzle/AStar-Studio/internal/runner"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	http.ServeFile(res, req, "index.html")
}

func pathfindingHandler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method == http.MethodPost {
		const oneMB = 1 << 20
		body, err := io.ReadAll(http.MaxBytesReader(res, req.Body, oneMB))
		if err != nil {
			http.Error(res, "Error reading body", http.StatusInternalServerError)
			return
		}
		wasSuccessful, pmapStr, timingStrs := runner.RunAStars(string(body))
		if wasSuccessful {
			fmt.Fprintf(res, "%d,%v,%v", 0, strings.Join(timingStrs, "&"), pmapStr)
		} else {
			fmt.Fprintf(res, "%d,%v", 1, pmapStr)
		}
	} else {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/find-me-a-path", pathfindingHandler)

	portNum := 8080
	fmt.Printf("Pathfinding server running on port %d\n", portNum)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
