package questlisting

import (
	"io"
	templateprovider "website/internal/templateProvider"
)

func RenderListing(writer io.Writer) error {
	tmpl, err := templateprovider.Provide("list.html")
	if err != nil {
		return err
	}

	tmpl.Execute(writer, provideQuests())

	return nil
}
