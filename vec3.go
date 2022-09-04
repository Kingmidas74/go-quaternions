package go_quaternions

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) Sub(arg *Vec3) *Vec3 {
	return &Vec3{
		X: v.X - arg.X,
		Y: v.Y - arg.Y,
		Z: v.Z - arg.Z,
	}
}

func (v *Vec3) RotateRad(axis *Vec3, angle float64) (*Vec3, error) {
	bSource := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, v.X, v.Y, v.Z),
	)

	bRotateReal, err := NewQuaternionByCoords(0, axis.X, axis.Y, axis.Z).ToRotateQuaternion(angle)
	if err != nil {
		return nil, err
	}

	bRotate := NewBQuaternion(
		bRotateReal,
		NewQuaternionByCoords(0, 0, 0, 0),
	)

	res := bRotate.Mul(bSource).Mul(bRotate.ComplexConjugate())

	return &Vec3{
		X: res.Q.I,
		Y: res.Q.J,
		Z: res.Q.K,
	}, nil
}

func (v *Vec3) StepTo(axis *Vec3) (*Vec3, error) {
	bSource := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, v.X, v.Y, v.Z),
	)

	bStepImage := NewQuaternionByCoords(0, axis.X/2, axis.Y/2, axis.Z/2)

	bStep := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		bStepImage,
	)

	res := bStep.Mul(bSource).Mul(bStep.Conjugate().ComplexConjugate())

	return &Vec3{
		X: res.Q.I,
		Y: res.Q.J,
		Z: res.Q.K,
	}, nil
}

func (v *Vec3) RotateAndStepTo(stepAxis *Vec3, rotateAxis *Vec3, angle float64) (*Vec3, error) {
	bSource := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, v.X, v.Y, v.Z),
	)

	bRotate, err := NewQuaternionByCoords(0, rotateAxis.X, rotateAxis.Y, rotateAxis.Z).ToRotateQuaternion(angle)
	if err != nil {
		return nil, err
	}

	bStep := NewQuaternionByCoords(0, stepAxis.X/2, stepAxis.Y/2, stepAxis.Z/2)

	bMove := NewBQuaternion(
		bRotate,
		bStep.MulByGrassmann(bRotate),
	)

	res := bMove.Mul(bSource).Mul(bMove.Conjugate().ComplexConjugate())

	return &Vec3{
		X: res.Q.I,
		Y: res.Q.J,
		Z: res.Q.K,
	}, nil
}

func (v *Vec3) StepAndRotateTo(stepAxis *Vec3, rotateAxis *Vec3, angle float64) (*Vec3, error) {
	bSource := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, v.X, v.Y, v.Z),
	)

	bRotate, err := NewQuaternionByCoords(0, rotateAxis.X, rotateAxis.Y, rotateAxis.Z).ToRotateQuaternion(angle)
	if err != nil {
		return nil, err
	}

	bStep := NewQuaternionByCoords(0, stepAxis.X/2, stepAxis.Y/2, stepAxis.Z/2)

	bMove := NewBQuaternion(
		bRotate,
		bRotate.MulByGrassmann(bStep),
	)

	res := bMove.Mul(bSource).Mul(bMove.Conjugate().ComplexConjugate())

	return &Vec3{
		X: res.Q.I,
		Y: res.Q.J,
		Z: res.Q.K,
	}, nil
}

func (v *Vec3) RotateAroundPoint(point *Vec3, rotateAxis *Vec3, angle float64) (*Vec3, error) {
	bSource := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, v.X, v.Y, v.Z),
	)

	bRotate, err := NewQuaternionByCoords(0, rotateAxis.X, rotateAxis.Y, rotateAxis.Z).ToRotateQuaternion(angle)
	if err != nil {
		return nil, err
	}

	t := NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, point.X/2, point.Y/2, point.Z/2),
	).Mul(NewBQuaternion(
		bRotate,
		NewQuaternionByCoords(0, 0, 0, 0),
	)).Mul(NewBQuaternion(
		NewQuaternionByCoords(1, 0, 0, 0),
		NewQuaternionByCoords(0, point.X/2, point.Y/2, point.Z/2).Conjugate(),
	))

	res := t.Mul(bSource).Mul(t.Conjugate().ComplexConjugate())

	return &Vec3{
		X: res.Q.I,
		Y: res.Q.J,
		Z: res.Q.K,
	}, nil

}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Equals(arg *Vec3, eps float64) bool {
	vr := v.Sub(arg)
	return math.Abs(vr.X) < eps && math.Abs(vr.Y) < eps && math.Abs(vr.Z) < eps
}

func (v *Vec3) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}
