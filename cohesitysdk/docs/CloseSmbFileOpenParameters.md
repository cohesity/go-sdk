# CloseSmbFileOpenParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | Pointer to **NullableString** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | [optional] 
**OpenId** | Pointer to **NullableInt64** | Specifies the id of the active open. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. | [optional] 

## Methods

### NewCloseSmbFileOpenParameters

`func NewCloseSmbFileOpenParameters() *CloseSmbFileOpenParameters`

NewCloseSmbFileOpenParameters instantiates a new CloseSmbFileOpenParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseSmbFileOpenParametersWithDefaults

`func NewCloseSmbFileOpenParametersWithDefaults() *CloseSmbFileOpenParameters`

NewCloseSmbFileOpenParametersWithDefaults instantiates a new CloseSmbFileOpenParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *CloseSmbFileOpenParameters) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CloseSmbFileOpenParameters) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CloseSmbFileOpenParameters) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *CloseSmbFileOpenParameters) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *CloseSmbFileOpenParameters) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *CloseSmbFileOpenParameters) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetOpenId

`func (o *CloseSmbFileOpenParameters) GetOpenId() int64`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *CloseSmbFileOpenParameters) GetOpenIdOk() (*int64, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *CloseSmbFileOpenParameters) SetOpenId(v int64)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *CloseSmbFileOpenParameters) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### SetOpenIdNil

`func (o *CloseSmbFileOpenParameters) SetOpenIdNil(b bool)`

 SetOpenIdNil sets the value for OpenId to be an explicit nil

### UnsetOpenId
`func (o *CloseSmbFileOpenParameters) UnsetOpenId()`

UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
### GetViewName

`func (o *CloseSmbFileOpenParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *CloseSmbFileOpenParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *CloseSmbFileOpenParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *CloseSmbFileOpenParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *CloseSmbFileOpenParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *CloseSmbFileOpenParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


