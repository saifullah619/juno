package block

import (
	"github.com/forbole/juno/v3/node"

	"github.com/forbole/juno/v3/logging"

	"github.com/forbole/juno/v3/database"
	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	node   node.Node
	db     database.Database
	logger logging.Logger
}

// NewModule builds a new Module instance
func NewModule(node node.Node, db database.Database, logger logging.Logger) *Module {
	return &Module{
		node:   node,
		db:     db,
		logger: logger,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "block"
}
