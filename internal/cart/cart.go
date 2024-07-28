package cart

import (
	"errors"
	"os"
)

type Cart interface{}

type GBCart struct{}

type GBACart struct{}

func LoadCart(cartPath string) (Cart, error) {
	if _, err := os.Stat(cartPath); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return GBCart{}, nil
}
