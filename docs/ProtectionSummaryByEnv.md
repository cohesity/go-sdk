# ProtectionSummaryByEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the type of environment of the source object like kSQL etc. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ProtectedCount** | Pointer to **NullableInt64** | Specifies the number of objects that are protected under the given entity. | [optional] 
**ProtectedSize** | Pointer to **NullableInt64** | Specifies the total size of the protected objects under the given entity. | [optional] 
**UnprotectedCount** | Pointer to **NullableInt64** | Specifies the number of objects that are not protected under the given entity. | [optional] 
**UnprotectedSize** | Pointer to **NullableInt64** | Specifies the total size of the unprotected objects under the given entity. | [optional] 

## Methods

### NewProtectionSummaryByEnv

`func NewProtectionSummaryByEnv() *ProtectionSummaryByEnv`

NewProtectionSummaryByEnv instantiates a new ProtectionSummaryByEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSummaryByEnvWithDefaults

`func NewProtectionSummaryByEnvWithDefaults() *ProtectionSummaryByEnv`

NewProtectionSummaryByEnvWithDefaults instantiates a new ProtectionSummaryByEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ProtectionSummaryByEnv) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionSummaryByEnv) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionSummaryByEnv) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionSummaryByEnv) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionSummaryByEnv) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionSummaryByEnv) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProtectedCount

`func (o *ProtectionSummaryByEnv) GetProtectedCount() int64`

GetProtectedCount returns the ProtectedCount field if non-nil, zero value otherwise.

### GetProtectedCountOk

`func (o *ProtectionSummaryByEnv) GetProtectedCountOk() (*int64, bool)`

GetProtectedCountOk returns a tuple with the ProtectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedCount

`func (o *ProtectionSummaryByEnv) SetProtectedCount(v int64)`

SetProtectedCount sets ProtectedCount field to given value.

### HasProtectedCount

`func (o *ProtectionSummaryByEnv) HasProtectedCount() bool`

HasProtectedCount returns a boolean if a field has been set.

### SetProtectedCountNil

`func (o *ProtectionSummaryByEnv) SetProtectedCountNil(b bool)`

 SetProtectedCountNil sets the value for ProtectedCount to be an explicit nil

### UnsetProtectedCount
`func (o *ProtectionSummaryByEnv) UnsetProtectedCount()`

UnsetProtectedCount ensures that no value is present for ProtectedCount, not even an explicit nil
### GetProtectedSize

`func (o *ProtectionSummaryByEnv) GetProtectedSize() int64`

GetProtectedSize returns the ProtectedSize field if non-nil, zero value otherwise.

### GetProtectedSizeOk

`func (o *ProtectionSummaryByEnv) GetProtectedSizeOk() (*int64, bool)`

GetProtectedSizeOk returns a tuple with the ProtectedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSize

`func (o *ProtectionSummaryByEnv) SetProtectedSize(v int64)`

SetProtectedSize sets ProtectedSize field to given value.

### HasProtectedSize

`func (o *ProtectionSummaryByEnv) HasProtectedSize() bool`

HasProtectedSize returns a boolean if a field has been set.

### SetProtectedSizeNil

`func (o *ProtectionSummaryByEnv) SetProtectedSizeNil(b bool)`

 SetProtectedSizeNil sets the value for ProtectedSize to be an explicit nil

### UnsetProtectedSize
`func (o *ProtectionSummaryByEnv) UnsetProtectedSize()`

UnsetProtectedSize ensures that no value is present for ProtectedSize, not even an explicit nil
### GetUnprotectedCount

`func (o *ProtectionSummaryByEnv) GetUnprotectedCount() int64`

GetUnprotectedCount returns the UnprotectedCount field if non-nil, zero value otherwise.

### GetUnprotectedCountOk

`func (o *ProtectionSummaryByEnv) GetUnprotectedCountOk() (*int64, bool)`

GetUnprotectedCountOk returns a tuple with the UnprotectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedCount

`func (o *ProtectionSummaryByEnv) SetUnprotectedCount(v int64)`

SetUnprotectedCount sets UnprotectedCount field to given value.

### HasUnprotectedCount

`func (o *ProtectionSummaryByEnv) HasUnprotectedCount() bool`

HasUnprotectedCount returns a boolean if a field has been set.

### SetUnprotectedCountNil

`func (o *ProtectionSummaryByEnv) SetUnprotectedCountNil(b bool)`

 SetUnprotectedCountNil sets the value for UnprotectedCount to be an explicit nil

### UnsetUnprotectedCount
`func (o *ProtectionSummaryByEnv) UnsetUnprotectedCount()`

UnsetUnprotectedCount ensures that no value is present for UnprotectedCount, not even an explicit nil
### GetUnprotectedSize

`func (o *ProtectionSummaryByEnv) GetUnprotectedSize() int64`

GetUnprotectedSize returns the UnprotectedSize field if non-nil, zero value otherwise.

### GetUnprotectedSizeOk

`func (o *ProtectionSummaryByEnv) GetUnprotectedSizeOk() (*int64, bool)`

GetUnprotectedSizeOk returns a tuple with the UnprotectedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSize

`func (o *ProtectionSummaryByEnv) SetUnprotectedSize(v int64)`

SetUnprotectedSize sets UnprotectedSize field to given value.

### HasUnprotectedSize

`func (o *ProtectionSummaryByEnv) HasUnprotectedSize() bool`

HasUnprotectedSize returns a boolean if a field has been set.

### SetUnprotectedSizeNil

`func (o *ProtectionSummaryByEnv) SetUnprotectedSizeNil(b bool)`

 SetUnprotectedSizeNil sets the value for UnprotectedSize to be an explicit nil

### UnsetUnprotectedSize
`func (o *ProtectionSummaryByEnv) UnsetUnprotectedSize()`

UnsetUnprotectedSize ensures that no value is present for UnprotectedSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


