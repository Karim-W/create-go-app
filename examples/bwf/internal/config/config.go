package config

import (
	"github.com/karim-w/cafe"
)

// config is the app config scoped to this package.
// It is initialized by InitOrDie.
// refer to https://github.com/karim-w/cafe for more details on how to use it.
var config *cafe.Cafe

// Cafe is my personal package for config management and can be swapped out
// for any other config management package since the getters are the only
// thing that is used in the code.

// InitRestConfigOrDie initializes the config and panics if it fails.
func InitRestConfigOrDie() {
	var err error
	config, err = cafe.New(cafe.Schema{
		"SERVER_PORT": cafe.String("PORT").Default("8080"),
		// TODO: Add your config here
	})
	if err != nil {
		panic(err)
	}
}

// GetServerPort returns the port the server will listen on
// It is used to initialize the server adapter
func GetServerPort() (string, error) {
	port, err := config.GetString("SERVER_PORT")
	if err != nil {
		return "", err
	}
	return ":" + port, nil
}

// TODO: Add your config getters here
// func GetFoo() (string,error) {
// 	return Config.GetString("foo")
// }
