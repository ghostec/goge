package store

type Key string

type Store struct {
	data map[Key]interface{}
}

func New() *Store {
	return &Store{data: map[Key]interface{}{}}
}

func (s Store) Get(key Key) interface{} {
	return s.data[key]
}

func (s *Store) Set(key Key, value interface{}) {
	s.data[key] = value
}

func (s *Store) Unset(key Key) {
	delete(s.data, key)
}

func (s Store) All() []interface{} {
	all := make([]interface{}, 0, len(s.data))
	for _, v := range s.data {
		all = append(all, v)
	}
	return all
}
