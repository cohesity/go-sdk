# RegisteredAppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationErrorMessage** | Pointer to **NullableString** | Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful. | [optional] 
**AuthenticationStatus** | Pointer to **NullableString** | Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is &#39;kFinished&#39; and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. &#39;kPending&#39; indicates the authentication is in progress. &#39;kScheduled&#39; indicates the authentication is scheduled. &#39;kFinished&#39; indicates the authentication is completed. &#39;kRefreshInProgress&#39; indicates the refresh is in progress. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the application environment. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**HostSettingsCheckResults** | Pointer to [**[]HostSettingsCheckResult**](HostSettingsCheckResult.md) | Specifies the list of check results internally performed to verify status of various services such as &#39;AgnetRunning&#39;, &#39;SQLWriterRunning&#39; etc. | [optional] 
**RefreshErrorMessage** | Pointer to **NullableString** | Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset. | [optional] 

## Methods

### NewRegisteredAppInfo

`func NewRegisteredAppInfo() *RegisteredAppInfo`

NewRegisteredAppInfo instantiates a new RegisteredAppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredAppInfoWithDefaults

`func NewRegisteredAppInfoWithDefaults() *RegisteredAppInfo`

NewRegisteredAppInfoWithDefaults instantiates a new RegisteredAppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationErrorMessage

`func (o *RegisteredAppInfo) GetAuthenticationErrorMessage() string`

GetAuthenticationErrorMessage returns the AuthenticationErrorMessage field if non-nil, zero value otherwise.

### GetAuthenticationErrorMessageOk

`func (o *RegisteredAppInfo) GetAuthenticationErrorMessageOk() (*string, bool)`

GetAuthenticationErrorMessageOk returns a tuple with the AuthenticationErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationErrorMessage

`func (o *RegisteredAppInfo) SetAuthenticationErrorMessage(v string)`

SetAuthenticationErrorMessage sets AuthenticationErrorMessage field to given value.

### HasAuthenticationErrorMessage

`func (o *RegisteredAppInfo) HasAuthenticationErrorMessage() bool`

HasAuthenticationErrorMessage returns a boolean if a field has been set.

### SetAuthenticationErrorMessageNil

`func (o *RegisteredAppInfo) SetAuthenticationErrorMessageNil(b bool)`

 SetAuthenticationErrorMessageNil sets the value for AuthenticationErrorMessage to be an explicit nil

### UnsetAuthenticationErrorMessage
`func (o *RegisteredAppInfo) UnsetAuthenticationErrorMessage()`

UnsetAuthenticationErrorMessage ensures that no value is present for AuthenticationErrorMessage, not even an explicit nil
### GetAuthenticationStatus

`func (o *RegisteredAppInfo) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *RegisteredAppInfo) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *RegisteredAppInfo) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *RegisteredAppInfo) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### SetAuthenticationStatusNil

`func (o *RegisteredAppInfo) SetAuthenticationStatusNil(b bool)`

 SetAuthenticationStatusNil sets the value for AuthenticationStatus to be an explicit nil

### UnsetAuthenticationStatus
`func (o *RegisteredAppInfo) UnsetAuthenticationStatus()`

UnsetAuthenticationStatus ensures that no value is present for AuthenticationStatus, not even an explicit nil
### GetEnvironment

`func (o *RegisteredAppInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RegisteredAppInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RegisteredAppInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RegisteredAppInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RegisteredAppInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RegisteredAppInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetHostSettingsCheckResults

`func (o *RegisteredAppInfo) GetHostSettingsCheckResults() []HostSettingsCheckResult`

GetHostSettingsCheckResults returns the HostSettingsCheckResults field if non-nil, zero value otherwise.

### GetHostSettingsCheckResultsOk

`func (o *RegisteredAppInfo) GetHostSettingsCheckResultsOk() (*[]HostSettingsCheckResult, bool)`

GetHostSettingsCheckResultsOk returns a tuple with the HostSettingsCheckResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSettingsCheckResults

`func (o *RegisteredAppInfo) SetHostSettingsCheckResults(v []HostSettingsCheckResult)`

SetHostSettingsCheckResults sets HostSettingsCheckResults field to given value.

### HasHostSettingsCheckResults

`func (o *RegisteredAppInfo) HasHostSettingsCheckResults() bool`

HasHostSettingsCheckResults returns a boolean if a field has been set.

### SetHostSettingsCheckResultsNil

`func (o *RegisteredAppInfo) SetHostSettingsCheckResultsNil(b bool)`

 SetHostSettingsCheckResultsNil sets the value for HostSettingsCheckResults to be an explicit nil

### UnsetHostSettingsCheckResults
`func (o *RegisteredAppInfo) UnsetHostSettingsCheckResults()`

UnsetHostSettingsCheckResults ensures that no value is present for HostSettingsCheckResults, not even an explicit nil
### GetRefreshErrorMessage

`func (o *RegisteredAppInfo) GetRefreshErrorMessage() string`

GetRefreshErrorMessage returns the RefreshErrorMessage field if non-nil, zero value otherwise.

### GetRefreshErrorMessageOk

`func (o *RegisteredAppInfo) GetRefreshErrorMessageOk() (*string, bool)`

GetRefreshErrorMessageOk returns a tuple with the RefreshErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshErrorMessage

`func (o *RegisteredAppInfo) SetRefreshErrorMessage(v string)`

SetRefreshErrorMessage sets RefreshErrorMessage field to given value.

### HasRefreshErrorMessage

`func (o *RegisteredAppInfo) HasRefreshErrorMessage() bool`

HasRefreshErrorMessage returns a boolean if a field has been set.

### SetRefreshErrorMessageNil

`func (o *RegisteredAppInfo) SetRefreshErrorMessageNil(b bool)`

 SetRefreshErrorMessageNil sets the value for RefreshErrorMessage to be an explicit nil

### UnsetRefreshErrorMessage
`func (o *RegisteredAppInfo) UnsetRefreshErrorMessage()`

UnsetRefreshErrorMessage ensures that no value is present for RefreshErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


