package smallsh

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCD(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		want     string
		relative bool
	}{
		{
			name:     "relative",
			args:     []string{"testdata"},
			want:     "testdata",
			relative: true,
		},
		{
			name:     "relative_extra_args",
			args:     []string{"testdata", "extra", "args", "are", "ignored"},
			want:     "testdata",
			relative: true,
		},
		{
			name:     "absolute",
			args:     []string{"/"},
			want:     "/",
			relative: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, err := os.Getwd()
			if err != nil {
				t.Fatalf("Could not get working directory before cd: %v", err)
			}
			s := new(Smallsh)
			if err := s.cd(test.args); err != nil {
				t.Logf("With cd(%+v):", test.args)
				t.Errorf("Should not have gotten error: %v", err)
			}
			got, err := os.Getwd()
			if err != nil {
				t.Fatalf("Could not get working directory after cd: %v", err)
			}
			var want string
			if test.relative {
				want = filepath.Join(before, test.want)
			} else {
				want = test.want
			}

			if got != want {
				t.Logf("After cd(%+v):", test.args)
				t.Errorf("Cwd is now %v, want %v", got, want)
			}
			if err := os.Chdir(before); err != nil {
				t.Fatalf("Could not change dir back to %v: %v", before, err)
			}
		})
	}
}

func TestCDError(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "invalid_folder",
			args: []string{"doesnt_exist"},
			want: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, err := os.Getwd()
			if err != nil {
				t.Fatalf("Could not get working directory before: %v", err)
			}
			s := new(Smallsh)
			if err := s.cd(test.args); err == nil {
				t.Logf("With cd(%+v):", test.args)
				t.Errorf("Should have gotten an error.")
			}
			after, err := os.Getwd()
			if err != nil {
				t.Errorf("Could not get working directory after: %v", err)
			}
			if before != after {
				t.Errorf("Cwd is now %v, should not have changed from %v.", after, before)
			}
		})
	}
}

func TestCDWithNoArguments(t *testing.T) {
	want, ok := os.LookupEnv("HOME")
	if !ok {
		t.Fatalf("Could not find $HOME.")
	}
	before, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get working directory before cd: %v", err)
	}
	s := new(Smallsh)
	s.cd(nil)
	got, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get working directory after cd: %v", err)
	}
	if got != want {
		t.Errorf("Directory = %v, want %v", got, want)
	}
	if err := os.Chdir(before); err != nil {
		t.Fatalf("Could not change dir back: %v", err)
	}
}
