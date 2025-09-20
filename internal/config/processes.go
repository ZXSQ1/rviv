package config

import (
	"github.com/spf13/viper"
)

func LoadProcesses() ([]ProcessGroup, error) {
	for field, validations := range MainProcessesValidation.Validations() {
		val := viper.Get(string(field))

		for _, validation := range validations {
			if err := validation(val, ""); err != nil {
				return nil, err
			}
		}
	}

	processes := []ProcessGroup{}
	processesRaw := viper.Get("processes").(map[string]any)

	for groupname, groupinfo := range processesRaw {
		groupinfo := groupinfo.(map[string]any)
		aliases := []string{}

		for _, alias := range groupinfo["aliases"].([]any) {
			alias := alias.(string)
			aliases = append(aliases, alias)
		}

		processGroup := ProcessGroup{}
		processGroup.Groupname = groupname
		processGroup.Aliases = aliases

		subprocesses := groupinfo["subprocesses"].([]any)

		for _, subprocess := range subprocesses {
			subprocess := subprocess.(map[string]any)
			process := Process{}
			process.Kind = subprocess["type"].(string)

			switch process.Kind {
			case "copy":
				process.Options = LoadCopyOpts(subprocess)
			case "move":
				process.Options = LoadMoveOpts(subprocess)
			case "mkdir":
				process.Options = LoadMkdirOpts(subprocess)
			case "remove":
				process.Options = LoadRemoveOpts(subprocess)
			case "sync":
				process.Options = LoadSyncOpts(subprocess)
			case "organize":
				process.Options = LoadOrganizeOpts(subprocess)
			}

			processGroup.Processes = append(processGroup.Processes, process)
		}

		processes = append(processes, processGroup)
	}

	return processes, nil
}
