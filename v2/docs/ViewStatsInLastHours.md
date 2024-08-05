# ViewStatsInLastHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHours** | Pointer to **NullableInt64** | Specifies the time range. | [optional] 
**NfsProtocolValue** | Pointer to **NullableInt64** | Specifies the stats value for NFS protocol. | [optional] 
**S3ProtocolValue** | Pointer to **NullableInt64** | Specifies the stats value for S3 protocol. | [optional] 
**SmbProtocolValue** | Pointer to **NullableInt64** | Specifies the stats value for SMB protocol. | [optional] 
**Value** | Pointer to **NullableInt64** | Specifies the stats value for any protocols. | [optional] 

## Methods

### NewViewStatsInLastHours

`func NewViewStatsInLastHours() *ViewStatsInLastHours`

NewViewStatsInLastHours instantiates a new ViewStatsInLastHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsInLastHoursWithDefaults

`func NewViewStatsInLastHoursWithDefaults() *ViewStatsInLastHours`

NewViewStatsInLastHoursWithDefaults instantiates a new ViewStatsInLastHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastHours

`func (o *ViewStatsInLastHours) GetLastHours() int64`

GetLastHours returns the LastHours field if non-nil, zero value otherwise.

### GetLastHoursOk

`func (o *ViewStatsInLastHours) GetLastHoursOk() (*int64, bool)`

GetLastHoursOk returns a tuple with the LastHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHours

`func (o *ViewStatsInLastHours) SetLastHours(v int64)`

SetLastHours sets LastHours field to given value.

### HasLastHours

`func (o *ViewStatsInLastHours) HasLastHours() bool`

HasLastHours returns a boolean if a field has been set.

### SetLastHoursNil

`func (o *ViewStatsInLastHours) SetLastHoursNil(b bool)`

 SetLastHoursNil sets the value for LastHours to be an explicit nil

### UnsetLastHours
`func (o *ViewStatsInLastHours) UnsetLastHours()`

UnsetLastHours ensures that no value is present for LastHours, not even an explicit nil
### GetNfsProtocolValue

`func (o *ViewStatsInLastHours) GetNfsProtocolValue() int64`

GetNfsProtocolValue returns the NfsProtocolValue field if non-nil, zero value otherwise.

### GetNfsProtocolValueOk

`func (o *ViewStatsInLastHours) GetNfsProtocolValueOk() (*int64, bool)`

GetNfsProtocolValueOk returns a tuple with the NfsProtocolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsProtocolValue

`func (o *ViewStatsInLastHours) SetNfsProtocolValue(v int64)`

SetNfsProtocolValue sets NfsProtocolValue field to given value.

### HasNfsProtocolValue

`func (o *ViewStatsInLastHours) HasNfsProtocolValue() bool`

HasNfsProtocolValue returns a boolean if a field has been set.

### SetNfsProtocolValueNil

`func (o *ViewStatsInLastHours) SetNfsProtocolValueNil(b bool)`

 SetNfsProtocolValueNil sets the value for NfsProtocolValue to be an explicit nil

### UnsetNfsProtocolValue
`func (o *ViewStatsInLastHours) UnsetNfsProtocolValue()`

UnsetNfsProtocolValue ensures that no value is present for NfsProtocolValue, not even an explicit nil
### GetS3ProtocolValue

`func (o *ViewStatsInLastHours) GetS3ProtocolValue() int64`

GetS3ProtocolValue returns the S3ProtocolValue field if non-nil, zero value otherwise.

### GetS3ProtocolValueOk

`func (o *ViewStatsInLastHours) GetS3ProtocolValueOk() (*int64, bool)`

GetS3ProtocolValueOk returns a tuple with the S3ProtocolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ProtocolValue

`func (o *ViewStatsInLastHours) SetS3ProtocolValue(v int64)`

SetS3ProtocolValue sets S3ProtocolValue field to given value.

### HasS3ProtocolValue

`func (o *ViewStatsInLastHours) HasS3ProtocolValue() bool`

HasS3ProtocolValue returns a boolean if a field has been set.

### SetS3ProtocolValueNil

`func (o *ViewStatsInLastHours) SetS3ProtocolValueNil(b bool)`

 SetS3ProtocolValueNil sets the value for S3ProtocolValue to be an explicit nil

### UnsetS3ProtocolValue
`func (o *ViewStatsInLastHours) UnsetS3ProtocolValue()`

UnsetS3ProtocolValue ensures that no value is present for S3ProtocolValue, not even an explicit nil
### GetSmbProtocolValue

`func (o *ViewStatsInLastHours) GetSmbProtocolValue() int64`

GetSmbProtocolValue returns the SmbProtocolValue field if non-nil, zero value otherwise.

### GetSmbProtocolValueOk

`func (o *ViewStatsInLastHours) GetSmbProtocolValueOk() (*int64, bool)`

GetSmbProtocolValueOk returns a tuple with the SmbProtocolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbProtocolValue

`func (o *ViewStatsInLastHours) SetSmbProtocolValue(v int64)`

SetSmbProtocolValue sets SmbProtocolValue field to given value.

### HasSmbProtocolValue

`func (o *ViewStatsInLastHours) HasSmbProtocolValue() bool`

HasSmbProtocolValue returns a boolean if a field has been set.

### SetSmbProtocolValueNil

`func (o *ViewStatsInLastHours) SetSmbProtocolValueNil(b bool)`

 SetSmbProtocolValueNil sets the value for SmbProtocolValue to be an explicit nil

### UnsetSmbProtocolValue
`func (o *ViewStatsInLastHours) UnsetSmbProtocolValue()`

UnsetSmbProtocolValue ensures that no value is present for SmbProtocolValue, not even an explicit nil
### GetValue

`func (o *ViewStatsInLastHours) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewStatsInLastHours) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewStatsInLastHours) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewStatsInLastHours) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ViewStatsInLastHours) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ViewStatsInLastHours) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


