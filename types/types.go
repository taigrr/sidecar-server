package types

type Config struct {
	Matches []URLMatch `yaml:"matches"`
}

type URLMatch struct {
	URLPattern  string `yaml:"url-pattern"`
	ShouldClose bool   `yaml:"should-close"`
	Action      string `yaml:"action"`
}
