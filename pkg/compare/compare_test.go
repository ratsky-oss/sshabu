package compare

import (
	"reflect"
	"testing"
	"fmt"
)

func Test_transformDifferencesToReadableFormat(t *testing.T) {
	type args struct {
		differences []Difference
		firstBites  Bites
		secondBites Bites
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "No differences",
			args: args{
				differences: []Difference{},
				firstBites:  Bites{Content: []string{"A"}},
				secondBites: Bites{Content: []string{"A"}},
			},
			want: []string{ fmt.Sprintf("%d: %s%s%s", 1, White, "A", White)},
		},
		{
			name: "Addition in the second file",
			args: args{
				differences: []Difference{{lineNumber: 2, line: "B", Added: true}},
				firstBites:  Bites{Content: []string{"A"}},
				secondBites: Bites{Content: []string{"A","B"}},
			},
			want: []string{ 
							fmt.Sprintf("%d: %s%s%s", 1, White, "A", White), 
							fmt.Sprintf("%d: %s%s%s", 2, Green, "B", White), 
						  },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformDifferencesToReadableFormat(tt.args.differences, tt.args.firstBites, tt.args.secondBites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformDifferencesToReadableFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diffBites(t *testing.T) {
	type args struct {
		bites1 Bites
		bites2 Bites
	}
	tests := []struct {
		name string
		args args
		want []Difference
	}{
		{
			name: "Add one line",
			args: args{
				bites1: Bites{Content: []string{"A","B", "C"}},
				bites2: Bites{Content: []string{"A","B", "C", "D"}},
			},
			want: []Difference{{lineNumber: 4, line: "D", Added: true}},
		},
		{
			name: "Del one line",
			args: args{
				bites1: Bites{Content: []string{"A","B", "C"}},
				bites2: Bites{Content: []string{"A","B"}},
			},
			want: []Difference{{lineNumber: 3, line: "C", Added: false}},
		},
		{
			name: "Multy edit",
			args: args{
				bites1: Bites{Content: []string{"B","E","D","E"}},
				bites2: Bites{Content: []string{"A","B","D","E","D","E","C"}},
			},
			want: []Difference{ {lineNumber: 1, line: "A", Added: true},
								{lineNumber: 3, line: "D", Added: true},
								{lineNumber: 7, line: "C", Added: true},
							  },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffBites(tt.args.bites1, tt.args.bites2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffBites() = %v, want %v", got, tt.want)
			}
		})
	}
}
