# RegisteredSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessInfo** | Pointer to [**ConnectorParameters**](ConnectorParameters.md) |  | [optional] 
**AuthenticationErrorMessage** | Pointer to **NullableString** | Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful. | [optional] 
**AuthenticationStatus** | Pointer to **NullableString** | Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster. If the status is &#39;kFinished&#39; and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. &#39;kPending&#39; indicates the authentication is in progress. &#39;kScheduled&#39; indicates the authentication is scheduled. &#39;kFinished&#39; indicates the authentication is completed. &#39;kRefreshInProgress&#39; indicates the refresh is in progress. | [optional] 
**BlacklistedIpAddresses** | Pointer to **[]string** | Specifies the list of IP Addresses on the registered source to be blacklisted for doing any type of IO operations. | [optional] 
**CassandraParams** | Pointer to [**CassandraConnectParams**](CassandraConnectParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseConnectParams**](CouchbaseConnectParams.md) |  | [optional] 
**Environments** | Pointer to **[]string** | Specifies a list of applications environment that are registered with this Protection Source such as &#39;kSQL&#39;. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**HbaseParams** | Pointer to [**HBaseConnectParams**](HBaseConnectParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsConnectParams**](HdfsConnectParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveConnectParams**](HiveConnectParams.md) |  | [optional] 
**IsDbAuthenticated** | Pointer to **NullableBool** | Specifies if application entity dbAuthenticated or not. ex: oracle database. | [optional] 
**MinimumFreeSpaceGB** | Pointer to **NullableInt64** | Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to &#39;kVMware&#39; type of environments. | [optional] 
**MongodbParams** | Pointer to [**MongoDBConnectParams**](MongoDBConnectParams.md) |  | [optional] 
**NasMountCredentials** | Pointer to [**NullableNasMountCredentialParams**](NasMountCredentialParams.md) | Specifies the credentials required to mount directories on the NetApp server if given. | [optional] 
**Office365CredentialsList** | Pointer to [**[]Office365Credentials**](Office365Credentials.md) | Office365 Source Credentials.  Specifies credentials needed to authenticate &amp; authorize user for Office365. | [optional] 
**Office365Region** | Pointer to **NullableString** | Specifies the region for Office365. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**RefreshErrorMessage** | Pointer to **NullableString** | Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset. | [optional] 
**RefreshTimeUsecs** | Pointer to **NullableInt64** | Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built. | [optional] 
**RegisteredAppsInfo** | Pointer to [**[]RegisteredAppInfo**](RegisteredAppInfo.md) | Specifies information of the applications registered on this protection source. | [optional] 
**RegistrationTimeUsecs** | Pointer to **NullableInt64** | Specifies the Unix epoch time (in microseconds) when the Protection Source was registered. | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration. | [optional] 
**ThrottlingPolicy** | Pointer to [**ThrottlingPolicyParameters**](ThrottlingPolicyParameters.md) |  | [optional] 
**ThrottlingPolicyOverrides** | Pointer to [**[]ThrottlingPolicyOverride**](ThrottlingPolicyOverride.md) | Array of Throttling Policy Overrides for Datastores.  Specifies a list of Throttling Policy for datastores that override the common throttling policy specified for the registered Protection Source. For datastores not in this list, common policy will still apply. | [optional] 
**UseOAuthForExchangeOnline** | Pointer to **NullableBool** | Specifies whether OAuth should be used for authentication in case of Exchange Online. | [optional] 
**UseVmBiosUuid** | Pointer to **NullableBool** | Specifies if registered vCenter is using BIOS UUID to track virtual machines. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 
**VlanParams** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 
**WarningMessages** | Pointer to **[]string** | Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing. | [optional] 

## Methods

### NewRegisteredSourceInfo

`func NewRegisteredSourceInfo() *RegisteredSourceInfo`

NewRegisteredSourceInfo instantiates a new RegisteredSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredSourceInfoWithDefaults

`func NewRegisteredSourceInfoWithDefaults() *RegisteredSourceInfo`

NewRegisteredSourceInfoWithDefaults instantiates a new RegisteredSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessInfo

`func (o *RegisteredSourceInfo) GetAccessInfo() ConnectorParameters`

GetAccessInfo returns the AccessInfo field if non-nil, zero value otherwise.

### GetAccessInfoOk

`func (o *RegisteredSourceInfo) GetAccessInfoOk() (*ConnectorParameters, bool)`

GetAccessInfoOk returns a tuple with the AccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInfo

`func (o *RegisteredSourceInfo) SetAccessInfo(v ConnectorParameters)`

SetAccessInfo sets AccessInfo field to given value.

### HasAccessInfo

`func (o *RegisteredSourceInfo) HasAccessInfo() bool`

HasAccessInfo returns a boolean if a field has been set.

### GetAuthenticationErrorMessage

`func (o *RegisteredSourceInfo) GetAuthenticationErrorMessage() string`

GetAuthenticationErrorMessage returns the AuthenticationErrorMessage field if non-nil, zero value otherwise.

### GetAuthenticationErrorMessageOk

`func (o *RegisteredSourceInfo) GetAuthenticationErrorMessageOk() (*string, bool)`

GetAuthenticationErrorMessageOk returns a tuple with the AuthenticationErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationErrorMessage

`func (o *RegisteredSourceInfo) SetAuthenticationErrorMessage(v string)`

SetAuthenticationErrorMessage sets AuthenticationErrorMessage field to given value.

### HasAuthenticationErrorMessage

`func (o *RegisteredSourceInfo) HasAuthenticationErrorMessage() bool`

HasAuthenticationErrorMessage returns a boolean if a field has been set.

### SetAuthenticationErrorMessageNil

`func (o *RegisteredSourceInfo) SetAuthenticationErrorMessageNil(b bool)`

 SetAuthenticationErrorMessageNil sets the value for AuthenticationErrorMessage to be an explicit nil

### UnsetAuthenticationErrorMessage
`func (o *RegisteredSourceInfo) UnsetAuthenticationErrorMessage()`

UnsetAuthenticationErrorMessage ensures that no value is present for AuthenticationErrorMessage, not even an explicit nil
### GetAuthenticationStatus

`func (o *RegisteredSourceInfo) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *RegisteredSourceInfo) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *RegisteredSourceInfo) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *RegisteredSourceInfo) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### SetAuthenticationStatusNil

`func (o *RegisteredSourceInfo) SetAuthenticationStatusNil(b bool)`

 SetAuthenticationStatusNil sets the value for AuthenticationStatus to be an explicit nil

### UnsetAuthenticationStatus
`func (o *RegisteredSourceInfo) UnsetAuthenticationStatus()`

UnsetAuthenticationStatus ensures that no value is present for AuthenticationStatus, not even an explicit nil
### GetBlacklistedIpAddresses

`func (o *RegisteredSourceInfo) GetBlacklistedIpAddresses() []string`

GetBlacklistedIpAddresses returns the BlacklistedIpAddresses field if non-nil, zero value otherwise.

### GetBlacklistedIpAddressesOk

`func (o *RegisteredSourceInfo) GetBlacklistedIpAddressesOk() (*[]string, bool)`

GetBlacklistedIpAddressesOk returns a tuple with the BlacklistedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIpAddresses

`func (o *RegisteredSourceInfo) SetBlacklistedIpAddresses(v []string)`

SetBlacklistedIpAddresses sets BlacklistedIpAddresses field to given value.

### HasBlacklistedIpAddresses

`func (o *RegisteredSourceInfo) HasBlacklistedIpAddresses() bool`

HasBlacklistedIpAddresses returns a boolean if a field has been set.

### SetBlacklistedIpAddressesNil

`func (o *RegisteredSourceInfo) SetBlacklistedIpAddressesNil(b bool)`

 SetBlacklistedIpAddressesNil sets the value for BlacklistedIpAddresses to be an explicit nil

### UnsetBlacklistedIpAddresses
`func (o *RegisteredSourceInfo) UnsetBlacklistedIpAddresses()`

UnsetBlacklistedIpAddresses ensures that no value is present for BlacklistedIpAddresses, not even an explicit nil
### GetCassandraParams

`func (o *RegisteredSourceInfo) GetCassandraParams() CassandraConnectParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *RegisteredSourceInfo) GetCassandraParamsOk() (*CassandraConnectParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *RegisteredSourceInfo) SetCassandraParams(v CassandraConnectParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *RegisteredSourceInfo) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *RegisteredSourceInfo) GetCouchbaseParams() CouchbaseConnectParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *RegisteredSourceInfo) GetCouchbaseParamsOk() (*CouchbaseConnectParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *RegisteredSourceInfo) SetCouchbaseParams(v CouchbaseConnectParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *RegisteredSourceInfo) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetEnvironments

`func (o *RegisteredSourceInfo) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *RegisteredSourceInfo) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *RegisteredSourceInfo) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *RegisteredSourceInfo) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### SetEnvironmentsNil

`func (o *RegisteredSourceInfo) SetEnvironmentsNil(b bool)`

 SetEnvironmentsNil sets the value for Environments to be an explicit nil

### UnsetEnvironments
`func (o *RegisteredSourceInfo) UnsetEnvironments()`

UnsetEnvironments ensures that no value is present for Environments, not even an explicit nil
### GetHbaseParams

`func (o *RegisteredSourceInfo) GetHbaseParams() HBaseConnectParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *RegisteredSourceInfo) GetHbaseParamsOk() (*HBaseConnectParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *RegisteredSourceInfo) SetHbaseParams(v HBaseConnectParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *RegisteredSourceInfo) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *RegisteredSourceInfo) GetHdfsParams() HdfsConnectParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *RegisteredSourceInfo) GetHdfsParamsOk() (*HdfsConnectParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *RegisteredSourceInfo) SetHdfsParams(v HdfsConnectParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *RegisteredSourceInfo) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *RegisteredSourceInfo) GetHiveParams() HiveConnectParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *RegisteredSourceInfo) GetHiveParamsOk() (*HiveConnectParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *RegisteredSourceInfo) SetHiveParams(v HiveConnectParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *RegisteredSourceInfo) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetIsDbAuthenticated

`func (o *RegisteredSourceInfo) GetIsDbAuthenticated() bool`

GetIsDbAuthenticated returns the IsDbAuthenticated field if non-nil, zero value otherwise.

### GetIsDbAuthenticatedOk

`func (o *RegisteredSourceInfo) GetIsDbAuthenticatedOk() (*bool, bool)`

GetIsDbAuthenticatedOk returns a tuple with the IsDbAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDbAuthenticated

`func (o *RegisteredSourceInfo) SetIsDbAuthenticated(v bool)`

SetIsDbAuthenticated sets IsDbAuthenticated field to given value.

### HasIsDbAuthenticated

`func (o *RegisteredSourceInfo) HasIsDbAuthenticated() bool`

HasIsDbAuthenticated returns a boolean if a field has been set.

### SetIsDbAuthenticatedNil

`func (o *RegisteredSourceInfo) SetIsDbAuthenticatedNil(b bool)`

 SetIsDbAuthenticatedNil sets the value for IsDbAuthenticated to be an explicit nil

### UnsetIsDbAuthenticated
`func (o *RegisteredSourceInfo) UnsetIsDbAuthenticated()`

UnsetIsDbAuthenticated ensures that no value is present for IsDbAuthenticated, not even an explicit nil
### GetMinimumFreeSpaceGB

`func (o *RegisteredSourceInfo) GetMinimumFreeSpaceGB() int64`

GetMinimumFreeSpaceGB returns the MinimumFreeSpaceGB field if non-nil, zero value otherwise.

### GetMinimumFreeSpaceGBOk

`func (o *RegisteredSourceInfo) GetMinimumFreeSpaceGBOk() (*int64, bool)`

GetMinimumFreeSpaceGBOk returns a tuple with the MinimumFreeSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFreeSpaceGB

`func (o *RegisteredSourceInfo) SetMinimumFreeSpaceGB(v int64)`

SetMinimumFreeSpaceGB sets MinimumFreeSpaceGB field to given value.

### HasMinimumFreeSpaceGB

`func (o *RegisteredSourceInfo) HasMinimumFreeSpaceGB() bool`

HasMinimumFreeSpaceGB returns a boolean if a field has been set.

### SetMinimumFreeSpaceGBNil

`func (o *RegisteredSourceInfo) SetMinimumFreeSpaceGBNil(b bool)`

 SetMinimumFreeSpaceGBNil sets the value for MinimumFreeSpaceGB to be an explicit nil

### UnsetMinimumFreeSpaceGB
`func (o *RegisteredSourceInfo) UnsetMinimumFreeSpaceGB()`

UnsetMinimumFreeSpaceGB ensures that no value is present for MinimumFreeSpaceGB, not even an explicit nil
### GetMongodbParams

`func (o *RegisteredSourceInfo) GetMongodbParams() MongoDBConnectParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *RegisteredSourceInfo) GetMongodbParamsOk() (*MongoDBConnectParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *RegisteredSourceInfo) SetMongodbParams(v MongoDBConnectParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *RegisteredSourceInfo) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetNasMountCredentials

`func (o *RegisteredSourceInfo) GetNasMountCredentials() NasMountCredentialParams`

GetNasMountCredentials returns the NasMountCredentials field if non-nil, zero value otherwise.

### GetNasMountCredentialsOk

`func (o *RegisteredSourceInfo) GetNasMountCredentialsOk() (*NasMountCredentialParams, bool)`

GetNasMountCredentialsOk returns a tuple with the NasMountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasMountCredentials

`func (o *RegisteredSourceInfo) SetNasMountCredentials(v NasMountCredentialParams)`

SetNasMountCredentials sets NasMountCredentials field to given value.

### HasNasMountCredentials

`func (o *RegisteredSourceInfo) HasNasMountCredentials() bool`

HasNasMountCredentials returns a boolean if a field has been set.

### SetNasMountCredentialsNil

`func (o *RegisteredSourceInfo) SetNasMountCredentialsNil(b bool)`

 SetNasMountCredentialsNil sets the value for NasMountCredentials to be an explicit nil

### UnsetNasMountCredentials
`func (o *RegisteredSourceInfo) UnsetNasMountCredentials()`

UnsetNasMountCredentials ensures that no value is present for NasMountCredentials, not even an explicit nil
### GetOffice365CredentialsList

`func (o *RegisteredSourceInfo) GetOffice365CredentialsList() []Office365Credentials`

GetOffice365CredentialsList returns the Office365CredentialsList field if non-nil, zero value otherwise.

### GetOffice365CredentialsListOk

`func (o *RegisteredSourceInfo) GetOffice365CredentialsListOk() (*[]Office365Credentials, bool)`

GetOffice365CredentialsListOk returns a tuple with the Office365CredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365CredentialsList

`func (o *RegisteredSourceInfo) SetOffice365CredentialsList(v []Office365Credentials)`

SetOffice365CredentialsList sets Office365CredentialsList field to given value.

### HasOffice365CredentialsList

`func (o *RegisteredSourceInfo) HasOffice365CredentialsList() bool`

HasOffice365CredentialsList returns a boolean if a field has been set.

### SetOffice365CredentialsListNil

`func (o *RegisteredSourceInfo) SetOffice365CredentialsListNil(b bool)`

 SetOffice365CredentialsListNil sets the value for Office365CredentialsList to be an explicit nil

### UnsetOffice365CredentialsList
`func (o *RegisteredSourceInfo) UnsetOffice365CredentialsList()`

UnsetOffice365CredentialsList ensures that no value is present for Office365CredentialsList, not even an explicit nil
### GetOffice365Region

`func (o *RegisteredSourceInfo) GetOffice365Region() string`

GetOffice365Region returns the Office365Region field if non-nil, zero value otherwise.

### GetOffice365RegionOk

`func (o *RegisteredSourceInfo) GetOffice365RegionOk() (*string, bool)`

GetOffice365RegionOk returns a tuple with the Office365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Region

`func (o *RegisteredSourceInfo) SetOffice365Region(v string)`

SetOffice365Region sets Office365Region field to given value.

### HasOffice365Region

`func (o *RegisteredSourceInfo) HasOffice365Region() bool`

HasOffice365Region returns a boolean if a field has been set.

### SetOffice365RegionNil

`func (o *RegisteredSourceInfo) SetOffice365RegionNil(b bool)`

 SetOffice365RegionNil sets the value for Office365Region to be an explicit nil

### UnsetOffice365Region
`func (o *RegisteredSourceInfo) UnsetOffice365Region()`

UnsetOffice365Region ensures that no value is present for Office365Region, not even an explicit nil
### GetPassword

`func (o *RegisteredSourceInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisteredSourceInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisteredSourceInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisteredSourceInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisteredSourceInfo) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisteredSourceInfo) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRefreshErrorMessage

`func (o *RegisteredSourceInfo) GetRefreshErrorMessage() string`

GetRefreshErrorMessage returns the RefreshErrorMessage field if non-nil, zero value otherwise.

### GetRefreshErrorMessageOk

`func (o *RegisteredSourceInfo) GetRefreshErrorMessageOk() (*string, bool)`

GetRefreshErrorMessageOk returns a tuple with the RefreshErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshErrorMessage

`func (o *RegisteredSourceInfo) SetRefreshErrorMessage(v string)`

SetRefreshErrorMessage sets RefreshErrorMessage field to given value.

### HasRefreshErrorMessage

`func (o *RegisteredSourceInfo) HasRefreshErrorMessage() bool`

HasRefreshErrorMessage returns a boolean if a field has been set.

### SetRefreshErrorMessageNil

`func (o *RegisteredSourceInfo) SetRefreshErrorMessageNil(b bool)`

 SetRefreshErrorMessageNil sets the value for RefreshErrorMessage to be an explicit nil

### UnsetRefreshErrorMessage
`func (o *RegisteredSourceInfo) UnsetRefreshErrorMessage()`

UnsetRefreshErrorMessage ensures that no value is present for RefreshErrorMessage, not even an explicit nil
### GetRefreshTimeUsecs

`func (o *RegisteredSourceInfo) GetRefreshTimeUsecs() int64`

GetRefreshTimeUsecs returns the RefreshTimeUsecs field if non-nil, zero value otherwise.

### GetRefreshTimeUsecsOk

`func (o *RegisteredSourceInfo) GetRefreshTimeUsecsOk() (*int64, bool)`

GetRefreshTimeUsecsOk returns a tuple with the RefreshTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTimeUsecs

`func (o *RegisteredSourceInfo) SetRefreshTimeUsecs(v int64)`

SetRefreshTimeUsecs sets RefreshTimeUsecs field to given value.

### HasRefreshTimeUsecs

`func (o *RegisteredSourceInfo) HasRefreshTimeUsecs() bool`

HasRefreshTimeUsecs returns a boolean if a field has been set.

### SetRefreshTimeUsecsNil

`func (o *RegisteredSourceInfo) SetRefreshTimeUsecsNil(b bool)`

 SetRefreshTimeUsecsNil sets the value for RefreshTimeUsecs to be an explicit nil

### UnsetRefreshTimeUsecs
`func (o *RegisteredSourceInfo) UnsetRefreshTimeUsecs()`

UnsetRefreshTimeUsecs ensures that no value is present for RefreshTimeUsecs, not even an explicit nil
### GetRegisteredAppsInfo

`func (o *RegisteredSourceInfo) GetRegisteredAppsInfo() []RegisteredAppInfo`

GetRegisteredAppsInfo returns the RegisteredAppsInfo field if non-nil, zero value otherwise.

### GetRegisteredAppsInfoOk

`func (o *RegisteredSourceInfo) GetRegisteredAppsInfoOk() (*[]RegisteredAppInfo, bool)`

GetRegisteredAppsInfoOk returns a tuple with the RegisteredAppsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAppsInfo

`func (o *RegisteredSourceInfo) SetRegisteredAppsInfo(v []RegisteredAppInfo)`

SetRegisteredAppsInfo sets RegisteredAppsInfo field to given value.

### HasRegisteredAppsInfo

`func (o *RegisteredSourceInfo) HasRegisteredAppsInfo() bool`

HasRegisteredAppsInfo returns a boolean if a field has been set.

### SetRegisteredAppsInfoNil

`func (o *RegisteredSourceInfo) SetRegisteredAppsInfoNil(b bool)`

 SetRegisteredAppsInfoNil sets the value for RegisteredAppsInfo to be an explicit nil

### UnsetRegisteredAppsInfo
`func (o *RegisteredSourceInfo) UnsetRegisteredAppsInfo()`

UnsetRegisteredAppsInfo ensures that no value is present for RegisteredAppsInfo, not even an explicit nil
### GetRegistrationTimeUsecs

`func (o *RegisteredSourceInfo) GetRegistrationTimeUsecs() int64`

GetRegistrationTimeUsecs returns the RegistrationTimeUsecs field if non-nil, zero value otherwise.

### GetRegistrationTimeUsecsOk

`func (o *RegisteredSourceInfo) GetRegistrationTimeUsecsOk() (*int64, bool)`

GetRegistrationTimeUsecsOk returns a tuple with the RegistrationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeUsecs

`func (o *RegisteredSourceInfo) SetRegistrationTimeUsecs(v int64)`

SetRegistrationTimeUsecs sets RegistrationTimeUsecs field to given value.

### HasRegistrationTimeUsecs

`func (o *RegisteredSourceInfo) HasRegistrationTimeUsecs() bool`

HasRegistrationTimeUsecs returns a boolean if a field has been set.

### SetRegistrationTimeUsecsNil

`func (o *RegisteredSourceInfo) SetRegistrationTimeUsecsNil(b bool)`

 SetRegistrationTimeUsecsNil sets the value for RegistrationTimeUsecs to be an explicit nil

### UnsetRegistrationTimeUsecs
`func (o *RegisteredSourceInfo) UnsetRegistrationTimeUsecs()`

UnsetRegistrationTimeUsecs ensures that no value is present for RegistrationTimeUsecs, not even an explicit nil
### GetSubnets

`func (o *RegisteredSourceInfo) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *RegisteredSourceInfo) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *RegisteredSourceInfo) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *RegisteredSourceInfo) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *RegisteredSourceInfo) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *RegisteredSourceInfo) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetThrottlingPolicy

`func (o *RegisteredSourceInfo) GetThrottlingPolicy() ThrottlingPolicyParameters`

GetThrottlingPolicy returns the ThrottlingPolicy field if non-nil, zero value otherwise.

### GetThrottlingPolicyOk

`func (o *RegisteredSourceInfo) GetThrottlingPolicyOk() (*ThrottlingPolicyParameters, bool)`

GetThrottlingPolicyOk returns a tuple with the ThrottlingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicy

`func (o *RegisteredSourceInfo) SetThrottlingPolicy(v ThrottlingPolicyParameters)`

SetThrottlingPolicy sets ThrottlingPolicy field to given value.

### HasThrottlingPolicy

`func (o *RegisteredSourceInfo) HasThrottlingPolicy() bool`

HasThrottlingPolicy returns a boolean if a field has been set.

### GetThrottlingPolicyOverrides

`func (o *RegisteredSourceInfo) GetThrottlingPolicyOverrides() []ThrottlingPolicyOverride`

GetThrottlingPolicyOverrides returns the ThrottlingPolicyOverrides field if non-nil, zero value otherwise.

### GetThrottlingPolicyOverridesOk

`func (o *RegisteredSourceInfo) GetThrottlingPolicyOverridesOk() (*[]ThrottlingPolicyOverride, bool)`

GetThrottlingPolicyOverridesOk returns a tuple with the ThrottlingPolicyOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicyOverrides

`func (o *RegisteredSourceInfo) SetThrottlingPolicyOverrides(v []ThrottlingPolicyOverride)`

SetThrottlingPolicyOverrides sets ThrottlingPolicyOverrides field to given value.

### HasThrottlingPolicyOverrides

`func (o *RegisteredSourceInfo) HasThrottlingPolicyOverrides() bool`

HasThrottlingPolicyOverrides returns a boolean if a field has been set.

### SetThrottlingPolicyOverridesNil

`func (o *RegisteredSourceInfo) SetThrottlingPolicyOverridesNil(b bool)`

 SetThrottlingPolicyOverridesNil sets the value for ThrottlingPolicyOverrides to be an explicit nil

### UnsetThrottlingPolicyOverrides
`func (o *RegisteredSourceInfo) UnsetThrottlingPolicyOverrides()`

UnsetThrottlingPolicyOverrides ensures that no value is present for ThrottlingPolicyOverrides, not even an explicit nil
### GetUseOAuthForExchangeOnline

`func (o *RegisteredSourceInfo) GetUseOAuthForExchangeOnline() bool`

GetUseOAuthForExchangeOnline returns the UseOAuthForExchangeOnline field if non-nil, zero value otherwise.

### GetUseOAuthForExchangeOnlineOk

`func (o *RegisteredSourceInfo) GetUseOAuthForExchangeOnlineOk() (*bool, bool)`

GetUseOAuthForExchangeOnlineOk returns a tuple with the UseOAuthForExchangeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOAuthForExchangeOnline

`func (o *RegisteredSourceInfo) SetUseOAuthForExchangeOnline(v bool)`

SetUseOAuthForExchangeOnline sets UseOAuthForExchangeOnline field to given value.

### HasUseOAuthForExchangeOnline

`func (o *RegisteredSourceInfo) HasUseOAuthForExchangeOnline() bool`

HasUseOAuthForExchangeOnline returns a boolean if a field has been set.

### SetUseOAuthForExchangeOnlineNil

`func (o *RegisteredSourceInfo) SetUseOAuthForExchangeOnlineNil(b bool)`

 SetUseOAuthForExchangeOnlineNil sets the value for UseOAuthForExchangeOnline to be an explicit nil

### UnsetUseOAuthForExchangeOnline
`func (o *RegisteredSourceInfo) UnsetUseOAuthForExchangeOnline()`

UnsetUseOAuthForExchangeOnline ensures that no value is present for UseOAuthForExchangeOnline, not even an explicit nil
### GetUseVmBiosUuid

`func (o *RegisteredSourceInfo) GetUseVmBiosUuid() bool`

GetUseVmBiosUuid returns the UseVmBiosUuid field if non-nil, zero value otherwise.

### GetUseVmBiosUuidOk

`func (o *RegisteredSourceInfo) GetUseVmBiosUuidOk() (*bool, bool)`

GetUseVmBiosUuidOk returns a tuple with the UseVmBiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVmBiosUuid

`func (o *RegisteredSourceInfo) SetUseVmBiosUuid(v bool)`

SetUseVmBiosUuid sets UseVmBiosUuid field to given value.

### HasUseVmBiosUuid

`func (o *RegisteredSourceInfo) HasUseVmBiosUuid() bool`

HasUseVmBiosUuid returns a boolean if a field has been set.

### SetUseVmBiosUuidNil

`func (o *RegisteredSourceInfo) SetUseVmBiosUuidNil(b bool)`

 SetUseVmBiosUuidNil sets the value for UseVmBiosUuid to be an explicit nil

### UnsetUseVmBiosUuid
`func (o *RegisteredSourceInfo) UnsetUseVmBiosUuid()`

UnsetUseVmBiosUuid ensures that no value is present for UseVmBiosUuid, not even an explicit nil
### GetUsername

`func (o *RegisteredSourceInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisteredSourceInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisteredSourceInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisteredSourceInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RegisteredSourceInfo) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RegisteredSourceInfo) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetVlanParams

`func (o *RegisteredSourceInfo) GetVlanParams() VlanParameters`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *RegisteredSourceInfo) GetVlanParamsOk() (*VlanParameters, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *RegisteredSourceInfo) SetVlanParams(v VlanParameters)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *RegisteredSourceInfo) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.

### GetWarningMessages

`func (o *RegisteredSourceInfo) GetWarningMessages() []string`

GetWarningMessages returns the WarningMessages field if non-nil, zero value otherwise.

### GetWarningMessagesOk

`func (o *RegisteredSourceInfo) GetWarningMessagesOk() (*[]string, bool)`

GetWarningMessagesOk returns a tuple with the WarningMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessages

`func (o *RegisteredSourceInfo) SetWarningMessages(v []string)`

SetWarningMessages sets WarningMessages field to given value.

### HasWarningMessages

`func (o *RegisteredSourceInfo) HasWarningMessages() bool`

HasWarningMessages returns a boolean if a field has been set.

### SetWarningMessagesNil

`func (o *RegisteredSourceInfo) SetWarningMessagesNil(b bool)`

 SetWarningMessagesNil sets the value for WarningMessages to be an explicit nil

### UnsetWarningMessages
`func (o *RegisteredSourceInfo) UnsetWarningMessages()`

UnsetWarningMessages ensures that no value is present for WarningMessages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


