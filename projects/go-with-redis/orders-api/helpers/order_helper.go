package helpers

import "fmt"

func OrderIDKey(id uint) string {
	return fmt.Sprintf("order:%d", id)
}
