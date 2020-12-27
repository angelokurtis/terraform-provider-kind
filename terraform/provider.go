package terraform

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewKindProvider(c KindCluster) *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kind_cluster": {
				CreateContext: c.Create,
				ReadContext:   c.Read,
				UpdateContext: c.Update,
				DeleteContext: c.Delete,
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Description: "The kind name that is given to the created cluster.",
						Required:    true,
						ForceNew:    true,
					},
					"nodes": {
						Type: schema.TypeList,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"role":                   {Type: schema.TypeString},
								"image":                  {Type: schema.TypeString},
								"kubeadm_config_patches": {Type: schema.TypeMap},
								"extra_port_mappings":    {Type: schema.TypeList},
							},
						},
						Optional: true,
					},
				},
			},
		},
	}
}

type KindCluster interface {
	Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
	Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
	Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
	Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics
}
