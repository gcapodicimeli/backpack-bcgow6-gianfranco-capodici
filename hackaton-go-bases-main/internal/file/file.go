package file

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() (lines [][]string, err error) {
	file, err := os.Open(f.Path)
	if err != nil || file == nil {
		return nil, fmt.Errorf("No se pudo abrir el archivo %v", f.Path)
	}
	lines, err = csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return
}

func (f *File) Write(service.Ticket) error {
	return nil
}
