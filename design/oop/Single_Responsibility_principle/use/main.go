package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Report is responsible for generating reports.
type Report struct {
	Title   string
	Content string
}

func (r *Report) Generate() {
	r.Title = "Monthly Report"
	r.Content = "This is the content of the monthly report."
}

// ReportSaver is responsible for saving reports.
type ReportSaver struct{}

func (rs *ReportSaver) SaveToFile(r *Report, filename string) {
	content := r.Title + "\n" + r.Content
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Report saved to", filename)
}

func main() {
	report := &Report{}
	report.Generate()

	saver := &ReportSaver{}
	saver.SaveToFile(report, "report.txt")
}
