# GetFailoverOpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedFailoverOperations** | Pointer to [**[]AllowedFailoverOperation**](AllowedFailoverOperation.md) | Failover operations that can be performed corresponding to the view id. | [optional] 

## Methods

### NewGetFailoverOpsResponse

`func NewGetFailoverOpsResponse() *GetFailoverOpsResponse`

NewGetFailoverOpsResponse instantiates a new GetFailoverOpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFailoverOpsResponseWithDefaults

`func NewGetFailoverOpsResponseWithDefaults() *GetFailoverOpsResponse`

NewGetFailoverOpsResponseWithDefaults instantiates a new GetFailoverOpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedFailoverOperations

`func (o *GetFailoverOpsResponse) GetAllowedFailoverOperations() []AllowedFailoverOperation`

GetAllowedFailoverOperations returns the AllowedFailoverOperations field if non-nil, zero value otherwise.

### GetAllowedFailoverOperationsOk

`func (o *GetFailoverOpsResponse) GetAllowedFailoverOperationsOk() (*[]AllowedFailoverOperation, bool)`

GetAllowedFailoverOperationsOk returns a tuple with the AllowedFailoverOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFailoverOperations

`func (o *GetFailoverOpsResponse) SetAllowedFailoverOperations(v []AllowedFailoverOperation)`

SetAllowedFailoverOperations sets AllowedFailoverOperations field to given value.

### HasAllowedFailoverOperations

`func (o *GetFailoverOpsResponse) HasAllowedFailoverOperations() bool`

HasAllowedFailoverOperations returns a boolean if a field has been set.

### SetAllowedFailoverOperationsNil

`func (o *GetFailoverOpsResponse) SetAllowedFailoverOperationsNil(b bool)`

 SetAllowedFailoverOperationsNil sets the value for AllowedFailoverOperations to be an explicit nil

### UnsetAllowedFailoverOperations
`func (o *GetFailoverOpsResponse) UnsetAllowedFailoverOperations()`

UnsetAllowedFailoverOperations ensures that no value is present for AllowedFailoverOperations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


