package main

import (
	"bytes"
	"reflect"
	"sort"
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

func Test_read(t *testing.T) {
	csv := `"Bart Beatty","98","130","158","129","128"
"Marc Murphy","156","93","131","112","118"
"Briana Bauch","108","133","168","121","144"
"Gerda Rosenbaum","109","166","95","159","139"
"Guido Witting","120","123","165","107","135"`
	rows, cols := 5, 5
	sr := strings.NewReader(csv)
	ch, errch := read(sr, rows)

	// read from the channels and test each row.
	for {
		select {
		case row, ok := <-ch:
			if !ok {
				// Channel got closed - this is expected.
				return
			}
			// Check properties rather than values
			if row.Name == "" || len(row.Hrate) != cols {
				t.Errorf("read(): invalid row %#v", row)
			}
		case err, ok := <-errch:
			if ok { // errch has not been closed
				t.Errorf("read(): got error %v", err)
			}
		}
	}
}

func Test_process(t *testing.T) {
	tests := []struct {
		name  string
		bufsz int
		data  []Row
		want  []Row
	}{
		{
			name:  "process1",
			bufsz: 5,
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
			in := make(chan Row, tt.bufsz)

			// feed the process goroutine with rows.
			go func() {
				for _, row := range tt.data {
					in <- row
				}
				close(in)
			}()

			out := process(in, tt.bufsz)

			// Collect rows from the output channel.
			got := []Row{}
			for row, ok := <-out; ok; row, ok = <-out {
				got = append(got, row)
			}

			// The concurrent version does not enforce to preserve the
			// original sequence of rows.
			// Therefore, we need to sort the output from process() before
			// comparing against the wanted result.
			byName := func(i, j int) bool {
				if got[i].Name < got[j].Name {
					return true
				}
				return false
			}
			sort.Slice(got, byName)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = \n%#v\nwant = \n%#v", got, tt.want)
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
	rows := []Row{
		{Name: "Bart Beatty", Hrate: []int{128, 98, 158}},
		{Name: "Alejandra Kunde", Hrate: []int{146, 127, 161}},
		{Name: "Sunny Gerlach", Hrate: []int{170, 165, 174}},
		{Name: "Jarod Wolff", Hrate: []int{136, 125, 145}},
		{Name: "Verla Abshire", Hrate: []int{135, 114, 154}},
	}
	tests := []struct {
		name    string
		rows    []Row
		wantErr bool
	}{
		{name: "write1", rows: rows, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan Row)
			go func() {
				for _, row := range tt.rows {
					in <- row
				}
				close(in)
			}()
			if err := write(in, &out); (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out.String() != csv {
				t.Errorf("write() = [%v]\nwant = [%v]", out.String(), csv)
			}
		})
	}
}
