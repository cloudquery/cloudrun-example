package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("received a request")

	cloudQueryAPIKey := os.Getenv("CLOUDQUERY_API_KEY")
	// Command to execute
	cmd := exec.Command("/app/cloudquery", "sync", "/secrets/config.yaml", "--log-console")
	cmd.Env = []string{fmt.Sprintf("CLOUDQUERY_API_KEY=%s", cloudQueryAPIKey)}

	// Create pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Create goroutines to stream stdout and stderr
	go streamOutput(stdoutPipe, "stdout")
	go streamOutput(stderrPipe, "stderr")

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("Command execution completed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Command execution completed"))
}

// streamOutput reads from the pipe and prints the output
func streamOutput(pipe io.Reader, name string) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", name, scanner.Text())
	}
}

func main() {
	if os.Getenv("CLOUDQUERY_API_KEY") == "" {
		log.Fatal("CLOUDQUERY_API_KEY is required")
	}
	log.Print("starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
