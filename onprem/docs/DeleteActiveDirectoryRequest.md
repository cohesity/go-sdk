# DeleteActiveDirectoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryAdminParams** | [**NullableActiveDirectoryAdminParams**](ActiveDirectoryAdminParams.md) | Specifies the params of a user with administrative privilege of this Active Directory. This field is mandatory if machine accounts are updated. | 

## Methods

### NewDeleteActiveDirectoryRequest

`func NewDeleteActiveDirectoryRequest(activeDirectoryAdminParams NullableActiveDirectoryAdminParams, ) *DeleteActiveDirectoryRequest`

NewDeleteActiveDirectoryRequest instantiates a new DeleteActiveDirectoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteActiveDirectoryRequestWithDefaults

`func NewDeleteActiveDirectoryRequestWithDefaults() *DeleteActiveDirectoryRequest`

NewDeleteActiveDirectoryRequestWithDefaults instantiates a new DeleteActiveDirectoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryAdminParams

`func (o *DeleteActiveDirectoryRequest) GetActiveDirectoryAdminParams() ActiveDirectoryAdminParams`

GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field if non-nil, zero value otherwise.

### GetActiveDirectoryAdminParamsOk

`func (o *DeleteActiveDirectoryRequest) GetActiveDirectoryAdminParamsOk() (*ActiveDirectoryAdminParams, bool)`

GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryAdminParams

`func (o *DeleteActiveDirectoryRequest) SetActiveDirectoryAdminParams(v ActiveDirectoryAdminParams)`

SetActiveDirectoryAdminParams sets ActiveDirectoryAdminParams field to given value.


### SetActiveDirectoryAdminParamsNil

`func (o *DeleteActiveDirectoryRequest) SetActiveDirectoryAdminParamsNil(b bool)`

 SetActiveDirectoryAdminParamsNil sets the value for ActiveDirectoryAdminParams to be an explicit nil

### UnsetActiveDirectoryAdminParams
`func (o *DeleteActiveDirectoryRequest) UnsetActiveDirectoryAdminParams()`

UnsetActiveDirectoryAdminParams ensures that no value is present for ActiveDirectoryAdminParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


