/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	guardadhocquery "github.com/oracle/provider-oci/internal/controller/cloudguard/guardadhocquery"
	guardcloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/cloudguard/guardcloudguardconfiguration"
	guarddatamaskrule "github.com/oracle/provider-oci/internal/controller/cloudguard/guarddatamaskrule"
	guarddatasource "github.com/oracle/provider-oci/internal/controller/cloudguard/guarddatasource"
	guarddetectorrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guarddetectorrecipe"
	guardmanagedlist "github.com/oracle/provider-oci/internal/controller/cloudguard/guardmanagedlist"
	guardresponderrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guardresponderrecipe"
	guardsavedquery "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsavedquery"
	guardsecurityrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsecurityrecipe"
	guardsecurityzone "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsecurityzone"
	guardtarget "github.com/oracle/provider-oci/internal/controller/cloudguard/guardtarget"
	guardwlpagent "github.com/oracle/provider-oci/internal/controller/cloudguard/guardwlpagent"
)

// Setup_cloudguard creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudguard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		guardadhocquery.Setup,
		guardcloudguardconfiguration.Setup,
		guarddatamaskrule.Setup,
		guarddatasource.Setup,
		guarddetectorrecipe.Setup,
		guardmanagedlist.Setup,
		guardresponderrecipe.Setup,
		guardsavedquery.Setup,
		guardsecurityrecipe.Setup,
		guardsecurityzone.Setup,
		guardtarget.Setup,
		guardwlpagent.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
