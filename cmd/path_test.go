package cmd

import "testing"

// disabled, FindProjectRoot uses os.Stat
func testFindProjectroot(t *testing.T) {
	tests := []struct {
		path   string
		gopath []string
		want   string
		err    error
	}{{
		path: "/home/foo/work/project/src",
		want: "/home/foo/work/project",
	}}

	for _, tt := range tests {
		got, err := FindProjectroot(tt.path, tt.gopath)
		if got != tt.want || err != tt.err {
			t.Errorf("FindProjectroot(%v, %v): want: %v, %v, got %v, %v", tt.path, tt.gopath, tt.want, tt.err, got, err)
		}
	}
}
