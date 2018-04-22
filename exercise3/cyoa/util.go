package cyoa

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func GetStoryline(r io.Reader) (Story, error) {
	chapters := make(Story)

	readFile, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(readFile, &chapters)
	if err != nil {
		return nil, err
	}

	return chapters, nil
}
