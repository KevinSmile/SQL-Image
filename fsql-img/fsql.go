package fsql

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/KevinSmile/SQL-Image/fsql-img/parser"
	"github.com/KevinSmile/SQL-Image/image"
)

func init() {
	CurrentImage = "ubuntu:latest1234"
}

var CurrentImage string

// Run parses the input and executes the resultant query.
func Run(input string) error {
	input = strings.TrimSpace(input)
	if strings.EqualFold(input, "SHOW IMAGES") {
		for _, imageTag := range image.GetImageTags() {
			fmt.Println(imageTag)
		}
		return nil
	}
	if strings.HasPrefix(strings.ToUpper(input), "USE") {
		CurrentImage = strings.Split(input, " ")[1]
		fmt.Println(CurrentImage)
		return nil

	}
	if strings.EqualFold(input, "SHOW LAYERS") {
		for _, imageTag := range image.GetImageLayers(CurrentImage) {
			fmt.Println(imageTag)
		}
		return nil
	}

	q, err := parser.Run(input)
	if err != nil {
		return err
	}

	// Find length of the longest name to normalize name output.
	var max = 0
	var results = make([]map[string]interface{}, 0)

	err = q.Execute(
		func(path string, info os.FileInfo, result map[string]interface{}) {
			result["name"] = path
			results = append(results, result)
			if !q.HasAttribute("name") {
				return
			}
			if s, ok := result["name"].(string); ok && len(s) > max {
				max = len(s)
			}
		},
	)
	if err != nil {
		return err
	}

	for _, result := range results {
		var buf bytes.Buffer
		for j, attribute := range q.Attributes {
			// If the current attribute is "name", pad the output string by `max`
			// spaces.
			format := "%v"
			if attribute == "name" {
				format = fmt.Sprintf("%%-%ds", max)
			}
			buf.WriteString(fmt.Sprintf(format, result[attribute]))
			if j != len(q.Attributes)-1 {
				buf.WriteString("\t")
			}
		}
		fmt.Printf("%s\n", buf.String())
	}

	return nil
}
