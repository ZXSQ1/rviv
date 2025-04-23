
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
   - `-c`: coloring and colored output.
   - `<config>`: the configuration file to base your options on.
      - the configuration file set through environment variables.
      - there must be a default configuration directory
7. The system must provide customizable options & operations through config.
