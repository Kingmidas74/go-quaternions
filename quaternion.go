package go_quaternions

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)

type Quaternion struct {
	W, I, J, K float64
}

var (
	AllComponentsEqualsToZeroError = errors.WithStack(errors.New("All components equals to zero"))
)

func NewQuaternionByCoords(W, I, J, K float64) *Quaternion {
	return &Quaternion{
		W: W,
		I: I,
		J: J,
		K: K,
	}
}

func (q *Quaternion) Add(arg *Quaternion) *Quaternion {
	return NewQuaternionByCoords(q.W+arg.W, q.I+arg.I, q.J+arg.J, q.K+arg.K)
}

func (q *Quaternion) Sub(arg *Quaternion) *Quaternion {
	return NewQuaternionByCoords(q.W-arg.W, q.I-arg.I, q.J-arg.J, q.K-arg.K)
}

func (q *Quaternion) MulByGrassmann(arg *Quaternion) *Quaternion {
	return NewQuaternionByCoords(q.W*arg.W-q.I*arg.I-q.J*arg.J-q.K*arg.K,
		q.W*arg.I+q.I*arg.W+q.J*arg.K-q.K*arg.J,
		q.W*arg.J-q.I*arg.K+q.J*arg.W+q.K*arg.I,
		q.W*arg.K+q.I*arg.J-q.J*arg.I+q.K*arg.W)
}

func (q *Quaternion) MulByEuclid(arg *Quaternion) *Quaternion {
	conjQ1 := q.Conjugate()

	return NewQuaternionByCoords(conjQ1.W*arg.W-conjQ1.I*arg.I-conjQ1.J*arg.J-conjQ1.K*arg.K,
		conjQ1.W*arg.I+conjQ1.I*arg.W+conjQ1.J*arg.K-conjQ1.K*arg.J,
		conjQ1.W*arg.J-conjQ1.I*arg.K+conjQ1.J*arg.W+conjQ1.K*arg.I,
		conjQ1.W*arg.K+conjQ1.I*arg.J-conjQ1.J*arg.I+conjQ1.K*arg.W)
}

func (q *Quaternion) MulScalar(arg *Quaternion) float64 {
	return q.Conjugate().MulByGrassmann(arg).Add(arg.Conjugate().MulByGrassmann(q)).W / 2
}

func (q *Quaternion) MulOuter(arg *Quaternion) float64 {
	return q.Conjugate().MulByGrassmann(arg).Sub(arg.Conjugate().MulByGrassmann(q)).W / 2
}

func (q *Quaternion) MulByNumber(n float64) *Quaternion {
	return q.MulByGrassmann(NewQuaternionByCoords(n, 0, 0, 0))
}

func (q *Quaternion) MulVector(arg *Quaternion) *Quaternion {
	d := q.MulByGrassmann(arg).Sub(arg.MulByGrassmann(q))

	return NewQuaternionByCoords(d.W/2, d.I/2, d.J/2, d.K/2)
}

func (q *Quaternion) Reverse() (*Quaternion, error) {
	conjugateQ := q.Conjugate()

	norm := q.Norm()
	if norm == 0 {
		return nil, AllComponentsEqualsToZeroError
	}

	return NewQuaternionByCoords(conjugateQ.W/norm, conjugateQ.I/norm, conjugateQ.J/norm, conjugateQ.K/norm), nil
}

func (q *Quaternion) Conjugate() *Quaternion {
	return NewQuaternionByCoords(q.W, q.I*-1, q.J*-1, q.K*-1)
}

func (q *Quaternion) Norm() float64 {
	return q.W*q.W + q.I*q.I + q.J*q.J + q.K*q.K
}

func (q *Quaternion) Equals(q2 *Quaternion, eps float64) bool {
	return q.equalsByCoords(q2, eps) && q.equalsByNorm(q2, eps)
}

func (q *Quaternion) Normalize() (*Quaternion, error) {
	norm := q.Norm()

	if norm == 0 {
		return nil, AllComponentsEqualsToZeroError
	}

	normSqrt := math.Sqrt(norm)

	return NewQuaternionByCoords(q.W/normSqrt, q.I/normSqrt, q.J/normSqrt, q.K/normSqrt), nil
}

func (q *Quaternion) ToRotateQuaternion(angle float64) (*Quaternion, error) {
	normQ, err := q.Normalize()
	if err != nil {
		return nil, err
	}

	normQ.W = math.Cos(angle / 2)
	normQ.I = normQ.I * math.Sin(angle/2)
	normQ.J = normQ.J * math.Sin(angle/2)
	normQ.K = normQ.K * math.Sin(angle/2)

	return normQ, nil
}

func (q *Quaternion) String() string {
	return fmt.Sprintf("(%v)+(%v)i+(%v)j+(%v)k", q.W, q.I, q.J, q.K)
}

func (q *Quaternion) equalsByCoords(q2 *Quaternion, eps float64) bool {
	sub := q.Sub(q2)
	return math.Abs(sub.W) < eps && math.Abs(sub.I) < eps && math.Abs(sub.J) < eps && math.Abs(sub.K) < eps
}

func (q *Quaternion) equalsByNorm(arg *Quaternion, eps float64) bool {
	return math.Abs(q.Norm()-arg.Norm()) < eps
}
