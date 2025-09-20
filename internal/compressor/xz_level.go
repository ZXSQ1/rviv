package compressor

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ulikunitz/xz"
	"github.com/ulikunitz/xz/lzma"
)

func (compressor *XzCompressor) Level(level CompressionLevel) xz.WriterConfig {
	if level < 0 {
		level = 0
	} else if level > 9 {
		level = 9
	}

	properties := &lzma.Properties{LC: 3, LP: 0, PB: 2}
	dictcap := 1 << (16 + level)
	matcher := lzma.HashTable4

	if level > 4 {
		matcher = lzma.BinaryTree
	}

	return xz.WriterConfig{
		Properties: properties,
		DictCap:    dictcap,
		BufSize:    filesystem.BufferSize,
		CheckSum:   0x01,
		NoCheckSum: false,
		Matcher:    matcher,
	}
}
