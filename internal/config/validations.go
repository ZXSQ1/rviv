package config

var (
	DeviceValidations = []TypeBasedFieldValidation{
		LocalDeviceValidation{}, FtpDeviceValidation{}, SFtpDeviceValidation{},
		WebDavDeviceValidation{},
	}

	MainDevicesValidation = DeviceValidation{}
)

var (
	ProcessVerifications = []TypeBasedFieldValidation{
		CopyProcessValidation{}, MkdirProcessValidation{},
		MoveProcessValidation{}, RemoveProcessValidation{},
		SyncProcessValidation{}, OrganizeProcessValidation{},
	}

	MainProcessesValidation = ProcessValidation{}
)
