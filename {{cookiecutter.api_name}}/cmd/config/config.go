package config

import (
	"gitlab.falabella.tech/fif/integracion/forthehorde/commons/naga"
)

// configEntries retorna configuraciones iniciales (env y flags)
func configEntries() []naga.ConfigEntry {
	//TODO: Agregar entradas de configuracion adicionales requeridas para el servicio
	return []naga.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Puerto a utilizar",
			Shortcut:     "p",
			DefaultValue: ":8080",
		},
		{
			VariableName: "logging_level",
			Description:  "Level de detalle de logs",
			Shortcut:     "l",
			DefaultValue: "info",
		},
		{
			VariableName: "tracing_enabled",
			Description:  "Especifica si se debe configurar tracing",
			Shortcut:     "t",
			DefaultValue: false,
		},
		{
			VariableName: "metrics_enabled",
			Description:  "Especifica si se debe configurar metrics",
			Shortcut:     "m",
			DefaultValue: true,
		},
		{
			VariableName: "timeout",
			Description:  "timeout por defecto ",
			DefaultValue: 30,
		},
		{
			VariableName: "uri_prefix",
			Description:  "Prefijo de URL con version",
			DefaultValue: "/fifpe/v1",
		},
		{
			VariableName: "dd_tracing_enabled",
			Description:  "Especifica si se debe configurar tracing para datadog",
			Shortcut:     "d",
			DefaultValue: true,
		},
		{
			VariableName: "dd_profile_enabled",
			Description:  "Especifica si está activado el profiler de datadog",
			DefaultValue: false,
		},
	}
}

// APIConfig struct
type APIConfig struct {
	*naga.CfgBase
	DDTraceEnabled   bool
	DDProfileEnabled bool
}

// GetAPIConfig obtiene configuracion API
func GetAPIConfig() *APIConfig {
	typeResolver := naga.NewVariableTypeResolver()
	flagConfigurator := naga.NewFlagConfigurator(typeResolver)
	configurator := naga.NewConfigurator(flagConfigurator, typeResolver)
	cfg, _ := configurator.Configure("", configEntries())

	return &APIConfig{
		CfgBase:          naga.GetBaseCfg(cfg),
		DDTraceEnabled:   cfg["dd_tracing_enabled"].(bool),
		DDProfileEnabled: cfg["dd_profile_enabled"].(bool),
	}
}
