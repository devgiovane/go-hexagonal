package domain_test

import (
	"github.com/stretchr/testify/require"
	"hexagonal/domain"
	"testing"
)

func TestProduct_Validate(t *testing.T) {
	product := domain.NewProduct("Beer", 3.99)
	isValid, err := product.Validate()
	require.True(t, isValid)
	require.Nil(t, err)
}
