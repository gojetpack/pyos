package pyos

import (
	"testing"
)

const (
	NonExistTestFile = "testdata/this_file_doesnt_exist.txt"
	ExistingTestFile = "testdata/test_file.txt"
	ExistingTestDir  = "testdata/empty_dir"
)

func Test_path_Exist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "does not exist",
			args: args{
				path: NonExistTestFile,
			},
			want: false,
		},
		{
			name: "exist",
			args: args{
				path: ExistingTestFile,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Path.Exist(tt.args.path); got != tt.want {
				t.Errorf("FileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_path_IsDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is file",
			args: args{
				path: ExistingTestFile,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := path{}
			if got := p.IsDir(tt.args.path); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_path_IsFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is file",
			args: args{
				path: ExistingTestFile,
			},
			want: true,
		},
		{
			name: "is dir",
			args: args{
				path: ExistingTestDir,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := path{}
			if got := p.IsFile(tt.args.path); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
