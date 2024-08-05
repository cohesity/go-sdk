# LifecycleRuleExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateInUsecs** | Pointer to **NullableInt64** | Specifies the Timestamp in Usecs for the date when the object is subject to the rule. | [optional] 
**Days** | Pointer to **NullableInt32** | Specifies the Lifetime in days of the objects that are subject to the rule. | [optional] 
**ExpiredObjectDeleteMarker** | Pointer to **NullableBool** | Specifies whether Amazon S3 will remove a delete marker with no non-current versions. If set, the delete marker will be expired. | [optional] 

## Methods

### NewLifecycleRuleExpiration

`func NewLifecycleRuleExpiration() *LifecycleRuleExpiration`

NewLifecycleRuleExpiration instantiates a new LifecycleRuleExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleExpirationWithDefaults

`func NewLifecycleRuleExpirationWithDefaults() *LifecycleRuleExpiration`

NewLifecycleRuleExpirationWithDefaults instantiates a new LifecycleRuleExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateInUsecs

`func (o *LifecycleRuleExpiration) GetDateInUsecs() int64`

GetDateInUsecs returns the DateInUsecs field if non-nil, zero value otherwise.

### GetDateInUsecsOk

`func (o *LifecycleRuleExpiration) GetDateInUsecsOk() (*int64, bool)`

GetDateInUsecsOk returns a tuple with the DateInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInUsecs

`func (o *LifecycleRuleExpiration) SetDateInUsecs(v int64)`

SetDateInUsecs sets DateInUsecs field to given value.

### HasDateInUsecs

`func (o *LifecycleRuleExpiration) HasDateInUsecs() bool`

HasDateInUsecs returns a boolean if a field has been set.

### SetDateInUsecsNil

`func (o *LifecycleRuleExpiration) SetDateInUsecsNil(b bool)`

 SetDateInUsecsNil sets the value for DateInUsecs to be an explicit nil

### UnsetDateInUsecs
`func (o *LifecycleRuleExpiration) UnsetDateInUsecs()`

UnsetDateInUsecs ensures that no value is present for DateInUsecs, not even an explicit nil
### GetDays

`func (o *LifecycleRuleExpiration) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LifecycleRuleExpiration) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LifecycleRuleExpiration) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *LifecycleRuleExpiration) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *LifecycleRuleExpiration) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *LifecycleRuleExpiration) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetExpiredObjectDeleteMarker

`func (o *LifecycleRuleExpiration) GetExpiredObjectDeleteMarker() bool`

GetExpiredObjectDeleteMarker returns the ExpiredObjectDeleteMarker field if non-nil, zero value otherwise.

### GetExpiredObjectDeleteMarkerOk

`func (o *LifecycleRuleExpiration) GetExpiredObjectDeleteMarkerOk() (*bool, bool)`

GetExpiredObjectDeleteMarkerOk returns a tuple with the ExpiredObjectDeleteMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredObjectDeleteMarker

`func (o *LifecycleRuleExpiration) SetExpiredObjectDeleteMarker(v bool)`

SetExpiredObjectDeleteMarker sets ExpiredObjectDeleteMarker field to given value.

### HasExpiredObjectDeleteMarker

`func (o *LifecycleRuleExpiration) HasExpiredObjectDeleteMarker() bool`

HasExpiredObjectDeleteMarker returns a boolean if a field has been set.

### SetExpiredObjectDeleteMarkerNil

`func (o *LifecycleRuleExpiration) SetExpiredObjectDeleteMarkerNil(b bool)`

 SetExpiredObjectDeleteMarkerNil sets the value for ExpiredObjectDeleteMarker to be an explicit nil

### UnsetExpiredObjectDeleteMarker
`func (o *LifecycleRuleExpiration) UnsetExpiredObjectDeleteMarker()`

UnsetExpiredObjectDeleteMarker ensures that no value is present for ExpiredObjectDeleteMarker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


