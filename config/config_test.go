package config

var testConfig = `
{
	"devices": {
		"drive": {
      	"type": "ftp",
      	"ip": "127.0.0.1",
      	"port": 312,
      	"user": "$USER",
      	"pass": "$PASS"
    	},

    	"drive1": {
      	"type": "ssh",
      	"ip": "lan",
      	"port": 312,
      	"user": "$USER",
      	"pass": "$PASS"
    	},

    	"drive2": {
      	"type": "webdav",
      	"ip": "127.0.0.1",
      	"port": 312,
      	"user": "$USER",
      	"pass": "$PASS"
    	},
        
    	"dev": {
      	"type": "local",
      	"prefix": "$HOME"
    	}
  	},

  "processes": {
    	"rsync": {
      	"aliases": ["rs132_3"],
      	"subprocesses": [
            {
               "type": "copy",
               "srcs": [
                  "dev::$HOME/Downloads",
                  "drive2::/Download"
               ],

               "dest": "dev::$HOME/.config",
               "method": "dd",
               "necessary": true
            }
      	]
    	}
  	}
}
`
