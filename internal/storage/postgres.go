package storage

import (
	"SHORTNERED_URL/internal/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres - хранилище, работающее через PostgreSQL
type Postgres struct {
	pool *pgxpool.Pool // pool соединений
}

// Создание нового объекта Postgres
// pool должен быть предварительно инициализирован
func NewPostgres(pool *pgxpool.Pool) *Postgres {
	return &Postgres{pool: pool}
}

// Close - закрывает соединение
func (s *Postgres) Close() {
	s.pool.Close()
}

// Get - поиск сокращенного URL по его идентификатору
// Return struct Shortening or error, if id not found
func (s *Postgres) Get(identifier string) (*model.Shortening, error) {
	query := `
	SELECT identifier, original_url, visits
	FROM shortenings
	WHERE identifier = $1
	`
	//Возвращаем один результат или error. Данные читаются только через Scan
	row := s.pool.QueryRow(context.Background(), query, identifier)

	var sh model.Shortening
	//Считываем результат запроса в поля структуры sh
	//If строки нет, return error: pgx.ErrNoRows
	err := row.Scan(&sh.Identifier, &sh.OriginalURL, &sh.Visits)
	if err != nil {
		return nil, fmt.Errorf("shortening not found: %w", err)
	}

	return &sh, nil
}

// Put добавляет новую запись в таблицу shortenings
// If such an identifier already exists, Insert не выполняется (ON CONLFLICT DO NOTHING)
// Возвращаем вставленную строку (RETURNING), чтобы сразу получить объект
func (s *Postgres) Put(shortening model.Shortening) (*model.Shortening, error) {
	query := `
	INSERT INTO shortenings (identifier, original_url, visits)
	VALUES($1, $2, $3)
	ON CONFLICT (identifier) DO NOTHING 
	RETURNING identifier, original_url, visits
	`

	//Выполняем запрос
	row := s.pool.QueryRow(context.Background(), query, shortening.Identifier, shortening.OriginalURL, shortening.Visits)

	var sh model.Shortening
	//Помещаем результат в структуру sh
	//pgx.ErrNoRows если строка не вставилась(например такой id существует)
	err := row.Scan(&sh.Identifier, &sh.OriginalURL, &sh.Visits)
	if err != nil {
		return nil, fmt.Errorf("couldn't insert shortening: %w", err)
	}

	return &sh, nil
}

func (s *Postgres) IncrementVisits(identifier string) error {
	query := `
	UPDATE shortenings
	SET visits = visits + 1
	WHERE identifier = $1
	`
	// Делаем запрос через Exec, который возвращает CommandTag, который в свою очередь содержит кол-во затронутых строку
	ct, err := s.pool.Exec(context.Background(), query, identifier)
	if err != nil {
		return fmt.Errorf("failed to increment visits: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return fmt.Errorf("shortening not found")
	}

	return nil
}
