package utils

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/wouter173/mailer/structs"
)

//Read the keys from keys.toml.
func Readkeys() structs.Config {
	data, err := os.ReadFile("keys.toml")
	if err != nil {
		panic(err)
	}

	var keys structs.Config
	err = toml.Unmarshal(data, &keys)

	if err != nil {
		panic(err)
	}

	return keys
}
