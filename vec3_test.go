package glx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3_BasicMath(t *testing.T) {
	tests := []struct {
		name    string
		input   Vec3
		operand Vec3
		wantAdd Vec3
		wantSub Vec3
		wantMul Vec3
		wantDiv Vec3
	}{
		{
			name:    "pos",
			input:   Vec3{5, 10, 4},
			operand: Vec3{2, 3, 1},
			//
			wantAdd: Vec3{7, 13, 5},
			wantSub: Vec3{3, 7, 3},
			wantMul: Vec3{10, 30, 4},
			wantDiv: Vec3{2.5, 3.3333333, 4},
		},
		{
			name:    "neg",
			input:   Vec3{-5, -10, -4},
			operand: Vec3{-2, -3, -1},
			//
			wantAdd: Vec3{-7, -13, -5},
			wantSub: Vec3{-3, -7, -3},
			wantMul: Vec3{10, 30, 4},
			wantDiv: Vec3{2.5, 3.3333333, 4},
		},
		{
			name:    "cross",
			input:   Vec3{-5, 10, 4},
			operand: Vec3{2, -3, 1},
			//
			wantAdd: Vec3{-3, 7, 5},
			wantSub: Vec3{-7, 13, 3},
			wantMul: Vec3{-10, -30, 4},
			wantDiv: Vec3{-2.5, -3.3333333, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantAdd, tt.input.Add(tt.operand), "add")
			assert.Equal(t, tt.wantSub, tt.input.Sub(tt.operand), "sub")
			assert.Equal(t, tt.wantMul, tt.input.Mul(tt.operand), "mul")
			assert.Equal(t, tt.wantDiv, tt.input.Div(tt.operand), "div")
		})
	}
}
