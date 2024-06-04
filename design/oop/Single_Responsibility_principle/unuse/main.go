package main

import (
	"fmt"
	"log"
	"os"
)

// Report is responsible for generating and saving reports.
type Report struct {
	Title   string
	Content string
}

func (r *Report) Generate() {
	r.Title = "Monthly Report"
	r.Content = "This is the content of the monthly report."
}

func (r *Report) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(r.Title + "\n" + r.Content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Report saved to", filename)
}

func main() {
	report := &Report{}
	report.Generate()
	report.SaveToFile("report.txt")
}
