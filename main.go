package main

import (
	"vibe/domain"
	"vibe/workers"
)

const url = "google.com"

func main() {
	workers := workers.Workers{Count: 10}

	workers.StartupWorkers(domain.LookupA, "hello.txt")
}
