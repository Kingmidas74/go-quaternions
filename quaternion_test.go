package go_quaternions

import (
	"math"
	"math/rand"
	"testing"
)

var EqualsEpsilon = math.Pow10(-15)

func TestQuaternion_Add_ShouldPassForPreparedValues(t *testing.T) {
	tests := []struct {
		name string
		q1   *Quaternion
		q2   *Quaternion
		want *Quaternion
		eps  float64
	}{
		{
			name: "add zero to zero",
			q1:   NewQuaternionByCoords(0, 0, 0, 0),
			q2:   NewQuaternionByCoords(0, 0, 0, 0),
			want: NewQuaternionByCoords(0, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
		{
			name: "add zero to one",
			q1:   NewQuaternionByCoords(0, 0, 0, 0),
			q2:   NewQuaternionByCoords(1, 0, 0, 0),
			want: NewQuaternionByCoords(1, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
		{
			name: "add one to zero",
			q1:   NewQuaternionByCoords(1, 0, 0, 0),
			q2:   NewQuaternionByCoords(0, 0, 0, 0),
			want: NewQuaternionByCoords(1, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
		{
			name: "add one to one",
			q1:   NewQuaternionByCoords(1, 0, 0, 0),
			q2:   NewQuaternionByCoords(1, 0, 0, 0),
			want: NewQuaternionByCoords(2, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
		{
			name: "add negative one to one",
			q1:   NewQuaternionByCoords(-1, 0, 0, 0),
			q2:   NewQuaternionByCoords(1, 0, 0, 0),
			want: NewQuaternionByCoords(0, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
		{
			name: "add negative one to negative one",
			q1:   NewQuaternionByCoords(-1, 0, 0, 0),
			q2:   NewQuaternionByCoords(-1, 0, 0, 0),
			want: NewQuaternionByCoords(-2, 0, 0, 0),
			eps:  EqualsEpsilon,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q1.Add(tt.q2); !got.Equals(tt.want, tt.eps) {
				t.Errorf("Wrong result of sum for %v and %v. Expected %v, got %v", tt.q1, tt.q2, tt.want, got)
			}
		})
	}
}

func TestQuaternion_Add_ShouldPassForRandomValuesByCommutative(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
	}{
		{
			name: "add random to random with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
		},
		{
			name: "add random to random with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
		{
			name: "add random to random with bound from -1 to 0",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  0,
		},
		{
			name: "add random to random with bound from -1 to 1",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  1,
		},
		{
			name: "add random to random with bound from random to random",
			eps:  EqualsEpsilon,
			min:  rand.Float64(),
			max:  rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q1 := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))
			q2 := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))
			if !q1.Add(q2).Equals(q2.Add(q1), tt.eps) {
				t.Errorf("Wrong result of sum for %v and %v", q1, q2)
			}
		})
	}
}

func TestQuaternion_MulByGrassmann_ShouldPassForRandomValuesByCommutativeConjugate(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
	}{
		{
			name: "mul random to random with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
		},
		{
			name: "mul random to random with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
		{
			name: "mul random to random with bound from -1 to 0",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  0,
		},
		{
			name: "mul random to random with bound from -1 to 1",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  1,
		},
		{
			name: "mul random to random with bound from random to random",
			eps:  EqualsEpsilon,
			min:  rand.Float64(),
			max:  rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q1 := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))
			q2 := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))

			if !q1.MulByGrassmann(q2).Conjugate().Equals(q2.Conjugate().MulByGrassmann(q1.Conjugate()), tt.eps) {
				t.Errorf("Wrong result of mul for %v and %v", q1, q2)
			}
		})
	}
}

func TestQuaternion_MulByGrassmann_ShouldPassForRandomValuesByReverse(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
		err  bool
	}{
		{
			name: "mul random to random with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
			err:  true,
		},
		{
			name: "mul random to random with bound 1",
			eps:  EqualsEpsilon,
			min:  1,
			max:  1,
		},
		{
			name: "mul random to random with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
		{
			name: "mul random to random with bound from -1 to 0",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  0,
		},
		{
			name: "mul random to random with bound from -1 to 1",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  1,
		},
		{
			name: "mul random to random with bound from random to random",
			eps:  EqualsEpsilon,
			min:  rand.Float64(),
			max:  rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))
			reverse, err := q.Reverse()
			if tt.err && err == nil {
				t.Errorf("Wrong result of mul for %v and %v", q, reverse)
			}
			if !tt.err && !q.MulByGrassmann(reverse).Equals(NewQuaternionByCoords(1, 0, 0, 0), tt.eps) {
				t.Errorf("Wrong result of mul for %v and %v", q, reverse)
			}
		})
	}
}

func TestQuaternion_MulScalar(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
		err  bool
	}{
		{
			name: "mul random to random with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
			err:  true,
		},
		{
			name: "mul random to random with bound 1",
			eps:  EqualsEpsilon,
			min:  1,
			max:  1,
		},
		{
			name: "mul random to random with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
		{
			name: "mul random to random with bound from -1 to 0",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  0,
		},
		{
			name: "mul random to random with bound from -1 to 1",
			eps:  EqualsEpsilon,
			min:  -1,
			max:  1,
		},
		{
			name: "mul random to random with bound from random to random",
			eps:  EqualsEpsilon,
			min:  rand.Float64(),
			max:  rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))
			q := NewQuaternionByCoords(0, 1, 0, 0)

			r := p.MulScalar(q)

			if !tt.err && r != p.I {
				t.Errorf("Wrong result of mul for %v and %v", q, r)
			}
		})
	}
}

func TestQuaternion_MulByNumber(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
		num  float64
		err  bool
	}{
		{
			name: "mul random to 0 with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
			num:  0,
			err:  true,
		},
		{
			name: "mul random to 1 with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
			num:  1,
		},
		{
			name: "mul random to -1 with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
			num:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min))

			r := p.MulByNumber(tt.num)

			if !tt.err && math.Abs(p.Norm()-r.Norm()) > tt.eps {
				t.Errorf("Wrong result of mul %v by %v. Norm of p = %v, result = %v", p, tt.num, p.Norm(), r.Norm())
			}
		})
	}
}

func TestQuaternion_MulConjugate(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
		err  bool
	}{
		{
			name: "mul random to 0 with bound 0",
			eps:  EqualsEpsilon,
			min:  0,
			max:  0,
			err:  true,
		},
		{
			name: "mul random to 1 with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
		{
			name: "mul random to -1 with bound from 0 to 1",
			eps:  EqualsEpsilon,
			min:  0,
			max:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewQuaternionByCoords(tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min), tt.min+rand.Float64()*(tt.max-tt.min)).Normalize()
			if tt.err && err == nil {
				t.Errorf("Wrong result of mul for %v by conjugate", p)
			}
			if tt.err && err != nil {
				return
			}
			r := p.MulByGrassmann(p.Conjugate())

			expectedResult := NewQuaternionByCoords(1, 0, 0, 0)

			if !tt.err && !r.Equals(expectedResult, tt.eps) {
				t.Errorf("Wrong result of mul %v by conjugate", p)
			}
		})
	}
}
