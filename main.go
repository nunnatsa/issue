package main

import (
	"github.com/nunnatsa/issue/cmd/cmda"
	"github.com/nunnatsa/issue/cmd/cmdb"
	"github.com/nunnatsa/issue/pkg/pkga"
	"github.com/nunnatsa/issue/pkg/pkgb"
	"github.com/nunnatsa/issue/pkg/pkgc"
	"github.com/nunnatsa/issue/pkg/pkgd"
	"github.com/nunnatsa/issue/pkg/pkge"
)

func main() {
	cmdb.B()
	pkga.A()
	pkgc.C()
	pkge.E()
	cmda.A()
	pkgd.D()
	pkgb.B()
}
