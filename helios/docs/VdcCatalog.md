# VdcCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **NullableString** | Specifies the UUID of the catalog. | 
**Name** | Pointer to **NullableString** | Specifies the name of the catalog. | [optional] [readonly] 

## Methods

### NewVdcCatalog

`func NewVdcCatalog(uuid NullableString, ) *VdcCatalog`

NewVdcCatalog instantiates a new VdcCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdcCatalogWithDefaults

`func NewVdcCatalogWithDefaults() *VdcCatalog`

NewVdcCatalogWithDefaults instantiates a new VdcCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *VdcCatalog) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VdcCatalog) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VdcCatalog) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### SetUuidNil

`func (o *VdcCatalog) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *VdcCatalog) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *VdcCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdcCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdcCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdcCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VdcCatalog) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VdcCatalog) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


