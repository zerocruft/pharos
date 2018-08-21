package pharos

type Chapter struct {
	Sort       int    `toml:"sort"`
	OrderTitle string `toml:"ordertitle"`
	Title      string `toml:"title"`
	Subtitle   string `toml:"subtitle"`
	Target     string `toml:"target"`
}
