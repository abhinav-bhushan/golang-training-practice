package fileops

import (
	"errors"
	"fmt"
	"os"
)

func GetCharCount(filename *string) (int, error) {
	if filename == nil {
		return 0, errors.New("nil filename")
	}

	bytes, err := os.ReadFile(*filename)

	if err != nil {
		println(err)
		return 0, err
	}
	fmt.Println("Count of Chars in the file:", len(bytes))
	return len(bytes), nil
}
