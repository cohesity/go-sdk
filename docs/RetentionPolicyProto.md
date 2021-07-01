# RetentionPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDaysToKeep** | Pointer to **NullableInt64** | The number of days to keep the snapshots for a backup run. | [optional] 
**NumSecsToKeep** | Pointer to **NullableInt32** | The number of seconds to keep the snapshots for a backup run. | [optional] 
**WormRetention** | Pointer to [**WormRetentionProto**](WormRetentionProto.md) |  | [optional] 

## Methods

### NewRetentionPolicyProto

`func NewRetentionPolicyProto() *RetentionPolicyProto`

NewRetentionPolicyProto instantiates a new RetentionPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionPolicyProtoWithDefaults

`func NewRetentionPolicyProtoWithDefaults() *RetentionPolicyProto`

NewRetentionPolicyProtoWithDefaults instantiates a new RetentionPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDaysToKeep

`func (o *RetentionPolicyProto) GetNumDaysToKeep() int64`

GetNumDaysToKeep returns the NumDaysToKeep field if non-nil, zero value otherwise.

### GetNumDaysToKeepOk

`func (o *RetentionPolicyProto) GetNumDaysToKeepOk() (*int64, bool)`

GetNumDaysToKeepOk returns a tuple with the NumDaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysToKeep

`func (o *RetentionPolicyProto) SetNumDaysToKeep(v int64)`

SetNumDaysToKeep sets NumDaysToKeep field to given value.

### HasNumDaysToKeep

`func (o *RetentionPolicyProto) HasNumDaysToKeep() bool`

HasNumDaysToKeep returns a boolean if a field has been set.

### SetNumDaysToKeepNil

`func (o *RetentionPolicyProto) SetNumDaysToKeepNil(b bool)`

 SetNumDaysToKeepNil sets the value for NumDaysToKeep to be an explicit nil

### UnsetNumDaysToKeep
`func (o *RetentionPolicyProto) UnsetNumDaysToKeep()`

UnsetNumDaysToKeep ensures that no value is present for NumDaysToKeep, not even an explicit nil
### GetNumSecsToKeep

`func (o *RetentionPolicyProto) GetNumSecsToKeep() int32`

GetNumSecsToKeep returns the NumSecsToKeep field if non-nil, zero value otherwise.

### GetNumSecsToKeepOk

`func (o *RetentionPolicyProto) GetNumSecsToKeepOk() (*int32, bool)`

GetNumSecsToKeepOk returns a tuple with the NumSecsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSecsToKeep

`func (o *RetentionPolicyProto) SetNumSecsToKeep(v int32)`

SetNumSecsToKeep sets NumSecsToKeep field to given value.

### HasNumSecsToKeep

`func (o *RetentionPolicyProto) HasNumSecsToKeep() bool`

HasNumSecsToKeep returns a boolean if a field has been set.

### SetNumSecsToKeepNil

`func (o *RetentionPolicyProto) SetNumSecsToKeepNil(b bool)`

 SetNumSecsToKeepNil sets the value for NumSecsToKeep to be an explicit nil

### UnsetNumSecsToKeep
`func (o *RetentionPolicyProto) UnsetNumSecsToKeep()`

UnsetNumSecsToKeep ensures that no value is present for NumSecsToKeep, not even an explicit nil
### GetWormRetention

`func (o *RetentionPolicyProto) GetWormRetention() WormRetentionProto`

GetWormRetention returns the WormRetention field if non-nil, zero value otherwise.

### GetWormRetentionOk

`func (o *RetentionPolicyProto) GetWormRetentionOk() (*WormRetentionProto, bool)`

GetWormRetentionOk returns a tuple with the WormRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormRetention

`func (o *RetentionPolicyProto) SetWormRetention(v WormRetentionProto)`

SetWormRetention sets WormRetention field to given value.

### HasWormRetention

`func (o *RetentionPolicyProto) HasWormRetention() bool`

HasWormRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


