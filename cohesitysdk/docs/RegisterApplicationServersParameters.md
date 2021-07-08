# RegisterApplicationServersParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **[]string** | Specifies the types of applications such as &#39;kSQL&#39;, &#39;kExchange&#39;, &#39;kAD&#39; running on the Protection Source. overrideDescription: true Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**HasPersistentAgent** | Pointer to **NullableBool** | Set this to true if a persistent agent is running on the host. If this is specified, then credentials would not be used to log into the host environment. This mechanism may be used in environments such as VMware to get around UAC permission issues by running the agent as a service with the right credentials. If this field is not specified, credentials must be specified. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies the Id of the Protection Source that contains one or more Application Servers running on it. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 

## Methods

### NewRegisterApplicationServersParameters

`func NewRegisterApplicationServersParameters() *RegisterApplicationServersParameters`

NewRegisterApplicationServersParameters instantiates a new RegisterApplicationServersParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterApplicationServersParametersWithDefaults

`func NewRegisterApplicationServersParametersWithDefaults() *RegisterApplicationServersParameters`

NewRegisterApplicationServersParametersWithDefaults instantiates a new RegisterApplicationServersParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *RegisterApplicationServersParameters) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RegisterApplicationServersParameters) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RegisterApplicationServersParameters) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RegisterApplicationServersParameters) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *RegisterApplicationServersParameters) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *RegisterApplicationServersParameters) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetHasPersistentAgent

`func (o *RegisterApplicationServersParameters) GetHasPersistentAgent() bool`

GetHasPersistentAgent returns the HasPersistentAgent field if non-nil, zero value otherwise.

### GetHasPersistentAgentOk

`func (o *RegisterApplicationServersParameters) GetHasPersistentAgentOk() (*bool, bool)`

GetHasPersistentAgentOk returns a tuple with the HasPersistentAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersistentAgent

`func (o *RegisterApplicationServersParameters) SetHasPersistentAgent(v bool)`

SetHasPersistentAgent sets HasPersistentAgent field to given value.

### HasHasPersistentAgent

`func (o *RegisterApplicationServersParameters) HasHasPersistentAgent() bool`

HasHasPersistentAgent returns a boolean if a field has been set.

### SetHasPersistentAgentNil

`func (o *RegisterApplicationServersParameters) SetHasPersistentAgentNil(b bool)`

 SetHasPersistentAgentNil sets the value for HasPersistentAgent to be an explicit nil

### UnsetHasPersistentAgent
`func (o *RegisterApplicationServersParameters) UnsetHasPersistentAgent()`

UnsetHasPersistentAgent ensures that no value is present for HasPersistentAgent, not even an explicit nil
### GetPassword

`func (o *RegisterApplicationServersParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterApplicationServersParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterApplicationServersParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterApplicationServersParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisterApplicationServersParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisterApplicationServersParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProtectionSourceId

`func (o *RegisterApplicationServersParameters) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *RegisterApplicationServersParameters) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *RegisterApplicationServersParameters) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *RegisterApplicationServersParameters) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *RegisterApplicationServersParameters) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *RegisterApplicationServersParameters) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
### GetUsername

`func (o *RegisterApplicationServersParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterApplicationServersParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterApplicationServersParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterApplicationServersParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RegisterApplicationServersParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RegisterApplicationServersParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


