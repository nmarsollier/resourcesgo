package env

import (
	"cmp"
	"os"

	"github.com/nmarsollier/resourcesgo/internal/tools/strs"
)

// Configuration properties
type Configuration struct {
	ServerName  string `json:"serverName"`
	Port        int    `json:"port"`
	GqlPort     int    `json:"gqlPort"`
	PostgresURL string `json:"postgresUrl"`
}

var config *Configuration

// Get returns the singleton instance of the Configuration.
// If the Configuration is not already loaded, it loads the configuration
// by calling the load function and stores it in the config variable.
var Get = func() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

func load() *Configuration {
	return &Configuration{
		ServerName:  cmp.Or(os.Getenv("SERVER_NAME"), "resourcesgo"),
		Port:        cmp.Or(strs.AtoiZero(os.Getenv("PORT")), 3000),
		GqlPort:     cmp.Or(strs.AtoiZero(os.Getenv("GQL_PORT")), 4000),
		PostgresURL: cmp.Or(os.Getenv("POSTGRES_URL"), "postgres://postgres@localhost:5432/postgres"),
	}
}
