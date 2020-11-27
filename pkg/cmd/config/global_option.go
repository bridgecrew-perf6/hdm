package config

import "log"

var globalOption *GlobalOptions

type GlobalOptions struct {
	DryRun bool
	Config *Config
}

func newGlobalOption(dryrun bool, config *Config) {
	if globalOption != nil{
		log.Println("global option is already set.")
		return
	}

	globalOption = &GlobalOptions{
		DryRun: dryrun,
		Config: config,
	}
}

func GetGlobalOption() *GlobalOptions {
	if globalOption == nil {
		return nil
	}
	return globalOption
}