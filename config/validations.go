package config

var (
	DeviceValidations []TypeValidation = []TypeValidation{
		LocalDeviceValidation{}, FtpDeviceValidation{}, SFtpDeviceValidation{},
		WebDavDeviceValidation{},
	}

	MainDevicesValidation TypeValidation = DeviceValidation{}
)

var (
	ProcessVerifications []TypeValidation = []TypeValidation{
		CopyProcessValidation{}, MkdirProcessValidation{},
		MoveProcessValidation{}, RemoveProcessValidation{},
		SyncProcessValidation{}, ArchiveProcessValidation{},
	}

	MainProcessesValidation TypeValidation = ProcessValidation{}
)
