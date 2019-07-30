package types

type Type string

type HasType interface {
	Type() Type
}
