package main

import (
	"testing"
)

func TestReadLine (t *testing.T) {
	lines := []struct {
		line    string
		title   string
		url     string
		files   []string
		styles  []string
		scripts []string
	}{
		// valid tests
		{"Easy @ simple => file1", "Easy", "simple", []string{"file1"}, []string{}, []string{}},
		{"notitle => file2,  file3", "Aria's notes", "notitle", []string{"file2", "file3"},
			[]string{}, []string{},},
		{"Apple @ orange => file4\nstyles => style1.css, style2.css\nstyle3.css", "Apple",
			"orange", []string{"file4"}, []string{"style1.css", "style2.css", "style3.css"},
			[]string{}},
		// incorrect syntax
		{"Pomegranite @ peach @ strawberry => file5", "Pomegranite", "", []string{}, []string{},
			[]string{}},
	}
	for _, val := range(lines) {
		url, config := parseLine(val.line)
		if url != val.url {
			t.Errorf("Urls don't match, expected: %s, got: %s", val.url, url)
		}
		if config.title != val.title {
			t.Errorf("Titles don't match, expected: %s, got: %s", val.title, config.title)
		}
		
		/* check slices */
		// check files for actual page
		if len(config.files) == len(val.files) {
			for idx, file := range(config.files) {
				if file != val.files[idx] {
					t.Errorf("Files don't match: %s %s", file, val.files[idx])
				}
			}
		} else {
			t.Errorf("Different numbers of files")
		}
		
		// check styles
		if len(config.styles) == len(val.styles) {
			for idx, style := range(config.styles) {
				if style != val.styles[idx] {
					t.Errorf("Styles don't match: %s %s", style, val.styles[idx])
				}
			}
		} else {
			t.Errorf("Different numbers of styles")
		}
		
		// check scripts
		if len(config.scripts) == len(val.scripts) {
			for idx, script := range(config.scripts) {
				if script != val.scripts[idx] {
					t.Errorf("Scripts don't match: %s %s", script, val.scripts[idx])
				}
			}
		} else {
			t.Errorf("Different numbers of scripts")
		} // end else for checking scripts
	} // end for
}