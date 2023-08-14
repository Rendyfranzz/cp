package main

type Body struct {
	Name     string `yaml:"name"`
	Datatype string `yaml:"type"`
}

type Dictionary struct {
	Attribute string `yaml:"attribute"`
	Operator  string `yaml:"operator"`
	Label     string `yaml:"label"`
}

type Action struct {
	Attribute string `yaml:"attribute"`
	Label     string `yaml:"label"`
	Datatype  string `yaml:"type"`
}

type Rule struct {
	Value  []any `yaml:"value"`
	Action any   `yaml:"action"`
}
