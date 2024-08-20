package settings

type Settings struct {
	Theme  string `json:"theme"`
	Locked bool   `json:"locked"`
	Mode   string `json:"mode"`
}

func GetTheme() string {
	settings := GetSettings()
	return settings.Theme
}

func GetSettings() Settings {
	return Settings{
		Theme:  "dark",
		Locked: false,
		Mode:   "testing",
	}
}
