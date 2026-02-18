/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	api "github.com/oracle/provider-oci/internal/controller/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/apigateway/usageplan"
	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/autoscaling/autoscalingconfiguration"
	bootvolume "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroupbackup"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
	managementcabundle "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/managementcabundle"
	managementcertificate "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/managementcertificate"
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
	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/compute/shapemanagement"
	addon "github.com/oracle/provider-oci/internal/controller/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clustercompletecredentialrotationmanagement"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/containerengine/virtualnodepool"
	byoasn "github.com/oracle/provider-oci/internal/controller/core/byoasn"
	listingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/core/listingresourceversionagreement"
	virtualnetwork "github.com/oracle/provider-oci/internal/controller/core/virtualnetwork"
	applicationvip "github.com/oracle/provider-oci/internal/controller/database/applicationvip"
	autonomouscontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabase"
	autonomouscontainerdatabaseaddstandby "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabaseaddstandby"
	autonomouscontainerdatabasedataguardassociation "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardassociation"
	autonomouscontainerdatabasedataguardassociationoperation "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardassociationoperation"
	autonomouscontainerdatabasedataguardrolechange "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardrolechange"
	autonomouscontainerdatabasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasesnapshotstandby"
	autonomousdatabase "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabase"
	autonomousdatabasebackup "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasebackup"
	autonomousdatabaseinstancewalletmanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabaseinstancewalletmanagement"
	autonomousdatabaseregionalwalletmanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabaseregionalwalletmanagement"
	autonomousdatabasesaasadminuser "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasesaasadminuser"
	autonomousdatabasesoftwareimage "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasesoftwareimage"
	autonomousdatabasewallet "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasewallet"
	autonomousexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/autonomousexadatainfrastructure"
	autonomousvmcluster "github.com/oracle/provider-oci/internal/controller/database/autonomousvmcluster"
	autonomousvmclusterordscertificatemanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousvmclusterordscertificatemanagement"
	autonomousvmclustersslcertificatemanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousvmclustersslcertificatemanagement"
	backup "github.com/oracle/provider-oci/internal/controller/database/backup"
	backupcancelmanagement "github.com/oracle/provider-oci/internal/controller/database/backupcancelmanagement"
	backupdestination "github.com/oracle/provider-oci/internal/controller/database/backupdestination"
	cloudautonomousvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudautonomousvmcluster"
	clouddatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/clouddatabasemanagement"
	cloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/cloudexadatainfrastructure"
	cloudexadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/database/cloudexadatainfrastructureconfigureexascalemanagement"
	cloudvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudvmcluster"
	cloudvmclusteriormconfig "github.com/oracle/provider-oci/internal/controller/database/cloudvmclusteriormconfig"
	database "github.com/oracle/provider-oci/internal/controller/database/database"
	databasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/database/databasesnapshotstandby"
	databasesoftwareimage "github.com/oracle/provider-oci/internal/controller/database/databasesoftwareimage"
	databaseupgrade "github.com/oracle/provider-oci/internal/controller/database/databaseupgrade"
	dataguardassociation "github.com/oracle/provider-oci/internal/controller/database/dataguardassociation"
	dbhome "github.com/oracle/provider-oci/internal/controller/database/dbhome"
	dbnode "github.com/oracle/provider-oci/internal/controller/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsolehistory"
	dbnodesnapshot "github.com/oracle/provider-oci/internal/controller/database/dbnodesnapshot"
	dbnodesnapshotmanagement "github.com/oracle/provider-oci/internal/controller/database/dbnodesnapshotmanagement"
	dbsystem "github.com/oracle/provider-oci/internal/controller/database/dbsystem"
	dbsystemsupgrade "github.com/oracle/provider-oci/internal/controller/database/dbsystemsupgrade"
	exadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructure"
	exadatainfrastructurecompute "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructurecompute"
	exadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructureconfigureexascalemanagement"
	exadatainfrastructurestorage "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructurestorage"
	exadataiormconfig "github.com/oracle/provider-oci/internal/controller/database/exadataiormconfig"
	exadbvmcluster "github.com/oracle/provider-oci/internal/controller/database/exadbvmcluster"
	exascaledbstoragevault "github.com/oracle/provider-oci/internal/controller/database/exascaledbstoragevault"
	executionaction "github.com/oracle/provider-oci/internal/controller/database/executionaction"
	executionwindow "github.com/oracle/provider-oci/internal/controller/database/executionwindow"
	externalcontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabase"
	externalcontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabasemanagement"
	externalcontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabasesstackmonitoring"
	externaldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/externaldatabaseconnector"
	externalnoncontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabase"
	externalnoncontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabasemanagement"
	externalnoncontainerdatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabaseoperationsinsightsmanagement"
	externalnoncontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabasesstackmonitoring"
	externalpluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabase"
	externalpluggabledatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabasemanagement"
	externalpluggabledatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabaseoperationsinsightsmanagement"
	externalpluggabledatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabasesstackmonitoring"
	keystore "github.com/oracle/provider-oci/internal/controller/database/keystore"
	maintenancerun "github.com/oracle/provider-oci/internal/controller/database/maintenancerun"
	managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement"
	managementcloudasm "github.com/oracle/provider-oci/internal/controller/database/managementcloudasm"
	managementcloudasminstance "github.com/oracle/provider-oci/internal/controller/database/managementcloudasminstance"
	managementcloudcluster "github.com/oracle/provider-oci/internal/controller/database/managementcloudcluster"
	managementcloudclusterinstance "github.com/oracle/provider-oci/internal/controller/database/managementcloudclusterinstance"
	managementclouddbhome "github.com/oracle/provider-oci/internal/controller/database/managementclouddbhome"
	managementclouddbnode "github.com/oracle/provider-oci/internal/controller/database/managementclouddbnode"
	managementclouddbsystem "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystem"
	managementclouddbsystemclouddatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemclouddatabasemanagementsmanagement"
	managementclouddbsystemcloudstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemcloudstackmonitoringsmanagement"
	managementclouddbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemconnector"
	managementclouddbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemdiscovery"
	managementcloudlistener "github.com/oracle/provider-oci/internal/controller/database/managementcloudlistener"
	managementdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementdatabasedbmfeaturesmanagement"
	managementdbmanagementprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/managementdbmanagementprivateendpoint"
	managementexternalasm "github.com/oracle/provider-oci/internal/controller/database/managementexternalasm"
	managementexternalasminstance "github.com/oracle/provider-oci/internal/controller/database/managementexternalasminstance"
	managementexternalcluster "github.com/oracle/provider-oci/internal/controller/database/managementexternalcluster"
	managementexternalclusterinstance "github.com/oracle/provider-oci/internal/controller/database/managementexternalclusterinstance"
	managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement"
	managementexternaldbhome "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbhome"
	managementexternaldbnode "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbnode"
	managementexternaldbsystem "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystem"
	managementexternaldbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemconnector"
	managementexternaldbsystemdatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemdatabasemanagementsmanagement"
	managementexternaldbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemdiscovery"
	managementexternaldbsystemstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemstackmonitoringsmanagement"
	managementexternalexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatainfrastructure"
	managementexternalexadatainfrastructureexadatamanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatainfrastructureexadatamanagement"
	managementexternalexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastorageconnector"
	managementexternalexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastoragegrid"
	managementexternalexadatastorageserver "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastorageserver"
	managementexternallistener "github.com/oracle/provider-oci/internal/controller/database/managementexternallistener"
	managementexternalmysqldatabase "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabase"
	managementexternalmysqldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabaseconnector"
	managementexternalmysqldatabaseexternalmysqldatabasesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabaseexternalmysqldatabasesmanagement"
	managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement"
	managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement"
	managementmanageddatabase "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabase"
	managementmanageddatabasegroup "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabasegroup"
	managementmanageddatabaseschangedatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabaseschangedatabaseparameter"
	managementmanageddatabasesresetdatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabasesresetdatabaseparameter"
	managementnamedcredential "github.com/oracle/provider-oci/internal/controller/database/managementnamedcredential"
	managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement"
	migration "github.com/oracle/provider-oci/internal/controller/database/migration"
	migrationconnection "github.com/oracle/provider-oci/internal/controller/database/migrationconnection"
	migrationjob "github.com/oracle/provider-oci/internal/controller/database/migrationjob"
	migrationmigration "github.com/oracle/provider-oci/internal/controller/database/migrationmigration"
	oneoffpatch "github.com/oracle/provider-oci/internal/controller/database/oneoffpatch"
	pluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabase"
	pluggabledatabasepluggabledatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasepluggabledatabasemanagementsmanagement"
	pluggabledatabaseslocalclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabaseslocalclone"
	pluggabledatabasesnapshot "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasesnapshot"
	pluggabledatabasesremoteclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasesremoteclone"
	scheduledaction "github.com/oracle/provider-oci/internal/controller/database/scheduledaction"
	schedulingplan "github.com/oracle/provider-oci/internal/controller/database/schedulingplan"
	schedulingpolicy "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicy"
	schedulingpolicyschedulingwindow "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicyschedulingwindow"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsidentity "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsidentity"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsprivateendpoint"
	vmcluster "github.com/oracle/provider-oci/internal/controller/database/vmcluster"
	vmclusteraddvirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusteraddvirtualmachine"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/database/vmclusternetwork"
	vmclusterremovevirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusterremovevirtualmachine"
	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonestagednsseckeyversion"
	dkim "github.com/oracle/provider-oci/internal/controller/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/email/emaildomain"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/email/suppression"
	rule "github.com/oracle/provider-oci/internal/controller/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/filestorage/filesystem"
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
	storagefilesystemquotarule "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemquotarule"
	storagefilesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemsnapshotpolicy"
	storageoutboundconnector "github.com/oracle/provider-oci/internal/controller/filestorage/storageoutboundconnector"
	application "github.com/oracle/provider-oci/internal/controller/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/functions/invokefunction"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	apikey "github.com/oracle/provider-oci/internal/controller/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identity/customersecretkey"
	dataplanegeneratescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/identity/dataplanegeneratescopedaccesstoken"
	dbcredential "github.com/oracle/provider-oci/internal/controller/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/identity/domainreplicationtoregion"
	domainsaccountrecoverysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsaccountrecoverysetting"
	domainsapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsapikey"
	domainsapp "github.com/oracle/provider-oci/internal/controller/identity/domainsapp"
	domainsapprole "github.com/oracle/provider-oci/internal/controller/identity/domainsapprole"
	domainsapprovalworkflow "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflow"
	domainsapprovalworkflowassignment "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowassignment"
	domainsapprovalworkflowstep "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowstep"
	domainsauthenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsauthenticationfactorsetting"
	domainsauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsauthtoken"
	domainscloudgate "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgate"
	domainscloudgatemapping "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgatemapping"
	domainscloudgateserver "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgateserver"
	domainscondition "github.com/oracle/provider-oci/internal/controller/identity/domainscondition"
	domainscustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainscustomersecretkey"
	domainsdynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/identity/domainsdynamicresourcegroup"
	domainsgrant "github.com/oracle/provider-oci/internal/controller/identity/domainsgrant"
	domainsgroup "github.com/oracle/provider-oci/internal/controller/identity/domainsgroup"
	domainsidentitypropagationtrust "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitypropagationtrust"
	domainsidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainsidentityprovider"
	domainsidentitysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitysetting"
	domainskmsisetting "github.com/oracle/provider-oci/internal/controller/identity/domainskmsisetting"
	domainsmyapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsmyapikey"
	domainsmyauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsmyauthtoken"
	domainsmycustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainsmycustomersecretkey"
	domainsmyoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyoauth2clientcredential"
	domainsmyrequest "github.com/oracle/provider-oci/internal/controller/identity/domainsmyrequest"
	domainsmysmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmysmtpcredential"
	domainsmysupportaccount "github.com/oracle/provider-oci/internal/controller/identity/domainsmysupportaccount"
	domainsmyuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyuserdbcredential"
	domainsnetworkperimeter "github.com/oracle/provider-oci/internal/controller/identity/domainsnetworkperimeter"
	domainsnotificationsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsnotificationsetting"
	domainsoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsoauth2clientcredential"
	domainsoauthclientcertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthclientcertificate"
	domainsoauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthpartnercertificate"
	domainspasswordpolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspasswordpolicy"
	domainspolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspolicy"
	domainsrule "github.com/oracle/provider-oci/internal/controller/identity/domainsrule"
	domainssecurityquestion "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestion"
	domainssecurityquestionsetting "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestionsetting"
	domainsselfregistrationprofile "github.com/oracle/provider-oci/internal/controller/identity/domainsselfregistrationprofile"
	domainssetting "github.com/oracle/provider-oci/internal/controller/identity/domainssetting"
	domainssmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainssmtpcredential"
	domainssocialidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainssocialidentityprovider"
	domainsuser "github.com/oracle/provider-oci/internal/controller/identity/domainsuser"
	domainsuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsuserdbcredential"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/identity/usergroupmembership"
	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/kms/verify"
	backend "github.com/oracle/provider-oci/internal/controller/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/backendset"
	certificateloadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/loadbalancer/sslciphersuite"
	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/monitoring/vtap"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/mysql/mysqlbackup"
	mysqlchannel "github.com/oracle/provider-oci/internal/controller/mysql/mysqlchannel"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/mysql/mysqldbsystem"
	mysqlheatwavecluster "github.com/oracle/provider-oci/internal/controller/mysql/mysqlheatwavecluster"
	mysqlreplica "github.com/oracle/provider-oci/internal/controller/mysql/mysqlreplica"
	cpe "github.com/oracle/provider-oci/internal/controller/networkconnectivity/cpe"
	crossconnect "github.com/oracle/provider-oci/internal/controller/networkconnectivity/crossconnect"
	crossconnectgroup "github.com/oracle/provider-oci/internal/controller/networkconnectivity/crossconnectgroup"
	drg "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutetablerouterule"
	ipsec "github.com/oracle/provider-oci/internal/controller/networkconnectivity/ipsec"
	ipsecconnectiontunnelmanagement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/ipsecconnectiontunnelmanagement"
	virtualcircuit "github.com/oracle/provider-oci/internal/controller/networkconnectivity/virtualcircuit"
	networkfirewall "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicy"
	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/networking/vnicattachment"
	backendnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backend"
	backendsetnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backendset"
	listenernetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancersbackendsetsunified"
	configuration "github.com/oracle/provider-oci/internal/controller/nosql/configuration"
	index "github.com/oracle/provider-oci/internal/controller/nosql/index"
	table "github.com/oracle/provider-oci/internal/controller/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/nosql/tablereplica"
	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	namespacemetadata "github.com/oracle/provider-oci/internal/controller/objectstorage/namespacemetadata"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/objectstorage/preauthrequest"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/objectstorage/privateendpoint"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/replicationpolicy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/ons/subscription"
	managementhubevent "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubevent"
	managementhublifecycleenvironment "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecycleenvironment"
	managementhublifecyclestageattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestageattachmanagedinstancesmanagement"
	managementhublifecyclestagedetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagedetachmanagedinstancesmanagement"
	managementhublifecyclestagepromotesoftwaresourcemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagepromotesoftwaresourcemanagement"
	managementhublifecyclestagerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagerebootmanagement"
	managementhubmanagedinstance "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstance"
	managementhubmanagedinstanceattachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceattachprofilemanagement"
	managementhubmanagedinstancedetachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancedetachprofilemanagement"
	managementhubmanagedinstancegroup "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroup"
	managementhubmanagedinstancegroupattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupattachmanagedinstancesmanagement"
	managementhubmanagedinstancegroupattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupattachsoftwaresourcesmanagement"
	managementhubmanagedinstancegroupdetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupdetachmanagedinstancesmanagement"
	managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement"
	managementhubmanagedinstancegroupinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupinstallpackagesmanagement"
	managementhubmanagedinstancegroupinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupinstallwindowsupdatesmanagement"
	managementhubmanagedinstancegroupmanagemodulestreamsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupmanagemodulestreamsmanagement"
	managementhubmanagedinstancegrouprebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegrouprebootmanagement"
	managementhubmanagedinstancegroupremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupremovepackagesmanagement"
	managementhubmanagedinstancegroupupdateallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupupdateallpackagesmanagement"
	managementhubmanagedinstanceinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceinstallwindowsupdatesmanagement"
	managementhubmanagedinstancerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancerebootmanagement"
	managementhubmanagedinstanceupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceupdatepackagesmanagement"
	managementhubmanagementstation "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstation"
	managementhubmanagementstationassociatemanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationassociatemanagedinstancesmanagement"
	managementhubmanagementstationmirrorsynchronizemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationmirrorsynchronizemanagement"
	managementhubmanagementstationrefreshmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationrefreshmanagement"
	managementhubmanagementstationsynchronizemirrorsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationsynchronizemirrorsmanagement"
	managementhubprofile "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofile"
	managementhubprofileattachlifecyclestagemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachlifecyclestagemanagement"
	managementhubprofileattachmanagedinstancegroupmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachmanagedinstancegroupmanagement"
	managementhubprofileattachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachmanagementstationmanagement"
	managementhubprofileattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachsoftwaresourcesmanagement"
	managementhubprofiledetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofiledetachsoftwaresourcesmanagement"
	managementhubscheduledjob "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubscheduledjob"
	managementhubsoftwaresource "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresource"
	managementhubsoftwaresourceaddpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourceaddpackagesmanagement"
	managementhubsoftwaresourcechangeavailabilitymanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcechangeavailabilitymanagement"
	managementhubsoftwaresourcegeneratemetadatamanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcegeneratemetadatamanagement"
	managementhubsoftwaresourcemanifest "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcemanifest"
	managementhubsoftwaresourceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourceremovepackagesmanagement"
	managementhubsoftwaresourcereplacepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcereplacepackagesmanagement"
	managementhubworkrequestrerunmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubworkrequestrerunmanagement"
	providerconfig "github.com/oracle/provider-oci/internal/controller/providerconfig"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/psql/psqldbsystem"
	queue "github.com/oracle/provider-oci/internal/controller/queue/queue"
	ocicacheconfigset "github.com/oracle/provider-oci/internal/controller/redis/ocicacheconfigset"
	ocicacheconfigsetlistassociatedocicachecluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheconfigsetlistassociatedocicachecluster"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/redis/rediscluster"
	redisclusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclusterattachocicacheuser"
	redisclustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/redis/redisclustercreateidentitytoken"
	redisclusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclusterdetachocicacheuser"
	redisclustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclustergetocicacheuser"
	privateendpointresourcemanager "github.com/oracle/provider-oci/internal/controller/resourcemanager/privateendpoint"
	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
	secret "github.com/oracle/provider-oci/internal/controller/vault/secret"
	networkaddresslist "github.com/oracle/provider-oci/internal/controller/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/waf/webappfirewallpolicy"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		certificate.Setup,
		deployment.Setup,
		gateway.Setup,
		subscriber.Setup,
		usageplan.Setup,
		containerconfiguration.Setup,
		containerimagesignature.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
		autoscalingconfiguration.Setup,
		bootvolume.Setup,
		bootvolumebackup.Setup,
		volume.Setup,
		volumeattachment.Setup,
		volumebackup.Setup,
		volumebackuppolicy.Setup,
		volumebackuppolicyassignment.Setup,
		volumegroup.Setup,
		volumegroupbackup.Setup,
		certificateauthority.Setup,
		managementcabundle.Setup,
		managementcertificate.Setup,
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
		appcataloglistingresourceversionagreement.Setup,
		appcatalogsubscription.Setup,
		clusternetwork.Setup,
		computecapacityreport.Setup,
		computecapacityreservation.Setup,
		computecapacitytopology.Setup,
		computecluster.Setup,
		computegpumemorycluster.Setup,
		computegpumemoryfabric.Setup,
		computehost.Setup,
		computehostgroup.Setup,
		computeimagecapabilityschema.Setup,
		consolehistory.Setup,
		dedicatedvmhost.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		instanceconsoleconnection.Setup,
		instancemaintenanceevent.Setup,
		instancepool.Setup,
		instancepoolinstance.Setup,
		shapemanagement.Setup,
		addon.Setup,
		cluster.Setup,
		clustercompletecredentialrotationmanagement.Setup,
		clusterstartcredentialrotationmanagement.Setup,
		clusterworkloadmapping.Setup,
		nodepool.Setup,
		virtualnodepool.Setup,
		byoasn.Setup,
		listingresourceversionagreement.Setup,
		virtualnetwork.Setup,
		applicationvip.Setup,
		autonomouscontainerdatabase.Setup,
		autonomouscontainerdatabaseaddstandby.Setup,
		autonomouscontainerdatabasedataguardassociation.Setup,
		autonomouscontainerdatabasedataguardassociationoperation.Setup,
		autonomouscontainerdatabasedataguardrolechange.Setup,
		autonomouscontainerdatabasesnapshotstandby.Setup,
		autonomousdatabase.Setup,
		autonomousdatabasebackup.Setup,
		autonomousdatabaseinstancewalletmanagement.Setup,
		autonomousdatabaseregionalwalletmanagement.Setup,
		autonomousdatabasesaasadminuser.Setup,
		autonomousdatabasesoftwareimage.Setup,
		autonomousdatabasewallet.Setup,
		autonomousexadatainfrastructure.Setup,
		autonomousvmcluster.Setup,
		autonomousvmclusterordscertificatemanagement.Setup,
		autonomousvmclustersslcertificatemanagement.Setup,
		backup.Setup,
		backupcancelmanagement.Setup,
		backupdestination.Setup,
		cloudautonomousvmcluster.Setup,
		clouddatabasemanagement.Setup,
		cloudexadatainfrastructure.Setup,
		cloudexadatainfrastructureconfigureexascalemanagement.Setup,
		cloudvmcluster.Setup,
		cloudvmclusteriormconfig.Setup,
		database.Setup,
		databasesnapshotstandby.Setup,
		databasesoftwareimage.Setup,
		databaseupgrade.Setup,
		dataguardassociation.Setup,
		dbhome.Setup,
		dbnode.Setup,
		dbnodeconsoleconnection.Setup,
		dbnodeconsolehistory.Setup,
		dbnodesnapshot.Setup,
		dbnodesnapshotmanagement.Setup,
		dbsystem.Setup,
		dbsystemsupgrade.Setup,
		exadatainfrastructure.Setup,
		exadatainfrastructurecompute.Setup,
		exadatainfrastructureconfigureexascalemanagement.Setup,
		exadatainfrastructurestorage.Setup,
		exadataiormconfig.Setup,
		exadbvmcluster.Setup,
		exascaledbstoragevault.Setup,
		executionaction.Setup,
		executionwindow.Setup,
		externalcontainerdatabase.Setup,
		externalcontainerdatabasemanagement.Setup,
		externalcontainerdatabasesstackmonitoring.Setup,
		externaldatabaseconnector.Setup,
		externalnoncontainerdatabase.Setup,
		externalnoncontainerdatabasemanagement.Setup,
		externalnoncontainerdatabaseoperationsinsightsmanagement.Setup,
		externalnoncontainerdatabasesstackmonitoring.Setup,
		externalpluggabledatabase.Setup,
		externalpluggabledatabasemanagement.Setup,
		externalpluggabledatabaseoperationsinsightsmanagement.Setup,
		externalpluggabledatabasesstackmonitoring.Setup,
		keystore.Setup,
		maintenancerun.Setup,
		managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement.Setup,
		managementcloudasm.Setup,
		managementcloudasminstance.Setup,
		managementcloudcluster.Setup,
		managementcloudclusterinstance.Setup,
		managementclouddbhome.Setup,
		managementclouddbnode.Setup,
		managementclouddbsystem.Setup,
		managementclouddbsystemclouddatabasemanagementsmanagement.Setup,
		managementclouddbsystemcloudstackmonitoringsmanagement.Setup,
		managementclouddbsystemconnector.Setup,
		managementclouddbsystemdiscovery.Setup,
		managementcloudlistener.Setup,
		managementdatabasedbmfeaturesmanagement.Setup,
		managementdbmanagementprivateendpoint.Setup,
		managementexternalasm.Setup,
		managementexternalasminstance.Setup,
		managementexternalcluster.Setup,
		managementexternalclusterinstance.Setup,
		managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement.Setup,
		managementexternaldbhome.Setup,
		managementexternaldbnode.Setup,
		managementexternaldbsystem.Setup,
		managementexternaldbsystemconnector.Setup,
		managementexternaldbsystemdatabasemanagementsmanagement.Setup,
		managementexternaldbsystemdiscovery.Setup,
		managementexternaldbsystemstackmonitoringsmanagement.Setup,
		managementexternalexadatainfrastructure.Setup,
		managementexternalexadatainfrastructureexadatamanagement.Setup,
		managementexternalexadatastorageconnector.Setup,
		managementexternalexadatastoragegrid.Setup,
		managementexternalexadatastorageserver.Setup,
		managementexternallistener.Setup,
		managementexternalmysqldatabase.Setup,
		managementexternalmysqldatabaseconnector.Setup,
		managementexternalmysqldatabaseexternalmysqldatabasesmanagement.Setup,
		managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement.Setup,
		managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement.Setup,
		managementmanageddatabase.Setup,
		managementmanageddatabasegroup.Setup,
		managementmanageddatabaseschangedatabaseparameter.Setup,
		managementmanageddatabasesresetdatabaseparameter.Setup,
		managementnamedcredential.Setup,
		managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement.Setup,
		migration.Setup,
		migrationconnection.Setup,
		migrationjob.Setup,
		migrationmigration.Setup,
		oneoffpatch.Setup,
		pluggabledatabase.Setup,
		pluggabledatabasepluggabledatabasemanagementsmanagement.Setup,
		pluggabledatabaseslocalclone.Setup,
		pluggabledatabasesnapshot.Setup,
		pluggabledatabasesremoteclone.Setup,
		scheduledaction.Setup,
		schedulingplan.Setup,
		schedulingpolicy.Setup,
		schedulingpolicyschedulingwindow.Setup,
		toolsdatabasetoolsconnection.Setup,
		toolsdatabasetoolsidentity.Setup,
		toolsdatabasetoolsprivateendpoint.Setup,
		vmcluster.Setup,
		vmclusteraddvirtualmachine.Setup,
		vmclusternetwork.Setup,
		vmclusterremovevirtualmachine.Setup,
		actioncreatezonefromzonefile.Setup,
		record.Setup,
		resolver.Setup,
		resolverendpoint.Setup,
		rrset.Setup,
		steeringpolicy.Setup,
		steeringpolicyattachment.Setup,
		tsigkey.Setup,
		view.Setup,
		zone.Setup,
		zonepromotednsseckeyversion.Setup,
		zonestagednsseckeyversion.Setup,
		dkim.Setup,
		emaildomain.Setup,
		emailreturnpath.Setup,
		sender.Setup,
		suppression.Setup,
		rule.Setup,
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		mounttarget.Setup,
		replication.Setup,
		snapshot.Setup,
		storagefilesystemquotarule.Setup,
		storagefilesystemsnapshotpolicy.Setup,
		storageoutboundconnector.Setup,
		application.Setup,
		function.Setup,
		invokefunction.Setup,
		httpmonitor.Setup,
		pingmonitor.Setup,
		apikey.Setup,
		authenticationpolicy.Setup,
		authtoken.Setup,
		compartment.Setup,
		customersecretkey.Setup,
		dataplanegeneratescopedaccesstoken.Setup,
		dbcredential.Setup,
		domain.Setup,
		domainreplicationtoregion.Setup,
		domainsaccountrecoverysetting.Setup,
		domainsapikey.Setup,
		domainsapp.Setup,
		domainsapprole.Setup,
		domainsapprovalworkflow.Setup,
		domainsapprovalworkflowassignment.Setup,
		domainsapprovalworkflowstep.Setup,
		domainsauthenticationfactorsetting.Setup,
		domainsauthtoken.Setup,
		domainscloudgate.Setup,
		domainscloudgatemapping.Setup,
		domainscloudgateserver.Setup,
		domainscondition.Setup,
		domainscustomersecretkey.Setup,
		domainsdynamicresourcegroup.Setup,
		domainsgrant.Setup,
		domainsgroup.Setup,
		domainsidentitypropagationtrust.Setup,
		domainsidentityprovider.Setup,
		domainsidentitysetting.Setup,
		domainskmsisetting.Setup,
		domainsmyapikey.Setup,
		domainsmyauthtoken.Setup,
		domainsmycustomersecretkey.Setup,
		domainsmyoauth2clientcredential.Setup,
		domainsmyrequest.Setup,
		domainsmysmtpcredential.Setup,
		domainsmysupportaccount.Setup,
		domainsmyuserdbcredential.Setup,
		domainsnetworkperimeter.Setup,
		domainsnotificationsetting.Setup,
		domainsoauth2clientcredential.Setup,
		domainsoauthclientcertificate.Setup,
		domainsoauthpartnercertificate.Setup,
		domainspasswordpolicy.Setup,
		domainspolicy.Setup,
		domainsrule.Setup,
		domainssecurityquestion.Setup,
		domainssecurityquestionsetting.Setup,
		domainsselfregistrationprofile.Setup,
		domainssetting.Setup,
		domainssmtpcredential.Setup,
		domainssocialidentityprovider.Setup,
		domainsuser.Setup,
		domainsuserdbcredential.Setup,
		dynamicgroup.Setup,
		group.Setup,
		identityprovider.Setup,
		idpgroupmapping.Setup,
		importstandardtagsmanagement.Setup,
		networksource.Setup,
		policy.Setup,
		smtpcredential.Setup,
		tag.Setup,
		tagdefault.Setup,
		tagnamespace.Setup,
		uipassword.Setup,
		user.Setup,
		usercapabilitiesmanagement.Setup,
		usergroupmembership.Setup,
		ekmsprivateendpoint.Setup,
		encrypteddata.Setup,
		generatedkey.Setup,
		key.Setup,
		keyversion.Setup,
		sign.Setup,
		vault.Setup,
		vaultreplication.Setup,
		verify.Setup,
		backend.Setup,
		backendset.Setup,
		certificateloadbalancer.Setup,
		lbhostname.Setup,
		listener.Setup,
		loadbalancer.Setup,
		pathrouteset.Setup,
		routingpolicy.Setup,
		ruleset.Setup,
		sslciphersuite.Setup,
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
		alarm.Setup,
		alarmsuppression.Setup,
		capturefilter.Setup,
		vtap.Setup,
		mysqlbackup.Setup,
		mysqlchannel.Setup,
		mysqlconfiguration.Setup,
		mysqldbsystem.Setup,
		mysqlheatwavecluster.Setup,
		mysqlreplica.Setup,
		cpe.Setup,
		crossconnect.Setup,
		crossconnectgroup.Setup,
		drg.Setup,
		drgattachment.Setup,
		drgattachmentmanagement.Setup,
		drgattachmentslist.Setup,
		drgroutedistribution.Setup,
		drgroutedistributionstatement.Setup,
		drgroutetable.Setup,
		drgroutetablerouterule.Setup,
		ipsec.Setup,
		ipsecconnectiontunnelmanagement.Setup,
		virtualcircuit.Setup,
		networkfirewall.Setup,
		networkfirewallpolicy.Setup,
		defaultdhcpoptions.Setup,
		defaultroutetable.Setup,
		defaultsecuritylist.Setup,
		dhcpoptions.Setup,
		internetgateway.Setup,
		ipv6.Setup,
		localpeeringgateway.Setup,
		natgateway.Setup,
		networksecuritygroup.Setup,
		networksecuritygroupsecurityrule.Setup,
		privateip.Setup,
		publicip.Setup,
		publicippool.Setup,
		publicippoolcapacity.Setup,
		remotepeeringconnection.Setup,
		routetable.Setup,
		routetableattachment.Setup,
		securitylist.Setup,
		servicegateway.Setup,
		subnet.Setup,
		vcn.Setup,
		vlan.Setup,
		vnicattachment.Setup,
		backendnetworkloadbalancer.Setup,
		backendsetnetworkloadbalancer.Setup,
		listenernetworkloadbalancer.Setup,
		networkloadbalancer.Setup,
		networkloadbalancersbackendsetsunified.Setup,
		configuration.Setup,
		index.Setup,
		table.Setup,
		tablereplica.Setup,
		bucket.Setup,
		namespacemetadata.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		preauthrequest.Setup,
		privateendpoint.Setup,
		replicationpolicy.Setup,
		notificationtopic.Setup,
		subscription.Setup,
		managementhubevent.Setup,
		managementhublifecycleenvironment.Setup,
		managementhublifecyclestageattachmanagedinstancesmanagement.Setup,
		managementhublifecyclestagedetachmanagedinstancesmanagement.Setup,
		managementhublifecyclestagepromotesoftwaresourcemanagement.Setup,
		managementhublifecyclestagerebootmanagement.Setup,
		managementhubmanagedinstance.Setup,
		managementhubmanagedinstanceattachprofilemanagement.Setup,
		managementhubmanagedinstancedetachprofilemanagement.Setup,
		managementhubmanagedinstancegroup.Setup,
		managementhubmanagedinstancegroupattachmanagedinstancesmanagement.Setup,
		managementhubmanagedinstancegroupattachsoftwaresourcesmanagement.Setup,
		managementhubmanagedinstancegroupdetachmanagedinstancesmanagement.Setup,
		managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement.Setup,
		managementhubmanagedinstancegroupinstallpackagesmanagement.Setup,
		managementhubmanagedinstancegroupinstallwindowsupdatesmanagement.Setup,
		managementhubmanagedinstancegroupmanagemodulestreamsmanagement.Setup,
		managementhubmanagedinstancegrouprebootmanagement.Setup,
		managementhubmanagedinstancegroupremovepackagesmanagement.Setup,
		managementhubmanagedinstancegroupupdateallpackagesmanagement.Setup,
		managementhubmanagedinstanceinstallwindowsupdatesmanagement.Setup,
		managementhubmanagedinstancerebootmanagement.Setup,
		managementhubmanagedinstanceupdatepackagesmanagement.Setup,
		managementhubmanagementstation.Setup,
		managementhubmanagementstationassociatemanagedinstancesmanagement.Setup,
		managementhubmanagementstationmirrorsynchronizemanagement.Setup,
		managementhubmanagementstationrefreshmanagement.Setup,
		managementhubmanagementstationsynchronizemirrorsmanagement.Setup,
		managementhubprofile.Setup,
		managementhubprofileattachlifecyclestagemanagement.Setup,
		managementhubprofileattachmanagedinstancegroupmanagement.Setup,
		managementhubprofileattachmanagementstationmanagement.Setup,
		managementhubprofileattachsoftwaresourcesmanagement.Setup,
		managementhubprofiledetachsoftwaresourcesmanagement.Setup,
		managementhubscheduledjob.Setup,
		managementhubsoftwaresource.Setup,
		managementhubsoftwaresourceaddpackagesmanagement.Setup,
		managementhubsoftwaresourcechangeavailabilitymanagement.Setup,
		managementhubsoftwaresourcegeneratemetadatamanagement.Setup,
		managementhubsoftwaresourcemanifest.Setup,
		managementhubsoftwaresourceremovepackagesmanagement.Setup,
		managementhubsoftwaresourcereplacepackagesmanagement.Setup,
		managementhubworkrequestrerunmanagement.Setup,
		providerconfig.Setup,
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
		queue.Setup,
		ocicacheconfigset.Setup,
		ocicacheconfigsetlistassociatedocicachecluster.Setup,
		ocicacheuser.Setup,
		ocicacheusergetrediscluster.Setup,
		rediscluster.Setup,
		redisclusterattachocicacheuser.Setup,
		redisclustercreateidentitytoken.Setup,
		redisclusterdetachocicacheuser.Setup,
		redisclustergetocicacheuser.Setup,
		privateendpointresourcemanager.Setup,
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
		secret.Setup,
		networkaddresslist.Setup,
		webappfirewall.Setup,
		webappfirewallpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
