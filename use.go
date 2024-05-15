package gormcls

func Use[MOD columnsInterface[CLS], CLS any](a MOD) (MOD, CLS) {
	return a, a.Columns()
}

type columnsInterface[CLS any] interface {
	Columns() CLS
}
