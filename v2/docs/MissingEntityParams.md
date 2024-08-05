# MissingEntityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the parent source of the object. | [optional] [readonly] 
**ParentSourceName** | Pointer to **NullableString** | Specifies the name of the parent source of the object. | [optional] [readonly] 

## Methods

### NewMissingEntityParams

`func NewMissingEntityParams(id NullableInt64, ) *MissingEntityParams`

NewMissingEntityParams instantiates a new MissingEntityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissingEntityParamsWithDefaults

`func NewMissingEntityParamsWithDefaults() *MissingEntityParams`

NewMissingEntityParamsWithDefaults instantiates a new MissingEntityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MissingEntityParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MissingEntityParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MissingEntityParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MissingEntityParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MissingEntityParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *MissingEntityParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MissingEntityParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MissingEntityParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MissingEntityParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MissingEntityParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MissingEntityParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceId

`func (o *MissingEntityParams) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *MissingEntityParams) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *MissingEntityParams) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *MissingEntityParams) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *MissingEntityParams) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *MissingEntityParams) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetParentSourceName

`func (o *MissingEntityParams) GetParentSourceName() string`

GetParentSourceName returns the ParentSourceName field if non-nil, zero value otherwise.

### GetParentSourceNameOk

`func (o *MissingEntityParams) GetParentSourceNameOk() (*string, bool)`

GetParentSourceNameOk returns a tuple with the ParentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceName

`func (o *MissingEntityParams) SetParentSourceName(v string)`

SetParentSourceName sets ParentSourceName field to given value.

### HasParentSourceName

`func (o *MissingEntityParams) HasParentSourceName() bool`

HasParentSourceName returns a boolean if a field has been set.

### SetParentSourceNameNil

`func (o *MissingEntityParams) SetParentSourceNameNil(b bool)`

 SetParentSourceNameNil sets the value for ParentSourceName to be an explicit nil

### UnsetParentSourceName
`func (o *MissingEntityParams) UnsetParentSourceName()`

UnsetParentSourceName ensures that no value is present for ParentSourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


