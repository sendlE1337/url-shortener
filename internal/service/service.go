package shortener

import (
	"SHORTNERED_URL/internal/model"
	"math/rand"
)

// Storage интерфейс для работы с хранилищем ссылок.
// Такой подход помог нам не быть зависимым от какого-то определенного типа хранилища
// result - масштабируемость
type Storage interface {
	Get(identifier string) (*model.Shortening, error)
	Put(shortening model.Shortening) (*model.Shortening, error)
	IncrementVisits(identifier string) error
}

// Service содержит бизнес логику URL shortener
type Service struct {
	storage Storage
}

// NewService - конструктор
func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

// Shorten создает короткую ссылку для заданного URL
func (s *Service) Shorten(original_url string) (*model.Shortening, error) {
	var (
		id         = rand.Uint32() // рандомный ID
		identifier = Shortener(id) //преобразование в короткий идентификатор
	)

	shortening := model.Shortening{
		Identifier:  identifier,
		OriginalURL: original_url,
		Visits:      0,
	}

	_, err := s.storage.Put(shortening)
	if err != nil {
		return nil, err
	}
	return &shortening, nil
}

// Get возвращает объект Shortening по идентификатору
func (s *Service) Get(identifier string) (*model.Shortening, error) {
	return s.storage.Get(identifier)
}

// Redirect возвращает оригинальный url по короткому идентификатору
// и увеличивает счетчик Visits
func (s *Service) Redirect(identifier string) (string, error) {
	short, err := s.storage.Get(identifier)
	if err != nil {
		return "", err
	}

	_ = s.storage.IncrementVisits(identifier)
	return short.OriginalURL, nil
}
