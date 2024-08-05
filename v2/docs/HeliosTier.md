# HeliosTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefaultTier** | Pointer to **NullableBool** | Specifies whether the current tier will be the default tier for primary retention. | [optional] 
**MoveAfter** | Pointer to **NullableInt64** | Specifies the duration after which the backup will be moved to next tier. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the tier type. | [optional] 
**Unit** | Pointer to **NullableString** | Specificies the time unit after which backup will be moved to next tier. | [optional] 

## Methods

### NewHeliosTier

`func NewHeliosTier() *HeliosTier`

NewHeliosTier instantiates a new HeliosTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTierWithDefaults

`func NewHeliosTierWithDefaults() *HeliosTier`

NewHeliosTierWithDefaults instantiates a new HeliosTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefaultTier

`func (o *HeliosTier) GetIsDefaultTier() bool`

GetIsDefaultTier returns the IsDefaultTier field if non-nil, zero value otherwise.

### GetIsDefaultTierOk

`func (o *HeliosTier) GetIsDefaultTierOk() (*bool, bool)`

GetIsDefaultTierOk returns a tuple with the IsDefaultTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultTier

`func (o *HeliosTier) SetIsDefaultTier(v bool)`

SetIsDefaultTier sets IsDefaultTier field to given value.

### HasIsDefaultTier

`func (o *HeliosTier) HasIsDefaultTier() bool`

HasIsDefaultTier returns a boolean if a field has been set.

### SetIsDefaultTierNil

`func (o *HeliosTier) SetIsDefaultTierNil(b bool)`

 SetIsDefaultTierNil sets the value for IsDefaultTier to be an explicit nil

### UnsetIsDefaultTier
`func (o *HeliosTier) UnsetIsDefaultTier()`

UnsetIsDefaultTier ensures that no value is present for IsDefaultTier, not even an explicit nil
### GetMoveAfter

`func (o *HeliosTier) GetMoveAfter() int64`

GetMoveAfter returns the MoveAfter field if non-nil, zero value otherwise.

### GetMoveAfterOk

`func (o *HeliosTier) GetMoveAfterOk() (*int64, bool)`

GetMoveAfterOk returns a tuple with the MoveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAfter

`func (o *HeliosTier) SetMoveAfter(v int64)`

SetMoveAfter sets MoveAfter field to given value.

### HasMoveAfter

`func (o *HeliosTier) HasMoveAfter() bool`

HasMoveAfter returns a boolean if a field has been set.

### SetMoveAfterNil

`func (o *HeliosTier) SetMoveAfterNil(b bool)`

 SetMoveAfterNil sets the value for MoveAfter to be an explicit nil

### UnsetMoveAfter
`func (o *HeliosTier) UnsetMoveAfter()`

UnsetMoveAfter ensures that no value is present for MoveAfter, not even an explicit nil
### GetType

`func (o *HeliosTier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosTier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosTier) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeliosTier) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HeliosTier) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosTier) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUnit

`func (o *HeliosTier) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosTier) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosTier) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosTier) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosTier) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosTier) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


