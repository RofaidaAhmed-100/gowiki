package main

import (
    "fmt"
    "os"
)

type Page struct {
    Title string
    Body  []byte
}

// الميثود save
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

// الدالة loadPage
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// الدالة main
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    err := p1.save()
    if err != nil {
        fmt.Println("Error saving page:", err)
        return
    }

    p2, err := loadPage("TestPage")
    if err != nil {
        fmt.Println("Error loading page:", err)
        return
    }

    fmt.Println(string(p2.Body))
}
