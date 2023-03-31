package globals

import "g-case-study/settings"

var ApplicationSettings *settings.AppSettings

func SetAppSettings(appSettings *settings.AppSettings) {
	ApplicationSettings = appSettings
}
