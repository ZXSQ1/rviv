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
				srcs := []Path{}
				dest, _ := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					src, _ := NewPath(srcRaw.(string))
					srcs = append(srcs, src)
				}

				process.Options = CopyOpts{
					Srcs:      srcs,
					Dest:      dest,
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "move":
				srcs := []Path{}
				dest, _ := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					src, _ := NewPath(srcRaw.(string))
					srcs = append(srcs, src)
				}

				process.Options = MoveOpts{
					Srcs:      srcs,
					Dest:      dest,
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "mkdir":
				files := []Path{}

				for _, filename := range subprocess["paths"].([]any) {
					fileObj, _ := NewPath(filename.(string))
					files = append(files, fileObj)
				}

				process.Options = MkdirOpts{
					Filenames: files,
					Parent:    subprocess["parent"].(bool),
					Necessary: subprocess["necessary"].(bool),
				}
			case "remove":
				files := []Path{}

				for _, filename := range subprocess["paths"].([]any) {
					fileObj, _ := NewPath(filename.(string))
					files = append(files, fileObj)
				}

				process.Options = RemoveOpts{
					Filenames: files,
					Recursive: subprocess["recursive"].(bool),
					Necessary: subprocess["necessary"].(bool),
				}
			case "sync":
				src, _ := NewPath(subprocess["src"].(string))
				dest, _ := NewPath(subprocess["dest"].(string))

				process.Options = SyncOpts{
					Src:       src,
					Dest:      dest,
					Oneway:    subprocess["oneway"].(bool),
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "archive":
				archivename, _ := NewPath(subprocess["archive"].(string))
				entries := []Path{}
				compression := subprocess["compression"].(string)

				for _, entry := range subprocess["entries"].([]any) {
					entry, _ := NewPath(entry.(string))
					entries = append(entries, entry)
				}

				process.Options = ArchiveOpts{
					Archivename:  archivename,
					Entries:      entries,
					ExpiryInDays: subprocess["expirydays"].(float64),
					Compression:  compression,
					Safe:         subprocess["safe"].(bool),
					Necessary:    subprocess["necessary"].(bool),
				}
			}

			processGroup.Processes = append(processGroup.Processes, process)
		}

		processes = append(processes, processGroup)
	}

	return processes, nil
}
