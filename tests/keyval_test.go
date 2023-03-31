package tests

import (
	"g-case-study/dto"
	"g-case-study/inMemoryStore"
	"g-case-study/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	keyValRepo := repo.CreateKeyValRepository()
	keyVal := dto.KeyValDto{
		Key:   "test-key",
		Value: "test-value",
	}

	keyValRepo.Set(&keyVal)
	loadedVal, ok := inMemoryStore.InMemoryDb.Load(keyVal.Key)

	assert.True(t, ok)
	assert.Equal(t, keyVal.Value, loadedVal)
}

func TestGet(t *testing.T) {
	keyValRepo := repo.CreateKeyValRepository()
	keyVal := dto.KeyValDto{
		Key:   "test-key",
		Value: "test-value",
	}

	inMemoryStore.InMemoryDb.Store(keyVal.Key, keyVal.Value)

	result := keyValRepo.Get(keyVal.Key)
	assert.NotNil(t, result)
	assert.Equal(t, keyVal.Key, result.Key)
	assert.Equal(t, keyVal.Value, result.Value)
}
