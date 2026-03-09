package examples

import "github.com/andrey/gostructtohashmapgenerator/models"

type Account struct {
	ID      int         `structtomap:"id"`
	Owner   models.User `structtomap:"owner"`
	Balance float64     `structtomap:"balance"`
}
