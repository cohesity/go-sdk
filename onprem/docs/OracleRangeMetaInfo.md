# OracleRangeMetaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartOfRange** | Pointer to **NullableInt64** | Specifies starting value of the range in time (usecs), SCN or sequence no. | [optional] 
**EndOfRange** | Pointer to **NullableInt64** | Specifies ending value of the range in time (usecs), SCN or sequence no. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies id of the Protection Group corresponding to this oracle range | [optional] 
**ResetLogId** | Pointer to **NullableInt64** | Specifies resetlogs identifier associated with the oracle range. Only applicable for ranges of type SCN and sequence no. | [optional] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies incarnation id associated with the oracle db for which the restore range belongs. Only applicable for ranges of type SCN and sequence no. | [optional] 

## Methods

### NewOracleRangeMetaInfo

`func NewOracleRangeMetaInfo() *OracleRangeMetaInfo`

NewOracleRangeMetaInfo instantiates a new OracleRangeMetaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleRangeMetaInfoWithDefaults

`func NewOracleRangeMetaInfoWithDefaults() *OracleRangeMetaInfo`

NewOracleRangeMetaInfoWithDefaults instantiates a new OracleRangeMetaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartOfRange

`func (o *OracleRangeMetaInfo) GetStartOfRange() int64`

GetStartOfRange returns the StartOfRange field if non-nil, zero value otherwise.

### GetStartOfRangeOk

`func (o *OracleRangeMetaInfo) GetStartOfRangeOk() (*int64, bool)`

GetStartOfRangeOk returns a tuple with the StartOfRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfRange

`func (o *OracleRangeMetaInfo) SetStartOfRange(v int64)`

SetStartOfRange sets StartOfRange field to given value.

### HasStartOfRange

`func (o *OracleRangeMetaInfo) HasStartOfRange() bool`

HasStartOfRange returns a boolean if a field has been set.

### SetStartOfRangeNil

`func (o *OracleRangeMetaInfo) SetStartOfRangeNil(b bool)`

 SetStartOfRangeNil sets the value for StartOfRange to be an explicit nil

### UnsetStartOfRange
`func (o *OracleRangeMetaInfo) UnsetStartOfRange()`

UnsetStartOfRange ensures that no value is present for StartOfRange, not even an explicit nil
### GetEndOfRange

`func (o *OracleRangeMetaInfo) GetEndOfRange() int64`

GetEndOfRange returns the EndOfRange field if non-nil, zero value otherwise.

### GetEndOfRangeOk

`func (o *OracleRangeMetaInfo) GetEndOfRangeOk() (*int64, bool)`

GetEndOfRangeOk returns a tuple with the EndOfRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfRange

`func (o *OracleRangeMetaInfo) SetEndOfRange(v int64)`

SetEndOfRange sets EndOfRange field to given value.

### HasEndOfRange

`func (o *OracleRangeMetaInfo) HasEndOfRange() bool`

HasEndOfRange returns a boolean if a field has been set.

### SetEndOfRangeNil

`func (o *OracleRangeMetaInfo) SetEndOfRangeNil(b bool)`

 SetEndOfRangeNil sets the value for EndOfRange to be an explicit nil

### UnsetEndOfRange
`func (o *OracleRangeMetaInfo) UnsetEndOfRange()`

UnsetEndOfRange ensures that no value is present for EndOfRange, not even an explicit nil
### GetProtectionGroupId

`func (o *OracleRangeMetaInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *OracleRangeMetaInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *OracleRangeMetaInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *OracleRangeMetaInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *OracleRangeMetaInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *OracleRangeMetaInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetResetLogId

`func (o *OracleRangeMetaInfo) GetResetLogId() int64`

GetResetLogId returns the ResetLogId field if non-nil, zero value otherwise.

### GetResetLogIdOk

`func (o *OracleRangeMetaInfo) GetResetLogIdOk() (*int64, bool)`

GetResetLogIdOk returns a tuple with the ResetLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetLogId

`func (o *OracleRangeMetaInfo) SetResetLogId(v int64)`

SetResetLogId sets ResetLogId field to given value.

### HasResetLogId

`func (o *OracleRangeMetaInfo) HasResetLogId() bool`

HasResetLogId returns a boolean if a field has been set.

### SetResetLogIdNil

`func (o *OracleRangeMetaInfo) SetResetLogIdNil(b bool)`

 SetResetLogIdNil sets the value for ResetLogId to be an explicit nil

### UnsetResetLogId
`func (o *OracleRangeMetaInfo) UnsetResetLogId()`

UnsetResetLogId ensures that no value is present for ResetLogId, not even an explicit nil
### GetIncarnationId

`func (o *OracleRangeMetaInfo) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *OracleRangeMetaInfo) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *OracleRangeMetaInfo) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *OracleRangeMetaInfo) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *OracleRangeMetaInfo) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *OracleRangeMetaInfo) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


