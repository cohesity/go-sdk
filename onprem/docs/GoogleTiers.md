# GoogleTiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tiers** | [**[]GoogleTier**](GoogleTier.md) | Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier. | 

## Methods

### NewGoogleTiers

`func NewGoogleTiers(tiers []GoogleTier, ) *GoogleTiers`

NewGoogleTiers instantiates a new GoogleTiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleTiersWithDefaults

`func NewGoogleTiersWithDefaults() *GoogleTiers`

NewGoogleTiersWithDefaults instantiates a new GoogleTiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiers

`func (o *GoogleTiers) GetTiers() []GoogleTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *GoogleTiers) GetTiersOk() (*[]GoogleTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *GoogleTiers) SetTiers(v []GoogleTier)`

SetTiers sets Tiers field to given value.


### SetTiersNil

`func (o *GoogleTiers) SetTiersNil(b bool)`

 SetTiersNil sets the value for Tiers to be an explicit nil

### UnsetTiers
`func (o *GoogleTiers) UnsetTiers()`

UnsetTiers ensures that no value is present for Tiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


