# ExpirationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateInUsecs** | Pointer to **NullableInt64** | Specifies the Timestamp in Usecs for the date when the object is subject to the rule. | [optional] 
**Days** | Pointer to **NullableInt32** | Specifies the Lifetime in days of the objects that are subject to the rule. | [optional] 
**ExpiredObjectDeleteMarker** | Pointer to **NullableBool** | Specifies whether Amazon S3 will remove a delete marker with no non-current versions. If set, the delete marker will be expired. | [optional] 

## Methods

### NewExpirationAction

`func NewExpirationAction() *ExpirationAction`

NewExpirationAction instantiates a new ExpirationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationActionWithDefaults

`func NewExpirationActionWithDefaults() *ExpirationAction`

NewExpirationActionWithDefaults instantiates a new ExpirationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateInUsecs

`func (o *ExpirationAction) GetDateInUsecs() int64`

GetDateInUsecs returns the DateInUsecs field if non-nil, zero value otherwise.

### GetDateInUsecsOk

`func (o *ExpirationAction) GetDateInUsecsOk() (*int64, bool)`

GetDateInUsecsOk returns a tuple with the DateInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInUsecs

`func (o *ExpirationAction) SetDateInUsecs(v int64)`

SetDateInUsecs sets DateInUsecs field to given value.

### HasDateInUsecs

`func (o *ExpirationAction) HasDateInUsecs() bool`

HasDateInUsecs returns a boolean if a field has been set.

### SetDateInUsecsNil

`func (o *ExpirationAction) SetDateInUsecsNil(b bool)`

 SetDateInUsecsNil sets the value for DateInUsecs to be an explicit nil

### UnsetDateInUsecs
`func (o *ExpirationAction) UnsetDateInUsecs()`

UnsetDateInUsecs ensures that no value is present for DateInUsecs, not even an explicit nil
### GetDays

`func (o *ExpirationAction) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ExpirationAction) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ExpirationAction) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ExpirationAction) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *ExpirationAction) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *ExpirationAction) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetExpiredObjectDeleteMarker

`func (o *ExpirationAction) GetExpiredObjectDeleteMarker() bool`

GetExpiredObjectDeleteMarker returns the ExpiredObjectDeleteMarker field if non-nil, zero value otherwise.

### GetExpiredObjectDeleteMarkerOk

`func (o *ExpirationAction) GetExpiredObjectDeleteMarkerOk() (*bool, bool)`

GetExpiredObjectDeleteMarkerOk returns a tuple with the ExpiredObjectDeleteMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredObjectDeleteMarker

`func (o *ExpirationAction) SetExpiredObjectDeleteMarker(v bool)`

SetExpiredObjectDeleteMarker sets ExpiredObjectDeleteMarker field to given value.

### HasExpiredObjectDeleteMarker

`func (o *ExpirationAction) HasExpiredObjectDeleteMarker() bool`

HasExpiredObjectDeleteMarker returns a boolean if a field has been set.

### SetExpiredObjectDeleteMarkerNil

`func (o *ExpirationAction) SetExpiredObjectDeleteMarkerNil(b bool)`

 SetExpiredObjectDeleteMarkerNil sets the value for ExpiredObjectDeleteMarker to be an explicit nil

### UnsetExpiredObjectDeleteMarker
`func (o *ExpirationAction) UnsetExpiredObjectDeleteMarker()`

UnsetExpiredObjectDeleteMarker ensures that no value is present for ExpiredObjectDeleteMarker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


