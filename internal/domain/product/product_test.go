package product

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct(uuid.NewString(), "Product Name", 100.0, uuid.NewString())

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product Name", product.Name)
	assert.Equal(t, 100.0, product.Price)
}

func TestNewProductInvalidName(t *testing.T) {
	product, err := NewProduct(uuid.NewString(), "", 100.0, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(uuid.NewString(), " ", 100.0, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProductInvalidPrice(t *testing.T) {
	product, err := NewProduct(uuid.NewString(), "Product Name", -10.0, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(uuid.NewString(), "Product Name", 0.0, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProductInvalidMarkID(t *testing.T) {
	product, err := NewProduct(uuid.NewString(), "Product Name", 100.0, uuid.NewString())
	assert.Error(t, err)
	assert.Nil(t, product)

	product, err = NewProduct(uuid.NewString(), "Product Name", 100.0, uuid.NewString())
	assert.NoError(t, err)
	assert.NotNil(t, product)
}
