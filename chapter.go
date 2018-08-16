package pharos

type Chapter struct {
	Sort    int    `toml:"sort"`
	Numeral int    `toml:"numeral"`
	Source  string `toml:"source"`
}
