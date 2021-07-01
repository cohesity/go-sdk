# RestoreEnvStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment. | [optional] 
**ObjectCount** | Pointer to **NullableInt64** |  | [optional] 
**TotalBytes** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRestoreEnvStats

`func NewRestoreEnvStats() *RestoreEnvStats`

NewRestoreEnvStats instantiates a new RestoreEnvStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreEnvStatsWithDefaults

`func NewRestoreEnvStatsWithDefaults() *RestoreEnvStats`

NewRestoreEnvStatsWithDefaults instantiates a new RestoreEnvStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *RestoreEnvStats) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RestoreEnvStats) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RestoreEnvStats) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RestoreEnvStats) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RestoreEnvStats) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RestoreEnvStats) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectCount

`func (o *RestoreEnvStats) GetObjectCount() int64`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *RestoreEnvStats) GetObjectCountOk() (*int64, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *RestoreEnvStats) SetObjectCount(v int64)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *RestoreEnvStats) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.

### SetObjectCountNil

`func (o *RestoreEnvStats) SetObjectCountNil(b bool)`

 SetObjectCountNil sets the value for ObjectCount to be an explicit nil

### UnsetObjectCount
`func (o *RestoreEnvStats) UnsetObjectCount()`

UnsetObjectCount ensures that no value is present for ObjectCount, not even an explicit nil
### GetTotalBytes

`func (o *RestoreEnvStats) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *RestoreEnvStats) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *RestoreEnvStats) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *RestoreEnvStats) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### SetTotalBytesNil

`func (o *RestoreEnvStats) SetTotalBytesNil(b bool)`

 SetTotalBytesNil sets the value for TotalBytes to be an explicit nil

### UnsetTotalBytes
`func (o *RestoreEnvStats) UnsetTotalBytes()`

UnsetTotalBytes ensures that no value is present for TotalBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


