# UpdateClusterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppsSubnet** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 
**BannerEnabled** | Pointer to **NullableBool** | Specifies whether UI banner is enabled on the cluster or not. When banner is enabled, UI will make an additional API call to fetch the banner and show at the login page. | [optional] 
**ClusterAuditLogConfig** | Pointer to [**ClusterAuditLogConfiguration**](ClusterAuditLogConfiguration.md) |  | [optional] 
**DnsServerIps** | Pointer to **[]string** | Array of IP Addresses of DNS Servers.  Specifies the IP addresses of the DNS Servers used by the Cohesity Cluster. | [optional] 
**DomainNames** | Pointer to **[]string** | Array of Domain Names.  The first domain name specified in the array is the fully qualified domain name assigned to the Cohesity Cluster. Any additional domain names specified are used for the domain search list for hostname look-up. | [optional] 
**EnableActiveMonitoring** | Pointer to **NullableBool** | Specifies if Cohesity can receive monitoring information from the Cohesity Cluster. If &#39;true&#39;, remote monitoring of the Cohesity Cluster is allowed. | [optional] 
**EnableUpgradePkgPolling** | Pointer to **NullableBool** | If &#39;true&#39;, Cohesity&#39;s upgrade server is polled for new releases. | [optional] 
**EncryptionKeyRotationPeriodSecs** | Pointer to **NullableInt64** | Specifies the period of time (in seconds) when encryption keys are rotated. By default, the encryption keys are rotated every 77760000 seconds (30 days). | [optional] 
**FaultToleranceLevel** | Pointer to **NullableString** | Specifies the level which &#39;MetadataFaultToleranceFactor&#39; applies to. &#39;kNode&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Node level. &#39;kChassis&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Chassis level. &#39;kRack&#39; indicates &#39;MetadataFaultToleranceFactor&#39; applies to Rack level. | [optional] 
**FilerAuditLogConfig** | Pointer to [**FilerAuditLogConfiguration**](FilerAuditLogConfiguration.md) |  | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway IP address. | [optional] 
**GoogleAnalyticsEnabled** | Pointer to **NullableBool** | Specifies whether Google Analytics is enabled. | [optional] 
**IsDocumentationLocal** | Pointer to **NullableBool** | Specifies what version of the documentation is used. If &#39;true&#39;, the version of documentation stored locally on the Cohesity Cluster is used. If &#39;false&#39;, the documentation stored on a Cohesity Web Server is used. The default is &#39;false&#39;. Cohesity recommends accessing the Help from the Cohesity Web site which provides the newest and most complete version of Help. | [optional] 
**LanguageLocale** | Pointer to **NullableString** | Specifies the language and locale for this Cohesity Cluster. | [optional] 
**LocalAuthDomainName** | Pointer to **NullableString** | Domain name for SMB local authentication. | [optional] 
**LocalGroupsEnabled** | Pointer to **NullableBool** | Specifies whether to enable local groups on cluster. Once it is enabled, it cannot be disabled. | [optional] 
**MetadataFaultToleranceFactor** | Pointer to **NullableInt32** | Specifies metadata fault tolerance setting for the cluster. This denotes the number of simultaneous failures[node] supported by metadata services like gandalf and scribe. | [optional] 
**MultiTenancyEnabled** | Pointer to **NullableBool** | Specifies if multi tenancy is enabled in the cluster. Authentication &amp; Authorization will always use tenant_id, however, some UI elements may be disabled when multi tenancy is disabled. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Cohesity Cluster. | [optional] 
**NtpSettings** | Pointer to [**NtpSettingsConfig**](NtpSettingsConfig.md) |  | [optional] 
**PcieSsdTierRebalanceDelaySecs** | Pointer to **NullableInt32** | Specifies the rebalance delay in seconds for cluster PcieSSD storage tier. | [optional] 
**ReverseTunnelEnabled** | Pointer to **NullableBool** | If &#39;true&#39;, Cohesity&#39;s Remote Tunnel is enabled. Cohesity can access the Cluster and provide remote assistance via a Remote Tunnel. | [optional] 
**ReverseTunnelEndTimeMsecs** | Pointer to **NullableInt64** | ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since epoch until when the reverse tunnel will stay enabled. | [optional] 
**SmbAdDisabled** | Pointer to **NullableBool** | Specifies if Active Directory should be disabled for authentication of SMB shares. If &#39;true&#39;, Active Directory is disabled. | [optional] 
**SmbMultichannelEnabled** | Pointer to **NullableBool** | Specifies whether SMB multichannel is enabled on the cluster. When this is set to true, then any SMB3 multichannel enabled client can establish multiple TCP connection per session to the Server. | [optional] 
**StigMode** | Pointer to **NullableBool** | Specifies if STIG mode is enabled or not. | [optional] 
**SyslogServers** | Pointer to [**[]OldSyslogServer**](OldSyslogServer.md) | Syslog servers. | [optional] 
**TenantViewboxSharingEnabled** | Pointer to **NullableBool** | In case multi tenancy is enabled, this flag controls whether multiple tenants can be placed on the same viewbox. Once set to true, this flag should never become false. | [optional] 
**Timezone** | Pointer to **NullableString** | Specifies the timezone to use for showing time in emails, reports, filer audit logs, etc. | [optional] 
**TurboMode** | Pointer to **NullableBool** | Specifies if the cluster is in Turbo mode. | [optional] 

## Methods

### NewUpdateClusterParams

`func NewUpdateClusterParams() *UpdateClusterParams`

NewUpdateClusterParams instantiates a new UpdateClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterParamsWithDefaults

`func NewUpdateClusterParamsWithDefaults() *UpdateClusterParams`

NewUpdateClusterParamsWithDefaults instantiates a new UpdateClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsSubnet

`func (o *UpdateClusterParams) GetAppsSubnet() Subnet`

GetAppsSubnet returns the AppsSubnet field if non-nil, zero value otherwise.

### GetAppsSubnetOk

`func (o *UpdateClusterParams) GetAppsSubnetOk() (*Subnet, bool)`

GetAppsSubnetOk returns a tuple with the AppsSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsSubnet

`func (o *UpdateClusterParams) SetAppsSubnet(v Subnet)`

SetAppsSubnet sets AppsSubnet field to given value.

### HasAppsSubnet

`func (o *UpdateClusterParams) HasAppsSubnet() bool`

HasAppsSubnet returns a boolean if a field has been set.

### GetBannerEnabled

`func (o *UpdateClusterParams) GetBannerEnabled() bool`

GetBannerEnabled returns the BannerEnabled field if non-nil, zero value otherwise.

### GetBannerEnabledOk

`func (o *UpdateClusterParams) GetBannerEnabledOk() (*bool, bool)`

GetBannerEnabledOk returns a tuple with the BannerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerEnabled

`func (o *UpdateClusterParams) SetBannerEnabled(v bool)`

SetBannerEnabled sets BannerEnabled field to given value.

### HasBannerEnabled

`func (o *UpdateClusterParams) HasBannerEnabled() bool`

HasBannerEnabled returns a boolean if a field has been set.

### SetBannerEnabledNil

`func (o *UpdateClusterParams) SetBannerEnabledNil(b bool)`

 SetBannerEnabledNil sets the value for BannerEnabled to be an explicit nil

### UnsetBannerEnabled
`func (o *UpdateClusterParams) UnsetBannerEnabled()`

UnsetBannerEnabled ensures that no value is present for BannerEnabled, not even an explicit nil
### GetClusterAuditLogConfig

`func (o *UpdateClusterParams) GetClusterAuditLogConfig() ClusterAuditLogConfiguration`

GetClusterAuditLogConfig returns the ClusterAuditLogConfig field if non-nil, zero value otherwise.

### GetClusterAuditLogConfigOk

`func (o *UpdateClusterParams) GetClusterAuditLogConfigOk() (*ClusterAuditLogConfiguration, bool)`

GetClusterAuditLogConfigOk returns a tuple with the ClusterAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAuditLogConfig

`func (o *UpdateClusterParams) SetClusterAuditLogConfig(v ClusterAuditLogConfiguration)`

SetClusterAuditLogConfig sets ClusterAuditLogConfig field to given value.

### HasClusterAuditLogConfig

`func (o *UpdateClusterParams) HasClusterAuditLogConfig() bool`

HasClusterAuditLogConfig returns a boolean if a field has been set.

### GetDnsServerIps

`func (o *UpdateClusterParams) GetDnsServerIps() []string`

GetDnsServerIps returns the DnsServerIps field if non-nil, zero value otherwise.

### GetDnsServerIpsOk

`func (o *UpdateClusterParams) GetDnsServerIpsOk() (*[]string, bool)`

GetDnsServerIpsOk returns a tuple with the DnsServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerIps

`func (o *UpdateClusterParams) SetDnsServerIps(v []string)`

SetDnsServerIps sets DnsServerIps field to given value.

### HasDnsServerIps

`func (o *UpdateClusterParams) HasDnsServerIps() bool`

HasDnsServerIps returns a boolean if a field has been set.

### SetDnsServerIpsNil

`func (o *UpdateClusterParams) SetDnsServerIpsNil(b bool)`

 SetDnsServerIpsNil sets the value for DnsServerIps to be an explicit nil

### UnsetDnsServerIps
`func (o *UpdateClusterParams) UnsetDnsServerIps()`

UnsetDnsServerIps ensures that no value is present for DnsServerIps, not even an explicit nil
### GetDomainNames

`func (o *UpdateClusterParams) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *UpdateClusterParams) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *UpdateClusterParams) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *UpdateClusterParams) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *UpdateClusterParams) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *UpdateClusterParams) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetEnableActiveMonitoring

`func (o *UpdateClusterParams) GetEnableActiveMonitoring() bool`

GetEnableActiveMonitoring returns the EnableActiveMonitoring field if non-nil, zero value otherwise.

### GetEnableActiveMonitoringOk

`func (o *UpdateClusterParams) GetEnableActiveMonitoringOk() (*bool, bool)`

GetEnableActiveMonitoringOk returns a tuple with the EnableActiveMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableActiveMonitoring

`func (o *UpdateClusterParams) SetEnableActiveMonitoring(v bool)`

SetEnableActiveMonitoring sets EnableActiveMonitoring field to given value.

### HasEnableActiveMonitoring

`func (o *UpdateClusterParams) HasEnableActiveMonitoring() bool`

HasEnableActiveMonitoring returns a boolean if a field has been set.

### SetEnableActiveMonitoringNil

`func (o *UpdateClusterParams) SetEnableActiveMonitoringNil(b bool)`

 SetEnableActiveMonitoringNil sets the value for EnableActiveMonitoring to be an explicit nil

### UnsetEnableActiveMonitoring
`func (o *UpdateClusterParams) UnsetEnableActiveMonitoring()`

UnsetEnableActiveMonitoring ensures that no value is present for EnableActiveMonitoring, not even an explicit nil
### GetEnableUpgradePkgPolling

`func (o *UpdateClusterParams) GetEnableUpgradePkgPolling() bool`

GetEnableUpgradePkgPolling returns the EnableUpgradePkgPolling field if non-nil, zero value otherwise.

### GetEnableUpgradePkgPollingOk

`func (o *UpdateClusterParams) GetEnableUpgradePkgPollingOk() (*bool, bool)`

GetEnableUpgradePkgPollingOk returns a tuple with the EnableUpgradePkgPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpgradePkgPolling

`func (o *UpdateClusterParams) SetEnableUpgradePkgPolling(v bool)`

SetEnableUpgradePkgPolling sets EnableUpgradePkgPolling field to given value.

### HasEnableUpgradePkgPolling

`func (o *UpdateClusterParams) HasEnableUpgradePkgPolling() bool`

HasEnableUpgradePkgPolling returns a boolean if a field has been set.

### SetEnableUpgradePkgPollingNil

`func (o *UpdateClusterParams) SetEnableUpgradePkgPollingNil(b bool)`

 SetEnableUpgradePkgPollingNil sets the value for EnableUpgradePkgPolling to be an explicit nil

### UnsetEnableUpgradePkgPolling
`func (o *UpdateClusterParams) UnsetEnableUpgradePkgPolling()`

UnsetEnableUpgradePkgPolling ensures that no value is present for EnableUpgradePkgPolling, not even an explicit nil
### GetEncryptionKeyRotationPeriodSecs

`func (o *UpdateClusterParams) GetEncryptionKeyRotationPeriodSecs() int64`

GetEncryptionKeyRotationPeriodSecs returns the EncryptionKeyRotationPeriodSecs field if non-nil, zero value otherwise.

### GetEncryptionKeyRotationPeriodSecsOk

`func (o *UpdateClusterParams) GetEncryptionKeyRotationPeriodSecsOk() (*int64, bool)`

GetEncryptionKeyRotationPeriodSecsOk returns a tuple with the EncryptionKeyRotationPeriodSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyRotationPeriodSecs

`func (o *UpdateClusterParams) SetEncryptionKeyRotationPeriodSecs(v int64)`

SetEncryptionKeyRotationPeriodSecs sets EncryptionKeyRotationPeriodSecs field to given value.

### HasEncryptionKeyRotationPeriodSecs

`func (o *UpdateClusterParams) HasEncryptionKeyRotationPeriodSecs() bool`

HasEncryptionKeyRotationPeriodSecs returns a boolean if a field has been set.

### SetEncryptionKeyRotationPeriodSecsNil

`func (o *UpdateClusterParams) SetEncryptionKeyRotationPeriodSecsNil(b bool)`

 SetEncryptionKeyRotationPeriodSecsNil sets the value for EncryptionKeyRotationPeriodSecs to be an explicit nil

### UnsetEncryptionKeyRotationPeriodSecs
`func (o *UpdateClusterParams) UnsetEncryptionKeyRotationPeriodSecs()`

UnsetEncryptionKeyRotationPeriodSecs ensures that no value is present for EncryptionKeyRotationPeriodSecs, not even an explicit nil
### GetFaultToleranceLevel

`func (o *UpdateClusterParams) GetFaultToleranceLevel() string`

GetFaultToleranceLevel returns the FaultToleranceLevel field if non-nil, zero value otherwise.

### GetFaultToleranceLevelOk

`func (o *UpdateClusterParams) GetFaultToleranceLevelOk() (*string, bool)`

GetFaultToleranceLevelOk returns a tuple with the FaultToleranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultToleranceLevel

`func (o *UpdateClusterParams) SetFaultToleranceLevel(v string)`

SetFaultToleranceLevel sets FaultToleranceLevel field to given value.

### HasFaultToleranceLevel

`func (o *UpdateClusterParams) HasFaultToleranceLevel() bool`

HasFaultToleranceLevel returns a boolean if a field has been set.

### SetFaultToleranceLevelNil

`func (o *UpdateClusterParams) SetFaultToleranceLevelNil(b bool)`

 SetFaultToleranceLevelNil sets the value for FaultToleranceLevel to be an explicit nil

### UnsetFaultToleranceLevel
`func (o *UpdateClusterParams) UnsetFaultToleranceLevel()`

UnsetFaultToleranceLevel ensures that no value is present for FaultToleranceLevel, not even an explicit nil
### GetFilerAuditLogConfig

`func (o *UpdateClusterParams) GetFilerAuditLogConfig() FilerAuditLogConfiguration`

GetFilerAuditLogConfig returns the FilerAuditLogConfig field if non-nil, zero value otherwise.

### GetFilerAuditLogConfigOk

`func (o *UpdateClusterParams) GetFilerAuditLogConfigOk() (*FilerAuditLogConfiguration, bool)`

GetFilerAuditLogConfigOk returns a tuple with the FilerAuditLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerAuditLogConfig

`func (o *UpdateClusterParams) SetFilerAuditLogConfig(v FilerAuditLogConfiguration)`

SetFilerAuditLogConfig sets FilerAuditLogConfig field to given value.

### HasFilerAuditLogConfig

`func (o *UpdateClusterParams) HasFilerAuditLogConfig() bool`

HasFilerAuditLogConfig returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateClusterParams) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateClusterParams) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateClusterParams) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateClusterParams) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *UpdateClusterParams) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *UpdateClusterParams) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGoogleAnalyticsEnabled

`func (o *UpdateClusterParams) GetGoogleAnalyticsEnabled() bool`

GetGoogleAnalyticsEnabled returns the GoogleAnalyticsEnabled field if non-nil, zero value otherwise.

### GetGoogleAnalyticsEnabledOk

`func (o *UpdateClusterParams) GetGoogleAnalyticsEnabledOk() (*bool, bool)`

GetGoogleAnalyticsEnabledOk returns a tuple with the GoogleAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsEnabled

`func (o *UpdateClusterParams) SetGoogleAnalyticsEnabled(v bool)`

SetGoogleAnalyticsEnabled sets GoogleAnalyticsEnabled field to given value.

### HasGoogleAnalyticsEnabled

`func (o *UpdateClusterParams) HasGoogleAnalyticsEnabled() bool`

HasGoogleAnalyticsEnabled returns a boolean if a field has been set.

### SetGoogleAnalyticsEnabledNil

`func (o *UpdateClusterParams) SetGoogleAnalyticsEnabledNil(b bool)`

 SetGoogleAnalyticsEnabledNil sets the value for GoogleAnalyticsEnabled to be an explicit nil

### UnsetGoogleAnalyticsEnabled
`func (o *UpdateClusterParams) UnsetGoogleAnalyticsEnabled()`

UnsetGoogleAnalyticsEnabled ensures that no value is present for GoogleAnalyticsEnabled, not even an explicit nil
### GetIsDocumentationLocal

`func (o *UpdateClusterParams) GetIsDocumentationLocal() bool`

GetIsDocumentationLocal returns the IsDocumentationLocal field if non-nil, zero value otherwise.

### GetIsDocumentationLocalOk

`func (o *UpdateClusterParams) GetIsDocumentationLocalOk() (*bool, bool)`

GetIsDocumentationLocalOk returns a tuple with the IsDocumentationLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDocumentationLocal

`func (o *UpdateClusterParams) SetIsDocumentationLocal(v bool)`

SetIsDocumentationLocal sets IsDocumentationLocal field to given value.

### HasIsDocumentationLocal

`func (o *UpdateClusterParams) HasIsDocumentationLocal() bool`

HasIsDocumentationLocal returns a boolean if a field has been set.

### SetIsDocumentationLocalNil

`func (o *UpdateClusterParams) SetIsDocumentationLocalNil(b bool)`

 SetIsDocumentationLocalNil sets the value for IsDocumentationLocal to be an explicit nil

### UnsetIsDocumentationLocal
`func (o *UpdateClusterParams) UnsetIsDocumentationLocal()`

UnsetIsDocumentationLocal ensures that no value is present for IsDocumentationLocal, not even an explicit nil
### GetLanguageLocale

`func (o *UpdateClusterParams) GetLanguageLocale() string`

GetLanguageLocale returns the LanguageLocale field if non-nil, zero value otherwise.

### GetLanguageLocaleOk

`func (o *UpdateClusterParams) GetLanguageLocaleOk() (*string, bool)`

GetLanguageLocaleOk returns a tuple with the LanguageLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageLocale

`func (o *UpdateClusterParams) SetLanguageLocale(v string)`

SetLanguageLocale sets LanguageLocale field to given value.

### HasLanguageLocale

`func (o *UpdateClusterParams) HasLanguageLocale() bool`

HasLanguageLocale returns a boolean if a field has been set.

### SetLanguageLocaleNil

`func (o *UpdateClusterParams) SetLanguageLocaleNil(b bool)`

 SetLanguageLocaleNil sets the value for LanguageLocale to be an explicit nil

### UnsetLanguageLocale
`func (o *UpdateClusterParams) UnsetLanguageLocale()`

UnsetLanguageLocale ensures that no value is present for LanguageLocale, not even an explicit nil
### GetLocalAuthDomainName

`func (o *UpdateClusterParams) GetLocalAuthDomainName() string`

GetLocalAuthDomainName returns the LocalAuthDomainName field if non-nil, zero value otherwise.

### GetLocalAuthDomainNameOk

`func (o *UpdateClusterParams) GetLocalAuthDomainNameOk() (*string, bool)`

GetLocalAuthDomainNameOk returns a tuple with the LocalAuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuthDomainName

`func (o *UpdateClusterParams) SetLocalAuthDomainName(v string)`

SetLocalAuthDomainName sets LocalAuthDomainName field to given value.

### HasLocalAuthDomainName

`func (o *UpdateClusterParams) HasLocalAuthDomainName() bool`

HasLocalAuthDomainName returns a boolean if a field has been set.

### SetLocalAuthDomainNameNil

`func (o *UpdateClusterParams) SetLocalAuthDomainNameNil(b bool)`

 SetLocalAuthDomainNameNil sets the value for LocalAuthDomainName to be an explicit nil

### UnsetLocalAuthDomainName
`func (o *UpdateClusterParams) UnsetLocalAuthDomainName()`

UnsetLocalAuthDomainName ensures that no value is present for LocalAuthDomainName, not even an explicit nil
### GetLocalGroupsEnabled

`func (o *UpdateClusterParams) GetLocalGroupsEnabled() bool`

GetLocalGroupsEnabled returns the LocalGroupsEnabled field if non-nil, zero value otherwise.

### GetLocalGroupsEnabledOk

`func (o *UpdateClusterParams) GetLocalGroupsEnabledOk() (*bool, bool)`

GetLocalGroupsEnabledOk returns a tuple with the LocalGroupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupsEnabled

`func (o *UpdateClusterParams) SetLocalGroupsEnabled(v bool)`

SetLocalGroupsEnabled sets LocalGroupsEnabled field to given value.

### HasLocalGroupsEnabled

`func (o *UpdateClusterParams) HasLocalGroupsEnabled() bool`

HasLocalGroupsEnabled returns a boolean if a field has been set.

### SetLocalGroupsEnabledNil

`func (o *UpdateClusterParams) SetLocalGroupsEnabledNil(b bool)`

 SetLocalGroupsEnabledNil sets the value for LocalGroupsEnabled to be an explicit nil

### UnsetLocalGroupsEnabled
`func (o *UpdateClusterParams) UnsetLocalGroupsEnabled()`

UnsetLocalGroupsEnabled ensures that no value is present for LocalGroupsEnabled, not even an explicit nil
### GetMetadataFaultToleranceFactor

`func (o *UpdateClusterParams) GetMetadataFaultToleranceFactor() int32`

GetMetadataFaultToleranceFactor returns the MetadataFaultToleranceFactor field if non-nil, zero value otherwise.

### GetMetadataFaultToleranceFactorOk

`func (o *UpdateClusterParams) GetMetadataFaultToleranceFactorOk() (*int32, bool)`

GetMetadataFaultToleranceFactorOk returns a tuple with the MetadataFaultToleranceFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFaultToleranceFactor

`func (o *UpdateClusterParams) SetMetadataFaultToleranceFactor(v int32)`

SetMetadataFaultToleranceFactor sets MetadataFaultToleranceFactor field to given value.

### HasMetadataFaultToleranceFactor

`func (o *UpdateClusterParams) HasMetadataFaultToleranceFactor() bool`

HasMetadataFaultToleranceFactor returns a boolean if a field has been set.

### SetMetadataFaultToleranceFactorNil

`func (o *UpdateClusterParams) SetMetadataFaultToleranceFactorNil(b bool)`

 SetMetadataFaultToleranceFactorNil sets the value for MetadataFaultToleranceFactor to be an explicit nil

### UnsetMetadataFaultToleranceFactor
`func (o *UpdateClusterParams) UnsetMetadataFaultToleranceFactor()`

UnsetMetadataFaultToleranceFactor ensures that no value is present for MetadataFaultToleranceFactor, not even an explicit nil
### GetMultiTenancyEnabled

`func (o *UpdateClusterParams) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *UpdateClusterParams) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *UpdateClusterParams) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *UpdateClusterParams) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *UpdateClusterParams) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *UpdateClusterParams) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetName

`func (o *UpdateClusterParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateClusterParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateClusterParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNtpSettings

`func (o *UpdateClusterParams) GetNtpSettings() NtpSettingsConfig`

GetNtpSettings returns the NtpSettings field if non-nil, zero value otherwise.

### GetNtpSettingsOk

`func (o *UpdateClusterParams) GetNtpSettingsOk() (*NtpSettingsConfig, bool)`

GetNtpSettingsOk returns a tuple with the NtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpSettings

`func (o *UpdateClusterParams) SetNtpSettings(v NtpSettingsConfig)`

SetNtpSettings sets NtpSettings field to given value.

### HasNtpSettings

`func (o *UpdateClusterParams) HasNtpSettings() bool`

HasNtpSettings returns a boolean if a field has been set.

### GetPcieSsdTierRebalanceDelaySecs

`func (o *UpdateClusterParams) GetPcieSsdTierRebalanceDelaySecs() int32`

GetPcieSsdTierRebalanceDelaySecs returns the PcieSsdTierRebalanceDelaySecs field if non-nil, zero value otherwise.

### GetPcieSsdTierRebalanceDelaySecsOk

`func (o *UpdateClusterParams) GetPcieSsdTierRebalanceDelaySecsOk() (*int32, bool)`

GetPcieSsdTierRebalanceDelaySecsOk returns a tuple with the PcieSsdTierRebalanceDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSsdTierRebalanceDelaySecs

`func (o *UpdateClusterParams) SetPcieSsdTierRebalanceDelaySecs(v int32)`

SetPcieSsdTierRebalanceDelaySecs sets PcieSsdTierRebalanceDelaySecs field to given value.

### HasPcieSsdTierRebalanceDelaySecs

`func (o *UpdateClusterParams) HasPcieSsdTierRebalanceDelaySecs() bool`

HasPcieSsdTierRebalanceDelaySecs returns a boolean if a field has been set.

### SetPcieSsdTierRebalanceDelaySecsNil

`func (o *UpdateClusterParams) SetPcieSsdTierRebalanceDelaySecsNil(b bool)`

 SetPcieSsdTierRebalanceDelaySecsNil sets the value for PcieSsdTierRebalanceDelaySecs to be an explicit nil

### UnsetPcieSsdTierRebalanceDelaySecs
`func (o *UpdateClusterParams) UnsetPcieSsdTierRebalanceDelaySecs()`

UnsetPcieSsdTierRebalanceDelaySecs ensures that no value is present for PcieSsdTierRebalanceDelaySecs, not even an explicit nil
### GetReverseTunnelEnabled

`func (o *UpdateClusterParams) GetReverseTunnelEnabled() bool`

GetReverseTunnelEnabled returns the ReverseTunnelEnabled field if non-nil, zero value otherwise.

### GetReverseTunnelEnabledOk

`func (o *UpdateClusterParams) GetReverseTunnelEnabledOk() (*bool, bool)`

GetReverseTunnelEnabledOk returns a tuple with the ReverseTunnelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseTunnelEnabled

`func (o *UpdateClusterParams) SetReverseTunnelEnabled(v bool)`

SetReverseTunnelEnabled sets ReverseTunnelEnabled field to given value.

### HasReverseTunnelEnabled

`func (o *UpdateClusterParams) HasReverseTunnelEnabled() bool`

HasReverseTunnelEnabled returns a boolean if a field has been set.

### SetReverseTunnelEnabledNil

`func (o *UpdateClusterParams) SetReverseTunnelEnabledNil(b bool)`

 SetReverseTunnelEnabledNil sets the value for ReverseTunnelEnabled to be an explicit nil

### UnsetReverseTunnelEnabled
`func (o *UpdateClusterParams) UnsetReverseTunnelEnabled()`

UnsetReverseTunnelEnabled ensures that no value is present for ReverseTunnelEnabled, not even an explicit nil
### GetReverseTunnelEndTimeMsecs

`func (o *UpdateClusterParams) GetReverseTunnelEndTimeMsecs() int64`

GetReverseTunnelEndTimeMsecs returns the ReverseTunnelEndTimeMsecs field if non-nil, zero value otherwise.

### GetReverseTunnelEndTimeMsecsOk

`func (o *UpdateClusterParams) GetReverseTunnelEndTimeMsecsOk() (*int64, bool)`

GetReverseTunnelEndTimeMsecsOk returns a tuple with the ReverseTunnelEndTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseTunnelEndTimeMsecs

`func (o *UpdateClusterParams) SetReverseTunnelEndTimeMsecs(v int64)`

SetReverseTunnelEndTimeMsecs sets ReverseTunnelEndTimeMsecs field to given value.

### HasReverseTunnelEndTimeMsecs

`func (o *UpdateClusterParams) HasReverseTunnelEndTimeMsecs() bool`

HasReverseTunnelEndTimeMsecs returns a boolean if a field has been set.

### SetReverseTunnelEndTimeMsecsNil

`func (o *UpdateClusterParams) SetReverseTunnelEndTimeMsecsNil(b bool)`

 SetReverseTunnelEndTimeMsecsNil sets the value for ReverseTunnelEndTimeMsecs to be an explicit nil

### UnsetReverseTunnelEndTimeMsecs
`func (o *UpdateClusterParams) UnsetReverseTunnelEndTimeMsecs()`

UnsetReverseTunnelEndTimeMsecs ensures that no value is present for ReverseTunnelEndTimeMsecs, not even an explicit nil
### GetSmbAdDisabled

`func (o *UpdateClusterParams) GetSmbAdDisabled() bool`

GetSmbAdDisabled returns the SmbAdDisabled field if non-nil, zero value otherwise.

### GetSmbAdDisabledOk

`func (o *UpdateClusterParams) GetSmbAdDisabledOk() (*bool, bool)`

GetSmbAdDisabledOk returns a tuple with the SmbAdDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAdDisabled

`func (o *UpdateClusterParams) SetSmbAdDisabled(v bool)`

SetSmbAdDisabled sets SmbAdDisabled field to given value.

### HasSmbAdDisabled

`func (o *UpdateClusterParams) HasSmbAdDisabled() bool`

HasSmbAdDisabled returns a boolean if a field has been set.

### SetSmbAdDisabledNil

`func (o *UpdateClusterParams) SetSmbAdDisabledNil(b bool)`

 SetSmbAdDisabledNil sets the value for SmbAdDisabled to be an explicit nil

### UnsetSmbAdDisabled
`func (o *UpdateClusterParams) UnsetSmbAdDisabled()`

UnsetSmbAdDisabled ensures that no value is present for SmbAdDisabled, not even an explicit nil
### GetSmbMultichannelEnabled

`func (o *UpdateClusterParams) GetSmbMultichannelEnabled() bool`

GetSmbMultichannelEnabled returns the SmbMultichannelEnabled field if non-nil, zero value otherwise.

### GetSmbMultichannelEnabledOk

`func (o *UpdateClusterParams) GetSmbMultichannelEnabledOk() (*bool, bool)`

GetSmbMultichannelEnabledOk returns a tuple with the SmbMultichannelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMultichannelEnabled

`func (o *UpdateClusterParams) SetSmbMultichannelEnabled(v bool)`

SetSmbMultichannelEnabled sets SmbMultichannelEnabled field to given value.

### HasSmbMultichannelEnabled

`func (o *UpdateClusterParams) HasSmbMultichannelEnabled() bool`

HasSmbMultichannelEnabled returns a boolean if a field has been set.

### SetSmbMultichannelEnabledNil

`func (o *UpdateClusterParams) SetSmbMultichannelEnabledNil(b bool)`

 SetSmbMultichannelEnabledNil sets the value for SmbMultichannelEnabled to be an explicit nil

### UnsetSmbMultichannelEnabled
`func (o *UpdateClusterParams) UnsetSmbMultichannelEnabled()`

UnsetSmbMultichannelEnabled ensures that no value is present for SmbMultichannelEnabled, not even an explicit nil
### GetStigMode

`func (o *UpdateClusterParams) GetStigMode() bool`

GetStigMode returns the StigMode field if non-nil, zero value otherwise.

### GetStigModeOk

`func (o *UpdateClusterParams) GetStigModeOk() (*bool, bool)`

GetStigModeOk returns a tuple with the StigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStigMode

`func (o *UpdateClusterParams) SetStigMode(v bool)`

SetStigMode sets StigMode field to given value.

### HasStigMode

`func (o *UpdateClusterParams) HasStigMode() bool`

HasStigMode returns a boolean if a field has been set.

### SetStigModeNil

`func (o *UpdateClusterParams) SetStigModeNil(b bool)`

 SetStigModeNil sets the value for StigMode to be an explicit nil

### UnsetStigMode
`func (o *UpdateClusterParams) UnsetStigMode()`

UnsetStigMode ensures that no value is present for StigMode, not even an explicit nil
### GetSyslogServers

`func (o *UpdateClusterParams) GetSyslogServers() []OldSyslogServer`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *UpdateClusterParams) GetSyslogServersOk() (*[]OldSyslogServer, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *UpdateClusterParams) SetSyslogServers(v []OldSyslogServer)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *UpdateClusterParams) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### SetSyslogServersNil

`func (o *UpdateClusterParams) SetSyslogServersNil(b bool)`

 SetSyslogServersNil sets the value for SyslogServers to be an explicit nil

### UnsetSyslogServers
`func (o *UpdateClusterParams) UnsetSyslogServers()`

UnsetSyslogServers ensures that no value is present for SyslogServers, not even an explicit nil
### GetTenantViewboxSharingEnabled

`func (o *UpdateClusterParams) GetTenantViewboxSharingEnabled() bool`

GetTenantViewboxSharingEnabled returns the TenantViewboxSharingEnabled field if non-nil, zero value otherwise.

### GetTenantViewboxSharingEnabledOk

`func (o *UpdateClusterParams) GetTenantViewboxSharingEnabledOk() (*bool, bool)`

GetTenantViewboxSharingEnabledOk returns a tuple with the TenantViewboxSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantViewboxSharingEnabled

`func (o *UpdateClusterParams) SetTenantViewboxSharingEnabled(v bool)`

SetTenantViewboxSharingEnabled sets TenantViewboxSharingEnabled field to given value.

### HasTenantViewboxSharingEnabled

`func (o *UpdateClusterParams) HasTenantViewboxSharingEnabled() bool`

HasTenantViewboxSharingEnabled returns a boolean if a field has been set.

### SetTenantViewboxSharingEnabledNil

`func (o *UpdateClusterParams) SetTenantViewboxSharingEnabledNil(b bool)`

 SetTenantViewboxSharingEnabledNil sets the value for TenantViewboxSharingEnabled to be an explicit nil

### UnsetTenantViewboxSharingEnabled
`func (o *UpdateClusterParams) UnsetTenantViewboxSharingEnabled()`

UnsetTenantViewboxSharingEnabled ensures that no value is present for TenantViewboxSharingEnabled, not even an explicit nil
### GetTimezone

`func (o *UpdateClusterParams) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateClusterParams) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateClusterParams) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateClusterParams) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *UpdateClusterParams) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *UpdateClusterParams) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTurboMode

`func (o *UpdateClusterParams) GetTurboMode() bool`

GetTurboMode returns the TurboMode field if non-nil, zero value otherwise.

### GetTurboModeOk

`func (o *UpdateClusterParams) GetTurboModeOk() (*bool, bool)`

GetTurboModeOk returns a tuple with the TurboMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboMode

`func (o *UpdateClusterParams) SetTurboMode(v bool)`

SetTurboMode sets TurboMode field to given value.

### HasTurboMode

`func (o *UpdateClusterParams) HasTurboMode() bool`

HasTurboMode returns a boolean if a field has been set.

### SetTurboModeNil

`func (o *UpdateClusterParams) SetTurboModeNil(b bool)`

 SetTurboModeNil sets the value for TurboMode to be an explicit nil

### UnsetTurboMode
`func (o *UpdateClusterParams) UnsetTurboMode()`

UnsetTurboMode ensures that no value is present for TurboMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


