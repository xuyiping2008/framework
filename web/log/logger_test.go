package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		path    string
		logname string
		level   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"", "", 1}, true},
		{"test2", args{"./golog", "logs.log", 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.path, tt.args.logname, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
