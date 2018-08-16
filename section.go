package pharos

type Section interface {
	SortID() int
	Type() int
	Serialize() string
	Deserialize(string)
}
