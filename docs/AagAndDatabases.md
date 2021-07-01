# AagAndDatabases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aag** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Databases** | Pointer to [**[]ProtectionSource**](ProtectionSource.md) | Specifies databases found that are members of the AAG. | [optional] 

## Methods

### NewAagAndDatabases

`func NewAagAndDatabases() *AagAndDatabases`

NewAagAndDatabases instantiates a new AagAndDatabases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAagAndDatabasesWithDefaults

`func NewAagAndDatabasesWithDefaults() *AagAndDatabases`

NewAagAndDatabasesWithDefaults instantiates a new AagAndDatabases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAag

`func (o *AagAndDatabases) GetAag() ProtectionSource`

GetAag returns the Aag field if non-nil, zero value otherwise.

### GetAagOk

`func (o *AagAndDatabases) GetAagOk() (*ProtectionSource, bool)`

GetAagOk returns a tuple with the Aag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAag

`func (o *AagAndDatabases) SetAag(v ProtectionSource)`

SetAag sets Aag field to given value.

### HasAag

`func (o *AagAndDatabases) HasAag() bool`

HasAag returns a boolean if a field has been set.

### GetDatabases

`func (o *AagAndDatabases) GetDatabases() []ProtectionSource`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *AagAndDatabases) GetDatabasesOk() (*[]ProtectionSource, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *AagAndDatabases) SetDatabases(v []ProtectionSource)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *AagAndDatabases) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### SetDatabasesNil

`func (o *AagAndDatabases) SetDatabasesNil(b bool)`

 SetDatabasesNil sets the value for Databases to be an explicit nil

### UnsetDatabases
`func (o *AagAndDatabases) UnsetDatabases()`

UnsetDatabases ensures that no value is present for Databases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


