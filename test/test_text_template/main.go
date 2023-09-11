package main

import (
        "path"
        "text/template"
        "fmt"
        "os"
)

type Page struct {
    Title string
    Name  string
}

type Data struct {
    Name string
}

func test_html_tmpl(page Page) {

    templatePath := "/Users/k.song/src/nginx-template/test/test_text_template/template.tmpl"
    tmpl, err := template.New(path.Base(templatePath)).ParseFiles(templatePath)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    err = tmpl.Execute(os.Stdout, page)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
        return
    }

    fmt.Printf("\n")
}

func test_html() {

    page1 := Page{Title: "Test Page 1", Name:  "Kylin"}
    page2 := Page{Title: "Test Page 2", Name:  "Kylin"}

    test_html_tmpl(page1)
    test_html_tmpl(page2)

}

func test_plain_txt_tmpl(data Data) {

    textPath := "/Users/k.song/src/nginx-template/test/test_text_template/text.tmpl"
    tmpl, err := template.New(path.Base(textPath)).ParseFiles(textPath)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    err = tmpl.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }

    fmt.Printf("\n")
}

func test_plain_txt() {
    
    kylin := Data{Name: "Kylin"}
    soong := Data{Name: "Soong"}

    test_plain_txt_tmpl(kylin)
    test_plain_txt_tmpl(soong)
}


func main() {

    test_html()

    test_plain_txt()

}
