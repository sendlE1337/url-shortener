//In-memory используем по большой	 части не как настоязее хранилище, а для того чтобы упростить тестирование.

package storage

import (
	"SHORTNERED_URL/internal/model"
	"sync"
)

// inMemory - потокобезопасное хранилище ссылок
type inMemory struct {
	m sync.Map
}

// NewInMemory создает пустое хранилище
func NewInMemory() *inMemory {
	return &inMemory{}
}

// Get возвращает Shortening по идентификатору. При возникновении ошибки - возвращаем ErrNotFound.
func (s *inMemory) Get(identifier string) (*model.Shortening, error) {
	value, ok := s.m.Load(identifier)
	if !ok {
		return nil, model.ErrNotFound
	}
	shortening := value.(model.Shortening)
	return &shortening, nil
}

// Put добавляет новую короткую ссылку
// Возвращает ошибку ErrAlreadyExists, если такой идентификатор есть
func (s *inMemory) Put(shortening model.Shortening) (*model.Shortening, error) {
	if _, exists := s.m.Load(shortening.Identifier); exists {
		return nil, model.ErrAlreadyExists
	}

	s.m.Store(shortening.Identifier, shortening)
	return &shortening, nil
}

// IncreamentVisits увеличивает счетчик посещений по идентификатору
func (s *inMemory) IncrementVisits(identifier string) error {
	v, ok := s.m.Load(identifier)
	if !ok {
		return model.ErrNotFound
	}

	shortening := v.(model.Shortening)
	shortening.Visits++

	s.m.Store(identifier, shortening)
	return nil
}
