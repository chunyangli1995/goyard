package err

import "fmt"

type InvalidInput struct {
	Msg string
}

func NewInvalidInputError() *InvalidInput {
	msg := fmt.Sprintf("invalid input error")
	return &InvalidInput{
		Msg: msg,
	}
}

func (c *InvalidInput) Error() string {
	return c.Msg
}
