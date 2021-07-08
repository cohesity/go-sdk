# ADUpdateRestoreTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectAttributesParam** | Pointer to [**ADAttributeRestoreParam**](ADAttributeRestoreParam.md) |  | [optional] 
**ObjectParam** | Pointer to [**ADObjectRestoreParam**](ADObjectRestoreParam.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | Specifies the AD restore request type. | [optional] 

## Methods

### NewADUpdateRestoreTaskOptions

`func NewADUpdateRestoreTaskOptions() *ADUpdateRestoreTaskOptions`

NewADUpdateRestoreTaskOptions instantiates a new ADUpdateRestoreTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADUpdateRestoreTaskOptionsWithDefaults

`func NewADUpdateRestoreTaskOptionsWithDefaults() *ADUpdateRestoreTaskOptions`

NewADUpdateRestoreTaskOptionsWithDefaults instantiates a new ADUpdateRestoreTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectAttributesParam

`func (o *ADUpdateRestoreTaskOptions) GetObjectAttributesParam() ADAttributeRestoreParam`

GetObjectAttributesParam returns the ObjectAttributesParam field if non-nil, zero value otherwise.

### GetObjectAttributesParamOk

`func (o *ADUpdateRestoreTaskOptions) GetObjectAttributesParamOk() (*ADAttributeRestoreParam, bool)`

GetObjectAttributesParamOk returns a tuple with the ObjectAttributesParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAttributesParam

`func (o *ADUpdateRestoreTaskOptions) SetObjectAttributesParam(v ADAttributeRestoreParam)`

SetObjectAttributesParam sets ObjectAttributesParam field to given value.

### HasObjectAttributesParam

`func (o *ADUpdateRestoreTaskOptions) HasObjectAttributesParam() bool`

HasObjectAttributesParam returns a boolean if a field has been set.

### GetObjectParam

`func (o *ADUpdateRestoreTaskOptions) GetObjectParam() ADObjectRestoreParam`

GetObjectParam returns the ObjectParam field if non-nil, zero value otherwise.

### GetObjectParamOk

`func (o *ADUpdateRestoreTaskOptions) GetObjectParamOk() (*ADObjectRestoreParam, bool)`

GetObjectParamOk returns a tuple with the ObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectParam

`func (o *ADUpdateRestoreTaskOptions) SetObjectParam(v ADObjectRestoreParam)`

SetObjectParam sets ObjectParam field to given value.

### HasObjectParam

`func (o *ADUpdateRestoreTaskOptions) HasObjectParam() bool`

HasObjectParam returns a boolean if a field has been set.

### GetType

`func (o *ADUpdateRestoreTaskOptions) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ADUpdateRestoreTaskOptions) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ADUpdateRestoreTaskOptions) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ADUpdateRestoreTaskOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ADUpdateRestoreTaskOptions) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ADUpdateRestoreTaskOptions) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


