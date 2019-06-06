package language

import (
	"encoding/json"
	"errors"

	"github.com/gobuffalo/packr/v2"
)

var box = packr.New("data", "../data")

var (
	ErrCountryNotFound = errors.New("country not found")
	ErrNoCountryRows   = errors.New("no country row")
)

type (
	Language struct {
		Name string `json:"English"`
		Alpha2 string `json:"alpha2"`
		Alpha3 string `json:"alpha3-b"`
	}

	Languages []Language
)

func List() (Languages, error) {
	var l Languages

	j, e := box.Find("language-codes-3b2.json")
	if e != nil {
		return nil, e
	}

	if e := json.Unmarshal(j, &l); e != nil {
		return nil, e
	}

	if len(l) == 0 {
		return nil, ErrNoCountryRows
	}

	return l, nil
}

func Find(query string) (*Language, error) {

	l, e := List()
	if e != nil {
		return nil, e
	}

	for _, v := range l {
		if v.Name == query ||
			v.Alpha3 == query ||
			v.Alpha2 == query {
			return &v, nil
		}
	}

	return nil, ErrCountryNotFound
}
