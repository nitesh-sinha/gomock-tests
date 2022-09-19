package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks -source=doer.go

type Doer interface {
	DoSomething(int, string) error
}
