package settings_test

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	. "github.com/julienmoumne/hotshell/cmd/hs/settings"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
	"testing"
)

func errMsg(msg string) string {
	return fmt.Sprintf("Error while reading ~/.hsrc.js\n%s", msg)
}

var (
	a      *assert.Assertions
	loader Loader
	fs     vfs.Filesystem
	tests  = []testCase{
		{
			in{fileNotFound: true},
			out{cfg: Defaults()},
		},
		{
			in{js: ``},
			out{cfg: Defaults()},
		},
		{
			in{js: `invalidStatement{}`},
			out{err: errMsg("(anonymous): Line 3:17 Unexpected token { (and 2 more errors)")},
		},
		{
			in{js: `
			var settings = require('hotshell-settings')
			var keys = settings.keys

			settings.set({
			  keys: {
			    bash: 2
			  }
			})
			`},
			out{err: errMsg("1 error(s) decoding:\n\n* 'Keys.Bash' expected type 'string', got unconvertible type 'int64'")},
		},
		{
			in{js: `
			var settings = require('hotshell-settings')
			var keys = settings.keys

			settings.set({
			  keys: {
			    back: keys.Return,
			    bash: keys.Backspace,
			    repeat: keys.Tab,
			    reload: keys.Space,
			  }
			})
			`},
			out{cfg: Settings{Keys: Keys{
				Back:   KeyCodes.Return,
				Bash:   KeyCodes.Backspace,
				Repeat: KeyCodes.Tab,
				Reload: KeyCodes.Space,
			}}},
		},
		{
			in{js: `
			var settings = require('hotshell-settings')
			var keys = settings.keys

			settings.set({
			  keys: {
			    repeat: keys.Tab,
			  }
			})
			`},
			out{cfg: Settings{Keys: Keys{
				Back:   KeyCodes.Space,
				Bash:   KeyCodes.Tab,
				Repeat: KeyCodes.Tab,
				Reload: KeyCodes.Backspace,
			}}},
		},
	}
)

type (
	in struct {
		fileNotFound bool
		js           string
	}
	out struct {
		cfg Settings
		err string
	}
	testCase struct {
		in
		out
	}
)

func TestSettings(t *testing.T) {
	a = assert.New(t)
	for _, t := range tests {
		runTest(t)
	}
}

func runTest(t testCase) {
	setupTest(t)
	validateTest(t)
}

func setupTest(t testCase) {
	fs = memfs.Create()
	loader = Loader{}
	if t.in.fileNotFound {
		return
	}
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	vfs.MkdirAll(fs, home, 0)
	vfs.WriteFile(fs, fmt.Sprintf("%s/.hsrc.js", home), []byte(t.in.js), 0)
}

func validateTest(t testCase) {
	cfg, err := loader.Load(fs)
	if t.out.err != "" {
		a.NotNil(err)
		a.Equal(t.out.err, err.Error())
	} else {
		a.Nil(err)
		a.Equal(t.out.cfg, cfg)
	}
}
