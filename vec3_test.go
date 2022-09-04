package go_quaternions

import (
	"math"
	"math/rand"
	"testing"
)

func TestVec3_RotateRad(t *testing.T) {
	tests := []struct {
		name  string
		eps   float64
		min   float64
		max   float64
		v     *Vec3
		angle float64
		err   bool
	}{
		{
			name: "rotate by X axis",
			v: &Vec3{
				X: 1,
				Y: 0,
				Z: 0,
			},
			angle: math.Pi,
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
		},
		{
			name: "rotate by Y axis",
			v: &Vec3{
				X: 0,
				Y: 1,
				Z: 0,
			},
			angle: math.Pi,
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
		},
		{
			name: "rotate by Z axis",
			v: &Vec3{
				X: 0,
				Y: 0,
				Z: 1,
			},
			angle: math.Pi,
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
		},
		{
			name: "rotate by union middle axis",
			v: &Vec3{
				X: 1,
				Y: 1,
				Z: 1,
			},
			angle: -2*math.Pi + rand.Float64()*(4*math.Pi),
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
		},
		{
			name: "rotate by zero",
			v: &Vec3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			angle: math.Pi,
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
			err:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			r, err := p.RotateRad(tt.v, tt.angle)

			if tt.err && err == nil {
				t.Errorf("Wrong result of rotate for %v around %v on angle %v", p, tt.v, tt.angle)
			}

			if !tt.err && math.Abs(p.Length()-r.Length()) > tt.eps {
				t.Errorf("Wrong result of rotate %v around %v. Norm of p = %v, result = %v", p, tt.v, p.Length(), r.Length())
			}
		})
	}
}

func TestVec3_StepTo(t *testing.T) {
	tests := []struct {
		name string
		eps  float64
		min  float64
		max  float64
		err  bool
	}{
		{
			name: "step into XY",
			eps:  EqualsEpsilon,
			min:  rand.Float64(),
			max:  rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			s := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			r, err := p.StepTo(s)

			if tt.err && err == nil {
				t.Errorf("Wrong result of step for %v on %v", p, s)
			}

			if !tt.err && !r.Sub(p).Equals(s, tt.eps) {
				t.Errorf("Wrong result of step for %v on %v", p, s)
			}
		})
	}
}

func TestVec3_RotateAroundPoint(t *testing.T) {
	tests := []struct {
		name  string
		eps   float64
		min   float64
		max   float64
		r     *Vec3
		angle float64
		err   bool
	}{
		{
			name:  "step around point",
			angle: math.Pi/0.1 + rand.Float64()*360,
			eps:   EqualsEpsilon,
			min:   rand.Float64(),
			max:   rand.Float64(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			s := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			r := &Vec3{
				X: tt.min + rand.Float64()*(tt.max-tt.min),
				Y: tt.min + rand.Float64()*(tt.max-tt.min),
				Z: tt.min + rand.Float64()*(tt.max-tt.min),
			}

			res, err := p.RotateAroundPoint(s, r, tt.angle)
			res, err = res.RotateAroundPoint(s, r, -tt.angle)

			if tt.err && err == nil {
				t.Errorf("Wrong result of step for %v on %v", p, s)
			}

			if !tt.err && res.Sub(p).Length() > tt.eps {
				t.Errorf("Wrong result of step for %v on %v", p, s)
			}
		})
	}
}
