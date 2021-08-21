package controlFlow

import (
	"errors"
	"fmt"
	"io"
)

type BadReader struct {
	err error
}

func ErrorHandling() error {
	var reader io.Reader = &BadReader{errors.New("My nonsense error.")}
	if _, err := reader.Read([]byte("test")); err != nil {
		fmt.Println("The error occurred.")
		return err
	}
	return nil
}

func (br *BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}
