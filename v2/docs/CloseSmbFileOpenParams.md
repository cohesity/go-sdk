# CloseSmbFileOpenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **NullableString** | Specifies the filepath in the Cohesity View relative to the root filesystem. If this field is specified, viewName field must also be specified. | 
**OpenId** | **NullableInt64** | Specifies the id of the active open. | 
**ViewName** | **NullableString** | Specifies the name of the Cohesity View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified. | 

## Methods

### NewCloseSmbFileOpenParams

`func NewCloseSmbFileOpenParams(filePath NullableString, openId NullableInt64, viewName NullableString, ) *CloseSmbFileOpenParams`

NewCloseSmbFileOpenParams instantiates a new CloseSmbFileOpenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseSmbFileOpenParamsWithDefaults

`func NewCloseSmbFileOpenParamsWithDefaults() *CloseSmbFileOpenParams`

NewCloseSmbFileOpenParamsWithDefaults instantiates a new CloseSmbFileOpenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *CloseSmbFileOpenParams) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CloseSmbFileOpenParams) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CloseSmbFileOpenParams) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### SetFilePathNil

`func (o *CloseSmbFileOpenParams) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *CloseSmbFileOpenParams) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetOpenId

`func (o *CloseSmbFileOpenParams) GetOpenId() int64`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *CloseSmbFileOpenParams) GetOpenIdOk() (*int64, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *CloseSmbFileOpenParams) SetOpenId(v int64)`

SetOpenId sets OpenId field to given value.


### SetOpenIdNil

`func (o *CloseSmbFileOpenParams) SetOpenIdNil(b bool)`

 SetOpenIdNil sets the value for OpenId to be an explicit nil

### UnsetOpenId
`func (o *CloseSmbFileOpenParams) UnsetOpenId()`

UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
### GetViewName

`func (o *CloseSmbFileOpenParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *CloseSmbFileOpenParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *CloseSmbFileOpenParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### SetViewNameNil

`func (o *CloseSmbFileOpenParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *CloseSmbFileOpenParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


