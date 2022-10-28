package pgx

type TestType string

func (t TestType) String() string {
	return string(t)
}

const (
	TestTypeOne   TestType = "one"
	TestTypeTwo   TestType = "two"
	TestTypeThree TestType = "three"
)
