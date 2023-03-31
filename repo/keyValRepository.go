package repo

import (
	"g-case-study/dto"
	"g-case-study/inMemoryStore"
)

type KeyValRepository struct {
}

func CreateKeyValRepository() KeyValRepository {
	return KeyValRepository{}
}

func (t *KeyValRepository) Set(requestDto *dto.KeyValDto) {
	inMemoryStore.InMemoryDb.Store(requestDto.Key, requestDto.Value)
}

func (t *KeyValRepository) Get(key string) *dto.KeyValDto {
	val, ok := inMemoryStore.InMemoryDb.Load(key)
	if !ok {
		return nil
	}

	result := dto.KeyValDto{
		Key:   key,
		Value: val.(string),
	}
	return &result
}
