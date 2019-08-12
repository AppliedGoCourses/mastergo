package tabledriven

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Single white spaces",
			args:    args{s: "a b c"},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "Multiple white spaces",
			args:    args{s: "a   b  \t  c"},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "Leading and trailing white spaces",
			args:    args{s: "    a   b  \t\n  c  \t "},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "Only white spaces",
			args:    args{s: "  \t  \n\t    \t"},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "Empty input",
			args:    args{s: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "No white spaces",
			args:    args{s: "abc"},
			want:    []string{"abc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Scan(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
