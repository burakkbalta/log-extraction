package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.CommandLine.String("path", "app.log", "The path to log file that should be analyzed.")
	level := flag.CommandLine.String("level", "ERROR", "Log level to serach for. Options are DEBUG, INFO, ERROR and CRITICAL.")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var analyzedLogs []string

	reader := bufio.NewReader(f)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			analyzedLogs = append(analyzedLogs, s)
		}
	}

	fmt.Println(analyzedLogs)
}
