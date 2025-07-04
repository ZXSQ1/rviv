package config

import (
	"github.com/spf13/viper"
)

var (
	Processes = []ProcessGroup{}
)

func LoadProcesses() error {
	if len(Processes) > 0 {
		return nil
	}

	for field, validations := range MainProcessesValidation.Validations() {
		for _, validation := range validations {
			if err := validation(viper.Get(string(field)), ""); err != nil {
				return err
			}
		}
	}

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
				dest := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					srcs = append(srcs, NewPath(srcRaw.(string)))
				}

				process.Options = CopyOpts{
					Srcs:      srcs,
					Dest:      dest,
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "move":
				srcs := []Path{}
				dest := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					srcs = append(srcs, NewPath(srcRaw.(string)))
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
					files = append(files, NewPath(filename.(string)))
				}

				process.Options = MkdirOpts{
					Filenames: files,
					Parent:    subprocess["parent"].(bool),
					Necessary: subprocess["necessary"].(bool),
				}
			case "remove":
				files := []Path{}

				for _, filename := range subprocess["paths"].([]any) {
					files = append(files, NewPath(filename.(string)))
				}

				process.Options = RemoveOpts{
					Filenames: files,
					Recursive: subprocess["recursive"].(bool),
					Necessary: subprocess["necessary"].(bool),
				}
			case "sync":
				src := NewPath(subprocess["src"].(string))
				dest := NewPath(subprocess["dest"].(string))

				process.Options = SyncOpts{
					Src:       src,
					Dest:      dest,
					Oneway:    subprocess["oneway"].(bool),
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "zip":
				archivename := NewPath(subprocess["archive"].(string))
				entries := []Path{}

				for _, entry := range subprocess["entries"].([]any) {
					entries = append(entries, NewPath(entry.(string)))
				}

				process.Options = ZipOpts{
					Archivename: archivename,
					Entries:     entries,
					Safe:        subprocess["safe"].(bool),
					Necessary:   subprocess["necessary"].(bool),
				}
			case "tar":
				archivename := NewPath(subprocess["archive"].(string))
				entries := []Path{}
				compression := subprocess["compression"].(string)

				for _, entry := range subprocess["entries"].([]any) {
					entries = append(entries, NewPath(entry.(string)))
				}

				process.Options = TarOpts{
					Archivename: archivename,
					Entries:     entries,
					Compression: compression,
					Safe:        subprocess["safe"].(bool),
					Necessary:   subprocess["necessary"].(bool),
				}
			}

			processGroup.Processes = append(processGroup.Processes, process)
		}

		Processes = append(Processes, processGroup)
	}

	return nil
}
