package config

import (
	"testing"

	upjetconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestStaticReferenceInjectorInjectReferences(t *testing.T) {
	injector := NewStaticReferenceInjector()

	resources := map[string]*upjetconfig.Resource{
		"resource": {
			TerraformResource: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"compartment_id": {Type: schema.TypeString},
					"bucket":         {Type: schema.TypeString},
					"name":           {Type: schema.TypeString},
				},
			},
			References: map[string]upjetconfig.Reference{},
		},
	}

	if err := injector.InjectReferences(resources); err != nil {
		t.Fatalf("InjectReferences returned error: %v", err)
	}

	res := resources["resource"]

	if ref, ok := res.References["compartment_id"]; !ok || ref.TerraformName != "oci_identity_compartment" {
		t.Fatalf("expected compartment_id to reference oci_identity_compartment, got %v", res.References["compartment_id"])
	}

	if ref, ok := res.References["bucket"]; !ok || ref.TerraformName != "oci_objectstorage_bucket" {
		t.Fatalf("expected bucket to reference oci_objectstorage_bucket, got %v", res.References["bucket"])
	}

	if _, ok := res.References["name"]; ok {
		t.Fatalf("did not expect reference for name attribute")
	}
}

func TestStaticReferenceInjectorDoesNotOverwriteWithoutFlag(t *testing.T) {
	injector := NewStaticReferenceInjector()

	custom := upjetconfig.Reference{TerraformName: "custom_resource"}

	resources := map[string]*upjetconfig.Resource{
		"resource": {
			TerraformResource: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"compartment_id": {Type: schema.TypeString},
				},
			},
			References: map[string]upjetconfig.Reference{
				"compartment_id": custom,
			},
		},
	}

	if err := injector.InjectReferences(resources); err != nil {
		t.Fatalf("InjectReferences returned error: %v", err)
	}

	if ref := resources["resource"].References["compartment_id"]; ref.TerraformName != custom.TerraformName {
		t.Fatalf("expected existing reference to remain, got %v", ref)
	}
}
