package main

import (
    "fmt"
    "io/ioutil"
)
/**
* Page definition
*/
type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}
