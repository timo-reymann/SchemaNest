package config

import (
	"fmt"
	"github.com/gobwas/glob"
)

type ApiKey struct {
	Identifier     string   `toml:"identifier"`
	Key            string   `toml:"key"`
	Patterns       []string `toml:"patterns"`
	parsedPatterns []glob.Glob
}

func (a *ApiKey) Validate() error {
	if a.Identifier == "" {
		return fmt.Errorf("api key identifier cannot be empty")
	}

	if a.Key == "" {
		return fmt.Errorf("api key cannot be empty")
	}

	if len(a.Patterns) == 0 {
		return fmt.Errorf("at least one pattern must be specified for api key %s", a.Identifier)
	}

	a.parsedPatterns = make([]glob.Glob, len(a.Patterns))
	for idx, pattern := range a.Patterns {
		compiled, err := glob.Compile(pattern)
		if err != nil {
			return fmt.Errorf("pattern #%d is invalid", idx)
		}

		a.parsedPatterns[idx] = compiled
	}

	return nil
}
