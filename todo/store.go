package todo

type Store interface {
	Load() ([]Item, error)
	Save([]Item) error
}
