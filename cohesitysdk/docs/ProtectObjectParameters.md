# ProtectObjectParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionSourceEnvironment** | Pointer to **NullableString** | Specifies the environment type of the Protection Source object. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ProtectionSourceIds** | **[]int64** | Specifies the ids of the Protection Sources to protect. | 
**RpoPolicyId** | **NullableString** | Specifies the Rpo policy id. | 

## Methods

### NewProtectObjectParameters

`func NewProtectObjectParameters(protectionSourceIds []int64, rpoPolicyId NullableString, ) *ProtectObjectParameters`

NewProtectObjectParameters instantiates a new ProtectObjectParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectObjectParametersWithDefaults

`func NewProtectObjectParametersWithDefaults() *ProtectObjectParameters`

NewProtectObjectParametersWithDefaults instantiates a new ProtectObjectParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionSourceEnvironment

`func (o *ProtectObjectParameters) GetProtectionSourceEnvironment() string`

GetProtectionSourceEnvironment returns the ProtectionSourceEnvironment field if non-nil, zero value otherwise.

### GetProtectionSourceEnvironmentOk

`func (o *ProtectObjectParameters) GetProtectionSourceEnvironmentOk() (*string, bool)`

GetProtectionSourceEnvironmentOk returns a tuple with the ProtectionSourceEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceEnvironment

`func (o *ProtectObjectParameters) SetProtectionSourceEnvironment(v string)`

SetProtectionSourceEnvironment sets ProtectionSourceEnvironment field to given value.

### HasProtectionSourceEnvironment

`func (o *ProtectObjectParameters) HasProtectionSourceEnvironment() bool`

HasProtectionSourceEnvironment returns a boolean if a field has been set.

### SetProtectionSourceEnvironmentNil

`func (o *ProtectObjectParameters) SetProtectionSourceEnvironmentNil(b bool)`

 SetProtectionSourceEnvironmentNil sets the value for ProtectionSourceEnvironment to be an explicit nil

### UnsetProtectionSourceEnvironment
`func (o *ProtectObjectParameters) UnsetProtectionSourceEnvironment()`

UnsetProtectionSourceEnvironment ensures that no value is present for ProtectionSourceEnvironment, not even an explicit nil
### GetProtectionSourceIds

`func (o *ProtectObjectParameters) GetProtectionSourceIds() []int64`

GetProtectionSourceIds returns the ProtectionSourceIds field if non-nil, zero value otherwise.

### GetProtectionSourceIdsOk

`func (o *ProtectObjectParameters) GetProtectionSourceIdsOk() (*[]int64, bool)`

GetProtectionSourceIdsOk returns a tuple with the ProtectionSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceIds

`func (o *ProtectObjectParameters) SetProtectionSourceIds(v []int64)`

SetProtectionSourceIds sets ProtectionSourceIds field to given value.


### SetProtectionSourceIdsNil

`func (o *ProtectObjectParameters) SetProtectionSourceIdsNil(b bool)`

 SetProtectionSourceIdsNil sets the value for ProtectionSourceIds to be an explicit nil

### UnsetProtectionSourceIds
`func (o *ProtectObjectParameters) UnsetProtectionSourceIds()`

UnsetProtectionSourceIds ensures that no value is present for ProtectionSourceIds, not even an explicit nil
### GetRpoPolicyId

`func (o *ProtectObjectParameters) GetRpoPolicyId() string`

GetRpoPolicyId returns the RpoPolicyId field if non-nil, zero value otherwise.

### GetRpoPolicyIdOk

`func (o *ProtectObjectParameters) GetRpoPolicyIdOk() (*string, bool)`

GetRpoPolicyIdOk returns a tuple with the RpoPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicyId

`func (o *ProtectObjectParameters) SetRpoPolicyId(v string)`

SetRpoPolicyId sets RpoPolicyId field to given value.


### SetRpoPolicyIdNil

`func (o *ProtectObjectParameters) SetRpoPolicyIdNil(b bool)`

 SetRpoPolicyIdNil sets the value for RpoPolicyId to be an explicit nil

### UnsetRpoPolicyId
`func (o *ProtectObjectParameters) UnsetRpoPolicyId()`

UnsetRpoPolicyId ensures that no value is present for RpoPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


