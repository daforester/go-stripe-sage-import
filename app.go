package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Test returns a test message
func (a *App) Test() string {
	return fmt.Sprintf("This is a test!")
}

func (a *App) OpenFile() {
	filePath, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
	})

	b, err := os.Open(filePath)
	defer func(b *os.File) {
		err = b.Close()
		if err != nil {
			panic(err)
		}
	}(b)
	if err != nil {
		panic(err)
	}

	lines, err := csv.NewReader(b).ReadAll()
	headers := make(map[string]int)

	for i, header := range lines[0] {
		headers[header] = i
	}

	output := make([][]string, 0)

	output = append(output, []string{
		"Date",
		"Description",
		"Money in",
		"Money out",
	})

	for _, line := range lines {
		if line[headers["Status"]] != "Paid" && line[headers["Status"]] != "Refunded" {
			continue
		}
		entry := make([]string, 0)
		date := strings.SplitN(line[headers["Created date (UTC)"]], " ", 2)[0]
		description := strings.Map(func(r rune) rune {
			if r > 127 {
				return -1
			}
			return r
		}, line[headers["Description"]])

		entry = append(entry, date)
		entry = append(entry, description)
		entry = append(entry, line[headers["Amount"]])
		entry = append(entry, "")
		output = append(output, entry)

		entry = make([]string, 0)
		entry = append(entry, date)
		entry = append(entry, "Processing Fee: "+description)
		entry = append(entry, "")
		entry = append(entry, line[headers["Fee"]])
		output = append(output, entry)

		if line[headers["Amount Refunded"]] != "0.00" {
			entry = make([]string, 0)
			date := strings.SplitN(line[headers["Refunded date (UTC)"]], " ", 2)[0]
			entry = append(entry, date)
			entry = append(entry, "Refund: "+description)
			entry = append(entry, "")
			entry = append(entry, line[headers["Amount Refunded"]])
			output = append(output, entry)
		}
	}

	dest, _ := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: "sage.csv",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "csv",
				Pattern:     "*.csv",
			},
		},
		Title: "Enter Destination File",
	})

	d, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(b *os.File) {
		err = d.Close()
		if err != nil {
			panic(err)
		}
	}(d)
	if err != nil {
		panic(err)
	}
	err = csv.NewWriter(d).WriteAll(output)
	if err != nil {
		panic(err)
	}
}
