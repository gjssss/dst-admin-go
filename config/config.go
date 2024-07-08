package config

type Config struct {
	Port           string `yaml:"port"`
	Path           string `yaml:"path"`
	Db             string `yaml:"database"`
	Steamcmd       string `yaml:"steamcmd"`
	SteamAPIKey    string `yaml:"steamAPIKey"`
	OPENAI_API_KEY string `yaml:"OPENAI_API_KEY"`
	Prompt         string `yaml:"prompt"`
	Flag           string `yaml:"flag"`
	WanIP          string `yaml:"wanip"`
	WhiteAdminIP   string `yaml:"whiteadminip"`
	Token          string `yaml:"token"`

	AutoUpdateModinfo struct {
		Enable              bool `yaml:"enable"`
		CheckInterval       int  `yaml:"checkInterval"`
		UpdateCheckInterval int  `yaml:"updateCheckInterval"`
	} `yaml:"autoUpdateModinfo"`

	DstCliPort string `yaml:"dstCliPort"`

	Master struct {
		Enable   bool   `yaml:"enable"`
		Pattern  string `yaml:"pattern"`
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Name     string `json:"name"`
	} `yaml:"master"`
}
