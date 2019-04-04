package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// save ...
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //filenameというp.Bodyという内容のfileを作る。最後のやつがわかんない。
}
