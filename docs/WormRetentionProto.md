# WormRetentionProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyType** | Pointer to **NullableInt32** | The type of WORM policy set on this run. This field is irrelevant while objects are on legal hold. Objects put on legal hold automatically raise the WORM level on the object behaviorally to the strictest level i.e. kCompliance. | [optional] 
**RetentionSecs** | Pointer to **NullableInt64** | Retention time for datalock in seconds. This will be always relative time. | [optional] 
**Version** | Pointer to **NullableInt32** | Version number for this proto. | [optional] 

## Methods

### NewWormRetentionProto

`func NewWormRetentionProto() *WormRetentionProto`

NewWormRetentionProto instantiates a new WormRetentionProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWormRetentionProtoWithDefaults

`func NewWormRetentionProtoWithDefaults() *WormRetentionProto`

NewWormRetentionProtoWithDefaults instantiates a new WormRetentionProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyType

`func (o *WormRetentionProto) GetPolicyType() int32`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *WormRetentionProto) GetPolicyTypeOk() (*int32, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *WormRetentionProto) SetPolicyType(v int32)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *WormRetentionProto) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### SetPolicyTypeNil

`func (o *WormRetentionProto) SetPolicyTypeNil(b bool)`

 SetPolicyTypeNil sets the value for PolicyType to be an explicit nil

### UnsetPolicyType
`func (o *WormRetentionProto) UnsetPolicyType()`

UnsetPolicyType ensures that no value is present for PolicyType, not even an explicit nil
### GetRetentionSecs

`func (o *WormRetentionProto) GetRetentionSecs() int64`

GetRetentionSecs returns the RetentionSecs field if non-nil, zero value otherwise.

### GetRetentionSecsOk

`func (o *WormRetentionProto) GetRetentionSecsOk() (*int64, bool)`

GetRetentionSecsOk returns a tuple with the RetentionSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSecs

`func (o *WormRetentionProto) SetRetentionSecs(v int64)`

SetRetentionSecs sets RetentionSecs field to given value.

### HasRetentionSecs

`func (o *WormRetentionProto) HasRetentionSecs() bool`

HasRetentionSecs returns a boolean if a field has been set.

### SetRetentionSecsNil

`func (o *WormRetentionProto) SetRetentionSecsNil(b bool)`

 SetRetentionSecsNil sets the value for RetentionSecs to be an explicit nil

### UnsetRetentionSecs
`func (o *WormRetentionProto) UnsetRetentionSecs()`

UnsetRetentionSecs ensures that no value is present for RetentionSecs, not even an explicit nil
### GetVersion

`func (o *WormRetentionProto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WormRetentionProto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WormRetentionProto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WormRetentionProto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WormRetentionProto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WormRetentionProto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


