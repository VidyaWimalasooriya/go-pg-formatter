package format

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestFormatSQLString(t *testing.T) {
	tests := []struct {
		name string
		args []string
		file string
	}{
		{
			"Default",
			[]string{},
			"default.sql",
		},
		{
			"Parameters: start with comma",
			[]string{"--comma-start"},
			"comma-start.sql",
		},
		{
			"Keyword Case: Capitalize",
			[]string{"--keyword-case", "3"},
			"keyword-capitalise-case.sql",
		},
		{
			"Spaces: 8",
			[]string{"--spaces", "10"},
			"spaces.sql",
		},
		{
			"No comments",
			[]string{"--nocomment"},
			"no-comments.sql",
		},
		{
			"Tabs + Keyword Case: Lowercase + Parameters: start with comma",
			[]string{"--keyword-case", "1", "--comma-start", "--tabs"},
			"composite.sql",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := filepath.Join("testdata/expected", test.file)

			str, err := FormatSQLString("testdata/default.sql", test.args...)
			require.NoError(t, err)

			data, err := os.ReadFile(expected)
			require.NoError(t, err)

			assert.Equal(t, string(data), str)
		})
	}
}
