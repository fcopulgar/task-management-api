package security

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBcryptHasherHashAndCompare(t *testing.T) {
	hasher := NewBcryptHasher()
	ctx := context.Background()

	hash, err := hasher.Hash(ctx, "my-secret-password")
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, "my-secret-password", hash)

	err = hasher.Compare(ctx, hash, "my-secret-password")
	assert.NoError(t, err)

	err = hasher.Compare(ctx, hash, "wrong-password")
	assert.Error(t, err)
}

func TestBcryptHasherHashProducesUniqueOutput(t *testing.T) {
	hasher := NewBcryptHasher()
	ctx := context.Background()

	hash1, _ := hasher.Hash(ctx, "password")
	hash2, _ := hasher.Hash(ctx, "password")
	assert.NotEqual(t, hash1, hash2)
}

func TestBcryptHasherCompareEmpty(t *testing.T) {
	hasher := NewBcryptHasher()
	ctx := context.Background()

	err := hasher.Compare(ctx, "", "anything")
	assert.Error(t, err)
}
