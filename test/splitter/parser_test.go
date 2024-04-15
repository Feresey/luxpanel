package parser_test

import (
	"context"
	"embed"
	"io/fs"
	"path/filepath"
)

//go:embed testdata
var daysFS embed.FS

func (s *Suite) TestParseGameLog() {
	r := s.Require()
	ctx := context.Background()

	days, err := daysFS.ReadDir("testdata")
	r.NoError(err)

	for _, day := range days {
		f, err := fs.Sub(daysFS, filepath.Join("testdata", day.Name()))
		r.NoError(err)

		levels, err := s.splitter.SplitLevels(ctx, f)
		r.NoError(err)
		r.NotEmpty(levels)
	}
}
