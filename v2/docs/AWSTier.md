# AWSTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveAfter** | Pointer to **NullableInt64** | Specifies the time period after which the backup will be moved from current tier to next tier. | [optional] 
**MoveAfterUnit** | Pointer to **NullableString** | Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the &#39;moveAfter&#39; field specified below. | [optional] 
**TierType** | **NullableString** | Specifies the AWS tier types. | 

## Methods

### NewAWSTier

`func NewAWSTier(tierType NullableString, ) *AWSTier`

NewAWSTier instantiates a new AWSTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSTierWithDefaults

`func NewAWSTierWithDefaults() *AWSTier`

NewAWSTierWithDefaults instantiates a new AWSTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveAfter

`func (o *AWSTier) GetMoveAfter() int64`

GetMoveAfter returns the MoveAfter field if non-nil, zero value otherwise.

### GetMoveAfterOk

`func (o *AWSTier) GetMoveAfterOk() (*int64, bool)`

GetMoveAfterOk returns a tuple with the MoveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAfter

`func (o *AWSTier) SetMoveAfter(v int64)`

SetMoveAfter sets MoveAfter field to given value.

### HasMoveAfter

`func (o *AWSTier) HasMoveAfter() bool`

HasMoveAfter returns a boolean if a field has been set.

### SetMoveAfterNil

`func (o *AWSTier) SetMoveAfterNil(b bool)`

 SetMoveAfterNil sets the value for MoveAfter to be an explicit nil

### UnsetMoveAfter
`func (o *AWSTier) UnsetMoveAfter()`

UnsetMoveAfter ensures that no value is present for MoveAfter, not even an explicit nil
### GetMoveAfterUnit

`func (o *AWSTier) GetMoveAfterUnit() string`

GetMoveAfterUnit returns the MoveAfterUnit field if non-nil, zero value otherwise.

### GetMoveAfterUnitOk

`func (o *AWSTier) GetMoveAfterUnitOk() (*string, bool)`

GetMoveAfterUnitOk returns a tuple with the MoveAfterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAfterUnit

`func (o *AWSTier) SetMoveAfterUnit(v string)`

SetMoveAfterUnit sets MoveAfterUnit field to given value.

### HasMoveAfterUnit

`func (o *AWSTier) HasMoveAfterUnit() bool`

HasMoveAfterUnit returns a boolean if a field has been set.

### SetMoveAfterUnitNil

`func (o *AWSTier) SetMoveAfterUnitNil(b bool)`

 SetMoveAfterUnitNil sets the value for MoveAfterUnit to be an explicit nil

### UnsetMoveAfterUnit
`func (o *AWSTier) UnsetMoveAfterUnit()`

UnsetMoveAfterUnit ensures that no value is present for MoveAfterUnit, not even an explicit nil
### GetTierType

`func (o *AWSTier) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *AWSTier) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *AWSTier) SetTierType(v string)`

SetTierType sets TierType field to given value.


### SetTierTypeNil

`func (o *AWSTier) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *AWSTier) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


