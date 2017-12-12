package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_fileName(t *testing.T) {
	type args struct {
		p string
		r int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test10x20",
			args: args{
				p: "test",
				r: 10,
				c: 20,
			},
			want: "test10x20.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileName(tt.args.p, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("fileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func verifyTableProperties(table Table, r, c int) error {
	if len(table) != r {
		return fmt.Errorf("makeTable(): rows = %v, want %v", len(table), r)
	}
	for i := 0; i < r; i++ {
		if len(table[i].Hrate) != c {
			return fmt.Errorf("makeTable(): cols = %v, want %v", len(table), c)
		}
	}
	return nil
}

func Test_makeTable(t *testing.T) {
	tests := []struct {
		name string
		r, c int
	}{
		{name: "10x10", r: 10, c: 10},
		{name: "1x10", r: 1, c: 10},
		{name: "10x1", r: 10, c: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := makeTable(tt.r, tt.c)
			err := verifyTableProperties(table, tt.r, tt.c)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_read(t *testing.T) {
	csv := `"Bart Beatty","98","130","158","129","128"
"Marc Murphy","156","93","131","112","118"
"Briana Bauch","108","133","168","121","144"
"Gerda Rosenbaum","109","166","95","159","139"
"Guido Witting","120","123","165","107","135"`
	rows, cols := 5, 5
	sr := strings.NewReader(csv)
	table, err := read(sr, rows, cols)
	if err != nil {
		t.Errorf("read() error = %v", err)
		return
	}
	err = verifyTableProperties(table, rows, cols)
	if err != nil {
		t.Error(err)
	}
}

func Test_process(t *testing.T) {
	tests := []struct {
		name string
		r, c int
		data Table
		want Table
	}{
		{
			name: "process1",
			r:    5,
			c:    3,
			data: []Row{
				{Name: "A", Hrate: []int{10, 20, 30}},
				{Name: "B", Hrate: []int{100, 100, 100}},
				{Name: "C", Hrate: []int{0, 100, 100}},
				{Name: "D", Hrate: []int{80, 110, 140}},
				{Name: "E", Hrate: []int{99, 100, 101}},
			},
			want: []Row{
				{Name: "A", Hrate: []int{20, 10, 30}},
				{Name: "B", Hrate: []int{100, 100, 100}},
				{Name: "C", Hrate: []int{66, 0, 100}},
				{Name: "D", Hrate: []int{110, 80, 140}},
				{Name: "E", Hrate: []int{100, 99, 101}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v\nwant %v", got, tt.want)
			}
		})
	}
}

func Test_write(t *testing.T) {
	*rows = 5
	*cols = 3
	csv := `Name,avg,min,max
Bart Beatty,128,98,158
Alejandra Kunde,146,127,161
Sunny Gerlach,170,165,174
Jarod Wolff,136,125,145
Verla Abshire,135,114,154
`
	var out bytes.Buffer // implements io.Writer
	table1 := []Row{
		{Name: "Bart Beatty", Hrate: []int{128, 98, 158}},
		{Name: "Alejandra Kunde", Hrate: []int{146, 127, 161}},
		{Name: "Sunny Gerlach", Hrate: []int{170, 165, 174}},
		{Name: "Jarod Wolff", Hrate: []int{136, 125, 145}},
		{Name: "Verla Abshire", Hrate: []int{135, 114, 154}},
	}
	tests := []struct {
		name    string
		t       Table
		wantErr bool
	}{
		{name: "write1", t: table1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := write(tt.t, &out); (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out.String() != csv {
				t.Errorf("write() = [%v]\nwant = [%v]", out.String(), csv)
			}
		})
	}
}
