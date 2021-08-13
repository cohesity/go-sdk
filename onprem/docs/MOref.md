# MOref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **NullableString** | Unique identifier for the object type. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of VMware object | [optional] 

## Methods

### NewMOref

`func NewMOref() *MOref`

NewMOref instantiates a new MOref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMOrefWithDefaults

`func NewMOrefWithDefaults() *MOref`

NewMOrefWithDefaults instantiates a new MOref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *MOref) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MOref) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *MOref) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *MOref) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *MOref) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *MOref) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetType

`func (o *MOref) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MOref) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MOref) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MOref) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MOref) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MOref) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


