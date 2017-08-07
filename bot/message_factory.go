package bot

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

// Generatea response from the json template
func genResponseTemplate(recipientId, message string, writer io.Writer) error {
	tm := loadTemplate("response.json")

	// when we execute the template we pass the data missing
	return tm.Execute(writer, map[string]string{
		"recipient_id": recipientId,
		"message":      message,
	})
}

// loads the template given the filename
func loadTemplate(filename string) *template.Template {
	template_path := fmt.Sprintf("messages/%v", filename)
	tm, err := template.ParseFiles(template_path)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return tm
}
