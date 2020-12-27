package kind

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Clusters struct {
}

func (c *Clusters) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("implement me")
}

func (c *Clusters) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("implement me")
}

func (c *Clusters) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("implement me")
}

func (c *Clusters) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("implement me")
}
