package config

import (
	"github.com/spf13/viper"
)

var (
	Processes = []ProcessGroup{}
)

func LoadProcesses() {
	if len(Processes) > 0 {
		return
	}

	for field, validations := range MainProcessesValidation.Validations() {
		for _, validation := range validations {
			validation(viper.Get(string(field)), "")
		}
	}

	processesRaw := viper.Get("processes").(map[string]map[string]any)

	for groupname, groupinfo := range processesRaw {
		processGroup := ProcessGroup{}
		processGroup.Groupname = groupname
		processGroup.Aliases = groupinfo["aliases"].([]string)

		prefix := "processes." + groupname
		subprocesses := groupinfo["subprocesses"].([]map[string]any)

		for _, subprocess := range subprocesses {
			process := Process{}
			prefix = prefix + ".subprocesses"

			process.Kind = subprocess["type"].(string)

			switch subprocess["type"].(string) {
			case "copy":
				srcs := []Path{}
				dest := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]string) {
					srcs = append(srcs, NewPath(srcRaw))
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

				for _, srcRaw := range subprocess["srcs"].([]string) {
					srcs = append(srcs, NewPath(srcRaw))
				}

				process.Options = MoveOpts{
					Srcs:      srcs,
					Dest:      dest,
					Method:    subprocess["method"].(string),
					Necessary: subprocess["necessary"].(bool),
				}
			case "mkdir":
				files := []Path{}

				for _, filename := range subprocess["paths"].([]string) {
					files = append(files, NewPath(filename))
				}

				process.Options = MkdirOpts{
					Filenames: files,
					Parent:    subprocess["parent"].(bool),
					Necessary: subprocess["necessary"].(bool),
				}
			case "remove":
				files := []Path{}

				for _, filename := range subprocess["paths"].([]string) {
					files = append(files, NewPath(filename))
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

				process.Options = ZipOpts{
					Archivename: archivename,
					Safe:        subprocess["safe"].(bool),
					Necessary:   subprocess["necessary"].(bool),
				}
			case "tar":
				archivename := NewPath(subprocess["archive"].(string))

				process.Options = TarOpts{
					Archivename: archivename,
					Safe:        subprocess["safe"].(bool),
					Necessary:   subprocess["necessary"].(bool),
				}
			}
		}

		Processes = append(Processes, processGroup)
	}
}
