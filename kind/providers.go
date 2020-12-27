package kind

import (
	"github.com/angelokurtis/terraform-provider-kind/terraform"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	wire.Bind(new(terraform.KindCluster), new(*Clusters)),
	wire.Struct(new(Clusters), "*"),
)
