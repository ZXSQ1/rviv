package config

import (
	"github.com/ZXSQ1/rviv/env"
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
					Srcs:    srcs,
					Dest:    dest,
					Method:  subprocess["method"].(string),
					Verbose: env.Verbose,
				}
			case "move":
				srcs := []Path{}
				dest, _ := NewPath(subprocess["dest"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					src, _ := NewPath(srcRaw.(string))
					srcs = append(srcs, src)
				}

				process.Options = MoveOpts{
					Srcs:    srcs,
					Dest:    dest,
					Method:  subprocess["method"].(string),
					Verbose: env.Verbose,
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
					Verbose:   env.Verbose,
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
					Verbose:   env.Verbose,
				}
			case "sync":
				src, _ := NewPath(subprocess["src"].(string))
				dest, _ := NewPath(subprocess["dest"].(string))

				process.Options = SyncOpts{
					Src:     src,
					Dest:    dest,
					Oneway:  subprocess["oneway"].(bool),
					Method:  subprocess["method"].(string),
					Verbose: env.Verbose,
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
					Archivename: archivename,
					Archivefmt: Path{
						Filename: subprocess["archive"].(string),
						Devname:  archivename.Devname,
						Devices:  archivename.Devices,
						Active:   archivename.Active,
						Fsys:     archivename.Fsys,
					},

					Entries:      entries,
					ExpiryInDays: int(subprocess["expirydays"].(float64)),
					Compression:  compression,
					Level:        int(subprocess["level"].(float64)),
					Safe:         subprocess["safe"].(bool),
					Verbose:      env.Verbose,
				}
			case "organize":
				srcs := []Path{}
				organizedir, _ := NewPath(subprocess["organizedir"].(string))

				for _, srcRaw := range subprocess["srcs"].([]any) {
					src, _ := NewPath(srcRaw.(string))
					srcs = append(srcs, src)
				}

				process.Options = OrganizeOpts{
					Srcs:        srcs,
					OrganizeDir: organizedir,
					Method:      subprocess["method"].(string),
					Datefmt:     subprocess["date"].(string),
					Verbose:     env.Verbose,
				}
			}

			processGroup.Processes = append(processGroup.Processes, process)
		}

		processes = append(processes, processGroup)
	}

	return processes, nil
}
