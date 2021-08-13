# HeliosRegConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **NullableString** | Specifies the type of entity that is registered on Helios. | [optional] 
**RigelRegConfig** | Pointer to [**RigelRegConfig**](RigelRegConfig.md) |  | [optional] 

## Methods

### NewHeliosRegConfig

`func NewHeliosRegConfig() *HeliosRegConfig`

NewHeliosRegConfig instantiates a new HeliosRegConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosRegConfigWithDefaults

`func NewHeliosRegConfigWithDefaults() *HeliosRegConfig`

NewHeliosRegConfigWithDefaults instantiates a new HeliosRegConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *HeliosRegConfig) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HeliosRegConfig) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HeliosRegConfig) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *HeliosRegConfig) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *HeliosRegConfig) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *HeliosRegConfig) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetRigelRegConfig

`func (o *HeliosRegConfig) GetRigelRegConfig() RigelRegConfig`

GetRigelRegConfig returns the RigelRegConfig field if non-nil, zero value otherwise.

### GetRigelRegConfigOk

`func (o *HeliosRegConfig) GetRigelRegConfigOk() (*RigelRegConfig, bool)`

GetRigelRegConfigOk returns a tuple with the RigelRegConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelRegConfig

`func (o *HeliosRegConfig) SetRigelRegConfig(v RigelRegConfig)`

SetRigelRegConfig sets RigelRegConfig field to given value.

### HasRigelRegConfig

`func (o *HeliosRegConfig) HasRigelRegConfig() bool`

HasRigelRegConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


