package config

type HelmConfig struct {
	HelmCommand string	`yaml:"command"`
	HelmVersion int		`yaml:"version"`
}
