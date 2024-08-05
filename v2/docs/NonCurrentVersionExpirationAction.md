# NonCurrentVersionExpirationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **NullableInt64** | Specifies the number of days an object is non-current before performing the associated action. | [optional] 

## Methods

### NewNonCurrentVersionExpirationAction

`func NewNonCurrentVersionExpirationAction() *NonCurrentVersionExpirationAction`

NewNonCurrentVersionExpirationAction instantiates a new NonCurrentVersionExpirationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonCurrentVersionExpirationActionWithDefaults

`func NewNonCurrentVersionExpirationActionWithDefaults() *NonCurrentVersionExpirationAction`

NewNonCurrentVersionExpirationActionWithDefaults instantiates a new NonCurrentVersionExpirationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *NonCurrentVersionExpirationAction) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *NonCurrentVersionExpirationAction) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *NonCurrentVersionExpirationAction) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *NonCurrentVersionExpirationAction) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *NonCurrentVersionExpirationAction) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *NonCurrentVersionExpirationAction) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


