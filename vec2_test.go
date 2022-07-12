package glx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2_BasicMath(t *testing.T) {
	tests := []struct {
		name    string
		input   Vec2
		operand Vec2
		wantAdd Vec2
		wantSub Vec2
		wantMul Vec2
		wantDiv Vec2
	}{
		{
			name:    "pos",
			input:   Vec2{5, 10},
			operand: Vec2{2, 3},
			//
			wantAdd: Vec2{7, 13},
			wantSub: Vec2{3, 7},
			wantMul: Vec2{10, 30},
			wantDiv: Vec2{2.5, 3.3333333},
		},
		{
			name:    "neg",
			input:   Vec2{-5, -10},
			operand: Vec2{-2, -3},
			//
			wantAdd: Vec2{-7, -13},
			wantSub: Vec2{-3, -7},
			wantMul: Vec2{10, 30},
			wantDiv: Vec2{2.5, 3.3333333},
		},
		{
			name:    "cross",
			input:   Vec2{-5, 10},
			operand: Vec2{2, -3},
			//
			wantAdd: Vec2{-3, 7},
			wantSub: Vec2{-7, 13},
			wantMul: Vec2{-10, -30},
			wantDiv: Vec2{-2.5, -3.3333333},
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
