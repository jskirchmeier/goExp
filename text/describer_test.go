package text

import (
	"testing"
)

func two(short bool) string {
	if short {
		return "short"
	}
	return "long"

}

func TestDescriberFunc_Description(t *testing.T) {
	tests := []struct {
		name  string
		h     DescriberFunc
		short bool
		want  string
	}{
		{"short", DescriberFunc(two), true, "short"},
		{"long", DescriberFunc(two), false, "long"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Description(tt.short); got != tt.want {
				t.Errorf("DescriberFunc.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleDescriber_Description(t *testing.T) {
	tests := []struct {
		name  string
		s     SimpleDescriber
		short bool
		want  string
	}{
		{"one", "one", true, "one"},
		{"one", "one", false, "one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Description(tt.short); got != tt.want {
				t.Errorf("SimpleDescriber.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDescriberHelper_Description(t *testing.T) {
	h := DescriberHelper{"short", "long"}
	tests := []struct {
		name  string
		h     DescriberHelper
		short bool
		want  string
	}{
		{"short", h, true, "short"},
		{"long", h, false, "long"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Description(tt.short); got != tt.want {
				t.Errorf("DescriberHelper.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplateHelper(t *testing.T) {
	o := struct {
		A string
		B int
	}{"one", 42}
	type args struct {
		shortTemplate string
		longTemplate  string
		o             interface{}
		short         bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"short", args{"{{.A}}", "{{.B}}", o, true}, "one"},
		{"long", args{"{{.A}}", "{{.B}}", o, false}, "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TemplateFunc(tt.args.shortTemplate, tt.args.longTemplate, tt.args.o).Description(tt.args.short); got != tt.want {
				t.Errorf("TemplateFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
