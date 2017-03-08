package gateway

//go:generate mockery -name IDGenerator -case underscore -outpkg mock -output ../idgen/mock
type IDGenerator interface {
	Generate() string
}
