//+build wireinject

package main

import (
	"github.com/angelokurtis/terraform-provider-kind/kind"
	"github.com/angelokurtis/terraform-provider-kind/terraform"
	"github.com/google/wire"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InitKindProvider() *schema.Provider { // Wire's Injector
	wire.Build(
		kind.Providers,
		terraform.Providers,
	)
	return nil
}
