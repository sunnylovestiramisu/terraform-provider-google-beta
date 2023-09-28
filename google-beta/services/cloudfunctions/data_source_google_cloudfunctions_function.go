// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package cloudfunctions

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleCloudFunctionsFunction() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceCloudFunctionsFunction().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project", "region")

	return &schema.Resource{
		Read:   dataSourceGoogleCloudFunctionsFunctionRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleCloudFunctionsFunctionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}

	cloudFuncId := &CloudFunctionId{
		Project: project,
		Region:  region,
		Name:    d.Get("name").(string),
	}

	d.SetId(cloudFuncId.CloudFunctionId())

	err = resourceCloudFunctionsRead(d, meta)
	if err != nil {
		return err
	}

	if err := tpgresource.SetDataSourceLabels(d); err != nil {
		return err
	}

	if d.Id() == "" {
		return fmt.Errorf("%s not found", cloudFuncId.CloudFunctionId())
	}

	return nil
}
