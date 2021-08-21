package controlFlow

import (
	"errors"
	"fmt"
	"io"
)

type SimpleReader struct {
	count int
}

var errCatasthropicReader = errors.New("Catastrhophe.")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	println(br.count)
	if br.count == 2 {
		panic(errCatasthropicReader)
	}
	if br.count > 3 {
		return 0, errors.New("Random error.")
		//return 0, io.EOF // Common EOF for readers.
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	fmt.Println("Closing reader.")
	return nil
}

func ReadFullFile() (err error) {
	var reader io.ReadCloser = &SimpleReader{}

	// If bad anything happens during the reading the file or at the end of reading,
	// the reader will be closed.

	defer func(reader io.ReadCloser) {
		fmt.Println("First defer function.")
		_ = reader.Close()
		if p := recover(); p == errCatasthropicReader {
			println(p)
			err = errors.New("Panic occurred, but we continue.")
		} else if p != nil {
			panic("Unexpected error.")
		}

	}(reader)

	defer func() {
		fmt.Println("Second defer function.")
	}()

	for {
		value, readerErr := reader.Read([]byte("Okaay."))
		if readerErr == io.EOF {
			println("Finished reading file.")
			//break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		fmt.Println("FOR: ", value)
	}
	return
}
