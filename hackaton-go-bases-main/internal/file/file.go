package file

import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	//tickets, err = os.ReadFile(f.path)
	return
}

func (f *File) Write(service.Ticket) error {
	return nil
}
