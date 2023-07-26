package entity

type Clause func(Prober) Expression

// Starter of a new expression.
type Prober interface {
	Field(name string) Field
	Group(Clause) Expression
	NotGroup(Clause) Expression
}

// Finished expression.
type Expression interface {
	And() Prober
	Or() Prober
}

// Field for logical comparison.
type Field interface {
	Is(operator string, value any) Expression
	IsLike(pattern string) Expression
	IsIn(values ...any) Expression
	IsBetween(start, end any) Expression
}

func init() {
	var _ Clause = func(p Prober) Expression {
		return p.Field("myField1").Is(OperatorEquals, "myValue").And().
			Field("myField2").IsBetween(0, 2).And().
			Group(func(s Prober) Expression {
				return s.Field("myField3").Is(OperatorLT, "2022-01-01").Or().
					Field("myField4").IsLike("prefix_%")
			}).And().
			Field("myField5").Is(OperatorGTE, -1)
	}
}
