package diary

import (
	"flag"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"
)

type Subcommand struct {
	Action string
}

func (r *Subcommand) Name() string {
	return "diary"
}

func (r *Subcommand) Parse(args []string) error {
	fs := flag.NewFlagSet(r.Name(), flag.ContinueOnError)
	fs.StringVar(&r.Action, "action", "", "Diary action")

	return fs.Parse(args)
}

func (r *Subcommand) Run() error {
	if r.Action == "add_diary" {
		return AddDiary()
	}

	return nil
}
