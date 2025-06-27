
In this project, I tried to build a backup system to backup to remote storage
systems that is easily and extensively configurable.

- primary programming language: go
- status: unfinished

---- PRE-REVIEW ----

Before getting into the code, a quick rundown:

The `filesystem` package contains the interface for all file operations; try to
check that out first to understand other packages like `ftpfs`, `localfs` and
`sftpfs`.

The `processes` package contains code for the filesystem-independent processes
that would be done in order in the configuration. The `processes` package has
user-displayed errors, so it uses the `info` package to report warnings and
errors. The `compressor` package contains code for compression algorithms for
single files.

The `config` package regards dealing with the configuration and turning it into
structures that are used throughout the program for running the processes and as
a configuration of the devices.

*Example Config*
```
{
   "devices": {
      "dev1": {
         "type": "ftp",
         "ip": "lan",
         "port": 2121,
         "pass": "$(ENVIRONMENT_VAR_PASS)", //environment variable substitution
         "user": "$(USERNAME)"
      },

      "dev2": {
         "type": "local",
         "path": "$(HOME)",
      }
   },

   "processes": {
      "group1": {
         "name": "process1",
         "aliases": ["p1", "proc1"],
         "subprocesses": [
            {
               "type": "copy",
               "necessary": true,
               "src": "srcpath",
               "dest": "destpath"
               // other command type-specific options
            },

            {
               "type": "sync",
               "necessary": false,
               "src": "srcpath",
               "dest": [
                  "${dev2}/dest1", "${dev1}/dest2" //specifying a path for a device
               ],
            }  
         ]
      },

      "group2": {
         "name": "process2",
         "aliases": ["p2", "proc2"],
         "subprocesses": [
            "#group1",
            {
               // other processes
            }
         ]
      }
   }
}
```

4. The `logging` package contains functions for reporting errors throughout
   functions. The convention used is to use the `ReportErr` function after
   every statement that contains a `, err`.


---- FEATURES ----
1. The system must provide multiple protocols for transfer (FTP, HTTP WebDAV,
   SFTP, Through Mounting).
2. The system must provide multiple formats for the archive (TAR (GZ, XZ, LZ4,
   BZ2), ZIP).
3. The system must provide configuration for the remote storage systems.
4. The system must provide configuration for the archives and file organization
   operations.
5. The system must provide a recommendation system for the selection of IPs (the
   connection to remote storage servers must have an IP; IPs aren't constant,
   and so the user will provide, in the configuration, a set of IPs that will be
   selected from).
6. The system must provide many options for operation.
   - `-c`: coloring and colored output.
   - `<config>`: the configuration file to base your options on.
      - the configuration file set through environment variables.
      - there must be a default configuration directory
7. The system must provide customizable options & operations through config.
