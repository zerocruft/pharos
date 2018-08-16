package pharos

type Manifest struct {
	Title    string     `toml:"title"`
	Author   string     `toml:"author"`
	Template string     `toml:"template"`
	Chapters []Chapter  `toml:"chapter"`
	Contents []Contents `toml:"contents"`
}
