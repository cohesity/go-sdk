# ClearNlmLocksParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the id of the client, related NLM locks needs to be clear. | [optional] 
**FilePath** | Pointer to **NullableString** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. | [optional] 

## Methods

### NewClearNlmLocksParameters

`func NewClearNlmLocksParameters() *ClearNlmLocksParameters`

NewClearNlmLocksParameters instantiates a new ClearNlmLocksParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearNlmLocksParametersWithDefaults

`func NewClearNlmLocksParametersWithDefaults() *ClearNlmLocksParameters`

NewClearNlmLocksParametersWithDefaults instantiates a new ClearNlmLocksParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClearNlmLocksParameters) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClearNlmLocksParameters) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClearNlmLocksParameters) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClearNlmLocksParameters) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *ClearNlmLocksParameters) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *ClearNlmLocksParameters) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetFilePath

`func (o *ClearNlmLocksParameters) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ClearNlmLocksParameters) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ClearNlmLocksParameters) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ClearNlmLocksParameters) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *ClearNlmLocksParameters) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *ClearNlmLocksParameters) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetViewName

`func (o *ClearNlmLocksParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ClearNlmLocksParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ClearNlmLocksParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ClearNlmLocksParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ClearNlmLocksParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ClearNlmLocksParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


