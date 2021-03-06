package markdown

import (
	"io/ioutil"
	"strings"
)

var readFileFn = ioutil.ReadFile

// Markdown is markdown document
type Markdown struct {
	elm    []ElementInterface
	cursor int
	length int
}

// NewMD creates a markdown document
func NewMD() *Markdown {
	return &Markdown{
		elm:    []ElementInterface{},
		cursor: -1,
		length: 0,
	}
}

// ParseString to parse string content
func (md *Markdown) ParseString(content string) error {
	lines := md.toLines(content)

	return md.parse(lines)
}

// Parse md file
func (md *Markdown) Parse(file string) error {
	lines, e := md.read(file)
	if e != nil {
		return e
	}

	return md.parse(lines)
}

func (md *Markdown) parse(lines []string) error {
	iterator := NewIterator(lines)
	for iterator.Next() {
		md.elm = append(md.elm, NewElement(iterator.Value()))
	}

	md.length = len(md.elm)

	return nil
}

func (md *Markdown) toLines(contents string) []string {
	linebreak := "\n"
	if strings.Contains(contents, "\r\n") {
		linebreak = "\r\n"
	}

	lines := strings.Split(contents, linebreak)
	output := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			output = append(output, line)
		}
	}
	return output
}

func (md *Markdown) read(file string) ([]string, error) {
	b, err := readFileFn(file)
	if err != nil {
		return []string{}, err
	}
	contents := string(b)
	output := md.toLines(contents)

	return output, nil
}

// Len returns number of elements
func (md *Markdown) Len() int {
	return md.length
}

// Next move cursor to next position
func (md *Markdown) Next() bool {
	md.cursor++
	return md.cursor < md.length
}

// Value get current value
func (md *Markdown) Value() ElementInterface {
	if md.cursor >= md.length {
		return nil
	}
	return md.elm[md.cursor]
}

// Reset cursor of iterator to beginning
func (md *Markdown) Reset() {
	md.cursor = -1
}
