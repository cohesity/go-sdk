# VaultRunInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt64** | Specifies the count of runs that ended in the specified state between the start time passed in and the current timestamp. | [optional] 
**Timestamp** | Pointer to **NullableInt64** | Specifies the Unix timestamp at which the run entered the specified state. | [optional] 

## Methods

### NewVaultRunInfo

`func NewVaultRunInfo() *VaultRunInfo`

NewVaultRunInfo instantiates a new VaultRunInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultRunInfoWithDefaults

`func NewVaultRunInfoWithDefaults() *VaultRunInfo`

NewVaultRunInfoWithDefaults instantiates a new VaultRunInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VaultRunInfo) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VaultRunInfo) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VaultRunInfo) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VaultRunInfo) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *VaultRunInfo) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *VaultRunInfo) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetTimestamp

`func (o *VaultRunInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VaultRunInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VaultRunInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VaultRunInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *VaultRunInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VaultRunInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


