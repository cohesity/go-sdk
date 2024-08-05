# ViewObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 
**Uid** | Pointer to **NullableString** | Specifies a distinct value that&#39;s unique to a source. | [optional] 

## Methods

### NewViewObjectParams

`func NewViewObjectParams() *ViewObjectParams`

NewViewObjectParams instantiates a new ViewObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewObjectParamsWithDefaults

`func NewViewObjectParamsWithDefaults() *ViewObjectParams`

NewViewObjectParamsWithDefaults instantiates a new ViewObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ViewObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ViewObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ViewObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUid

`func (o *ViewObjectParams) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ViewObjectParams) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ViewObjectParams) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ViewObjectParams) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ViewObjectParams) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ViewObjectParams) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


