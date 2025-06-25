package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct(1, "Product Name", 100.0, 1)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, uint64(1), product.ID)
	assert.Equal(t, "Product Name", product.Name)
	assert.Equal(t, 100.0, product.Price)
	assert.Equal(t, uint64(1), product.MarkID)
}

func TestNewProductInvalidName(t *testing.T) {
	product, err := NewProduct(1, "", 100.0, 1)
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(1, " ", 100.0, 1)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProductInvalidPrice(t *testing.T) {
	product, err := NewProduct(1, "Product Name", -10.0, 1)
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(1, "Product Name", 0.0, 1)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProductInvalidMarkID(t *testing.T) {
	product, err := NewProduct(1, "Product Name", 100.0, 0)
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(1, "Product Name", 100.0, 2)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, uint64(2), product.MarkID)
}
