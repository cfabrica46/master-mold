package config

import (
	"time"

	apiconfig "github.com/cfabrica46/api-config"
)

func configEntries() []apiconfig.ConfigEntry {
	return []apiconfig.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Puerto a utilizar",
			Shortcut:     "p",
			DefaultValue: ":8080",
		},
		{
			VariableName: "timeout",
			Description:  "timeout por defecto ",
			DefaultValue: 30,
		},
		{
			VariableName: "uri_prefix",
			Description:  "Prefijo de URL con version",
			DefaultValue: "/api/v1",
		},
	}
}

type APIConfig struct {
	*apiconfig.CfgBase
}

func GetAPIConfig() (*APIConfig, error) {
	typeResolver := apiconfig.NewVariableTypeResolver()
	flagConfigurator := apiconfig.NewFlagConfigurator(typeResolver)
	configurator := apiconfig.NewConfigurator(flagConfigurator, typeResolver)

	cfg, err := configurator.Configure(configEntries())
	if err != nil {
		return nil, err
	}

	return &APIConfig{
		CfgBase: &apiconfig.CfgBase{
			Port:      cfg["port"].(string),
			Timeout:   time.Duration(cfg["timeout"].(int)) * time.Second,
			URIPrefix: cfg["uri_prefix"].(string),
		},
	}, nil
}
