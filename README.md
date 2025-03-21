
In this project, I tried to build a backup system to backup
to remote storage systems that is easily and extensively
configurable.

- primary programming language: go
- status: unfinished

---- FEATURES ----
1. The system must provide multiple protocols for transfer 
   (FTP, HTTP WebDAV, SFTP, Through Mounting).
2. The system must provide multiple formats for the archive
   (TAR (GZ, XZ, LZ4, BZ2), ZIP).
3. The system must provide configuration for the remote storage systems.
4. The system must provide configuration for the archives and file
   organization operations.
5. The system must provide a recommendation system for the selection of IPs
   (the connection to remote storage servers must have an IP; IPs aren't
   constant, and so the user will provide, in the configuration, a set of IPs
   that will be selected from).
6. The system must provide many options for operation.
    - `-j` option for setting the number of jobs.
    - `-c` option to set the path to the configuration directory.
    - `-a <archive>` for archiving archives.
        - Archives all archives if no archive was provided.
        - Archives the given `archive` if it was provided.
    - `-ba` resyncs the archives.
    - `-bf` resyncs the files after some operations.
    - `-p` for progress indicators.
    - `-v` for whether to show output or not.
    - `-d` for specifying the server or device to backup to.
        - Name is specified in the configuration.
