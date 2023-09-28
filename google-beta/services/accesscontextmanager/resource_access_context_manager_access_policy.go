// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package accesscontextmanager

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceAccessContextManagerAccessPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAccessPolicyCreate,
		Read:   resourceAccessContextManagerAccessPolicyRead,
		Update: resourceAccessContextManagerAccessPolicyUpdate,
		Delete: resourceAccessContextManagerAccessPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerAccessPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The parent of this AccessPolicy in the Cloud Resource Hierarchy.
Format: organizations/{organization_id}`,
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human readable title. Does not affect behavior.`,
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Folder or project on which this policy is applicable.
Format: folders/{{folder_id}} or projects/{{project_id}}`,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was created in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Resource name of the AccessPolicy. Format: {policy_id}`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerAccessPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAccessPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !tpgresource.IsEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	scopesProp, err := expandAccessContextManagerAccessPolicyScopes(d.Get("scopes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scopes"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopesProp)) && (ok || !reflect.DeepEqual(v, scopesProp)) {
		obj["scopes"] = scopesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessPolicy: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating AccessPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = AccessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating AccessPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AccessPolicy: %s", err)
	}

	if err := d.Set("name", flattenAccessContextManagerAccessPolicyName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// The operation for this resource contains the generated name that we need
	// in order to perform a READ. We need to access the object inside of it as
	// a map[string]interface, so let's do that.

	resp := res["response"].(map[string]interface{})
	name := tpgresource.GetResourceNameFromSelfLink(resp["name"].(string))
	log.Printf("[DEBUG] Setting AccessPolicy name, id to %s", name)
	if err := d.Set("name", name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)

	log.Printf("[DEBUG] Finished creating AccessPolicy %q: %#v", d.Id(), res)

	return resourceAccessContextManagerAccessPolicyRead(d, meta)
}

func resourceAccessContextManagerAccessPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAccessPolicy %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessContextManagerAccessPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("create_time", flattenAccessContextManagerAccessPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenAccessContextManagerAccessPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("parent", flattenAccessContextManagerAccessPolicyParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("title", flattenAccessContextManagerAccessPolicyTitle(res["title"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}
	if err := d.Set("scopes", flattenAccessContextManagerAccessPolicyScopes(res["scopes"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessPolicy: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAccessPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	scopesProp, err := expandAccessContextManagerAccessPolicyScopes(d.Get("scopes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scopes"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scopesProp)) {
		obj["scopes"] = scopesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AccessPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}

	if d.HasChange("scopes") {
		updateMask = append(updateMask, "scopes")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating AccessPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AccessPolicy %q: %#v", d.Id(), res)
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Updating AccessPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerAccessPolicyRead(d, meta)
}

func resourceAccessContextManagerAccessPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AccessPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AccessPolicy")
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Deleting AccessPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AccessPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerAccessPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerAccessPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenAccessContextManagerAccessPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyParent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAccessPolicyScopes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAccessContextManagerAccessPolicyParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyScopes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
