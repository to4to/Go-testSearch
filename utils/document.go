package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// document represents a document with its title, URL, text, and ID.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	Id    int
}

// LoadDocuments loads documents from the specified path.
// It reads an XML file, unmarshals it into a slice of documents,
// and assigns an ID to each document.
// The path parameter specifies the path to the XML file.
// It returns the loaded documents and any error encountered.
func LoadDocuments(path string) ([]document, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	gz, err := gzip.NewReader(f)

	if err != nil {
		return nil, err
	}

	defer gz.Close()

	dec := xml.NewDecoder(gz)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents

	for i := range docs {
		docs[i].Id = i
	}
	return docs, nil
}
