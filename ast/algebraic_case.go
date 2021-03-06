package ast

import "github.com/raviqqe/stg/types"

// AlgebraicCase is an algebraic case expression.
type AlgebraicCase struct {
	abstractCase
	alternatives []AlgebraicAlternative
}

// NewAlgebraicCase creates an algebraic case expression.
func NewAlgebraicCase(e Expression, t types.Type, as []AlgebraicAlternative, a DefaultAlternative) AlgebraicCase {
	return AlgebraicCase{newAbstractCase(e, t, a), as}
}

// NewAlgebraicCaseWithoutDefault creates an algebraic case expression.
func NewAlgebraicCaseWithoutDefault(e Expression, t types.Type, as []AlgebraicAlternative) AlgebraicCase {
	return AlgebraicCase{newAbstractCase(e, t, DefaultAlternative{}), as}
}

// Alternatives returns alternatives.
func (c AlgebraicCase) Alternatives() []AlgebraicAlternative {
	return c.alternatives
}
