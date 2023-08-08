package collector

import (
	"github.com/antonmedv/expr"
	"log"
)

type Collector struct {
	expr string
}

func NewCollector(
	expr string,
) *Collector {
	return &Collector{
		expr: expr,
	}
}

func (col Collector) IsCollectable(it Item) bool {
	if len(col.expr) == 0 {
		return true
	}

	env := map[string]interface{}{"log": it}
	program, err := expr.Compile(col.expr, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, _ := expr.Run(program, env)

	return output.(bool)
}

func (col Collector) Collect(it Item) {
	if !col.IsCollectable(it) {
		return
	}

	log.Println(it)
}
