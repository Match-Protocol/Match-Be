package config

type Toml struct {
	Mysql struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Db       string `toml:"db"`
	}

	Redis struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Password string `toml:"password"`
		Db       int    `toml:"db"`
	}
	// 这个rpc干春啊?
	Rpc struct {
		Url string `toml:"url"`
		P   string `toml:"p"`
		Sol string `toml:"sol"`
	}
	Sms struct {
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
	}
}
