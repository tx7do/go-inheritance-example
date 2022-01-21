package fruit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApple(t *testing.T) {
	{
		var fruitApple Fruit
		fruitApple = NewApple("big")
		assert.NotNil(t, fruitApple)

		assert.Equal(t, fruitApple.GetType(), AppleType)
		assert.Equal(t, fruitApple.GetName(), "big")

		fruitApple.SetName("middle")
		assert.Equal(t, fruitApple.GetName(), "middle")
	}

	{
		apple := NewApple("little")
		assert.NotNil(t, apple)

		assert.Equal(t, apple.GetType(), AppleType)
		assert.Equal(t, apple.GetName(), "little")
	}
}

func TestBanana(t *testing.T) {
	{
		var fruitBanana Fruit
		fruitBanana = NewBanana("big")
		assert.NotNil(t, fruitBanana)

		assert.Equal(t, fruitBanana.GetType(), BananaType)
		assert.Equal(t, fruitBanana.GetName(), "big")

		fruitBanana.SetName("middle")
		assert.Equal(t, fruitBanana.GetName(), "middle")

		// fruitBanana.SetEnergy(2)
	}

	{
		banana := NewBanana("little")
		assert.NotNil(t, banana)

		assert.Equal(t, banana.GetType(), BananaType)
		assert.Equal(t, banana.GetName(), "little")

		banana.SetEnergy(2)
		assert.Equal(t, banana.GetEnergy(), float32(2))
	}
}
