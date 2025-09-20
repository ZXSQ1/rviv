package config

import "github.com/ZXSQ1/rviv/internal/filesystem"

// the field config
type Field string

// the config validation function
type Validation func(val any, parent Field) error

// the field validation map for config
type FieldValidationMap map[Field][]Validation

// the object for validations based on the type of the operation or device;
// configuration would usually have a field like this: {"type": "kind", ...};
// this object will have validation for all of the other fields depending on the
// "type" field
type TypeBasedFieldValidation interface {
	// the name of the type (the "kind" in the given example)
	Name() string

	// the validations for fields of the type
	Validations() FieldValidationMap
}

// the standard filesystem-independent path
type Path struct {
	// the pointer to devices
	Devices []Device

	// the filename
	Filename string

	// the device name
	Devname string

	// the filesystem
	Fsys filesystem.Filesystem

	// is the filesystem active or not
	Active bool
}
