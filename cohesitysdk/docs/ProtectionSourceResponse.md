# ProtectionSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ProtectionJobSummary**](ProtectionJobSummary.md) | Specifies the list of Protection Jobs that protect the object. | [optional] 
**LogicalSizeInBytes** | Pointer to **NullableInt64** | Specifies the logical size of Protection Source in bytes. | [optional] 
**ParentSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**ProtectionSourceUidList** | Pointer to [**[]ProtectionSourceUid**](ProtectionSourceUid.md) | Specifies the list of universal ids of the Protection Source. | [optional] 
**Source** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the unique id of the Protection Source. | [optional] 

## Methods

### NewProtectionSourceResponse

`func NewProtectionSourceResponse() *ProtectionSourceResponse`

NewProtectionSourceResponse instantiates a new ProtectionSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceResponseWithDefaults

`func NewProtectionSourceResponseWithDefaults() *ProtectionSourceResponse`

NewProtectionSourceResponseWithDefaults instantiates a new ProtectionSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ProtectionSourceResponse) GetJobs() []ProtectionJobSummary`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ProtectionSourceResponse) GetJobsOk() (*[]ProtectionJobSummary, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ProtectionSourceResponse) SetJobs(v []ProtectionJobSummary)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ProtectionSourceResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### SetJobsNil

`func (o *ProtectionSourceResponse) SetJobsNil(b bool)`

 SetJobsNil sets the value for Jobs to be an explicit nil

### UnsetJobs
`func (o *ProtectionSourceResponse) UnsetJobs()`

UnsetJobs ensures that no value is present for Jobs, not even an explicit nil
### GetLogicalSizeInBytes

`func (o *ProtectionSourceResponse) GetLogicalSizeInBytes() int64`

GetLogicalSizeInBytes returns the LogicalSizeInBytes field if non-nil, zero value otherwise.

### GetLogicalSizeInBytesOk

`func (o *ProtectionSourceResponse) GetLogicalSizeInBytesOk() (*int64, bool)`

GetLogicalSizeInBytesOk returns a tuple with the LogicalSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeInBytes

`func (o *ProtectionSourceResponse) SetLogicalSizeInBytes(v int64)`

SetLogicalSizeInBytes sets LogicalSizeInBytes field to given value.

### HasLogicalSizeInBytes

`func (o *ProtectionSourceResponse) HasLogicalSizeInBytes() bool`

HasLogicalSizeInBytes returns a boolean if a field has been set.

### SetLogicalSizeInBytesNil

`func (o *ProtectionSourceResponse) SetLogicalSizeInBytesNil(b bool)`

 SetLogicalSizeInBytesNil sets the value for LogicalSizeInBytes to be an explicit nil

### UnsetLogicalSizeInBytes
`func (o *ProtectionSourceResponse) UnsetLogicalSizeInBytes()`

UnsetLogicalSizeInBytes ensures that no value is present for LogicalSizeInBytes, not even an explicit nil
### GetParentSource

`func (o *ProtectionSourceResponse) GetParentSource() ProtectionSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *ProtectionSourceResponse) GetParentSourceOk() (*ProtectionSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *ProtectionSourceResponse) SetParentSource(v ProtectionSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *ProtectionSourceResponse) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetProtectionSourceUidList

`func (o *ProtectionSourceResponse) GetProtectionSourceUidList() []ProtectionSourceUid`

GetProtectionSourceUidList returns the ProtectionSourceUidList field if non-nil, zero value otherwise.

### GetProtectionSourceUidListOk

`func (o *ProtectionSourceResponse) GetProtectionSourceUidListOk() (*[]ProtectionSourceUid, bool)`

GetProtectionSourceUidListOk returns a tuple with the ProtectionSourceUidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceUidList

`func (o *ProtectionSourceResponse) SetProtectionSourceUidList(v []ProtectionSourceUid)`

SetProtectionSourceUidList sets ProtectionSourceUidList field to given value.

### HasProtectionSourceUidList

`func (o *ProtectionSourceResponse) HasProtectionSourceUidList() bool`

HasProtectionSourceUidList returns a boolean if a field has been set.

### SetProtectionSourceUidListNil

`func (o *ProtectionSourceResponse) SetProtectionSourceUidListNil(b bool)`

 SetProtectionSourceUidListNil sets the value for ProtectionSourceUidList to be an explicit nil

### UnsetProtectionSourceUidList
`func (o *ProtectionSourceResponse) UnsetProtectionSourceUidList()`

UnsetProtectionSourceUidList ensures that no value is present for ProtectionSourceUidList, not even an explicit nil
### GetSource

`func (o *ProtectionSourceResponse) GetSource() ProtectionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProtectionSourceResponse) GetSourceOk() (*ProtectionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProtectionSourceResponse) SetSource(v ProtectionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProtectionSourceResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUuid

`func (o *ProtectionSourceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProtectionSourceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProtectionSourceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProtectionSourceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProtectionSourceResponse) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProtectionSourceResponse) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


