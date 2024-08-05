# GoogleTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveAfter** | Pointer to **NullableInt64** | Specifies the time period after which the backup will be moved from current tier to next tier. | [optional] 
**MoveAfterUnit** | Pointer to **NullableString** | Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the &#39;moveAfter&#39; field specified below. | [optional] 
**TierType** | **NullableString** | Specifies the Google tier types. | 

## Methods

### NewGoogleTier

`func NewGoogleTier(tierType NullableString, ) *GoogleTier`

NewGoogleTier instantiates a new GoogleTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleTierWithDefaults

`func NewGoogleTierWithDefaults() *GoogleTier`

NewGoogleTierWithDefaults instantiates a new GoogleTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveAfter

`func (o *GoogleTier) GetMoveAfter() int64`

GetMoveAfter returns the MoveAfter field if non-nil, zero value otherwise.

### GetMoveAfterOk

`func (o *GoogleTier) GetMoveAfterOk() (*int64, bool)`

GetMoveAfterOk returns a tuple with the MoveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAfter

`func (o *GoogleTier) SetMoveAfter(v int64)`

SetMoveAfter sets MoveAfter field to given value.

### HasMoveAfter

`func (o *GoogleTier) HasMoveAfter() bool`

HasMoveAfter returns a boolean if a field has been set.

### SetMoveAfterNil

`func (o *GoogleTier) SetMoveAfterNil(b bool)`

 SetMoveAfterNil sets the value for MoveAfter to be an explicit nil

### UnsetMoveAfter
`func (o *GoogleTier) UnsetMoveAfter()`

UnsetMoveAfter ensures that no value is present for MoveAfter, not even an explicit nil
### GetMoveAfterUnit

`func (o *GoogleTier) GetMoveAfterUnit() string`

GetMoveAfterUnit returns the MoveAfterUnit field if non-nil, zero value otherwise.

### GetMoveAfterUnitOk

`func (o *GoogleTier) GetMoveAfterUnitOk() (*string, bool)`

GetMoveAfterUnitOk returns a tuple with the MoveAfterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAfterUnit

`func (o *GoogleTier) SetMoveAfterUnit(v string)`

SetMoveAfterUnit sets MoveAfterUnit field to given value.

### HasMoveAfterUnit

`func (o *GoogleTier) HasMoveAfterUnit() bool`

HasMoveAfterUnit returns a boolean if a field has been set.

### SetMoveAfterUnitNil

`func (o *GoogleTier) SetMoveAfterUnitNil(b bool)`

 SetMoveAfterUnitNil sets the value for MoveAfterUnit to be an explicit nil

### UnsetMoveAfterUnit
`func (o *GoogleTier) UnsetMoveAfterUnit()`

UnsetMoveAfterUnit ensures that no value is present for MoveAfterUnit, not even an explicit nil
### GetTierType

`func (o *GoogleTier) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *GoogleTier) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *GoogleTier) SetTierType(v string)`

SetTierType sets TierType field to given value.


### SetTierTypeNil

`func (o *GoogleTier) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *GoogleTier) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


