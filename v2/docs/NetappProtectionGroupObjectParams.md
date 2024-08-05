# NetappProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 

## Methods

### NewNetappProtectionGroupObjectParams

`func NewNetappProtectionGroupObjectParams(id NullableInt64, ) *NetappProtectionGroupObjectParams`

NewNetappProtectionGroupObjectParams instantiates a new NetappProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappProtectionGroupObjectParamsWithDefaults

`func NewNetappProtectionGroupObjectParamsWithDefaults() *NetappProtectionGroupObjectParams`

NewNetappProtectionGroupObjectParamsWithDefaults instantiates a new NetappProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetappProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetappProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetappProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NetappProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NetappProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *NetappProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetappProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetappProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetappProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NetappProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NetappProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


