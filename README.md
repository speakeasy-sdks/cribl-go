# cribl

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/cribl-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppscopeLibEntries.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppScopeLibEntries != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppscopeLibEntries](docs/sdks/appscopelibentries/README.md)

* [Get](docs/sdks/appscopelibentries/README.md#get) - Get a list of AppscopeLibEntry objects

### [AppscopeLibEntry](docs/sdks/appscopelibentry/README.md)

* [Create](docs/sdks/appscopelibentry/README.md#create) - Create AppscopeLibEntry
* [Delete](docs/sdks/appscopelibentry/README.md#delete) - Delete AppscopeLibEntry
* [Get](docs/sdks/appscopelibentry/README.md#get) - Get AppscopeLibEntry by ID
* [Update](docs/sdks/appscopelibentry/README.md#update) - Update AppscopeLibEntry

### [AuthToken](docs/sdks/authtoken/README.md)

* [Login](docs/sdks/authtoken/README.md#login) - Log in and obtain Auth token

### [AuthenticationSettings](docs/sdks/authenticationsettings/README.md)

* [Get](docs/sdks/authenticationsettings/README.md#get) - Get authentication settings
* [Update](docs/sdks/authenticationsettings/README.md#update) - Update authentication settings

### [Authorizations](docs/sdks/authorizations/README.md)

* [Get](docs/sdks/authorizations/README.md#get) - get the client's authorization policy

### [Branches](docs/sdks/branches/README.md)

* [Get](docs/sdks/branches/README.md#get) - get the list of branches

### [BulletinMessage](docs/sdks/bulletinmessage/README.md)

* [Create](docs/sdks/bulletinmessage/README.md#create) - Create BulletinMessage
* [Delete](docs/sdks/bulletinmessage/README.md#delete) - Delete BulletinMessage
* [Get](docs/sdks/bulletinmessage/README.md#get) - Get BulletinMessage by ID

### [BulletinMessages](docs/sdks/bulletinmessages/README.md)

* [Get](docs/sdks/bulletinmessages/README.md#get) - Get a list of BulletinMessage objects

### [Bytes](docs/sdks/bytes/README.md)

* [Get](docs/sdks/bytes/README.md#get) - Get some number of bytes from the file at the given path

### [CancelRunningGroup](docs/sdks/cancelrunninggroup/README.md)

* [Post](docs/sdks/cancelrunninggroup/README.md#post) - Cancel a running group upgrade

### [Certificate](docs/sdks/certificate/README.md)

* [Create](docs/sdks/certificate/README.md#create) - Create Certificate
* [Delete](docs/sdks/certificate/README.md#delete) - Delete Certificate
* [Get](docs/sdks/certificate/README.md#get) - Get Certificate by ID
* [Update](docs/sdks/certificate/README.md#update) - Update Certificate

### [Certificates](docs/sdks/certificates/README.md)

* [Get](docs/sdks/certificates/README.md#get) - Get a list of Certificate objects

### [ChangedFiles](docs/sdks/changedfiles/README.md)

* [Get](docs/sdks/changedfiles/README.md#get) - get the files changed

### [ChangelogViewState](docs/sdks/changelogviewstate/README.md)

* [Update](docs/sdks/changelogviewstate/README.md#update) - Update changelog viewed state

### [Changelogs](docs/sdks/changelogs/README.md)

* [Get](docs/sdks/changelogs/README.md#get) - Get changelog viewed state

### [ClientRoles](docs/sdks/clientroles/README.md)

* [Get](docs/sdks/clientroles/README.md#get) - get the client's roles

### [Cluis](docs/sdks/cluis/README.md)

* [Get](docs/sdks/cluis/README.md#get) - Get CLUI search results

### [Collector](docs/sdks/collector/README.md)

* [Get](docs/sdks/collector/README.md#get) - Get Collector by ID

### [CollectorObject](docs/sdks/collectorobject/README.md)

* [Get](docs/sdks/collectorobject/README.md#get) - Get a list of Collector objects

### [Commit](docs/sdks/commit/README.md)

* [Create](docs/sdks/commit/README.md#create) - create a new commit containing the current configs the given log message describing the changes.

### [Condition](docs/sdks/condition/README.md)

* [Get](docs/sdks/condition/README.md#get) - Get Condition by ID

### [Conditions](docs/sdks/conditions/README.md)

* [Get](docs/sdks/conditions/README.md#get) - Get a list of Condition objects

### [ConfigGroup](docs/sdks/configgroup/README.md)

* [Create](docs/sdks/configgroup/README.md#create) - Create ConfigGroup
* [Delete](docs/sdks/configgroup/README.md#delete) - Delete ConfigGroup
* [Get](docs/sdks/configgroup/README.md#get) - Get a specific ConfigGroup object
* [Update](docs/sdks/configgroup/README.md#update) - Update ConfigGroup

### [ConfiguredCollectors](docs/sdks/configuredcollectors/README.md)

* [Get](docs/sdks/configuredcollectors/README.md#get) - Get list of configured collectors

### [Container](docs/sdks/container/README.md)

* [Get](docs/sdks/container/README.md#get) - Get details for a single container on the edge host. Add stream=true to get a stream instead.

### [CountFile](docs/sdks/countfile/README.md)

* [Get](docs/sdks/countfile/README.md#get) - get the count of files of changed

### [CreatePipeline](docs/sdks/createpipeline/README.md)

* [Post](docs/sdks/createpipeline/README.md#post) - Create Pipeline

### [CriblMetadata](docs/sdks/criblmetadata/README.md)

* [Get](docs/sdks/criblmetadata/README.md#get) - Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider

### [CriblSystemSettings](docs/sdks/criblsystemsettings/README.md)

* [Get](docs/sdks/criblsystemsettings/README.md#get) - Get Cribl system settings
* [Update](docs/sdks/criblsystemsettings/README.md#update) - Update Cribl system settings

### [CriblsSettings](docs/sdks/criblssettings/README.md)

* [~~Get~~](docs/sdks/criblssettings/README.md#get) - Get Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/redis-cache-limits, /system/services-limits, /system/settings/git-settings, and /system/settings/conf respectively :warning: **Deprecated**
* [~~Update~~](docs/sdks/criblssettings/README.md#update) - Update Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/settings/git-settings, /system/settings/auth and /system/settings/conf respectively :warning: **Deprecated**

### [CurrentConfig](docs/sdks/currentconfig/README.md)

* [Push](docs/sdks/currentconfig/README.md#push) - push the current configs to the remote repository.

### [DataSample](docs/sdks/datasample/README.md)

* [Post](docs/sdks/datasample/README.md#post) - Create DataSample

### [DataSampleID](docs/sdks/datasampleid/README.md)

* [Delete](docs/sdks/datasampleid/README.md#delete) - Delete DataSample
* [Get](docs/sdks/datasampleid/README.md#get) - Get DataSample by ID
* [Update](docs/sdks/datasampleid/README.md#update) - Update DataSample

### [DatabaseConnection](docs/sdks/databaseconnection/README.md)

* [Post](docs/sdks/databaseconnection/README.md#post) - Create DatabaseConnectionConfig

### [DatabaseConnectionConfigID](docs/sdks/databaseconnectionconfigid/README.md)

* [Delete](docs/sdks/databaseconnectionconfigid/README.md#delete) - Delete DatabaseConnectionConfig
* [Get](docs/sdks/databaseconnectionconfigid/README.md#get) - Get DatabaseConnectionConfig by ID
* [Update](docs/sdks/databaseconnectionconfigid/README.md#update) - Update DatabaseConnectionConfig

### [Dataset](docs/sdks/dataset/README.md)

* [Create](docs/sdks/dataset/README.md#create) - Create DatasetProviderType
* [Delete](docs/sdks/dataset/README.md#delete) - Delete DatasetProviderType
* [Get](docs/sdks/dataset/README.md#get) - Get DatasetProviderType by ID
* [Update](docs/sdks/dataset/README.md#update) - Update DatasetProviderType

### [DatasetObject](docs/sdks/datasetobject/README.md)

* [Create](docs/sdks/datasetobject/README.md#create) - Create Dataset
* [Delete](docs/sdks/datasetobject/README.md#delete) - Delete Dataset
* [Get](docs/sdks/datasetobject/README.md#get) - Get Dataset by ID
* [Update](docs/sdks/datasetobject/README.md#update) - Update Dataset

### [DatasetObjects](docs/sdks/datasetobjects/README.md)

* [Get](docs/sdks/datasetobjects/README.md#get) - Get a list of Dataset objects

### [Datasets](docs/sdks/datasets/README.md)

* [Get](docs/sdks/datasets/README.md#get) - Get a list of DatasetProviderType objects

### [DestinationQueue](docs/sdks/destinationqueue/README.md)

* [Delete](docs/sdks/destinationqueue/README.md#delete) - Delete destination persistent queue

### [DiagBundle](docs/sdks/diagbundle/README.md)

* [Delete](docs/sdks/diagbundle/README.md#delete) - Remove diag bundle
* [Get](docs/sdks/diagbundle/README.md#get) - Returns a diag bundle as a tar.gz archive
* [Send](docs/sdks/diagbundle/README.md#send) - Send a diag bundle (tar.gz archive) to Cribl

### [DiagBundles](docs/sdks/diagbundles/README.md)

* [Get](docs/sdks/diagbundles/README.md#get) - Get list of existing diag bundles

### [DistributedDeployment](docs/sdks/distributeddeployment/README.md)

* [Get](docs/sdks/distributeddeployment/README.md#get) - Get summary of Distributed deployment

### [EdgeHostFiles](docs/sdks/edgehostfiles/README.md)

* [Get](docs/sdks/edgehostfiles/README.md#get) - Get details about a file on the edge host.

### [EdgeListing](docs/sdks/edgelisting/README.md)

* [Get](docs/sdks/edgelisting/README.md#get) - Get a directory listing of the given path

### [EventBreaker](docs/sdks/eventbreaker/README.md)

* [Delete](docs/sdks/eventbreaker/README.md#delete) - Delete Event Breaker Ruleset
* [Post](docs/sdks/eventbreaker/README.md#post) - Create Event Breaker Ruleset
* [Update](docs/sdks/eventbreaker/README.md#update) - Update Event Breaker Ruleset

### [EventBreakerID](docs/sdks/eventbreakerid/README.md)

* [Get](docs/sdks/eventbreakerid/README.md#get) - Get Event Breaker Ruleset by ID

### [EventBreakerOnData](docs/sdks/eventbreakerondata/README.md)

* [Post](docs/sdks/eventbreakerondata/README.md#post) - Runs an event breaker rule on the specified data

### [Events](docs/sdks/events/README.md)

* [Get](docs/sdks/events/README.md#get) - Get events generated by a specified source

### [ExecuteDistributedUpgrade](docs/sdks/executedistributedupgrade/README.md)

* [Post](docs/sdks/executedistributedupgrade/README.md#post) - Execute distributed group upgrade

### [ExecutorID](docs/sdks/executorid/README.md)

* [Get](docs/sdks/executorid/README.md#get) - Get Executor by ID

### [ExecutorObject](docs/sdks/executorobject/README.md)

* [Get](docs/sdks/executorobject/README.md#get) - Get a list of Executor objects

### [ExistingDiagBundles](docs/sdks/existingdiagbundles/README.md)

* [Get](docs/sdks/existingdiagbundles/README.md#get) - Get list of existing diag bundles

### [Feature](docs/sdks/feature/README.md)

* [Get](docs/sdks/feature/README.md#get) - Get feature by Id

### [Features](docs/sdks/features/README.md)

* [Get](docs/sdks/features/README.md#get) - List all features

### [FieldSummaries](docs/sdks/fieldsummaries/README.md)

* [Get](docs/sdks/fieldsummaries/README.md#get) - List field summaries

### [FleetMapping](docs/sdks/fleetmapping/README.md)

* [Create](docs/sdks/fleetmapping/README.md#create) - Create MappingRuleset

### [FleetMappings](docs/sdks/fleetmappings/README.md)

* [Get](docs/sdks/fleetmappings/README.md#get) - Get a list of MappingRuleset objects

### [FleetOrWorkerGroup](docs/sdks/fleetorworkergroup/README.md)

* [Deploy](docs/sdks/fleetorworkergroup/README.md#deploy) - Deploy commits for a Fleet or Worker Group

### [FunctionID](docs/sdks/functionid/README.md)

* [Get](docs/sdks/functionid/README.md#get) - Get Function by ID

### [GitSettings](docs/sdks/gitsettings/README.md)

* [Get](docs/sdks/gitsettings/README.md#get) - Get git settings
* [Update](docs/sdks/gitsettings/README.md#update) - Update git settings

### [GiveCriblVersion](docs/sdks/givecriblversion/README.md)

* [Post](docs/sdks/givecriblversion/README.md#post) - Upgrade Cribl to a given version

### [GlobalVariable](docs/sdks/globalvariable/README.md)

* [Post](docs/sdks/globalvariable/README.md#post) - Create Global Variable

### [GlobalVariableID](docs/sdks/globalvariableid/README.md)

* [Delete](docs/sdks/globalvariableid/README.md#delete) - Delete Global Variable
* [Get](docs/sdks/globalvariableid/README.md#get) - Get Global Variable by ID
* [Update](docs/sdks/globalvariableid/README.md#update) - Update Global Variable

### [GrokFile](docs/sdks/grokfile/README.md)

* [Create](docs/sdks/grokfile/README.md#create) - Create GrokFile
* [Delete](docs/sdks/grokfile/README.md#delete) - Delete GrokFile
* [Get](docs/sdks/grokfile/README.md#get) - Get GrokFile by ID
* [Update](docs/sdks/grokfile/README.md#update) - Update GrokFile

### [GrokFiles](docs/sdks/grokfiles/README.md)

* [Get](docs/sdks/grokfiles/README.md#get) - Get a list of GrokFile objects

### [GroupBundle](docs/sdks/groupbundle/README.md)

* [Get](docs/sdks/groupbundle/README.md#get) - Get effective bundle version for given Group

### [Groups](docs/sdks/groups/README.md)

* [Get](docs/sdks/groups/README.md#get) - Get a list of ConfigGroup objects

### [HealthInfo](docs/sdks/healthinfo/README.md)

* [Get](docs/sdks/healthinfo/README.md#get) - Provides health info for REST server

### [HostMetadataStructure](docs/sdks/hostmetadatastructure/README.md)

* [Get](docs/sdks/hostmetadatastructure/README.md#get) - Get the host's metadata structure

### [IDPAuth](docs/sdks/idpauth/README.md)

* [Get](docs/sdks/idpauth/README.md#get) - Get IDP used for an authorization code callback

### [IDPUserAuth](docs/sdks/idpuserauth/README.md)

* [Logout](docs/sdks/idpuserauth/README.md#logout) - Accepts a logout request from an IDP and logs out the user

### [InputID](docs/sdks/inputid/README.md)

* [Delete](docs/sdks/inputid/README.md#delete) - Delete Input
* [Get](docs/sdks/inputid/README.md#get) - Get Input by ID
* [Update](docs/sdks/inputid/README.md#update) - Update Input

### [InputObject](docs/sdks/inputobject/README.md)

* [Get](docs/sdks/inputobject/README.md#get) - Get a list of Input objects
* [Post](docs/sdks/inputobject/README.md#post) - Create Input

### [InputStatus](docs/sdks/inputstatus/README.md)

* [Get](docs/sdks/inputstatus/README.md#get) - Get a list of InputStatus objects

### [InputStatusID](docs/sdks/inputstatusid/README.md)

* [Get](docs/sdks/inputstatusid/README.md#get) - Get InputStatus by ID

### [InternalSystemMetrics](docs/sdks/internalsystemmetrics/README.md)

* [Post](docs/sdks/internalsystemmetrics/README.md#post) - Aggregate raw internal system metrics

### [JavascriptExpression](docs/sdks/javascriptexpression/README.md)

* [Post](docs/sdks/javascriptexpression/README.md#post) - Evaluate JavaScript expression

### [Job](docs/sdks/job/README.md)

* [Cancel](docs/sdks/job/README.md#cancel) - Cancel a job by instance id
* [Delete](docs/sdks/job/README.md#delete) - Remove job from job inspector by instance id
* [Get](docs/sdks/job/README.md#get) - Get job info by instance id
* [PauseJob](docs/sdks/job/README.md#pausejob) - Pause a job by instance id
* [Prevent](docs/sdks/job/README.md#prevent) - prevent job from being deleted automatically
* [Resume](docs/sdks/job/README.md#resume) - Resume a job by instance id
* [RunJob](docs/sdks/job/README.md#runjob) - Run or schedule a job

### [JobInfos](docs/sdks/jobinfos/README.md)

* [Get](docs/sdks/jobinfos/README.md#get) - Get info on jobs

### [JobResult](docs/sdks/jobresult/README.md)

* [Get](docs/sdks/jobresult/README.md#get) - Get results for a discover job by instance id

### [JobResults](docs/sdks/jobresults/README.md)

* [Get](docs/sdks/jobresults/README.md#get) - Get results for a discover job by instance id

### [JobStatus](docs/sdks/jobstatus/README.md)

* [Get](docs/sdks/jobstatus/README.md#get) - Get job status

### [KMSConfig](docs/sdks/kmsconfig/README.md)

* [Get](docs/sdks/kmsconfig/README.md#get) - Get Cribl KMS config
* [Update](docs/sdks/kmsconfig/README.md#update) - Update Cribl KMS config

### [KMSHealth](docs/sdks/kmshealth/README.md)

* [Get](docs/sdks/kmshealth/README.md#get) - Get Cribl KMS health

### [KeyMetadataEntities](docs/sdks/keymetadataentities/README.md)

* [Get](docs/sdks/keymetadataentities/README.md#get) - Get a list of KeyMetadataEntity objects

### [KeyMetadataEntity](docs/sdks/keymetadataentity/README.md)

* [Create](docs/sdks/keymetadataentity/README.md#create) - Create KeyMetadataEntity
* [Delete](docs/sdks/keymetadataentity/README.md#delete) - Delete KeyMetadataEntity
* [Get](docs/sdks/keymetadataentity/README.md#get) - Get KeyMetadataEntity by ID
* [Update](docs/sdks/keymetadataentity/README.md#update) - Update KeyMetadataEntity

### [LatestPQ](docs/sdks/latestpq/README.md)

* [Get](docs/sdks/latestpq/README.md#get) - Get status of latest clear PQ job for an output

### [License](docs/sdks/license/README.md)

* [Create](docs/sdks/license/README.md#create) - Create License
* [Delete](docs/sdks/license/README.md#delete) - Delete License
* [Get](docs/sdks/license/README.md#get) - Get License by ID

### [LicenseUsageMetrics](docs/sdks/licenseusagemetrics/README.md)

* [Get](docs/sdks/licenseusagemetrics/README.md#get) - Get license usage metrics, aggregated by day, up to last 90 days

### [Licenses](docs/sdks/licenses/README.md)

* [Get](docs/sdks/licenses/README.md#get) - Get a list of License objects

### [ListAuthGroup](docs/sdks/listauthgroup/README.md)

* [Get](docs/sdks/listauthgroup/README.md#get) - List the external authentication system's groups

### [ListContainerDetail](docs/sdks/listcontainerdetail/README.md)

* [Get](docs/sdks/listcontainerdetail/README.md#get) - Get a detailed list of containers running on the edge host.

### [ListCriblVersion](docs/sdks/listcriblversion/README.md)

* [Get](docs/sdks/listcriblversion/README.md#get) - Get a list of Cribl versions available for upgrade

### [ListDataSample](docs/sdks/listdatasample/README.md)

* [Get](docs/sdks/listdatasample/README.md#get) - Get a list of DataSample objects

### [ListDatabaseConnection](docs/sdks/listdatabaseconnection/README.md)

* [Get](docs/sdks/listdatabaseconnection/README.md#get) - Get a list of DatabaseConnection objects

### [ListEventBreaker](docs/sdks/listeventbreaker/README.md)

* [Get](docs/sdks/listeventbreaker/README.md#get) - Get a list of Event Breaker Ruleset objects

### [ListGlobalVariable](docs/sdks/listglobalvariable/README.md)

* [Get](docs/sdks/listglobalvariable/README.md#get) - Get a list of Global Variable objects

### [ListParser](docs/sdks/listparser/README.md)

* [Get](docs/sdks/listparser/README.md#get) - Get a list of Parser objects

### [ListProcessRunningDetail](docs/sdks/listprocessrunningdetail/README.md)

* [Get](docs/sdks/listprocessrunningdetail/README.md#get) - Get a detailed list of processes running on the edge host

### [ListSchema](docs/sdks/listschema/README.md)

* [Get](docs/sdks/listschema/README.md#get) - Get a list of Schema objects

### [LiveData](docs/sdks/livedata/README.md)

* [Post](docs/sdks/livedata/README.md#post) - Capture live incoming data

### [LogFileContent](docs/sdks/logfilecontent/README.md)

* [Get](docs/sdks/logfilecontent/README.md#get) - Get contents of the log file

### [LogFileContents](docs/sdks/logfilecontents/README.md)

* [Get](docs/sdks/logfilecontents/README.md#get) - Get contents of the log file

### [LogFileList](docs/sdks/logfilelist/README.md)

* [Get](docs/sdks/logfilelist/README.md#get) - list log files

### [LogFiles](docs/sdks/logfiles/README.md)

* [Get](docs/sdks/logfiles/README.md#get) - Get a list of log files

### [LogFilesContent](docs/sdks/logfilescontent/README.md)

* [Get](docs/sdks/logfilescontent/README.md#get) - Get contents of the log file

### [LogandTextual](docs/sdks/logandtextual/README.md)

* [Get](docs/sdks/logandtextual/README.md#get) - get the log message and textual diff for given commit

### [LoggerConfig](docs/sdks/loggerconfig/README.md)

* [Delete](docs/sdks/loggerconfig/README.md#delete) - Delete LoggerConfig
* [Get](docs/sdks/loggerconfig/README.md#get) - Get LoggerConfig by ID
* [Update](docs/sdks/loggerconfig/README.md#update) - Update LoggerConfig

### [LoggerConfigs](docs/sdks/loggerconfigs/README.md)

* [Get](docs/sdks/loggerconfigs/README.md#get) - Get a list of LoggerConfig objects

### [Lookup](docs/sdks/lookup/README.md)

* [Create](docs/sdks/lookup/README.md#create) - Create LookupFile
* [Delete](docs/sdks/lookup/README.md#delete) - Delete LookupFile
* [Get](docs/sdks/lookup/README.md#get) - Get LookupFile by ID
* [Update](docs/sdks/lookup/README.md#update) - Update LookupFile
* [Upload](docs/sdks/lookup/README.md#upload) - Upload LookupFile

### [Lookups](docs/sdks/lookups/README.md)

* [Get](docs/sdks/lookups/README.md#get) - Get a list of LookupFile objects

### [MappingRuleset](docs/sdks/mappingruleset/README.md)

* [Create](docs/sdks/mappingruleset/README.md#create) - Create MappingRuleset
* [Delete](docs/sdks/mappingruleset/README.md#delete) - Delete MappingRuleset
* [Get](docs/sdks/mappingruleset/README.md#get) - Get MappingRuleset by ID
* [Update](docs/sdks/mappingruleset/README.md#update) - Update MappingRuleset

### [MappingRulesetID](docs/sdks/mappingrulesetid/README.md)

* [Get](docs/sdks/mappingrulesetid/README.md#get) - Get MappingRuleset by ID

### [MappingRulesets](docs/sdks/mappingrulesets/README.md)

* [Delete](docs/sdks/mappingrulesets/README.md#delete) - Delete MappingRuleset
* [Get](docs/sdks/mappingrulesets/README.md#get) - Get a list of MappingRuleset objects
* [Update](docs/sdks/mappingrulesets/README.md#update) - Update MappingRuleset

### [MasterNodePackage](docs/sdks/masternodepackage/README.md)

* [Post](docs/sdks/masternodepackage/README.md#post) - Upgrade master node with the provided package

### [Metrics](docs/sdks/metrics/README.md)

* [Post](docs/sdks/metrics/README.md#post) - Enumerate all internal system metrics
* [Query](docs/sdks/metrics/README.md#query) - Query raw internal system metrics

### [NotificationTarget](docs/sdks/notificationtarget/README.md)

* [Create](docs/sdks/notificationtarget/README.md#create) - Create NotificationTarget
* [Delete](docs/sdks/notificationtarget/README.md#delete) - Delete NotificationTarget
* [Get](docs/sdks/notificationtarget/README.md#get) - Get NotificationTarget by ID
* [Update](docs/sdks/notificationtarget/README.md#update) - Update NotificationTarget

### [NotificationTargets](docs/sdks/notificationtargets/README.md)

* [Get](docs/sdks/notificationtargets/README.md#get) - Get a list of NotificationTarget objects

### [ObjectFunction](docs/sdks/objectfunction/README.md)

* [Get](docs/sdks/objectfunction/README.md#get) - Get a list of Function objects

### [OutputID](docs/sdks/outputid/README.md)

* [Delete](docs/sdks/outputid/README.md#delete) - Delete Output
* [Get](docs/sdks/outputid/README.md#get) - Get Output by ID
* [Update](docs/sdks/outputid/README.md#update) - Update Output

### [OutputObject](docs/sdks/outputobject/README.md)

* [Create](docs/sdks/outputobject/README.md#create) - Create Output

### [OutputObjects](docs/sdks/outputobjects/README.md)

* [Get](docs/sdks/outputobjects/README.md#get) - Get a list of Output objects

### [OutputStatus](docs/sdks/outputstatus/README.md)

* [Get](docs/sdks/outputstatus/README.md#get) - Get a list of OutputStatus objects

### [OutputStatusID](docs/sdks/outputstatusid/README.md)

* [Get](docs/sdks/outputstatusid/README.md#get) - Get OutputStatus by ID

### [Pack](docs/sdks/pack/README.md)

* [Clone](docs/sdks/pack/README.md#clone) - Clone Pack
* [Export](docs/sdks/pack/README.md#export) - Export Pack
* [Install](docs/sdks/pack/README.md#install) - Install Pack
* [Uninstall](docs/sdks/pack/README.md#uninstall) - Uninstall Pack from the system
* [Upgrade](docs/sdks/pack/README.md#upgrade) - Upgrade Pack
* [Upload](docs/sdks/pack/README.md#upload) - Upload Pack

### [Packs](docs/sdks/packs/README.md)

* [Get](docs/sdks/packs/README.md#get) - Get info on packs

### [ParserID](docs/sdks/parserid/README.md)

* [Delete](docs/sdks/parserid/README.md#delete) - Delete Parser
* [Get](docs/sdks/parserid/README.md#get) - Get Parser by ID
* [Update](docs/sdks/parserid/README.md#update) - Update Parser

### [ParserObject](docs/sdks/parserobject/README.md)

* [Post](docs/sdks/parserobject/README.md#post) - Create Parser

### [PipelineID](docs/sdks/pipelineid/README.md)

* [Delete](docs/sdks/pipelineid/README.md#delete) - Delete Pipeline
* [Get](docs/sdks/pipelineid/README.md#get) - Get Pipeline by ID
* [Update](docs/sdks/pipelineid/README.md#update) - Update Pipeline

### [PipelineObject](docs/sdks/pipelineobject/README.md)

* [Get](docs/sdks/pipelineobject/README.md#get) - Get a list of Pipeline objects

### [PolicyRule](docs/sdks/policyrule/README.md)

* [Create](docs/sdks/policyrule/README.md#create) - Create PolicyRule
* [Delete](docs/sdks/policyrule/README.md#delete) - Delete PolicyRule
* [Get](docs/sdks/policyrule/README.md#get) - Get PolicyRule by ID
* [Update](docs/sdks/policyrule/README.md#update) - Update PolicyRule

### [PolicyRules](docs/sdks/policyrules/README.md)

* [Get](docs/sdks/policyrules/README.md#get) - Get a list of PolicyRule objects

### [PreviousCriblPackage](docs/sdks/previouscriblpackage/README.md)

* [Get](docs/sdks/previouscriblpackage/README.md#get) - Get the previously downloaded Cribl package

### [ProcessRunningDetail](docs/sdks/processrunningdetail/README.md)

* [Get](docs/sdks/processrunningdetail/README.md#get) - Get details of a process running on the edge host

### [Processes](docs/sdks/processes/README.md)

* [Get](docs/sdks/processes/README.md#get) - Get a list of processes under management

### [Profiler](docs/sdks/profiler/README.md)

* [Create](docs/sdks/profiler/README.md#create) - Create ProfilerItem
* [Delete](docs/sdks/profiler/README.md#delete) - Delete ProfilerItem
* [Get](docs/sdks/profiler/README.md#get) - Get ProfilerItem by ID
* [Update](docs/sdks/profiler/README.md#update) - Update ProfilerItem

### [Profilers](docs/sdks/profilers/README.md)

* [Get](docs/sdks/profilers/README.md#get) - Get a list of ProfilerItem objects

### [QuerySnippet](docs/sdks/querysnippet/README.md)

* [Apply](docs/sdks/querysnippet/README.md#apply) - Applies a query snippet on a set of input events for preview

### [RedirectInfo](docs/sdks/redirectinfo/README.md)

* [Get](docs/sdks/redirectinfo/README.md#get) - Obtain redirect information

### [RedirectUserAuth](docs/sdks/redirectuserauth/README.md)

* [Logout](docs/sdks/redirectuserauth/README.md#logout) - Redirect user to IDP with logout request

### [RegexLibEntry](docs/sdks/regexlibentry/README.md)

* [Delete](docs/sdks/regexlibentry/README.md#delete) - Delete RegexLibEntry
* [Post](docs/sdks/regexlibentry/README.md#post) - Create RegexLibEntry
* [Update](docs/sdks/regexlibentry/README.md#update) - Update RegexLibEntry

### [RegexLibEntryID](docs/sdks/regexlibentryid/README.md)

* [Get](docs/sdks/regexlibentryid/README.md#get) - Get RegexLibEntry by ID

### [RegexLibEntryObject](docs/sdks/regexlibentryobject/README.md)

* [Get](docs/sdks/regexlibentryobject/README.md#get) - Get a list of RegexLibEntry objects

### [ReloadCriblSettings](docs/sdks/reloadcriblsettings/README.md)

* [Post](docs/sdks/reloadcriblsettings/README.md#post) - Reload Cribl settings from the filesystem

### [RemoteRepo](docs/sdks/remoterepo/README.md)

* [Sync](docs/sdks/remoterepo/README.md#sync) - syncs with remote repo via POST requests

### [RequestAuth](docs/sdks/requestauth/README.md)

* [Get](docs/sdks/requestauth/README.md#get) - Accepts an authentication request from an IDP and authenticates the user
* [Post](docs/sdks/requestauth/README.md#post) - API call that the IDP should use for an authentication request

### [RequestUserAuth](docs/sdks/requestuserauth/README.md)

* [Logout](docs/sdks/requestuserauth/README.md#logout) - API call that the IDP should use for a logout request

### [RestSecret](docs/sdks/restsecret/README.md)

* [Create](docs/sdks/restsecret/README.md#create) - Create RestSecret
* [Delete](docs/sdks/restsecret/README.md#delete) - Delete RestSecret
* [Get](docs/sdks/restsecret/README.md#get) - Get RestSecret by ID
* [Update](docs/sdks/restsecret/README.md#update) - Update RestSecret

### [RestSecrets](docs/sdks/restsecrets/README.md)

* [Get](docs/sdks/restsecrets/README.md#get) - Get a list of RestSecret objects

### [RestartCriblSettings](docs/sdks/restartcriblsettings/README.md)

* [Post](docs/sdks/restartcriblsettings/README.md#post) - Restart Cribl server

### [Role](docs/sdks/role/README.md)

* [Create](docs/sdks/role/README.md#create) - Create Role
* [Delete](docs/sdks/role/README.md#delete) - Delete Role
* [Get](docs/sdks/role/README.md#get) - Get Role by ID
* [Update](docs/sdks/role/README.md#update) - Update Role

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get a list of Role objects

### [RouteListID](docs/sdks/routelistid/README.md)

* [Get](docs/sdks/routelistid/README.md#get) - List all routes by id

### [RouteLists](docs/sdks/routelists/README.md)

* [Get](docs/sdks/routelists/README.md#get) - List all routes

### [RouteObject](docs/sdks/routeobject/README.md)

* [Update](docs/sdks/routeobject/README.md#update) - Add, delete or update the routes with the required content.

### [SampleContent](docs/sdks/samplecontent/README.md)

* [Get](docs/sdks/samplecontent/README.md#get) - Get sample content by ID

### [SampleEvents](docs/sdks/sampleevents/README.md)

* [Post](docs/sdks/sampleevents/README.md#post) - Sends sample events through a pipeline and returns the results

### [SampleOutput](docs/sdks/sampleoutput/README.md)

* [Post](docs/sdks/sampleoutput/README.md#post) - Send sample data to an output to validate configuration or test connectivity

### [SavedJob](docs/sdks/savedjob/README.md)

* [Delete](docs/sdks/savedjob/README.md#delete) - Delete SavedJob
* [Get](docs/sdks/savedjob/README.md#get) - Get SavedJob by ID
* [Update](docs/sdks/savedjob/README.md#update) - Update SavedJob

### [SavedJobs](docs/sdks/savedjobs/README.md)

* [Create](docs/sdks/savedjobs/README.md#create) - Create SavedJob
* [Get](docs/sdks/savedjobs/README.md#get) - Get a list of SavedJob objects

### [SavedQueries](docs/sdks/savedqueries/README.md)

* [Create](docs/sdks/savedqueries/README.md#create) - Create SavedQuery
* [Delete](docs/sdks/savedqueries/README.md#delete) - Delete SavedQuery
* [Get](docs/sdks/savedqueries/README.md#get) - Get a list of SavedQuery objects
* [Update](docs/sdks/savedqueries/README.md#update) - Update SavedQuery

### [SavedQuery](docs/sdks/savedquery/README.md)

* [Get](docs/sdks/savedquery/README.md#get) - Get SavedQuery by ID

### [Schema](docs/sdks/schema/README.md)

* [Create](docs/sdks/schema/README.md#create) - Create Schema
* [Delete](docs/sdks/schema/README.md#delete) - Delete Schema
* [Get](docs/sdks/schema/README.md#get) - Get Schema by ID
* [Post](docs/sdks/schema/README.md#post) - Create Schema
* [Update](docs/sdks/schema/README.md#update) - Update Schema

### [SchemaID](docs/sdks/schemaid/README.md)

* [Delete](docs/sdks/schemaid/README.md#delete) - Delete Schema
* [Get](docs/sdks/schemaid/README.md#get) - Get Schema by ID
* [Update](docs/sdks/schemaid/README.md#update) - Update Schema

### [Schemas](docs/sdks/schemas/README.md)

* [Get](docs/sdks/schemas/README.md#get) - Get a list of Schema objects

### [Script](docs/sdks/script/README.md)

* [Create](docs/sdks/script/README.md#create) - Create Script
* [Delete](docs/sdks/script/README.md#delete) - Delete Script
* [Get](docs/sdks/script/README.md#get) - Get Script by ID
* [Update](docs/sdks/script/README.md#update) - Update Script

### [Scripts](docs/sdks/scripts/README.md)

* [Get](docs/sdks/scripts/README.md#get) - Get a list of Script objects

### [Search](docs/sdks/search/README.md)

* [DispatchSearch](docs/sdks/search/README.md#dispatchsearch) - Dispatch search *id* to worker nodes filtered by worker node filter using RPC

### [SearchDoc](docs/sdks/searchdoc/README.md)

* [Get](docs/sdks/searchdoc/README.md#get) - Get Search documentation

### [SearchJob](docs/sdks/searchjob/README.md)

* [Create](docs/sdks/searchjob/README.md#create) - Create SearchJob
* [Delete](docs/sdks/searchjob/README.md#delete) - Delete SearchJob
* [Get](docs/sdks/searchjob/README.md#get) - Get SearchJob by ID
* [Update](docs/sdks/searchjob/README.md#update) - Update SearchJob

### [SearchJobMetrics](docs/sdks/searchjobmetrics/README.md)

* [Get](docs/sdks/searchjobmetrics/README.md#get) - Get search job metrics

### [SearchJobs](docs/sdks/searchjobs/README.md)

* [Get](docs/sdks/searchjobs/README.md#get) - Get a list of SearchJob objects

### [SearchLimits](docs/sdks/searchlimits/README.md)

* [Get](docs/sdks/searchlimits/README.md#get) - Get search limits

### [SearchLogs](docs/sdks/searchlogs/README.md)

* [Get](docs/sdks/searchlogs/README.md#get) - Get search logs

### [SearchTimeline](docs/sdks/searchtimeline/README.md)

* [Get](docs/sdks/searchtimeline/README.md#get) - Get search timeline

### [SpecifiedOutput](docs/sdks/specifiedoutput/README.md)

* [Get](docs/sdks/specifiedoutput/README.md#get) - Get samples data for the specified output. Used to get sample data for the test action.

### [StageDistributedPackage](docs/sdks/stagedistributedpackage/README.md)

* [Post](docs/sdks/stagedistributedpackage/README.md#post) - Stage distributed group upgrade

### [SystemInfo](docs/sdks/systeminfo/README.md)

* [Get](docs/sdks/systeminfo/README.md#get) - Get basic system information

### [TaskError](docs/sdks/taskerror/README.md)

* [Get](docs/sdks/taskerror/README.md#get) - Get Task errors for a job by id

### [TaskErrors](docs/sdks/taskerrors/README.md)

* [Get](docs/sdks/taskerrors/README.md#get) - Get Task errors for a job by id

### [TestDatabaseConnection](docs/sdks/testdatabaseconnection/README.md)

* [Post](docs/sdks/testdatabaseconnection/README.md#post) - Test a database connection given a type and connectionString

### [TextualDiff](docs/sdks/textualdiff/README.md)

* [Get](docs/sdks/textualdiff/README.md#get) - get the textual diff for given commit

### [TokenMetadata](docs/sdks/tokenmetadata/README.md)

* [Post](docs/sdks/tokenmetadata/README.md#post) - Add token and optional metadata to an existing hec input
* [Update](docs/sdks/tokenmetadata/README.md#update) - Update token metadata on existing hec input

### [TrustPolicies](docs/sdks/trustpolicies/README.md)

* [Get](docs/sdks/trustpolicies/README.md#get) - Get a list of TrustPolicy objects

### [UIState](docs/sdks/uistate/README.md)

* [Get](docs/sdks/uistate/README.md#get) - Get UI state by key
* [Update](docs/sdks/uistate/README.md#update) - Update UI state by key

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create User

### [UserAuth](docs/sdks/userauth/README.md)

* [Logout](docs/sdks/userauth/README.md#logout) - Log current user out

### [UserID](docs/sdks/userid/README.md)

* [Delete](docs/sdks/userid/README.md#delete) - Delete User
* [Get](docs/sdks/userid/README.md#get) - Get User by ID

### [UserObject](docs/sdks/userobject/README.md)

* [Get](docs/sdks/userobject/README.md#get) - Get a list of User objects
* [Update](docs/sdks/userobject/README.md#update) - Update User except for their roles

### [UserProperties](docs/sdks/userproperties/README.md)

* [Update](docs/sdks/userproperties/README.md#update) - Update User properties – admin use only

### [Versioning](docs/sdks/versioning/README.md)

* [Get](docs/sdks/versioning/README.md#get) - Get info about versioning availability

### [WorkerEdgeNodes](docs/sdks/workeredgenodes/README.md)

* [Get](docs/sdks/workeredgenodes/README.md#get) - get worker and edge nodes
* [Restarts](docs/sdks/workeredgenodes/README.md#restarts) - restarts worker nodes

### [WorkerEdgeNodesCount](docs/sdks/workeredgenodescount/README.md)

* [Get](docs/sdks/workeredgenodescount/README.md#get) - get worker and edge nodes count

### [WorkingTree](docs/sdks/workingtree/README.md)

* [Get](docs/sdks/workingtree/README.md#get) - get the the working tree status
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
