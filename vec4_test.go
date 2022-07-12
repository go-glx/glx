package glx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec4_BasicMath(t *testing.T) {
	tests := []struct {
		name    string
		input   Vec4
		operand Vec4
		wantAdd Vec4
		wantSub Vec4
		wantMul Vec4
		wantDiv Vec4
	}{
		{
			name:    "pos",
			input:   Vec4{5, 10, 4, 5},
			operand: Vec4{2, 3, 1, 0},
			//
			wantAdd: Vec4{7, 13, 5, 5},
			wantSub: Vec4{3, 7, 3, 5},
			wantMul: Vec4{10, 30, 4, 0},
			wantDiv: Vec4{2.5, 3.3333333, 4, InfPositive()},
		},
		{
			name:    "neg",
			input:   Vec4{-5, -10, -4, 0},
			operand: Vec4{-2, -3, -1, 5},
			//
			wantAdd: Vec4{-7, -13, -5, 5},
			wantSub: Vec4{-3, -7, -3, -5},
			wantMul: Vec4{10, 30, 4, 0},
			wantDiv: Vec4{2.5, 3.3333333, 4, 0},
		},
		{
			name:    "cross",
			input:   Vec4{-5, 10, 4, 0.5},
			operand: Vec4{2, -3, 1, 0.1},
			//
			wantAdd: Vec4{-3, 7, 5, 0.6},
			wantSub: Vec4{-7, 13, 3, 0.4},
			wantMul: Vec4{-10, -30, 4, 0.05},
			wantDiv: Vec4{-2.5, -3.3333333, 4, 5},
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
