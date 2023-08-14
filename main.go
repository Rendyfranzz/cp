package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goccy/go-yaml"
)

type RuleEngine struct {
	Bodies       []Body       `yaml:"body"`
	Dictionaries []Dictionary `yaml:"dictionary"`
	Action       Action       `yaml:"action"`
	Rules        []Rule       `yaml:"rule"`
}

var engine RuleEngine

func main() {
	content := readYAML("rule.yaml")

	err := yaml.Unmarshal(content, &engine)
	if err != nil {
		log.Fatalf("cannot unmarshal: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", endpoint)

	s := &http.Server{
		Addr:    ":7000",
		Handler: mux,
	}

	s.ListenAndServe()
}

func readYAML(filename string) []byte {
	content, err := os.ReadFile("rule.yaml")
	if err != nil {
		log.Fatalf("cant read file: %v", err)
	}

	return content
}
