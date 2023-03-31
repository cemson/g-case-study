package services

import (
	"g-case-study/dto"
	"g-case-study/repo"
)

type KeyValService struct {
	keyValRepository repo.KeyValRepository
}

func CreateKeyValService() KeyValService {
	return KeyValService{
		keyValRepository: repo.CreateKeyValRepository(),
	}
}

func (t *KeyValService) Get(key string) *dto.KeyValDto {
	return t.keyValRepository.Get(key)
}

func (t *KeyValService) Set(dto *dto.KeyValDto) *dto.KeyValDto {
	t.keyValRepository.Set(dto)
	return dto
}
