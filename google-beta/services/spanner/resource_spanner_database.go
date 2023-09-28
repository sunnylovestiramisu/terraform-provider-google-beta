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

package spanner

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

// customizeDiff func for additional checks on google_spanner_database properties:
func resourceSpannerDBDdlCustomDiffFunc(diff tpgresource.TerraformResourceDiff) error {
	old, new := diff.GetChange("ddl")
	oldDdls := old.([]interface{})
	newDdls := new.([]interface{})
	var err error

	if len(newDdls) < len(oldDdls) {
		err = diff.ForceNew("ddl")
		if err != nil {
			return fmt.Errorf("ForceNew failed for ddl, old - %v and new - %v", oldDdls, newDdls)
		}
		return nil
	}

	for i := range oldDdls {
		if newDdls[i].(string) != oldDdls[i].(string) {
			err = diff.ForceNew("ddl")
			if err != nil {
				return fmt.Errorf("ForceNew failed for ddl, old - %v and new - %v", oldDdls, newDdls)
			}
			return nil
		}
	}
	return nil
}

func resourceSpannerDBDdlCustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	// separate func to allow unit testing
	return resourceSpannerDBDdlCustomDiffFunc(diff)
}

func ValidateDatabaseRetentionPeriod(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	valueError := fmt.Errorf("version_retention_period should be in range [1h, 7d], in a format resembling 1d, 24h, 1440m, or 86400s")

	r := regexp.MustCompile("^(\\d{1}d|\\d{1,3}h|\\d{2,5}m|\\d{4,6}s)$")
	if !r.MatchString(value) {
		errors = append(errors, valueError)
		return
	}

	unit := value[len(value)-1:]
	multiple := value[:len(value)-1]
	num, err := strconv.Atoi(multiple)
	if err != nil {
		errors = append(errors, valueError)
		return
	}

	if unit == "d" && (num < 1 || num > 7) {
		errors = append(errors, valueError)
		return
	}
	if unit == "h" && (num < 1 || num > 7*24) {
		errors = append(errors, valueError)
		return
	}
	if unit == "m" && (num < 1*60 || num > 7*24*60) {
		errors = append(errors, valueError)
		return
	}
	if unit == "s" && (num < 1*60*60 || num > 7*24*60*60) {
		errors = append(errors, valueError)
		return
	}

	return
}

func resourceSpannerDBVirtualUpdate(d *schema.ResourceData, resourceSchema map[string]*schema.Schema) bool {
	// deletion_protection is the only virtual field
	if d.HasChange("deletion_protection") {
		for field := range resourceSchema {
			if field == "deletion_protection" {
				continue
			}
			if d.HasChange(field) {
				return false
			}
		}
		return true
	}
	return false
}

func ResourceSpannerDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpannerDatabaseCreate,
		Read:   resourceSpannerDatabaseRead,
		Update: resourceSpannerDatabaseUpdate,
		Delete: resourceSpannerDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSpannerDatabaseImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			resourceSpannerDBDdlCustomDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The instance to create the database on.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^[a-z][a-z0-9_\-]*[a-z0-9]$`),
				Description: `A unique identifier for the database, which cannot be changed after
the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].`,
			},
			"database_dialect": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GOOGLE_STANDARD_SQL", "POSTGRESQL", ""}),
				Description: `The dialect of the Cloud Spanner Database.
If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values: ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]`,
			},
			"ddl": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `An optional list of DDL statements to run inside the newly created
database. Statements can create tables, indexes, etc. These statements
execute atomically with the creation of the database: if there is an
error in any statement, the database is not created.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enable_drop_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether drop protection is enabled for this database. Defaults to false.
Drop protection is different from
the "deletion_protection" attribute in the following ways:
(1) "deletion_protection" only protects the database from deletions in Terraform.
whereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.
(2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the database.
"deletion_protection" attribute does not provide protection against the deletion of the parent instance.`,
				Default: false,
			},
			"encryption_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Encryption configuration for the database`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Fully qualified name of the KMS key to use to encrypt this database. This key must exist
in the same location as the Spanner Database.`,
						},
					},
				},
			},
			"version_retention_period": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: ValidateDatabaseRetentionPeriod,
				Description: `The retention period for the database. The retention period must be between 1 hour
and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
If this property is used, you must avoid adding new DDL statements to 'ddl' that
update the database's version_retention_period.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `An explanation of the status of the database.`,
			},
			"deletion_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: `Whether or not to allow Terraform to destroy the database. Defaults to true. Unless this field is set to false
in Terraform state, a 'terraform destroy' or 'terraform apply' that would delete the database will fail.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSpannerDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSpannerDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	versionRetentionPeriodProp, err := expandSpannerDatabaseVersionRetentionPeriod(d.Get("version_retention_period"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version_retention_period"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionRetentionPeriodProp)) && (ok || !reflect.DeepEqual(v, versionRetentionPeriodProp)) {
		obj["versionRetentionPeriod"] = versionRetentionPeriodProp
	}
	extraStatementsProp, err := expandSpannerDatabaseDdl(d.Get("ddl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ddl"); !tpgresource.IsEmptyValue(reflect.ValueOf(extraStatementsProp)) && (ok || !reflect.DeepEqual(v, extraStatementsProp)) {
		obj["extraStatements"] = extraStatementsProp
	}
	encryptionConfigProp, err := expandSpannerDatabaseEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	databaseDialectProp, err := expandSpannerDatabaseDatabaseDialect(d.Get("database_dialect"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_dialect"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseDialectProp)) && (ok || !reflect.DeepEqual(v, databaseDialectProp)) {
		obj["databaseDialect"] = databaseDialectProp
	}
	enableDropProtectionProp, err := expandSpannerDatabaseEnableDropProtection(d.Get("enable_drop_protection"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_drop_protection"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableDropProtectionProp)) && (ok || !reflect.DeepEqual(v, enableDropProtectionProp)) {
		obj["enableDropProtection"] = enableDropProtectionProp
	}
	instanceProp, err := expandSpannerDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	obj, err = resourceSpannerDatabaseEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Database: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating Database: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{instance}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = SpannerOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Database", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Database: %s", err)
	}

	opRes, err = resourceSpannerDatabaseDecoder(d, meta, opRes)
	if err != nil {
		return fmt.Errorf("Error decoding response from operation: %s", err)
	}
	if opRes == nil {
		return fmt.Errorf("Error decoding response from operation, could not find object")
	}

	if err := d.Set("name", flattenSpannerDatabaseName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{instance}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Note: Databases that are created with POSTGRESQL dialect do not support extra DDL
	// statements at the time of database creation. To avoid users needing to run
	// `terraform apply` twice to get their desired outcome, the provider does not set
	// `extraStatements` in the call to the `create` endpoint and all DDL (other than
	//  <CREATE DATABASE>) is run post-create, by calling the `updateDdl` endpoint

	_, ok := opRes["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}

	retention, retentionPeriodOk := d.GetOk("version_retention_period")
	retentionPeriod := retention.(string)
	ddl, ddlOk := d.GetOk("ddl")
	ddlStatements := ddl.([]interface{})

	if retentionPeriodOk || ddlOk {

		obj := make(map[string]interface{})
		updateDdls := []string{}

		if ddlOk {
			for i := 0; i < len(ddlStatements); i++ {
				if ddlStatements[i] != nil {
					updateDdls = append(updateDdls, ddlStatements[i].(string))
				}
			}
		}

		if retentionPeriodOk {
			dbName := d.Get("name")
			retentionDdl := fmt.Sprintf("ALTER DATABASE `%s` SET OPTIONS (version_retention_period=\"%s\")", dbName, retentionPeriod)
			if dialect, ok := d.GetOk("database_dialect"); ok && dialect == "POSTGRESQL" {
				retentionDdl = fmt.Sprintf("ALTER DATABASE \"%s\" SET spanner.version_retention_period TO \"%s\"", dbName, retentionPeriod)
			}
			updateDdls = append(updateDdls, retentionDdl)
		}

		// Skip API call if there are no new ddl entries (due to ignoring nil values)
		if len(updateDdls) > 0 {
			log.Printf("[DEBUG] Applying extra DDL statements to the new Database: %#v", updateDdls)

			obj["statements"] = updateDdls

			url, err = tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}/ddl")
			if err != nil {
				return err
			}

			res, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "PATCH",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: userAgent,
				Body:      obj,
				Timeout:   d.Timeout(schema.TimeoutUpdate),
			})
			if err != nil {
				return fmt.Errorf("Error executing DDL statements on Database: %s", err)
			}

			// Use the resource in the operation response to populate
			// identity fields and d.Id() before read
			var opRes map[string]interface{}
			err = SpannerOperationWaitTimeWithResponse(
				config, res, &opRes, project, "Creating Database", userAgent,
				d.Timeout(schema.TimeoutCreate))
			if err != nil {
				// The resource didn't actually create
				d.SetId("")
				return fmt.Errorf("Error waiting to run DDL against newly-created Database: %s", err)
			}
		}
	}

	enableDropProtection, enableDropProtectionOk := d.GetOk("enable_drop_protection")
	dropProtection := enableDropProtection.(bool)
	if enableDropProtectionOk && dropProtection {
		updateMask := []string{"enableDropProtection"}
		url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
		if err != nil {
			return err
		}
		// updateMask is a URL parameter but not present in the schema, so ReplaceVars
		// won't set it
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
		if err != nil {
			return err
		}
		obj := map[string]interface{}{"enableDropProtection": dropProtection}
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
			return fmt.Errorf("Error updating enableDropDatabaseProtection on Database: %s", err)
		} else {
			log.Printf("[DEBUG] Finished updating enableDropDatabaseProtection %q: %#v", d.Id(), res)
		}
	}

	log.Printf("[DEBUG] Finished creating Database %q: %#v", d.Id(), res)

	return resourceSpannerDatabaseRead(d, meta)
}

func resourceSpannerDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SpannerDatabase %q", d.Id()))
	}

	res, err = resourceSpannerDatabaseDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SpannerDatabase because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_protection"); !ok {
		if err := d.Set("deletion_protection", true); err != nil {
			return fmt.Errorf("Error setting deletion_protection: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	if err := d.Set("name", flattenSpannerDatabaseName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("version_retention_period", flattenSpannerDatabaseVersionRetentionPeriod(res["versionRetentionPeriod"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("state", flattenSpannerDatabaseState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("encryption_config", flattenSpannerDatabaseEncryptionConfig(res["encryptionConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("database_dialect", flattenSpannerDatabaseDatabaseDialect(res["databaseDialect"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("enable_drop_protection", flattenSpannerDatabaseEnableDropProtection(res["enableDropProtection"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("instance", flattenSpannerDatabaseInstance(res["instance"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	return nil
}

func resourceSpannerDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	enableDropProtectionProp, err := expandSpannerDatabaseEnableDropProtection(d.Get("enable_drop_protection"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_drop_protection"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableDropProtectionProp)) {
		obj["enableDropProtection"] = enableDropProtectionProp
	}

	obj, err = resourceSpannerDatabaseUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Database %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("enable_drop_protection") {
		updateMask = append(updateMask, "enableDropProtection")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	if obj["statements"] != nil {
		if len(obj["statements"].([]string)) == 0 {
			// Return early to avoid making an API call that errors,
			// due to containing no DDL SQL statements
			d.Partial(false)
			return resourceSpannerDatabaseRead(d, meta)
		}
	}

	if resourceSpannerDBVirtualUpdate(d, ResourceSpannerDatabase().Schema) {
		if d.Get("deletion_protection") != nil {
			if err := d.Set("deletion_protection", d.Get("deletion_protection")); err != nil {
				return fmt.Errorf("Error reading Instance: %s", err)
			}
		}
		return nil
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating Database %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Database %q: %#v", d.Id(), res)
		}

		err = SpannerOperationWaitTime(
			config, res, project, "Updating Database", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}
	d.Partial(true)

	if d.HasChange("version_retention_period") || d.HasChange("ddl") {
		obj := make(map[string]interface{})

		versionRetentionPeriodProp, err := expandSpannerDatabaseVersionRetentionPeriod(d.Get("version_retention_period"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("version_retention_period"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, versionRetentionPeriodProp)) {
			obj["versionRetentionPeriod"] = versionRetentionPeriodProp
		}
		extraStatementsProp, err := expandSpannerDatabaseDdl(d.Get("ddl"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ddl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, extraStatementsProp)) {
			obj["extraStatements"] = extraStatementsProp
		}

		obj, err = resourceSpannerDatabaseUpdateEncoder(d, meta, obj)
		if err != nil {
			return err
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}/ddl")
		if err != nil {
			return err
		}

		if obj["statements"] != nil {
			if len(obj["statements"].([]string)) == 0 {
				// Return early to avoid making an API call that errors,
				// due to containing no DDL SQL statements
				d.Partial(false)
				return resourceSpannerDatabaseRead(d, meta)
			}
		}

		if resourceSpannerDBVirtualUpdate(d, ResourceSpannerDatabase().Schema) {
			if d.Get("deletion_protection") != nil {
				if err := d.Set("deletion_protection", d.Get("deletion_protection")); err != nil {
					return fmt.Errorf("Error reading Instance: %s", err)
				}
			}
			return nil
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
			return fmt.Errorf("Error updating Database %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Database %q: %#v", d.Id(), res)
		}

		err = SpannerOperationWaitTime(
			config, res, project, "Updating Database", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceSpannerDatabaseRead(d, meta)
}

func resourceSpannerDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	if d.Get("deletion_protection").(bool) {
		return fmt.Errorf("cannot destroy instance without setting deletion_protection=false and running `terraform apply`")
	}
	log.Printf("[DEBUG] Deleting Database %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Database")
	}

	err = SpannerOperationWaitTime(
		config, res, project, "Deleting Database", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Database %q: %#v", d.Id(), res)
	return nil
}

func resourceSpannerDatabaseImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)$",
		"^instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<instance>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{instance}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_protection", true); err != nil {
		return nil, fmt.Errorf("Error setting deletion_protection: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSpannerDatabaseName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSpannerDatabaseVersionRetentionPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerDatabaseState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerDatabaseEncryptionConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenSpannerDatabaseEncryptionConfigKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenSpannerDatabaseEncryptionConfigKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerDatabaseDatabaseDialect(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerDatabaseEnableDropProtection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerDatabaseInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandSpannerDatabaseName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseVersionRetentionPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDdl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSpannerDatabaseEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSpannerDatabaseEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDatabaseDialect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseEnableDropProtection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("instances", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for instance: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceSpannerDatabaseEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["createStatement"] = fmt.Sprintf("CREATE DATABASE `%s`", obj["name"])
	if dialect, ok := obj["databaseDialect"]; ok && dialect == "POSTGRESQL" {
		obj["createStatement"] = fmt.Sprintf("CREATE DATABASE \"%s\"", obj["name"])
	}

	// Extra DDL statements are removed from the create request and instead applied to the database in
	// a post-create action, to accommodate retrictions when creating PostgreSQL-enabled databases.
	// https://cloud.google.com/spanner/docs/create-manage-databases#create_a_database
	log.Printf("[DEBUG] Preparing to create new Database. Any extra DDL statements will be applied to the Database in a separate API call")

	delete(obj, "name")
	delete(obj, "instance")

	delete(obj, "versionRetentionPeriod")
	delete(obj, "extraStatements")
	delete(obj, "enableDropProtection")
	return obj, nil
}

func resourceSpannerDatabaseUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	if obj["versionRetentionPeriod"] != nil || obj["extraStatements"] != nil {
		old, new := d.GetChange("ddl")
		oldDdls := old.([]interface{})
		newDdls := new.([]interface{})
		updateDdls := []string{}

		//Only new ddl statments to be add to update call
		for i := len(oldDdls); i < len(newDdls); i++ {
			if newDdls[i] != nil {
				updateDdls = append(updateDdls, newDdls[i].(string))
			}
		}

		//Add statement to update version_retention_period property, if needed
		if d.HasChange("version_retention_period") {
			dbName := d.Get("name")
			retentionDdl := fmt.Sprintf("ALTER DATABASE `%s` SET OPTIONS (version_retention_period=\"%s\")", dbName, obj["versionRetentionPeriod"])
			if dialect, ok := d.GetOk("database_dialect"); ok && dialect == "POSTGRESQL" {
				retentionDdl = fmt.Sprintf("ALTER DATABASE \"%s\" SET spanner.version_retention_period TO \"%s\"", dbName, obj["versionRetentionPeriod"])
			}
			updateDdls = append(updateDdls, retentionDdl)
		}

		obj["statements"] = updateDdls
		delete(obj, "name")
		delete(obj, "versionRetentionPeriod")
		delete(obj, "instance")
		delete(obj, "extraStatements")
	}
	return obj, nil
}

func resourceSpannerDatabaseDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	d.SetId(res["name"].(string))
	if err := tpgresource.ParseImportId([]string{"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}
	res["project"] = d.Get("project").(string)
	res["instance"] = d.Get("instance").(string)
	res["name"] = d.Get("name").(string)
	id, err := tpgresource.ReplaceVars(d, config, "{{instance}}/{{name}}")
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	return res, nil
}
