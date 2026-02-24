package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"exa-cli/client"
)

func printResults(results []client.Result, jsonOut bool) {
	if jsonOut {
		b, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(b))
		return
	}
	for i, r := range results {
		fmt.Printf("\n--- Result %d ---\n", i+1)
		if r.Title != nil {
			fmt.Printf("Title:  %s\n", *r.Title)
		}
		fmt.Printf("URL:    %s\n", r.URL)
		if r.Score != nil {
			fmt.Printf("Score:  %.4f\n", *r.Score)
		}
		if r.PublishedDate != nil {
			fmt.Printf("Date:   %s\n", *r.PublishedDate)
		}
		if r.Author != nil {
			fmt.Printf("Author: %s\n", *r.Author)
		}
		if r.Summary != nil {
			fmt.Printf("Summary:\n%s\n", wrap(*r.Summary, 80))
		}
		if r.Text != nil {
			text := *r.Text
			if len(text) > 500 {
				text = text[:500] + "..."
			}
			fmt.Printf("Text:\n%s\n", wrap(text, 80))
		}
		if len(r.Highlights) > 0 {
			fmt.Println("Highlights:")
			for _, h := range r.Highlights {
				fmt.Printf("  • %s\n", h)
			}
		}
	}
}

func printAnswer(resp *client.AnswerResponse, jsonOut bool) {
	if jsonOut {
		b, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(b))
		return
	}
	fmt.Println("\n=== Answer ===")
	switch v := resp.Answer.(type) {
	case string:
		fmt.Println(wrap(v, 80))
	default:
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
	}
	if len(resp.Citations) > 0 {
		fmt.Println("\n=== Citations ===")
		for i, c := range resp.Citations {
			title := c.URL
			if c.Title != nil {
				title = *c.Title
			}
			fmt.Printf("[%d] %s\n    %s\n", i+1, title, c.URL)
		}
	}
}

func wrap(s string, width int) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	var lines []string
	line := words[0]
	for _, w := range words[1:] {
		if len(line)+1+len(w) > width {
			lines = append(lines, line)
			line = w
		} else {
			line += " " + w
		}
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}
