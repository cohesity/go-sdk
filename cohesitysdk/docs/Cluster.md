# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppsSubnet** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 
**AssignedRacksCount** | Pointer to **NullableInt32** | Specifies the number of racks in cluster with at least one rack assigned. | [optional] 
**AvailableMetadataSpace** | Pointer to **NullableInt64** | Information about storage available for metadata | [optional] 
**BannerEnabled** | Pointer to **NullableBool** | Specifies whether UI banner is enabled on the cluster or not. When banner is enabled, UI will make an additional API call to fetch the banner and show at the login page. | [optional] 
**ChassisCount** | Pointer to **NullableInt32** | Specifies the number of chassis in cluster. | [optional] 
**ClusterAuditLogConfig** | Pointer to [**ClusterAuditLogConfiguration**](ClusterAuditLogConfiguration.md) |  | [optional] 
**ClusterSoftwareVersion** | Pointer to **NullableString** | Specifies the current release of the Cohesity software running on this Cohesity Cluster. | [optional] 
**ClusterType** | Pointer to **NullableString** | Specifies the type of Cluster such as kPhysical. &#39;kPhysical&#39; indicates the Cohesity Cluster is hosted directly on hardware. &#39;kVirtualRobo&#39; indicates the Cohesity Cluster is hosted in a VM on a ESXi Host of a VMware vCenter Server using Cohesity&#39;s Virtual Edition. &#39;kMicrosoftCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Microsoft Azure using Cohesity&#39;s Cloud Edition. &#39;kAmazonCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Amazon S3 using Cohesity&#39;s Cloud Edition. &#39;kGoogleCloud&#39; indicates the Cohesity Cluster is hosted in a VM on Google Cloud Platform using Cohesity&#39;s Cloud Edition. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the Cohesity Cluster was created. This value is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**CurrentOpScheduledTimeSecs** | Pointer to **NullableInt64** | Specifies the time scheduled by the Cohesity Cluster to start the current running operation. | [optional] 
**CurrentOperation** | Pointer to **NullableString** | Specifies the current Cluster-level operation in progress. &#39;kUpgrade&#39; indicates the Cohesity Cluster is upgrading to a new release. &#39;kRemoveNode&#39; indicates the Cohesity Cluster is removing a Node from the Cluster. &#39;kNone&#39; indicates no action is occurring on the Cohesity Cluster. &#39;kDestroy&#39; indicates the Cohesity Cluster is getting destoryed. &#39;kClean&#39; indicates the Cohesity Cluster is getting cleaned. &#39;kRestartServices&#39; indicates the Cohesity Cluster is restarting the services. | [optional] 
**CurrentTimeMsecs** | Pointer to **NullableInt64** | Specifies the current system time on the Cohesity Cluster. This value is specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**DnsServerIps** | Pointer to **[]string** | Array of IP Addresses of DNS Servers.  Specifies the IP addresses of the DNS Servers used by the Cohesity Cluster. | [optional] 
**DomainNames** | Pointer to **[]string** | Array of Domain Names.  The first domain name specified in the array is the fully qualified domain name assigned to the Cohesity Cluster. Any additional domain names specified are used for the domain search list for hostname look-up. | [optional] 
**EnableActiveMonitoring** | Pointer to **NullableBool** | Specifies if Cohesity can receive monitoring information from the Cohesity Cluster. If &#39;true&#39;, remote monitoring of the Cohesity Cluster is allowed. | [optional] 
**EnableUpgradePkgPolling** | Pointer to **NullableBool** | If &#39;true&#39;, Cohesity&#39;s upgrade server is polled for new releases. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | If &#39;true&#39;, the entire Cohesity Cluster is encrypted including all View Boxes. | [optional] 
**EncryptionKeyRotationPeriodSecs** | Pointer to **NullableInt64** | Specifies the period of time (in seconds) when encryption keys are rotated. By default, the encryption keys are rotated every 77760000 seconds (30 days). | [optional] 
**EulaConfig** | Pointer to [**NullableEulaConfig**](EulaConfig.md) | Specifies the End User License Agreement (EULA) acceptance information. | [optional] 
**FaultToleranceLevel** | Pointer to **NullableString** | Specifies the level which &#39;MetadataFaultToleranceFactor&#39; applies to. &#39;kNode&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Node level. &#39;kChassis&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Chassis level. &#39;kRack&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Rack level. | [optional] 
**FilerAuditLogConfig** | Pointer to [**FilerAuditLogConfiguration**](FilerAuditLogConfiguration.md) |  | [optional] 
**FipsModeEnabled** | Pointer to **NullableBool** | Specifies if the Cohesity Cluster should operate in the FIPS mode, which is compliant with the Federal Information Processing Standard 140-2 certification. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway IP address. | [optional] 
**GoogleAnalyticsEnabled** | Pointer to **NullableBool** | Specifies whether Google Analytics is enabled. | [optional] 
**HardwareInfo** | Pointer to [**ClusterHardwareInfo**](ClusterHardwareInfo.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the unique id of Cohesity Cluster. | [optional] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the unique incarnation id of the Cohesity Cluster. | [optional] 
**IpPreference** | Pointer to **NullableInt32** | IP preference | [optional] 
**IsDocumentationLocal** | Pointer to **NullableBool** | Specifies what version of the documentation is used. If &#39;true&#39;, the version of documentation stored locally on the Cohesity Cluster is used. If &#39;false&#39;, the documentation stored on a Cohesity Web Server is used. The default is &#39;false&#39;. Cohesity recommends accessing the Help from the Cohesity Web site which provides the newest and most complete version of Help. | [optional] 
**LanguageLocale** | Pointer to **NullableString** | Specifies the language and locale for this Cohesity Cluster. | [optional] 
**LicenseState** | Pointer to [**NullableLicenseState**](LicenseState.md) | Specifies the Licensing State information. | [optional] 
**LocalAuthDomainName** | Pointer to **NullableString** | Domain name for SMB local authentication. | [optional] 
**LocalGroupsEnabled** | Pointer to **NullableBool** | Specifies whether to enable local groups on cluster. Once it is enabled, it cannot be disabled. | [optional] 
**MetadataFaultToleranceFactor** | Pointer to **NullableInt32** | Specifies metadata fault tolerance setting for the cluster. This denotes the number of simultaneous failures[node] supported by metadata services like gandalf and scribe. | [optional] 
**MultiTenancyEnabled** | Pointer to **NullableBool** | Specifies if multi tenancy is enabled in the cluster. Authentication &amp; Authorization will always use tenant_id, however, some UI elements may be disabled when multi tenancy is disabled. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Cohesity Cluster. | [optional] 
**NodeCount** | Pointer to **NullableInt64** | Specifies the number of Nodes in the Cohesity Cluster. | [optional] 
**NodeIps** | Pointer to **NullableString** | IP addresses of nodes in the cluster | [optional] 
**NtpSettings** | Pointer to [**NtpSettingsConfig**](NtpSettingsConfig.md) |  | [optional] 
**PcieSsdTierRebalanceDelaySecs** | Pointer to **NullableInt32** | Specifies the rebalance delay in seconds for cluster PcieSSD storage tier. | [optional] 
**ProxyVMSubnet** | Pointer to **NullableString** | The subnet reserved for ProxyVM | [optional] 
**ReverseTunnelEnabled** | Pointer to **NullableBool** | If &#39;true&#39;, Cohesity&#39;s Remote Tunnel is enabled. Cohesity can access the Cluster and provide remote assistance via a Remote Tunnel. | [optional] 
**ReverseTunnelEndTimeMsecs** | Pointer to **NullableInt64** | ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since epoch until when the reverse tunnel will stay enabled. | [optional] 
**SchemaInfoList** | Pointer to [**[]SchemaInfo**](SchemaInfo.md) | Specifies the time series schema info of the cluster. | [optional] 
**SmbAdDisabled** | Pointer to **NullableBool** | Specifies if Active Directory should be disabled for authentication of SMB shares. If &#39;true&#39;, Active Directory is disabled. | [optional] 
**SmbMultichannelEnabled** | Pointer to **NullableBool** | Specifies whether SMB multichannel is enabled on the cluster. When this is set to true, then any SMB3 multichannel enabled client can establish multiple TCP connection per session to the Server. | [optional] 
**Stats** | Pointer to [**ClusterStats**](ClusterStats.md) |  | [optional] 
**StigMode** | Pointer to **NullableBool** | Specifies if STIG mode is enabled or not. | [optional] 
**SupportedConfig** | Pointer to [**SupportedConfig**](SupportedConfig.md) |  | [optional] 
**SyslogServers** | Pointer to [**[]OldSyslogServer**](OldSyslogServer.md) | Syslog servers. | [optional] 
**TargetSoftwareVersion** | Pointer to **NullableString** | Specifies the Cohesity release that this Cluster is being upgraded to if an upgrade operation is in progress. | [optional] 
**TenantViewboxSharingEnabled** | Pointer to **NullableBool** | In case multi tenancy is enabled, this flag controls whether multiple tenants can be placed on the same viewbox. Once set to true, this flag should never become false. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies the timezone to use for showing time in emails, reports, filer audit logs, etc. | [optional] 
**TurboMode** | Pointer to **NullableBool** | Specifies if the cluster is in Turbo mode. | [optional] 
**UsedMetadataSpacePct** | Pointer to **NullableFloat64** | UsedMetadataSpacePct measures the percentage about storage used for metadata over the total storage available for metadata | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsSubnet

`func (o *Cluster) GetAppsSubnet() Subnet`

GetAppsSubnet returns the AppsSubnet field if non-nil, zero value otherwise.

### GetAppsSubnetOk

`func (o *Cluster) GetAppsSubnetOk() (*Subnet, bool)`

GetAppsSubnetOk returns a tuple with the AppsSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsSubnet

`func (o *Cluster) SetAppsSubnet(v Subnet)`

SetAppsSubnet sets AppsSubnet field to given value.

### HasAppsSubnet

`func (o *Cluster) HasAppsSubnet() bool`

HasAppsSubnet returns a boolean if a field has been set.

### GetAssignedRacksCount

`func (o *Cluster) GetAssignedRacksCount() int32`

GetAssignedRacksCount returns the AssignedRacksCount field if non-nil, zero value otherwise.

### GetAssignedRacksCountOk

`func (o *Cluster) GetAssignedRacksCountOk() (*int32, bool)`

GetAssignedRacksCountOk returns a tuple with the AssignedRacksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedRacksCount

`func (o *Cluster) SetAssignedRacksCount(v int32)`

SetAssignedRacksCount sets AssignedRacksCount field to given value.

### HasAssignedRacksCount

`func (o *Cluster) HasAssignedRacksCount() bool`

HasAssignedRacksCount returns a boolean if a field has been set.

### SetAssignedRacksCountNil

`func (o *Cluster) SetAssignedRacksCountNil(b bool)`

 SetAssignedRacksCountNil sets the value for AssignedRacksCount to be an explicit nil

### UnsetAssignedRacksCount
`func (o *Cluster) UnsetAssignedRacksCount()`

UnsetAssignedRacksCount ensures that no value is present for AssignedRacksCount, not even an explicit nil
### GetAvailableMetadataSpace

`func (o *Cluster) GetAvailableMetadataSpace() int64`

GetAvailableMetadataSpace returns the AvailableMetadataSpace field if non-nil, zero value otherwise.

### GetAvailableMetadataSpaceOk

`func (o *Cluster) GetAvailableMetadataSpaceOk() (*int64, bool)`

GetAvailableMetadataSpaceOk returns a tuple with the AvailableMetadataSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMetadataSpace

`func (o *Cluster) SetAvailableMetadataSpace(v int64)`

SetAvailableMetadataSpace sets AvailableMetadataSpace field to given value.

### HasAvailableMetadataSpace

`func (o *Cluster) HasAvailableMetadataSpace() bool`

HasAvailableMetadataSpace returns a boolean if a field has been set.

### SetAvailableMetadataSpaceNil

`func (o *Cluster) SetAvailableMetadataSpaceNil(b bool)`

 SetAvailableMetadataSpaceNil sets the value for AvailableMetadataSpace to be an explicit nil

### UnsetAvailableMetadataSpace
`func (o *Cluster) UnsetAvailableMetadataSpace()`

UnsetAvailableMetadataSpace ensures that no value is present for AvailableMetadataSpace, not even an explicit nil
### GetBannerEnabled

`func (o *Cluster) GetBannerEnabled() bool`

GetBannerEnabled returns the BannerEnabled field if non-nil, zero value otherwise.

### GetBannerEnabledOk

`func (o *Cluster) GetBannerEnabledOk() (*bool, bool)`

GetBannerEnabledOk returns a tuple with the BannerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerEnabled

`func (o *Cluster) SetBannerEnabled(v bool)`

SetBannerEnabled sets BannerEnabled field to given value.

### HasBannerEnabled

`func (o *Cluster) HasBannerEnabled() bool`

HasBannerEnabled returns a boolean if a field has been set.

### SetBannerEnabledNil

`func (o *Cluster) SetBannerEnabledNil(b bool)`

 SetBannerEnabledNil sets the value for BannerEnabled to be an explicit nil

### UnsetBannerEnabled
`func (o *Cluster) UnsetBannerEnabled()`

UnsetBannerEnabled ensures that no value is present for BannerEnabled, not even an explicit nil
### GetChassisCount

`func (o *Cluster) GetChassisCount() int32`

GetChassisCount returns the ChassisCount field if non-nil, zero value otherwise.

### GetChassisCountOk

`func (o *Cluster) GetChassisCountOk() (*int32, bool)`

GetChassisCountOk returns a tuple with the ChassisCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisCount

`func (o *Cluster) SetChassisCount(v int32)`

SetChassisCount sets ChassisCount field to given value.

### HasChassisCount

`func (o *Cluster) HasChassisCount() bool`

HasChassisCount returns a boolean if a field has been set.

### SetChassisCountNil

`func (o *Cluster) SetChassisCountNil(b bool)`

 SetChassisCountNil sets the value for ChassisCount to be an explicit nil

### UnsetChassisCount
`func (o *Cluster) UnsetChassisCount()`

UnsetChassisCount ensures that no value is present for ChassisCount, not even an explicit nil
### GetClusterAuditLogConfig

`func (o *Cluster) GetClusterAuditLogConfig() ClusterAuditLogConfiguration`

GetClusterAuditLogConfig returns the ClusterAuditLogConfig field if non-nil, zero value otherwise.

### GetClusterAuditLogConfigOk

`func (o *Cluster) GetClusterAuditLogConfigOk() (*ClusterAuditLogConfiguration, bool)`

GetClusterAuditLogConfigOk returns a tuple with the ClusterAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAuditLogConfig

`func (o *Cluster) SetClusterAuditLogConfig(v ClusterAuditLogConfiguration)`

SetClusterAuditLogConfig sets ClusterAuditLogConfig field to given value.

### HasClusterAuditLogConfig

`func (o *Cluster) HasClusterAuditLogConfig() bool`

HasClusterAuditLogConfig returns a boolean if a field has been set.

### GetClusterSoftwareVersion

`func (o *Cluster) GetClusterSoftwareVersion() string`

GetClusterSoftwareVersion returns the ClusterSoftwareVersion field if non-nil, zero value otherwise.

### GetClusterSoftwareVersionOk

`func (o *Cluster) GetClusterSoftwareVersionOk() (*string, bool)`

GetClusterSoftwareVersionOk returns a tuple with the ClusterSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSoftwareVersion

`func (o *Cluster) SetClusterSoftwareVersion(v string)`

SetClusterSoftwareVersion sets ClusterSoftwareVersion field to given value.

### HasClusterSoftwareVersion

`func (o *Cluster) HasClusterSoftwareVersion() bool`

HasClusterSoftwareVersion returns a boolean if a field has been set.

### SetClusterSoftwareVersionNil

`func (o *Cluster) SetClusterSoftwareVersionNil(b bool)`

 SetClusterSoftwareVersionNil sets the value for ClusterSoftwareVersion to be an explicit nil

### UnsetClusterSoftwareVersion
`func (o *Cluster) UnsetClusterSoftwareVersion()`

UnsetClusterSoftwareVersion ensures that no value is present for ClusterSoftwareVersion, not even an explicit nil
### GetClusterType

`func (o *Cluster) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *Cluster) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *Cluster) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *Cluster) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### SetClusterTypeNil

`func (o *Cluster) SetClusterTypeNil(b bool)`

 SetClusterTypeNil sets the value for ClusterType to be an explicit nil

### UnsetClusterType
`func (o *Cluster) UnsetClusterType()`

UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *Cluster) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *Cluster) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *Cluster) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *Cluster) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *Cluster) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *Cluster) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetCurrentOpScheduledTimeSecs

`func (o *Cluster) GetCurrentOpScheduledTimeSecs() int64`

GetCurrentOpScheduledTimeSecs returns the CurrentOpScheduledTimeSecs field if non-nil, zero value otherwise.

### GetCurrentOpScheduledTimeSecsOk

`func (o *Cluster) GetCurrentOpScheduledTimeSecsOk() (*int64, bool)`

GetCurrentOpScheduledTimeSecsOk returns a tuple with the CurrentOpScheduledTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOpScheduledTimeSecs

`func (o *Cluster) SetCurrentOpScheduledTimeSecs(v int64)`

SetCurrentOpScheduledTimeSecs sets CurrentOpScheduledTimeSecs field to given value.

### HasCurrentOpScheduledTimeSecs

`func (o *Cluster) HasCurrentOpScheduledTimeSecs() bool`

HasCurrentOpScheduledTimeSecs returns a boolean if a field has been set.

### SetCurrentOpScheduledTimeSecsNil

`func (o *Cluster) SetCurrentOpScheduledTimeSecsNil(b bool)`

 SetCurrentOpScheduledTimeSecsNil sets the value for CurrentOpScheduledTimeSecs to be an explicit nil

### UnsetCurrentOpScheduledTimeSecs
`func (o *Cluster) UnsetCurrentOpScheduledTimeSecs()`

UnsetCurrentOpScheduledTimeSecs ensures that no value is present for CurrentOpScheduledTimeSecs, not even an explicit nil
### GetCurrentOperation

`func (o *Cluster) GetCurrentOperation() string`

GetCurrentOperation returns the CurrentOperation field if non-nil, zero value otherwise.

### GetCurrentOperationOk

`func (o *Cluster) GetCurrentOperationOk() (*string, bool)`

GetCurrentOperationOk returns a tuple with the CurrentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOperation

`func (o *Cluster) SetCurrentOperation(v string)`

SetCurrentOperation sets CurrentOperation field to given value.

### HasCurrentOperation

`func (o *Cluster) HasCurrentOperation() bool`

HasCurrentOperation returns a boolean if a field has been set.

### SetCurrentOperationNil

`func (o *Cluster) SetCurrentOperationNil(b bool)`

 SetCurrentOperationNil sets the value for CurrentOperation to be an explicit nil

### UnsetCurrentOperation
`func (o *Cluster) UnsetCurrentOperation()`

UnsetCurrentOperation ensures that no value is present for CurrentOperation, not even an explicit nil
### GetCurrentTimeMsecs

`func (o *Cluster) GetCurrentTimeMsecs() int64`

GetCurrentTimeMsecs returns the CurrentTimeMsecs field if non-nil, zero value otherwise.

### GetCurrentTimeMsecsOk

`func (o *Cluster) GetCurrentTimeMsecsOk() (*int64, bool)`

GetCurrentTimeMsecsOk returns a tuple with the CurrentTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTimeMsecs

`func (o *Cluster) SetCurrentTimeMsecs(v int64)`

SetCurrentTimeMsecs sets CurrentTimeMsecs field to given value.

### HasCurrentTimeMsecs

`func (o *Cluster) HasCurrentTimeMsecs() bool`

HasCurrentTimeMsecs returns a boolean if a field has been set.

### SetCurrentTimeMsecsNil

`func (o *Cluster) SetCurrentTimeMsecsNil(b bool)`

 SetCurrentTimeMsecsNil sets the value for CurrentTimeMsecs to be an explicit nil

### UnsetCurrentTimeMsecs
`func (o *Cluster) UnsetCurrentTimeMsecs()`

UnsetCurrentTimeMsecs ensures that no value is present for CurrentTimeMsecs, not even an explicit nil
### GetDnsServerIps

`func (o *Cluster) GetDnsServerIps() []string`

GetDnsServerIps returns the DnsServerIps field if non-nil, zero value otherwise.

### GetDnsServerIpsOk

`func (o *Cluster) GetDnsServerIpsOk() (*[]string, bool)`

GetDnsServerIpsOk returns a tuple with the DnsServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerIps

`func (o *Cluster) SetDnsServerIps(v []string)`

SetDnsServerIps sets DnsServerIps field to given value.

### HasDnsServerIps

`func (o *Cluster) HasDnsServerIps() bool`

HasDnsServerIps returns a boolean if a field has been set.

### SetDnsServerIpsNil

`func (o *Cluster) SetDnsServerIpsNil(b bool)`

 SetDnsServerIpsNil sets the value for DnsServerIps to be an explicit nil

### UnsetDnsServerIps
`func (o *Cluster) UnsetDnsServerIps()`

UnsetDnsServerIps ensures that no value is present for DnsServerIps, not even an explicit nil
### GetDomainNames

`func (o *Cluster) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Cluster) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Cluster) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *Cluster) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *Cluster) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *Cluster) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetEnableActiveMonitoring

`func (o *Cluster) GetEnableActiveMonitoring() bool`

GetEnableActiveMonitoring returns the EnableActiveMonitoring field if non-nil, zero value otherwise.

### GetEnableActiveMonitoringOk

`func (o *Cluster) GetEnableActiveMonitoringOk() (*bool, bool)`

GetEnableActiveMonitoringOk returns a tuple with the EnableActiveMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableActiveMonitoring

`func (o *Cluster) SetEnableActiveMonitoring(v bool)`

SetEnableActiveMonitoring sets EnableActiveMonitoring field to given value.

### HasEnableActiveMonitoring

`func (o *Cluster) HasEnableActiveMonitoring() bool`

HasEnableActiveMonitoring returns a boolean if a field has been set.

### SetEnableActiveMonitoringNil

`func (o *Cluster) SetEnableActiveMonitoringNil(b bool)`

 SetEnableActiveMonitoringNil sets the value for EnableActiveMonitoring to be an explicit nil

### UnsetEnableActiveMonitoring
`func (o *Cluster) UnsetEnableActiveMonitoring()`

UnsetEnableActiveMonitoring ensures that no value is present for EnableActiveMonitoring, not even an explicit nil
### GetEnableUpgradePkgPolling

`func (o *Cluster) GetEnableUpgradePkgPolling() bool`

GetEnableUpgradePkgPolling returns the EnableUpgradePkgPolling field if non-nil, zero value otherwise.

### GetEnableUpgradePkgPollingOk

`func (o *Cluster) GetEnableUpgradePkgPollingOk() (*bool, bool)`

GetEnableUpgradePkgPollingOk returns a tuple with the EnableUpgradePkgPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpgradePkgPolling

`func (o *Cluster) SetEnableUpgradePkgPolling(v bool)`

SetEnableUpgradePkgPolling sets EnableUpgradePkgPolling field to given value.

### HasEnableUpgradePkgPolling

`func (o *Cluster) HasEnableUpgradePkgPolling() bool`

HasEnableUpgradePkgPolling returns a boolean if a field has been set.

### SetEnableUpgradePkgPollingNil

`func (o *Cluster) SetEnableUpgradePkgPollingNil(b bool)`

 SetEnableUpgradePkgPollingNil sets the value for EnableUpgradePkgPolling to be an explicit nil

### UnsetEnableUpgradePkgPolling
`func (o *Cluster) UnsetEnableUpgradePkgPolling()`

UnsetEnableUpgradePkgPolling ensures that no value is present for EnableUpgradePkgPolling, not even an explicit nil
### GetEncryptionEnabled

`func (o *Cluster) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *Cluster) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *Cluster) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *Cluster) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *Cluster) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *Cluster) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetEncryptionKeyRotationPeriodSecs

`func (o *Cluster) GetEncryptionKeyRotationPeriodSecs() int64`

GetEncryptionKeyRotationPeriodSecs returns the EncryptionKeyRotationPeriodSecs field if non-nil, zero value otherwise.

### GetEncryptionKeyRotationPeriodSecsOk

`func (o *Cluster) GetEncryptionKeyRotationPeriodSecsOk() (*int64, bool)`

GetEncryptionKeyRotationPeriodSecsOk returns a tuple with the EncryptionKeyRotationPeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyRotationPeriodSecs

`func (o *Cluster) SetEncryptionKeyRotationPeriodSecs(v int64)`

SetEncryptionKeyRotationPeriodSecs sets EncryptionKeyRotationPeriodSecs field to given value.

### HasEncryptionKeyRotationPeriodSecs

`func (o *Cluster) HasEncryptionKeyRotationPeriodSecs() bool`

HasEncryptionKeyRotationPeriodSecs returns a boolean if a field has been set.

### SetEncryptionKeyRotationPeriodSecsNil

`func (o *Cluster) SetEncryptionKeyRotationPeriodSecsNil(b bool)`

 SetEncryptionKeyRotationPeriodSecsNil sets the value for EncryptionKeyRotationPeriodSecs to be an explicit nil

### UnsetEncryptionKeyRotationPeriodSecs
`func (o *Cluster) UnsetEncryptionKeyRotationPeriodSecs()`

UnsetEncryptionKeyRotationPeriodSecs ensures that no value is present for EncryptionKeyRotationPeriodSecs, not even an explicit nil
### GetEulaConfig

`func (o *Cluster) GetEulaConfig() EulaConfig`

GetEulaConfig returns the EulaConfig field if non-nil, zero value otherwise.

### GetEulaConfigOk

`func (o *Cluster) GetEulaConfigOk() (*EulaConfig, bool)`

GetEulaConfigOk returns a tuple with the EulaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaConfig

`func (o *Cluster) SetEulaConfig(v EulaConfig)`

SetEulaConfig sets EulaConfig field to given value.

### HasEulaConfig

`func (o *Cluster) HasEulaConfig() bool`

HasEulaConfig returns a boolean if a field has been set.

### SetEulaConfigNil

`func (o *Cluster) SetEulaConfigNil(b bool)`

 SetEulaConfigNil sets the value for EulaConfig to be an explicit nil

### UnsetEulaConfig
`func (o *Cluster) UnsetEulaConfig()`

UnsetEulaConfig ensures that no value is present for EulaConfig, not even an explicit nil
### GetFaultToleranceLevel

`func (o *Cluster) GetFaultToleranceLevel() string`

GetFaultToleranceLevel returns the FaultToleranceLevel field if non-nil, zero value otherwise.

### GetFaultToleranceLevelOk

`func (o *Cluster) GetFaultToleranceLevelOk() (*string, bool)`

GetFaultToleranceLevelOk returns a tuple with the FaultToleranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultToleranceLevel

`func (o *Cluster) SetFaultToleranceLevel(v string)`

SetFaultToleranceLevel sets FaultToleranceLevel field to given value.

### HasFaultToleranceLevel

`func (o *Cluster) HasFaultToleranceLevel() bool`

HasFaultToleranceLevel returns a boolean if a field has been set.

### SetFaultToleranceLevelNil

`func (o *Cluster) SetFaultToleranceLevelNil(b bool)`

 SetFaultToleranceLevelNil sets the value for FaultToleranceLevel to be an explicit nil

### UnsetFaultToleranceLevel
`func (o *Cluster) UnsetFaultToleranceLevel()`

UnsetFaultToleranceLevel ensures that no value is present for FaultToleranceLevel, not even an explicit nil
### GetFilerAuditLogConfig

`func (o *Cluster) GetFilerAuditLogConfig() FilerAuditLogConfiguration`

GetFilerAuditLogConfig returns the FilerAuditLogConfig field if non-nil, zero value otherwise.

### GetFilerAuditLogConfigOk

`func (o *Cluster) GetFilerAuditLogConfigOk() (*FilerAuditLogConfiguration, bool)`

GetFilerAuditLogConfigOk returns a tuple with the FilerAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerAuditLogConfig

`func (o *Cluster) SetFilerAuditLogConfig(v FilerAuditLogConfiguration)`

SetFilerAuditLogConfig sets FilerAuditLogConfig field to given value.

### HasFilerAuditLogConfig

`func (o *Cluster) HasFilerAuditLogConfig() bool`

HasFilerAuditLogConfig returns a boolean if a field has been set.

### GetFipsModeEnabled

`func (o *Cluster) GetFipsModeEnabled() bool`

GetFipsModeEnabled returns the FipsModeEnabled field if non-nil, zero value otherwise.

### GetFipsModeEnabledOk

`func (o *Cluster) GetFipsModeEnabledOk() (*bool, bool)`

GetFipsModeEnabledOk returns a tuple with the FipsModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsModeEnabled

`func (o *Cluster) SetFipsModeEnabled(v bool)`

SetFipsModeEnabled sets FipsModeEnabled field to given value.

### HasFipsModeEnabled

`func (o *Cluster) HasFipsModeEnabled() bool`

HasFipsModeEnabled returns a boolean if a field has been set.

### SetFipsModeEnabledNil

`func (o *Cluster) SetFipsModeEnabledNil(b bool)`

 SetFipsModeEnabledNil sets the value for FipsModeEnabled to be an explicit nil

### UnsetFipsModeEnabled
`func (o *Cluster) UnsetFipsModeEnabled()`

UnsetFipsModeEnabled ensures that no value is present for FipsModeEnabled, not even an explicit nil
### GetGateway

`func (o *Cluster) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Cluster) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Cluster) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Cluster) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Cluster) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Cluster) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGoogleAnalyticsEnabled

`func (o *Cluster) GetGoogleAnalyticsEnabled() bool`

GetGoogleAnalyticsEnabled returns the GoogleAnalyticsEnabled field if non-nil, zero value otherwise.

### GetGoogleAnalyticsEnabledOk

`func (o *Cluster) GetGoogleAnalyticsEnabledOk() (*bool, bool)`

GetGoogleAnalyticsEnabledOk returns a tuple with the GoogleAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsEnabled

`func (o *Cluster) SetGoogleAnalyticsEnabled(v bool)`

SetGoogleAnalyticsEnabled sets GoogleAnalyticsEnabled field to given value.

### HasGoogleAnalyticsEnabled

`func (o *Cluster) HasGoogleAnalyticsEnabled() bool`

HasGoogleAnalyticsEnabled returns a boolean if a field has been set.

### SetGoogleAnalyticsEnabledNil

`func (o *Cluster) SetGoogleAnalyticsEnabledNil(b bool)`

 SetGoogleAnalyticsEnabledNil sets the value for GoogleAnalyticsEnabled to be an explicit nil

### UnsetGoogleAnalyticsEnabled
`func (o *Cluster) UnsetGoogleAnalyticsEnabled()`

UnsetGoogleAnalyticsEnabled ensures that no value is present for GoogleAnalyticsEnabled, not even an explicit nil
### GetHardwareInfo

`func (o *Cluster) GetHardwareInfo() ClusterHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *Cluster) GetHardwareInfoOk() (*ClusterHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *Cluster) SetHardwareInfo(v ClusterHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *Cluster) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### GetId

`func (o *Cluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Cluster) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Cluster) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncarnationId

`func (o *Cluster) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *Cluster) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *Cluster) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *Cluster) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *Cluster) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *Cluster) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
### GetIpPreference

`func (o *Cluster) GetIpPreference() int32`

GetIpPreference returns the IpPreference field if non-nil, zero value otherwise.

### GetIpPreferenceOk

`func (o *Cluster) GetIpPreferenceOk() (*int32, bool)`

GetIpPreferenceOk returns a tuple with the IpPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPreference

`func (o *Cluster) SetIpPreference(v int32)`

SetIpPreference sets IpPreference field to given value.

### HasIpPreference

`func (o *Cluster) HasIpPreference() bool`

HasIpPreference returns a boolean if a field has been set.

### SetIpPreferenceNil

`func (o *Cluster) SetIpPreferenceNil(b bool)`

 SetIpPreferenceNil sets the value for IpPreference to be an explicit nil

### UnsetIpPreference
`func (o *Cluster) UnsetIpPreference()`

UnsetIpPreference ensures that no value is present for IpPreference, not even an explicit nil
### GetIsDocumentationLocal

`func (o *Cluster) GetIsDocumentationLocal() bool`

GetIsDocumentationLocal returns the IsDocumentationLocal field if non-nil, zero value otherwise.

### GetIsDocumentationLocalOk

`func (o *Cluster) GetIsDocumentationLocalOk() (*bool, bool)`

GetIsDocumentationLocalOk returns a tuple with the IsDocumentationLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDocumentationLocal

`func (o *Cluster) SetIsDocumentationLocal(v bool)`

SetIsDocumentationLocal sets IsDocumentationLocal field to given value.

### HasIsDocumentationLocal

`func (o *Cluster) HasIsDocumentationLocal() bool`

HasIsDocumentationLocal returns a boolean if a field has been set.

### SetIsDocumentationLocalNil

`func (o *Cluster) SetIsDocumentationLocalNil(b bool)`

 SetIsDocumentationLocalNil sets the value for IsDocumentationLocal to be an explicit nil

### UnsetIsDocumentationLocal
`func (o *Cluster) UnsetIsDocumentationLocal()`

UnsetIsDocumentationLocal ensures that no value is present for IsDocumentationLocal, not even an explicit nil
### GetLanguageLocale

`func (o *Cluster) GetLanguageLocale() string`

GetLanguageLocale returns the LanguageLocale field if non-nil, zero value otherwise.

### GetLanguageLocaleOk

`func (o *Cluster) GetLanguageLocaleOk() (*string, bool)`

GetLanguageLocaleOk returns a tuple with the LanguageLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageLocale

`func (o *Cluster) SetLanguageLocale(v string)`

SetLanguageLocale sets LanguageLocale field to given value.

### HasLanguageLocale

`func (o *Cluster) HasLanguageLocale() bool`

HasLanguageLocale returns a boolean if a field has been set.

### SetLanguageLocaleNil

`func (o *Cluster) SetLanguageLocaleNil(b bool)`

 SetLanguageLocaleNil sets the value for LanguageLocale to be an explicit nil

### UnsetLanguageLocale
`func (o *Cluster) UnsetLanguageLocale()`

UnsetLanguageLocale ensures that no value is present for LanguageLocale, not even an explicit nil
### GetLicenseState

`func (o *Cluster) GetLicenseState() LicenseState`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *Cluster) GetLicenseStateOk() (*LicenseState, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *Cluster) SetLicenseState(v LicenseState)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *Cluster) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### SetLicenseStateNil

`func (o *Cluster) SetLicenseStateNil(b bool)`

 SetLicenseStateNil sets the value for LicenseState to be an explicit nil

### UnsetLicenseState
`func (o *Cluster) UnsetLicenseState()`

UnsetLicenseState ensures that no value is present for LicenseState, not even an explicit nil
### GetLocalAuthDomainName

`func (o *Cluster) GetLocalAuthDomainName() string`

GetLocalAuthDomainName returns the LocalAuthDomainName field if non-nil, zero value otherwise.

### GetLocalAuthDomainNameOk

`func (o *Cluster) GetLocalAuthDomainNameOk() (*string, bool)`

GetLocalAuthDomainNameOk returns a tuple with the LocalAuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuthDomainName

`func (o *Cluster) SetLocalAuthDomainName(v string)`

SetLocalAuthDomainName sets LocalAuthDomainName field to given value.

### HasLocalAuthDomainName

`func (o *Cluster) HasLocalAuthDomainName() bool`

HasLocalAuthDomainName returns a boolean if a field has been set.

### SetLocalAuthDomainNameNil

`func (o *Cluster) SetLocalAuthDomainNameNil(b bool)`

 SetLocalAuthDomainNameNil sets the value for LocalAuthDomainName to be an explicit nil

### UnsetLocalAuthDomainName
`func (o *Cluster) UnsetLocalAuthDomainName()`

UnsetLocalAuthDomainName ensures that no value is present for LocalAuthDomainName, not even an explicit nil
### GetLocalGroupsEnabled

`func (o *Cluster) GetLocalGroupsEnabled() bool`

GetLocalGroupsEnabled returns the LocalGroupsEnabled field if non-nil, zero value otherwise.

### GetLocalGroupsEnabledOk

`func (o *Cluster) GetLocalGroupsEnabledOk() (*bool, bool)`

GetLocalGroupsEnabledOk returns a tuple with the LocalGroupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupsEnabled

`func (o *Cluster) SetLocalGroupsEnabled(v bool)`

SetLocalGroupsEnabled sets LocalGroupsEnabled field to given value.

### HasLocalGroupsEnabled

`func (o *Cluster) HasLocalGroupsEnabled() bool`

HasLocalGroupsEnabled returns a boolean if a field has been set.

### SetLocalGroupsEnabledNil

`func (o *Cluster) SetLocalGroupsEnabledNil(b bool)`

 SetLocalGroupsEnabledNil sets the value for LocalGroupsEnabled to be an explicit nil

### UnsetLocalGroupsEnabled
`func (o *Cluster) UnsetLocalGroupsEnabled()`

UnsetLocalGroupsEnabled ensures that no value is present for LocalGroupsEnabled, not even an explicit nil
### GetMetadataFaultToleranceFactor

`func (o *Cluster) GetMetadataFaultToleranceFactor() int32`

GetMetadataFaultToleranceFactor returns the MetadataFaultToleranceFactor field if non-nil, zero value otherwise.

### GetMetadataFaultToleranceFactorOk

`func (o *Cluster) GetMetadataFaultToleranceFactorOk() (*int32, bool)`

GetMetadataFaultToleranceFactorOk returns a tuple with the MetadataFaultToleranceFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFaultToleranceFactor

`func (o *Cluster) SetMetadataFaultToleranceFactor(v int32)`

SetMetadataFaultToleranceFactor sets MetadataFaultToleranceFactor field to given value.

### HasMetadataFaultToleranceFactor

`func (o *Cluster) HasMetadataFaultToleranceFactor() bool`

HasMetadataFaultToleranceFactor returns a boolean if a field has been set.

### SetMetadataFaultToleranceFactorNil

`func (o *Cluster) SetMetadataFaultToleranceFactorNil(b bool)`

 SetMetadataFaultToleranceFactorNil sets the value for MetadataFaultToleranceFactor to be an explicit nil

### UnsetMetadataFaultToleranceFactor
`func (o *Cluster) UnsetMetadataFaultToleranceFactor()`

UnsetMetadataFaultToleranceFactor ensures that no value is present for MetadataFaultToleranceFactor, not even an explicit nil
### GetMultiTenancyEnabled

`func (o *Cluster) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *Cluster) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *Cluster) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *Cluster) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *Cluster) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *Cluster) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Cluster) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Cluster) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeCount

`func (o *Cluster) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Cluster) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Cluster) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Cluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### SetNodeCountNil

`func (o *Cluster) SetNodeCountNil(b bool)`

 SetNodeCountNil sets the value for NodeCount to be an explicit nil

### UnsetNodeCount
`func (o *Cluster) UnsetNodeCount()`

UnsetNodeCount ensures that no value is present for NodeCount, not even an explicit nil
### GetNodeIps

`func (o *Cluster) GetNodeIps() string`

GetNodeIps returns the NodeIps field if non-nil, zero value otherwise.

### GetNodeIpsOk

`func (o *Cluster) GetNodeIpsOk() (*string, bool)`

GetNodeIpsOk returns a tuple with the NodeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIps

`func (o *Cluster) SetNodeIps(v string)`

SetNodeIps sets NodeIps field to given value.

### HasNodeIps

`func (o *Cluster) HasNodeIps() bool`

HasNodeIps returns a boolean if a field has been set.

### SetNodeIpsNil

`func (o *Cluster) SetNodeIpsNil(b bool)`

 SetNodeIpsNil sets the value for NodeIps to be an explicit nil

### UnsetNodeIps
`func (o *Cluster) UnsetNodeIps()`

UnsetNodeIps ensures that no value is present for NodeIps, not even an explicit nil
### GetNtpSettings

`func (o *Cluster) GetNtpSettings() NtpSettingsConfig`

GetNtpSettings returns the NtpSettings field if non-nil, zero value otherwise.

### GetNtpSettingsOk

`func (o *Cluster) GetNtpSettingsOk() (*NtpSettingsConfig, bool)`

GetNtpSettingsOk returns a tuple with the NtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpSettings

`func (o *Cluster) SetNtpSettings(v NtpSettingsConfig)`

SetNtpSettings sets NtpSettings field to given value.

### HasNtpSettings

`func (o *Cluster) HasNtpSettings() bool`

HasNtpSettings returns a boolean if a field has been set.

### GetPcieSsdTierRebalanceDelaySecs

`func (o *Cluster) GetPcieSsdTierRebalanceDelaySecs() int32`

GetPcieSsdTierRebalanceDelaySecs returns the PcieSsdTierRebalanceDelaySecs field if non-nil, zero value otherwise.

### GetPcieSsdTierRebalanceDelaySecsOk

`func (o *Cluster) GetPcieSsdTierRebalanceDelaySecsOk() (*int32, bool)`

GetPcieSsdTierRebalanceDelaySecsOk returns a tuple with the PcieSsdTierRebalanceDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSsdTierRebalanceDelaySecs

`func (o *Cluster) SetPcieSsdTierRebalanceDelaySecs(v int32)`

SetPcieSsdTierRebalanceDelaySecs sets PcieSsdTierRebalanceDelaySecs field to given value.

### HasPcieSsdTierRebalanceDelaySecs

`func (o *Cluster) HasPcieSsdTierRebalanceDelaySecs() bool`

HasPcieSsdTierRebalanceDelaySecs returns a boolean if a field has been set.

### SetPcieSsdTierRebalanceDelaySecsNil

`func (o *Cluster) SetPcieSsdTierRebalanceDelaySecsNil(b bool)`

 SetPcieSsdTierRebalanceDelaySecsNil sets the value for PcieSsdTierRebalanceDelaySecs to be an explicit nil

### UnsetPcieSsdTierRebalanceDelaySecs
`func (o *Cluster) UnsetPcieSsdTierRebalanceDelaySecs()`

UnsetPcieSsdTierRebalanceDelaySecs ensures that no value is present for PcieSsdTierRebalanceDelaySecs, not even an explicit nil
### GetProxyVMSubnet

`func (o *Cluster) GetProxyVMSubnet() string`

GetProxyVMSubnet returns the ProxyVMSubnet field if non-nil, zero value otherwise.

### GetProxyVMSubnetOk

`func (o *Cluster) GetProxyVMSubnetOk() (*string, bool)`

GetProxyVMSubnetOk returns a tuple with the ProxyVMSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyVMSubnet

`func (o *Cluster) SetProxyVMSubnet(v string)`

SetProxyVMSubnet sets ProxyVMSubnet field to given value.

### HasProxyVMSubnet

`func (o *Cluster) HasProxyVMSubnet() bool`

HasProxyVMSubnet returns a boolean if a field has been set.

### SetProxyVMSubnetNil

`func (o *Cluster) SetProxyVMSubnetNil(b bool)`

 SetProxyVMSubnetNil sets the value for ProxyVMSubnet to be an explicit nil

### UnsetProxyVMSubnet
`func (o *Cluster) UnsetProxyVMSubnet()`

UnsetProxyVMSubnet ensures that no value is present for ProxyVMSubnet, not even an explicit nil
### GetReverseTunnelEnabled

`func (o *Cluster) GetReverseTunnelEnabled() bool`

GetReverseTunnelEnabled returns the ReverseTunnelEnabled field if non-nil, zero value otherwise.

### GetReverseTunnelEnabledOk

`func (o *Cluster) GetReverseTunnelEnabledOk() (*bool, bool)`

GetReverseTunnelEnabledOk returns a tuple with the ReverseTunnelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseTunnelEnabled

`func (o *Cluster) SetReverseTunnelEnabled(v bool)`

SetReverseTunnelEnabled sets ReverseTunnelEnabled field to given value.

### HasReverseTunnelEnabled

`func (o *Cluster) HasReverseTunnelEnabled() bool`

HasReverseTunnelEnabled returns a boolean if a field has been set.

### SetReverseTunnelEnabledNil

`func (o *Cluster) SetReverseTunnelEnabledNil(b bool)`

 SetReverseTunnelEnabledNil sets the value for ReverseTunnelEnabled to be an explicit nil

### UnsetReverseTunnelEnabled
`func (o *Cluster) UnsetReverseTunnelEnabled()`

UnsetReverseTunnelEnabled ensures that no value is present for ReverseTunnelEnabled, not even an explicit nil
### GetReverseTunnelEndTimeMsecs

`func (o *Cluster) GetReverseTunnelEndTimeMsecs() int64`

GetReverseTunnelEndTimeMsecs returns the ReverseTunnelEndTimeMsecs field if non-nil, zero value otherwise.

### GetReverseTunnelEndTimeMsecsOk

`func (o *Cluster) GetReverseTunnelEndTimeMsecsOk() (*int64, bool)`

GetReverseTunnelEndTimeMsecsOk returns a tuple with the ReverseTunnelEndTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseTunnelEndTimeMsecs

`func (o *Cluster) SetReverseTunnelEndTimeMsecs(v int64)`

SetReverseTunnelEndTimeMsecs sets ReverseTunnelEndTimeMsecs field to given value.

### HasReverseTunnelEndTimeMsecs

`func (o *Cluster) HasReverseTunnelEndTimeMsecs() bool`

HasReverseTunnelEndTimeMsecs returns a boolean if a field has been set.

### SetReverseTunnelEndTimeMsecsNil

`func (o *Cluster) SetReverseTunnelEndTimeMsecsNil(b bool)`

 SetReverseTunnelEndTimeMsecsNil sets the value for ReverseTunnelEndTimeMsecs to be an explicit nil

### UnsetReverseTunnelEndTimeMsecs
`func (o *Cluster) UnsetReverseTunnelEndTimeMsecs()`

UnsetReverseTunnelEndTimeMsecs ensures that no value is present for ReverseTunnelEndTimeMsecs, not even an explicit nil
### GetSchemaInfoList

`func (o *Cluster) GetSchemaInfoList() []SchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *Cluster) GetSchemaInfoListOk() (*[]SchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *Cluster) SetSchemaInfoList(v []SchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *Cluster) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *Cluster) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *Cluster) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil
### GetSmbAdDisabled

`func (o *Cluster) GetSmbAdDisabled() bool`

GetSmbAdDisabled returns the SmbAdDisabled field if non-nil, zero value otherwise.

### GetSmbAdDisabledOk

`func (o *Cluster) GetSmbAdDisabledOk() (*bool, bool)`

GetSmbAdDisabledOk returns a tuple with the SmbAdDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAdDisabled

`func (o *Cluster) SetSmbAdDisabled(v bool)`

SetSmbAdDisabled sets SmbAdDisabled field to given value.

### HasSmbAdDisabled

`func (o *Cluster) HasSmbAdDisabled() bool`

HasSmbAdDisabled returns a boolean if a field has been set.

### SetSmbAdDisabledNil

`func (o *Cluster) SetSmbAdDisabledNil(b bool)`

 SetSmbAdDisabledNil sets the value for SmbAdDisabled to be an explicit nil

### UnsetSmbAdDisabled
`func (o *Cluster) UnsetSmbAdDisabled()`

UnsetSmbAdDisabled ensures that no value is present for SmbAdDisabled, not even an explicit nil
### GetSmbMultichannelEnabled

`func (o *Cluster) GetSmbMultichannelEnabled() bool`

GetSmbMultichannelEnabled returns the SmbMultichannelEnabled field if non-nil, zero value otherwise.

### GetSmbMultichannelEnabledOk

`func (o *Cluster) GetSmbMultichannelEnabledOk() (*bool, bool)`

GetSmbMultichannelEnabledOk returns a tuple with the SmbMultichannelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMultichannelEnabled

`func (o *Cluster) SetSmbMultichannelEnabled(v bool)`

SetSmbMultichannelEnabled sets SmbMultichannelEnabled field to given value.

### HasSmbMultichannelEnabled

`func (o *Cluster) HasSmbMultichannelEnabled() bool`

HasSmbMultichannelEnabled returns a boolean if a field has been set.

### SetSmbMultichannelEnabledNil

`func (o *Cluster) SetSmbMultichannelEnabledNil(b bool)`

 SetSmbMultichannelEnabledNil sets the value for SmbMultichannelEnabled to be an explicit nil

### UnsetSmbMultichannelEnabled
`func (o *Cluster) UnsetSmbMultichannelEnabled()`

UnsetSmbMultichannelEnabled ensures that no value is present for SmbMultichannelEnabled, not even an explicit nil
### GetStats

`func (o *Cluster) GetStats() ClusterStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Cluster) GetStatsOk() (*ClusterStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Cluster) SetStats(v ClusterStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Cluster) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStigMode

`func (o *Cluster) GetStigMode() bool`

GetStigMode returns the StigMode field if non-nil, zero value otherwise.

### GetStigModeOk

`func (o *Cluster) GetStigModeOk() (*bool, bool)`

GetStigModeOk returns a tuple with the StigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStigMode

`func (o *Cluster) SetStigMode(v bool)`

SetStigMode sets StigMode field to given value.

### HasStigMode

`func (o *Cluster) HasStigMode() bool`

HasStigMode returns a boolean if a field has been set.

### SetStigModeNil

`func (o *Cluster) SetStigModeNil(b bool)`

 SetStigModeNil sets the value for StigMode to be an explicit nil

### UnsetStigMode
`func (o *Cluster) UnsetStigMode()`

UnsetStigMode ensures that no value is present for StigMode, not even an explicit nil
### GetSupportedConfig

`func (o *Cluster) GetSupportedConfig() SupportedConfig`

GetSupportedConfig returns the SupportedConfig field if non-nil, zero value otherwise.

### GetSupportedConfigOk

`func (o *Cluster) GetSupportedConfigOk() (*SupportedConfig, bool)`

GetSupportedConfigOk returns a tuple with the SupportedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedConfig

`func (o *Cluster) SetSupportedConfig(v SupportedConfig)`

SetSupportedConfig sets SupportedConfig field to given value.

### HasSupportedConfig

`func (o *Cluster) HasSupportedConfig() bool`

HasSupportedConfig returns a boolean if a field has been set.

### GetSyslogServers

`func (o *Cluster) GetSyslogServers() []OldSyslogServer`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *Cluster) GetSyslogServersOk() (*[]OldSyslogServer, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *Cluster) SetSyslogServers(v []OldSyslogServer)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *Cluster) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### SetSyslogServersNil

`func (o *Cluster) SetSyslogServersNil(b bool)`

 SetSyslogServersNil sets the value for SyslogServers to be an explicit nil

### UnsetSyslogServers
`func (o *Cluster) UnsetSyslogServers()`

UnsetSyslogServers ensures that no value is present for SyslogServers, not even an explicit nil
### GetTargetSoftwareVersion

`func (o *Cluster) GetTargetSoftwareVersion() string`

GetTargetSoftwareVersion returns the TargetSoftwareVersion field if non-nil, zero value otherwise.

### GetTargetSoftwareVersionOk

`func (o *Cluster) GetTargetSoftwareVersionOk() (*string, bool)`

GetTargetSoftwareVersionOk returns a tuple with the TargetSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSoftwareVersion

`func (o *Cluster) SetTargetSoftwareVersion(v string)`

SetTargetSoftwareVersion sets TargetSoftwareVersion field to given value.

### HasTargetSoftwareVersion

`func (o *Cluster) HasTargetSoftwareVersion() bool`

HasTargetSoftwareVersion returns a boolean if a field has been set.

### SetTargetSoftwareVersionNil

`func (o *Cluster) SetTargetSoftwareVersionNil(b bool)`

 SetTargetSoftwareVersionNil sets the value for TargetSoftwareVersion to be an explicit nil

### UnsetTargetSoftwareVersion
`func (o *Cluster) UnsetTargetSoftwareVersion()`

UnsetTargetSoftwareVersion ensures that no value is present for TargetSoftwareVersion, not even an explicit nil
### GetTenantViewboxSharingEnabled

`func (o *Cluster) GetTenantViewboxSharingEnabled() bool`

GetTenantViewboxSharingEnabled returns the TenantViewboxSharingEnabled field if non-nil, zero value otherwise.

### GetTenantViewboxSharingEnabledOk

`func (o *Cluster) GetTenantViewboxSharingEnabledOk() (*bool, bool)`

GetTenantViewboxSharingEnabledOk returns a tuple with the TenantViewboxSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantViewboxSharingEnabled

`func (o *Cluster) SetTenantViewboxSharingEnabled(v bool)`

SetTenantViewboxSharingEnabled sets TenantViewboxSharingEnabled field to given value.

### HasTenantViewboxSharingEnabled

`func (o *Cluster) HasTenantViewboxSharingEnabled() bool`

HasTenantViewboxSharingEnabled returns a boolean if a field has been set.

### SetTenantViewboxSharingEnabledNil

`func (o *Cluster) SetTenantViewboxSharingEnabledNil(b bool)`

 SetTenantViewboxSharingEnabledNil sets the value for TenantViewboxSharingEnabled to be an explicit nil

### UnsetTenantViewboxSharingEnabled
`func (o *Cluster) UnsetTenantViewboxSharingEnabled()`

UnsetTenantViewboxSharingEnabled ensures that no value is present for TenantViewboxSharingEnabled, not even an explicit nil
### GetTimezone

`func (o *Cluster) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Cluster) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Cluster) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Cluster) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *Cluster) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *Cluster) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTurboMode

`func (o *Cluster) GetTurboMode() bool`

GetTurboMode returns the TurboMode field if non-nil, zero value otherwise.

### GetTurboModeOk

`func (o *Cluster) GetTurboModeOk() (*bool, bool)`

GetTurboModeOk returns a tuple with the TurboMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboMode

`func (o *Cluster) SetTurboMode(v bool)`

SetTurboMode sets TurboMode field to given value.

### HasTurboMode

`func (o *Cluster) HasTurboMode() bool`

HasTurboMode returns a boolean if a field has been set.

### SetTurboModeNil

`func (o *Cluster) SetTurboModeNil(b bool)`

 SetTurboModeNil sets the value for TurboMode to be an explicit nil

### UnsetTurboMode
`func (o *Cluster) UnsetTurboMode()`

UnsetTurboMode ensures that no value is present for TurboMode, not even an explicit nil
### GetUsedMetadataSpacePct

`func (o *Cluster) GetUsedMetadataSpacePct() float64`

GetUsedMetadataSpacePct returns the UsedMetadataSpacePct field if non-nil, zero value otherwise.

### GetUsedMetadataSpacePctOk

`func (o *Cluster) GetUsedMetadataSpacePctOk() (*float64, bool)`

GetUsedMetadataSpacePctOk returns a tuple with the UsedMetadataSpacePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMetadataSpacePct

`func (o *Cluster) SetUsedMetadataSpacePct(v float64)`

SetUsedMetadataSpacePct sets UsedMetadataSpacePct field to given value.

### HasUsedMetadataSpacePct

`func (o *Cluster) HasUsedMetadataSpacePct() bool`

HasUsedMetadataSpacePct returns a boolean if a field has been set.

### SetUsedMetadataSpacePctNil

`func (o *Cluster) SetUsedMetadataSpacePctNil(b bool)`

 SetUsedMetadataSpacePctNil sets the value for UsedMetadataSpacePct to be an explicit nil

### UnsetUsedMetadataSpacePct
`func (o *Cluster) UnsetUsedMetadataSpacePct()`

UnsetUsedMetadataSpacePct ensures that no value is present for UsedMetadataSpacePct, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


