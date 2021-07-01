# VcdStorageProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the storage profile. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the storage profile as identified by the VCD. | [optional] 

## Methods

### NewVcdStorageProfile

`func NewVcdStorageProfile() *VcdStorageProfile`

NewVcdStorageProfile instantiates a new VcdStorageProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcdStorageProfileWithDefaults

`func NewVcdStorageProfileWithDefaults() *VcdStorageProfile`

NewVcdStorageProfileWithDefaults instantiates a new VcdStorageProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcdStorageProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcdStorageProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcdStorageProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcdStorageProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VcdStorageProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VcdStorageProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUuid

`func (o *VcdStorageProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VcdStorageProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VcdStorageProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VcdStorageProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *VcdStorageProfile) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *VcdStorageProfile) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


