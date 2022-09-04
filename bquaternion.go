package go_quaternions

import "fmt"

type BQuaternion struct {
	P *Quaternion
	Q *Quaternion
}

func NewBQuaternion(p *Quaternion, q *Quaternion) *BQuaternion {
	return &BQuaternion{
		P: p,
		Q: q,
	}
}

func (bq *BQuaternion) Add(arg *BQuaternion) *BQuaternion {
	return NewBQuaternion(
		bq.P.Add(arg.P),
		bq.Q.Add(arg.Q),
	)
}

func (bq *BQuaternion) Sub(arg *BQuaternion) *BQuaternion {
	return NewBQuaternion(
		bq.P.Sub(arg.P),
		bq.Q.Sub(arg.Q),
	)
}

func (bq *BQuaternion) Mul(arg *BQuaternion) *BQuaternion {
	return NewBQuaternion(
		bq.P.MulByGrassmann(arg.P),
		bq.P.MulByGrassmann(arg.Q).Add(bq.Q.MulByGrassmann(arg.P)),
	)
}

func (bq *BQuaternion) Conjugate() *BQuaternion {
	return NewBQuaternion(
		bq.P,
		bq.Q.MulByNumber(-1),
	)
}

func (bq *BQuaternion) ComplexConjugate() *BQuaternion {
	return NewBQuaternion(
		bq.P.Conjugate(),
		bq.Q.Conjugate(),
	)
}

func (bq *BQuaternion) Equals(arg *BQuaternion, eps float64) bool {
	return bq.P.Equals(arg.P, eps) && bq.Q.Equals(arg.Q, eps)
}

func (bq *BQuaternion) String() string {
	return fmt.Sprintf(
		"{\n(%v)+(%v)i+(%v)j+(%v)k,\n(%v)+(%v)i+(%v)j+(%v)k\n}",
		bq.P.W, bq.P.I, bq.P.J, bq.P.K,
		bq.Q.W, bq.Q.I, bq.Q.J, bq.Q.K,
	)
}
