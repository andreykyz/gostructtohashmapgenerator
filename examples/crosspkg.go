package examples

import "github.com/andreykyz/gostructtohashmapgenerator/examples/models"

type Account struct {
	ID      int         `structtomap:"id"`
	Owner   models.User `structtomap:"owner"`
	Balance float64     `structtomap:"balance"`
}
