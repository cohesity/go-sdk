# AdRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectAttributeParameters** | Pointer to [**AdObjectAttributeParameters**](AdObjectAttributeParameters.md) |  | [optional] 
**ObjectParameters** | Pointer to [**AdObjectRestoreParameters**](AdObjectRestoreParameters.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the AD restore request type. Specifies the action type of AD restore.  &#39;kNone&#39; specifies no special behaviour. &#39;kObjects&#39; specifies the user action to restore AD objects from a mounted AD snapshot database. &#39;kObjectAttributes&#39; specifies the user action to restore attributes of an AD object from a mounted AD snapshot database. | [optional] 

## Methods

### NewAdRestoreOptions

`func NewAdRestoreOptions() *AdRestoreOptions`

NewAdRestoreOptions instantiates a new AdRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdRestoreOptionsWithDefaults

`func NewAdRestoreOptionsWithDefaults() *AdRestoreOptions`

NewAdRestoreOptionsWithDefaults instantiates a new AdRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectAttributeParameters

`func (o *AdRestoreOptions) GetObjectAttributeParameters() AdObjectAttributeParameters`

GetObjectAttributeParameters returns the ObjectAttributeParameters field if non-nil, zero value otherwise.

### GetObjectAttributeParametersOk

`func (o *AdRestoreOptions) GetObjectAttributeParametersOk() (*AdObjectAttributeParameters, bool)`

GetObjectAttributeParametersOk returns a tuple with the ObjectAttributeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAttributeParameters

`func (o *AdRestoreOptions) SetObjectAttributeParameters(v AdObjectAttributeParameters)`

SetObjectAttributeParameters sets ObjectAttributeParameters field to given value.

### HasObjectAttributeParameters

`func (o *AdRestoreOptions) HasObjectAttributeParameters() bool`

HasObjectAttributeParameters returns a boolean if a field has been set.

### GetObjectParameters

`func (o *AdRestoreOptions) GetObjectParameters() AdObjectRestoreParameters`

GetObjectParameters returns the ObjectParameters field if non-nil, zero value otherwise.

### GetObjectParametersOk

`func (o *AdRestoreOptions) GetObjectParametersOk() (*AdObjectRestoreParameters, bool)`

GetObjectParametersOk returns a tuple with the ObjectParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectParameters

`func (o *AdRestoreOptions) SetObjectParameters(v AdObjectRestoreParameters)`

SetObjectParameters sets ObjectParameters field to given value.

### HasObjectParameters

`func (o *AdRestoreOptions) HasObjectParameters() bool`

HasObjectParameters returns a boolean if a field has been set.

### GetType

`func (o *AdRestoreOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdRestoreOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdRestoreOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdRestoreOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AdRestoreOptions) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdRestoreOptions) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


