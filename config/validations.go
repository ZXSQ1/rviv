package config

var (
	DeviceValidations []TypeBasedFieldValidation = []TypeBasedFieldValidation{
		LocalDeviceValidation{}, FtpDeviceValidation{}, SFtpDeviceValidation{},
		WebDavDeviceValidation{},
	}

	MainDevicesValidation TypeBasedFieldValidation = DeviceValidation{}
)

var (
	ProcessVerifications []TypeBasedFieldValidation = []TypeBasedFieldValidation{
		CopyProcessValidation{}, MkdirProcessValidation{},
		MoveProcessValidation{}, RemoveProcessValidation{},
		SyncProcessValidation{}, ArchiveProcessValidation{},
		OrganizeProcessValidation{},
	}

	MainProcessesValidation TypeBasedFieldValidation = ProcessValidation{}
)
