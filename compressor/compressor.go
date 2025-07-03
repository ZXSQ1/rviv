package compressor

import "github.com/ZXSQ1/rviv/config"

// the standard compressor type
type Compressor interface {
	// compresses the file using the levels specied into the outfile
	Compress(outfile config.Path, level CompressionLevel) error
}

// the standard compression level (0 - 9)
type CompressionLevel int
