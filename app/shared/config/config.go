package config

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}
