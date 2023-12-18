package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"strconv"
)

var (
	pins = make(map[int]string)
	logger = log.New(os.Stdout, "[doorbell] ", log.LstdFlags|log.Lmsgprefix)
)

func main() {
	loadAuthorizedPINs()

	http.HandleFunc("/", handler)
	logger.Fatal(http.ListenAndServe(":80", nil))
}

func loadAuthorizedPINs() {
	b, err := os.ReadFile("/etc/doorbell.passwd")
	if err != nil {
		logger.Fatal(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) > 0 {
			parts := strings.Split(line, ":")
			pin, err := strconv.Atoi(parts[0])
			if err != nil || len(parts) != 2 {
				logger.Printf("[ERROR] Unable to load pin (%s) %v", parts[0], err)
			} else {
				pins[pin] = parts[1]
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	defer w.Write([]byte(htmlIndex()))

	if r.Method == "POST" {
		var pinStr string

		for _, val := range r.Form["pin"] {
			pinStr = fmt.Sprintf("%s%s", pinStr, val)
		}

		pin, err := strconv.Atoi(pinStr)
		if err != nil {
			logger.Printf("[ERROR] %v", err)
			return
		}

		if name, ok := pins[pin]; ok {
			logger.Printf("Doorbell pressed by %s", name)
		} else {
			logger.Printf("[ERROR] Doorbell pressed by unknown PIN %d", pin)
		}
	}
}
