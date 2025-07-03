package info

func Notify(necessary bool, format string, objs ...any) {
	if necessary {
		Error(format, objs...)
	}

	Warning(format, objs...)
}
