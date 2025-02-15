// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package cloudbuildv2_test

import (
	"context"
	"fmt"
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudbuildv2 "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccCloudbuildv2Connection_GheCompleteConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GheCompleteConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GheConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GheConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GheConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GhePrivConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GhePrivConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GhePrivUpdateConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GhePrivUpdateConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GhePrivUpdateConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GithubConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GithubConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GithubConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GitlabConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GitlabConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GleConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GleConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GleConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GleOldConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GleOldConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GleOldConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GlePrivConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GlePrivConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}
func TestAccCloudbuildv2Connection_GlePrivUpdateConnection(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudbuildv2ConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2Connection_GlePrivUpdateConnection(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
			{
				Config: testAccCloudbuildv2Connection_GlePrivUpdateConnectionUpdate0(context),
			},
			{
				ResourceName:            "google_cloudbuildv2_connection.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations"},
			},
		},
	})
}

func testAccCloudbuildv2Connection_GheCompleteConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri                      = "https://ghe.proctor-staging-test.com"
    app_id                        = 516
    app_installation_id           = 243
    app_slug                      = "myapp"
    private_key_secret_version    = "projects/gcb-terraform-creds/secrets/ghe-private-key/versions/latest"
    webhook_secret_secret_version = "projects/gcb-terraform-creds/secrets/ghe-webhook-secret/versions/latest"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GheConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri = "https://ghe.proctor-staging-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GheConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri                      = "https://ghe.proctor-staging-test.com"
    app_id                        = 516
    app_installation_id           = 243
    app_slug                      = "myapp"
    private_key_secret_version    = "projects/gcb-terraform-creds/secrets/ghe-private-key/versions/latest"
    webhook_secret_secret_version = "projects/gcb-terraform-creds/secrets/ghe-webhook-secret/versions/latest"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GhePrivConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri = "https://ghe.proctor-private-ca.com"

    service_directory_config {
      service = "projects/gcb-terraform-creds/locations/%{region}/namespaces/myns/services/serv"
    }

    ssl_ca = "-----BEGIN CERTIFICATE-----\nMIIEXTCCA0WgAwIBAgIUANaBCc9j/xdKJHU0sgmv6yE2WCIwDQYJKoZIhvcNAQEL\nBQAwLDEUMBIGA1UEChMLUHJvY3RvciBFbmcxFDASBgNVBAMTC1Byb2N0b3ItZW5n\nMB4XDTIxMDcxNTIwMDcwMloXDTIyMDcxNTIwMDcwMVowADCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAMVel7I88DkhwW445BNPBZvJNTV1AreHdz4um4U1\nop2+4L7JeNrUs5SRc0fzeOyOmA9ZzTDu9hBC7zj/sVNUy6cIQGCj32sr5SCAEIat\nnFZlzmVqJPT4J5NAaE37KO5347myTJEBrvpq8az4CtvX0yUzPK0gbUmaSaztVi4o\ndbJLKyv575xCLC/Hu6fIHBDH19eG1Ath9VpuAOkttRRoxu2VqijJZrGqaS+0o+OX\nrLi5HMtZbZjgQB4mc1g3ZDKX/gynxr+CDNaqNOqxuog33Tl5OcOk9DrR3MInaE7F\nyQFuH9mzF64AqOoTf7Tr/eAIz5XVt8K51nk+fSybEfKVwtMCAwEAAaOCAaEwggGd\nMA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQU/9dYyqMz\nv9rOMwPZcoIRMDAQCjAfBgNVHSMEGDAWgBTkQGTiCkLCmv/Awxdz5TAVRmyFfDCB\njQYIKwYBBQUHAQEEgYAwfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1j\nb250ZW50LTYxYWEyYzA5LTAwMDAtMjJjMi05ZjYyLWQ0ZjU0N2Y4MDIwMC5zdG9y\nYWdlLmdvb2dsZWFwaXMuY29tLzQxNGU4ZTJjZjU2ZWEyYzQxNmM0L2NhLmNydDAo\nBgNVHREBAf8EHjAcghpnaGUucHJvY3Rvci1wcml2YXRlLWNhLmNvbTCBggYDVR0f\nBHsweTB3oHWgc4ZxaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYxYWEyYzA5LTAw\nMDAtMjJjMi05ZjYyLWQ0ZjU0N2Y4MDIwMC5zdG9yYWdlLmdvb2dsZWFwaXMuY29t\nLzQxNGU4ZTJjZjU2ZWEyYzQxNmM0L2NybC5jcmwwDQYJKoZIhvcNAQELBQADggEB\nABo6BQLEZZ+YNiDuv2sRvcxSopQQb7fZjqIA9XOA35pNSKay2SncODnNvfsdRnOp\ncoy25sQSIzWyJ9zWl8DZ6evoOu5csZ2PoFqx5LsIq37w+ZcwD6DM8Zm7JqASxmxx\nGqTF0nHC4Aw8q8aJBeRD3PsSkfN5Q3DP3nTDnLyd0l+yPIkHUbZMoiFHX3BkhCng\nG96mYy/y3t16ghfV9lZkXpD/JK5aiN0bTHCDRc69owgfYiAcAqzBJ9gfZ90MBgzv\ngTTQel5dHg49SYXfnUpTy0HdQLEcoggOF8Q8V+xKdKa6eVbrvjJrkEJmvIQI5iCR\nhNvKR25mx8JUopqEXmONmqU=\n-----END CERTIFICATE-----\n\n-----BEGIN CERTIFICATE-----\nMIIDSDCCAjCgAwIBAgITMwWN+62nLcgyLa7p+jD1K90g6TANBgkqhkiG9w0BAQsF\nADAsMRQwEgYDVQQKEwtQcm9jdG9yIEVuZzEUMBIGA1UEAxMLUHJvY3Rvci1lbmcw\nHhcNMjEwNzEyMTM1OTQ0WhcNMzEwNzEwMTM1OTQzWjAsMRQwEgYDVQQKEwtQcm9j\ndG9yIEVuZzEUMBIGA1UEAxMLUHJvY3Rvci1lbmcwggEiMA0GCSqGSIb3DQEBAQUA\nA4IBDwAwggEKAoIBAQCYqJP5Qt90jIbld2dtuUV/zIkBFsTe4fapJfhBji03xBpN\nO1Yxj/jPSZ67Kdeoy0lEwvc2hL5FQGhIjLMR0mzOyN4fk/DZiA/4tAVi7hJyqpUC\n71JSwp7MwXL1b26CSE1MhcoCqA/E4iZxfJfF/ef4lhmC24UEmu8FEbldoy+6OysB\nRu7dGDwicW5F9h7eSkpGAsCRdJHh65iUx/IH0C4Ux2UZRDZdj6wVbuVu9tb938xF\nyRuVClONoLSn/lwdzeV7hQmBSm8qmfgbNPbYRaNLz3hOpsT+27aDQp2/pxue8hFJ\nd7We3+Lr5O4IL45PBwhVEAiFZqde6d4qViNEB2qTAgMBAAGjYzBhMA4GA1UdDwEB\n/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTkQGTiCkLCmv/Awxdz\n5TAVRmyFfDAfBgNVHSMEGDAWgBTkQGTiCkLCmv/Awxdz5TAVRmyFfDANBgkqhkiG\n9w0BAQsFAAOCAQEAfy5BJsWdx0oWWi7SFg9MbryWjBVPJl93UqACgG0Cgh813O/x\nlDZQhGO/ZFVhHz/WgooE/HgVNoVJTubKLLzz+zCkOB0wa3GMqJDyFjhFmUtd/3VM\nZh0ZQ+JWYsAiZW4VITj5xEn/d/B3xCFWGC1vhvhptEJ8Fo2cE1yM2pzk08NqFWoY\n4FaH0sbxWgyCKwTmtcYDbnx4FYuddryGCIxbYizqUK1dr4DGKeHonhm/d234Ew3x\n3vIBPoHMOfBec/coP1xAf5o+F+MRMO/sQ3tTGgyOH18lwsHo9SmXCrmOwVQPKrEw\nm+A+5TjXLmenyaBhqXa0vkAZYJhWdROhWC0VTA==\n-----END CERTIFICATE-----\n"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GhePrivUpdateConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri = "https://ghe.proctor-staging-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GhePrivUpdateConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"

  github_enterprise_config {
    host_uri = "https://ghe.proctor-private-ca.com"

    service_directory_config {
      service = "projects/gcb-terraform-creds/locations/%{region}/namespaces/myns/services/serv"
    }

    ssl_ca = "-----BEGIN CERTIFICATE-----\nMIIEXTCCA0WgAwIBAgIUANaBCc9j/xdKJHU0sgmv6yE2WCIwDQYJKoZIhvcNAQEL\nBQAwLDEUMBIGA1UEChMLUHJvY3RvciBFbmcxFDASBgNVBAMTC1Byb2N0b3ItZW5n\nMB4XDTIxMDcxNTIwMDcwMloXDTIyMDcxNTIwMDcwMVowADCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAMVel7I88DkhwW445BNPBZvJNTV1AreHdz4um4U1\nop2+4L7JeNrUs5SRc0fzeOyOmA9ZzTDu9hBC7zj/sVNUy6cIQGCj32sr5SCAEIat\nnFZlzmVqJPT4J5NAaE37KO5347myTJEBrvpq8az4CtvX0yUzPK0gbUmaSaztVi4o\ndbJLKyv575xCLC/Hu6fIHBDH19eG1Ath9VpuAOkttRRoxu2VqijJZrGqaS+0o+OX\nrLi5HMtZbZjgQB4mc1g3ZDKX/gynxr+CDNaqNOqxuog33Tl5OcOk9DrR3MInaE7F\nyQFuH9mzF64AqOoTf7Tr/eAIz5XVt8K51nk+fSybEfKVwtMCAwEAAaOCAaEwggGd\nMA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQU/9dYyqMz\nv9rOMwPZcoIRMDAQCjAfBgNVHSMEGDAWgBTkQGTiCkLCmv/Awxdz5TAVRmyFfDCB\njQYIKwYBBQUHAQEEgYAwfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1j\nb250ZW50LTYxYWEyYzA5LTAwMDAtMjJjMi05ZjYyLWQ0ZjU0N2Y4MDIwMC5zdG9y\nYWdlLmdvb2dsZWFwaXMuY29tLzQxNGU4ZTJjZjU2ZWEyYzQxNmM0L2NhLmNydDAo\nBgNVHREBAf8EHjAcghpnaGUucHJvY3Rvci1wcml2YXRlLWNhLmNvbTCBggYDVR0f\nBHsweTB3oHWgc4ZxaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYxYWEyYzA5LTAw\nMDAtMjJjMi05ZjYyLWQ0ZjU0N2Y4MDIwMC5zdG9yYWdlLmdvb2dsZWFwaXMuY29t\nLzQxNGU4ZTJjZjU2ZWEyYzQxNmM0L2NybC5jcmwwDQYJKoZIhvcNAQELBQADggEB\nABo6BQLEZZ+YNiDuv2sRvcxSopQQb7fZjqIA9XOA35pNSKay2SncODnNvfsdRnOp\ncoy25sQSIzWyJ9zWl8DZ6evoOu5csZ2PoFqx5LsIq37w+ZcwD6DM8Zm7JqASxmxx\nGqTF0nHC4Aw8q8aJBeRD3PsSkfN5Q3DP3nTDnLyd0l+yPIkHUbZMoiFHX3BkhCng\nG96mYy/y3t16ghfV9lZkXpD/JK5aiN0bTHCDRc69owgfYiAcAqzBJ9gfZ90MBgzv\ngTTQel5dHg49SYXfnUpTy0HdQLEcoggOF8Q8V+xKdKa6eVbrvjJrkEJmvIQI5iCR\nhNvKR25mx8JUopqEXmONmqU=\n-----END CERTIFICATE-----\n\n-----BEGIN CERTIFICATE-----\nMIIDSDCCAjCgAwIBAgITMwWN+62nLcgyLa7p+jD1K90g6TANBgkqhkiG9w0BAQsF\nADAsMRQwEgYDVQQKEwtQcm9jdG9yIEVuZzEUMBIGA1UEAxMLUHJvY3Rvci1lbmcw\nHhcNMjEwNzEyMTM1OTQ0WhcNMzEwNzEwMTM1OTQzWjAsMRQwEgYDVQQKEwtQcm9j\ndG9yIEVuZzEUMBIGA1UEAxMLUHJvY3Rvci1lbmcwggEiMA0GCSqGSIb3DQEBAQUA\nA4IBDwAwggEKAoIBAQCYqJP5Qt90jIbld2dtuUV/zIkBFsTe4fapJfhBji03xBpN\nO1Yxj/jPSZ67Kdeoy0lEwvc2hL5FQGhIjLMR0mzOyN4fk/DZiA/4tAVi7hJyqpUC\n71JSwp7MwXL1b26CSE1MhcoCqA/E4iZxfJfF/ef4lhmC24UEmu8FEbldoy+6OysB\nRu7dGDwicW5F9h7eSkpGAsCRdJHh65iUx/IH0C4Ux2UZRDZdj6wVbuVu9tb938xF\nyRuVClONoLSn/lwdzeV7hQmBSm8qmfgbNPbYRaNLz3hOpsT+27aDQp2/pxue8hFJ\nd7We3+Lr5O4IL45PBwhVEAiFZqde6d4qViNEB2qTAgMBAAGjYzBhMA4GA1UdDwEB\n/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTkQGTiCkLCmv/Awxdz\n5TAVRmyFfDAfBgNVHSMEGDAWgBTkQGTiCkLCmv/Awxdz5TAVRmyFfDANBgkqhkiG\n9w0BAQsFAAOCAQEAfy5BJsWdx0oWWi7SFg9MbryWjBVPJl93UqACgG0Cgh813O/x\nlDZQhGO/ZFVhHz/WgooE/HgVNoVJTubKLLzz+zCkOB0wa3GMqJDyFjhFmUtd/3VM\nZh0ZQ+JWYsAiZW4VITj5xEn/d/B3xCFWGC1vhvhptEJ8Fo2cE1yM2pzk08NqFWoY\n4FaH0sbxWgyCKwTmtcYDbnx4FYuddryGCIxbYizqUK1dr4DGKeHonhm/d234Ew3x\n3vIBPoHMOfBec/coP1xAf5o+F+MRMO/sQ3tTGgyOH18lwsHo9SmXCrmOwVQPKrEw\nm+A+5TjXLmenyaBhqXa0vkAZYJhWdROhWC0VTA==\n-----END CERTIFICATE-----\n"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GithubConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"
  disabled = true

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }

  project = "%{project_name}"

  annotations = {
    somekey = "somevalue"
  }
}


`, context)
}

func testAccCloudbuildv2Connection_GithubConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "%{region}"
  name     = "tf-test-connection%{random_suffix}"
  disabled = false

  github_config {
    app_installation_id = 31300675

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/latest"
    }
  }

  project = "%{project_name}"

  annotations = {
    otherkey = "othervalue"

    somekey = "somevalue"
  }
}


`, context)
}

func testAccCloudbuildv2Connection_GitlabConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gitlab-api-pat/versions/latest"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gitlab-read-pat/versions/latest"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GleConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-api-token/versions/latest"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-read-token/versions/latest"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-us-central1.gcb-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GleConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-old-api-token/versions/2"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-old-read-token/versions/3"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-old.gcb-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GleOldConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-old-api-token/versions/2"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-old-read-token/versions/3"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-old.gcb-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GleOldConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-api-token/versions/latest"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-read-token/versions/latest"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-us-central1.gcb-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GlePrivConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-private-api/versions/1"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-private-read-token/versions/1"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-us.gle-us-private.com"

    service_directory_config {
      service = "projects/407304063574/locations/us-west1/namespaces/private-conn/services/gitlab-private"
    }

    ssl_ca = "-----BEGIN CERTIFICATE-----\nMIIDcTCCAlmgAwIBAgIUbxJ3jxaRf5IPcUiQWRPRqpLL4s4wDQYJKoZIhvcNAQEL\nBQAwTjELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkdvb2dsZSBDbG91ZCBCdWlsZDEi\nMCAGA1UEAwwZZ2xlLXVzLmdsZS11cy1wcml2YXRlLmNvbTAeFw0yMzEwMjAxNjQ3\nNDJaFw0yNDEwMTkxNjQ3NDJaME4xCzAJBgNVBAYTAlVTMRswGQYDVQQKDBJHb29n\nbGUgQ2xvdWQgQnVpbGQxIjAgBgNVBAMMGWdsZS11cy5nbGUtdXMtcHJpdmF0ZS5j\nb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDRE84wKcuzc+beQ323\nIsNVVOF1+WZ975LvXpIt8Mw1bcYeYUcvgBXSXAByHGMtef8OBb9BUvLOVZdT3Xow\nCUbhCiK3zQy29pCn0rsneIvzGUXQgQRXK/Zap1N/hif4E7CIgjuvCN0Mn6BfDV/H\nXFm6EV3YUJrRPBr1rZik7doaYYwaJshCSTBtZxXZdvsG/OBuAXbJ9GB0B62EiTBz\n5g6yRdATut+PbgfzaWlPsgL3TTH+HPNCMO+ULnFupfZwRCtV+dJng76QYGs8fmFo\ntiWeElcsU8W7aqmjOkKRWcFsHpxPNXp8GG+jsZrVAnMOR3QeRLvowysSQD99IrGH\nAhwlAgMBAAGjRzBFMCQGA1UdEQQdMBuCGWdsZS11cy5nbGUtdXMtcHJpdmF0ZS5j\nb20wHQYDVR0OBBYEFKCIp5BruT4fpeDFQ2bKgdUvpfbWMA0GCSqGSIb3DQEBCwUA\nA4IBAQAQ4pUQmmd7eNIu9MQGna9lHYRFL0/G3mrK6Dcfm2As9qdcRw3dph8/iute\nxKDdBsnt6jDHrsjN0Na7Eq0040oBJrxG/NtqGX0zHNdpAT61bQ6j9reAT+KOrHys\nDJXH2lPuFW183AU7mmvcbXTEwkKex1i+DNoEdGYUbBfnWWeuhGzFog+/f9mtjoHL\nplcmx0VWHBQ5KO9Aq4OR/86DSg5QRPk76W9k3cM2ShXMm3TmTBZ+taJFfjZo5jP+\n7PLt2z9grSvFSXh2jnMyAs2yP4c+WezOXZLijqROr378AGBaksQK0CP+CYjZRpn1\nvndr1njLbvSjIypwKgZROb4P6XVa\n-----END CERTIFICATE-----\n"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GlePrivUpdateConnection(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-api-token/versions/latest"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-read-token/versions/latest"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-us-central1.gcb-test.com"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCloudbuildv2Connection_GlePrivUpdateConnectionUpdate0(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "primary" {
  location = "us-west1"
  name     = "tf-test-connection%{random_suffix}"

  gitlab_config {
    authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-private-api/versions/1"
    }

    read_authorizer_credential {
      user_token_secret_version = "projects/407304063574/secrets/gle-private-read-token/versions/1"
    }

    webhook_secret_secret_version = "projects/407304063574/secrets/gle-webhook-secret/versions/latest"
    host_uri                      = "https://gle-us.gle-us-private.com"

    service_directory_config {
      service = "projects/407304063574/locations/us-west1/namespaces/private-conn/services/gitlab-private"
    }

    ssl_ca = "-----BEGIN CERTIFICATE-----\nMIIDcTCCAlmgAwIBAgIUbxJ3jxaRf5IPcUiQWRPRqpLL4s4wDQYJKoZIhvcNAQEL\nBQAwTjELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkdvb2dsZSBDbG91ZCBCdWlsZDEi\nMCAGA1UEAwwZZ2xlLXVzLmdsZS11cy1wcml2YXRlLmNvbTAeFw0yMzEwMjAxNjQ3\nNDJaFw0yNDEwMTkxNjQ3NDJaME4xCzAJBgNVBAYTAlVTMRswGQYDVQQKDBJHb29n\nbGUgQ2xvdWQgQnVpbGQxIjAgBgNVBAMMGWdsZS11cy5nbGUtdXMtcHJpdmF0ZS5j\nb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDRE84wKcuzc+beQ323\nIsNVVOF1+WZ975LvXpIt8Mw1bcYeYUcvgBXSXAByHGMtef8OBb9BUvLOVZdT3Xow\nCUbhCiK3zQy29pCn0rsneIvzGUXQgQRXK/Zap1N/hif4E7CIgjuvCN0Mn6BfDV/H\nXFm6EV3YUJrRPBr1rZik7doaYYwaJshCSTBtZxXZdvsG/OBuAXbJ9GB0B62EiTBz\n5g6yRdATut+PbgfzaWlPsgL3TTH+HPNCMO+ULnFupfZwRCtV+dJng76QYGs8fmFo\ntiWeElcsU8W7aqmjOkKRWcFsHpxPNXp8GG+jsZrVAnMOR3QeRLvowysSQD99IrGH\nAhwlAgMBAAGjRzBFMCQGA1UdEQQdMBuCGWdsZS11cy5nbGUtdXMtcHJpdmF0ZS5j\nb20wHQYDVR0OBBYEFKCIp5BruT4fpeDFQ2bKgdUvpfbWMA0GCSqGSIb3DQEBCwUA\nA4IBAQAQ4pUQmmd7eNIu9MQGna9lHYRFL0/G3mrK6Dcfm2As9qdcRw3dph8/iute\nxKDdBsnt6jDHrsjN0Na7Eq0040oBJrxG/NtqGX0zHNdpAT61bQ6j9reAT+KOrHys\nDJXH2lPuFW183AU7mmvcbXTEwkKex1i+DNoEdGYUbBfnWWeuhGzFog+/f9mtjoHL\nplcmx0VWHBQ5KO9Aq4OR/86DSg5QRPk76W9k3cM2ShXMm3TmTBZ+taJFfjZo5jP+\n7PLt2z9grSvFSXh2jnMyAs2yP4c+WezOXZLijqROr378AGBaksQK0CP+CYjZRpn1\nvndr1njLbvSjIypwKgZROb4P6XVa\n-----END CERTIFICATE-----\n"
  }

  project     = "%{project_name}"
  annotations = {}
}


`, context)
}

func testAccCheckCloudbuildv2ConnectionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_cloudbuildv2_connection" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &cloudbuildv2.Connection{
				Location:    dcl.String(rs.Primary.Attributes["location"]),
				Name:        dcl.String(rs.Primary.Attributes["name"]),
				Disabled:    dcl.Bool(rs.Primary.Attributes["disabled"] == "true"),
				Project:     dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime:  dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				Etag:        dcl.StringOrNil(rs.Primary.Attributes["etag"]),
				Reconciling: dcl.Bool(rs.Primary.Attributes["reconciling"] == "true"),
				UpdateTime:  dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
			}

			client := transport_tpg.NewDCLCloudbuildv2Client(config, config.UserAgent, billingProject, 0)
			_, err := client.GetConnection(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_cloudbuildv2_connection still exists %v", obj)
			}
		}
		return nil
	}
}
