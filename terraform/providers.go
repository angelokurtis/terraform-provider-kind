package terraform

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewKindProvider,
)
