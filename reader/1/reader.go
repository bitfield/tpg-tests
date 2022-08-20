package reader

import "io"

func ReadAll(r io.Reader) ([]byte, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
