package services
	
import (
    "fmt"
		"os"
)

func ReadDay(day string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("inputs/%s.txt", day))

	if err != nil {
		return "", err
	}

	return string(data), nil
}
