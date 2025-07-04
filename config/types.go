package config

import "github.com/ZXSQ1/rviv/filesystem"

// the field config
type Field string

// the config validation function
type Validation func(val any, parent Field) error

// the field validation map for config
type FieldValidationMap map[Field][]Validation

// the config type metadata
type TypeValidation interface {
	// the name of the type
	Name() string

	// the validations for fields of the type
	Validations() FieldValidationMap
}

// the standard filesystem-independent path
type Path struct {
	// the filename
	Filename string

	// the device name
	Devname string

	// the filesystem
	Fsys filesystem.Filesystem

	// is the filesystem active or not
	Active bool
}
