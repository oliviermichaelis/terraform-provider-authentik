package provider

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"goauthentik.io/api"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{}",
			},
		},
	}
}

func resourceGroupSchemaToModel(d *schema.ResourceData, c *APIClient) (*api.GroupRequest, diag.Diagnostics) {
	m := api.GroupRequest{
		Name:        d.Get("name").(string),
		IsSuperuser: boolToPointer(d.Get("is_superuser").(bool)),
	}

	if l, ok := d.Get("parent").(string); ok {
		m.Parent.Set(&l)
	}

	users := d.Get("users").([]interface{})
	m.Users = make([]int32, len(users))
	for i, prov := range users {
		m.Users[i] = int32(prov.(int))
	}

	attr := make(map[string]interface{})
	if l, ok := d.Get("attributes").(string); ok {
		if l != "" {
			err := json.NewDecoder(strings.NewReader(l)).Decode(&attr)
			if err != nil {
				return nil, diag.FromErr(err)
			}
		}
	}
	m.Attributes = &attr
	return &m, nil
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app, diags := resourceGroupSchemaToModel(d, c)
	if diags != nil {
		return diags
	}

	res, hr, err := c.client.CoreApi.CoreGroupsCreate(ctx).GroupRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.SetId(res.Pk)
	return resourceGroupRead(ctx, d, m)
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	res, hr, err := c.client.CoreApi.CoreGroupsRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.Set("name", res.Name)
	d.Set("is_superuser", res.IsSuperuser)
	b, err := json.Marshal(res.Attributes)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("attributes", string(b))
	d.Set("users", res.Users)
	return diags
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app, di := resourceGroupSchemaToModel(d, c)
	if di != nil {
		return di
	}
	res, hr, err := c.client.CoreApi.CoreGroupsUpdate(ctx, d.Id()).GroupRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.SetId(res.Pk)
	return resourceGroupRead(ctx, d, m)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)
	hr, err := c.client.CoreApi.CoreGroupsDestroy(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}
	return diag.Diagnostics{}
}
