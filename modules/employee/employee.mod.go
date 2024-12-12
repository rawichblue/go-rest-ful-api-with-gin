package employee

import (
	"github.com/uptrace/bun"
)

type EmployeeModule struct {
	Ctl *EmployeeController
	Svc *EmployeeService
}

func New(db *bun.DB) *EmployeeModule {
	svc := newService(db)
	return &EmployeeModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
