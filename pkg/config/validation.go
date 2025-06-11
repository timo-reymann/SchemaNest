package config

type Validatable interface {
	Validate() error
}
