# ClearNlmLockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Specifies the id of the client, related to which NLM locks needs to be clear. | [optional] 
**FilePath** | Pointer to **string** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | [optional] 
**ViewName** | Pointer to **string** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified. | [optional] 

## Methods

### NewClearNlmLockRequest

`func NewClearNlmLockRequest() *ClearNlmLockRequest`

NewClearNlmLockRequest instantiates a new ClearNlmLockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearNlmLockRequestWithDefaults

`func NewClearNlmLockRequestWithDefaults() *ClearNlmLockRequest`

NewClearNlmLockRequestWithDefaults instantiates a new ClearNlmLockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClearNlmLockRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClearNlmLockRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClearNlmLockRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClearNlmLockRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFilePath

`func (o *ClearNlmLockRequest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ClearNlmLockRequest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ClearNlmLockRequest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ClearNlmLockRequest) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetViewName

`func (o *ClearNlmLockRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ClearNlmLockRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ClearNlmLockRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ClearNlmLockRequest) HasViewName() bool`

HasViewName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


