package entitys

func GetDefaultLog(level Level) (log Log) {
	return Log{
		Logger: LOG_LOGGER,
		Level:  level.String(),
	}
}
