# AwsAuthenticationMethodsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **NullableString** | Specifies the AWS External Target Authentication type. | 
**IAmRoleParams** | Pointer to [**AwsIAmRoleParams**](AwsIAmRoleParams.md) |  | [optional] 
**IAmRolesAnywhereParams** | Pointer to [**AwsIAmRolesAnywhereParams**](AwsIAmRolesAnywhereParams.md) |  | [optional] 
**IAmUserParams** | Pointer to [**AwsIAmUserParams**](AwsIAmUserParams.md) |  | [optional] 
**UseSTSParams** | Pointer to [**AwsUseSTSParams**](AwsUseSTSParams.md) |  | [optional] 

## Methods

### NewAwsAuthenticationMethodsParams

`func NewAwsAuthenticationMethodsParams(authenticationType NullableString, ) *AwsAuthenticationMethodsParams`

NewAwsAuthenticationMethodsParams instantiates a new AwsAuthenticationMethodsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAuthenticationMethodsParamsWithDefaults

`func NewAwsAuthenticationMethodsParamsWithDefaults() *AwsAuthenticationMethodsParams`

NewAwsAuthenticationMethodsParamsWithDefaults instantiates a new AwsAuthenticationMethodsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AwsAuthenticationMethodsParams) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AwsAuthenticationMethodsParams) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AwsAuthenticationMethodsParams) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### SetAuthenticationTypeNil

`func (o *AwsAuthenticationMethodsParams) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *AwsAuthenticationMethodsParams) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetIAmRoleParams

`func (o *AwsAuthenticationMethodsParams) GetIAmRoleParams() AwsIAmRoleParams`

GetIAmRoleParams returns the IAmRoleParams field if non-nil, zero value otherwise.

### GetIAmRoleParamsOk

`func (o *AwsAuthenticationMethodsParams) GetIAmRoleParamsOk() (*AwsIAmRoleParams, bool)`

GetIAmRoleParamsOk returns a tuple with the IAmRoleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAmRoleParams

`func (o *AwsAuthenticationMethodsParams) SetIAmRoleParams(v AwsIAmRoleParams)`

SetIAmRoleParams sets IAmRoleParams field to given value.

### HasIAmRoleParams

`func (o *AwsAuthenticationMethodsParams) HasIAmRoleParams() bool`

HasIAmRoleParams returns a boolean if a field has been set.

### GetIAmRolesAnywhereParams

`func (o *AwsAuthenticationMethodsParams) GetIAmRolesAnywhereParams() AwsIAmRolesAnywhereParams`

GetIAmRolesAnywhereParams returns the IAmRolesAnywhereParams field if non-nil, zero value otherwise.

### GetIAmRolesAnywhereParamsOk

`func (o *AwsAuthenticationMethodsParams) GetIAmRolesAnywhereParamsOk() (*AwsIAmRolesAnywhereParams, bool)`

GetIAmRolesAnywhereParamsOk returns a tuple with the IAmRolesAnywhereParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAmRolesAnywhereParams

`func (o *AwsAuthenticationMethodsParams) SetIAmRolesAnywhereParams(v AwsIAmRolesAnywhereParams)`

SetIAmRolesAnywhereParams sets IAmRolesAnywhereParams field to given value.

### HasIAmRolesAnywhereParams

`func (o *AwsAuthenticationMethodsParams) HasIAmRolesAnywhereParams() bool`

HasIAmRolesAnywhereParams returns a boolean if a field has been set.

### GetIAmUserParams

`func (o *AwsAuthenticationMethodsParams) GetIAmUserParams() AwsIAmUserParams`

GetIAmUserParams returns the IAmUserParams field if non-nil, zero value otherwise.

### GetIAmUserParamsOk

`func (o *AwsAuthenticationMethodsParams) GetIAmUserParamsOk() (*AwsIAmUserParams, bool)`

GetIAmUserParamsOk returns a tuple with the IAmUserParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAmUserParams

`func (o *AwsAuthenticationMethodsParams) SetIAmUserParams(v AwsIAmUserParams)`

SetIAmUserParams sets IAmUserParams field to given value.

### HasIAmUserParams

`func (o *AwsAuthenticationMethodsParams) HasIAmUserParams() bool`

HasIAmUserParams returns a boolean if a field has been set.

### GetUseSTSParams

`func (o *AwsAuthenticationMethodsParams) GetUseSTSParams() AwsUseSTSParams`

GetUseSTSParams returns the UseSTSParams field if non-nil, zero value otherwise.

### GetUseSTSParamsOk

`func (o *AwsAuthenticationMethodsParams) GetUseSTSParamsOk() (*AwsUseSTSParams, bool)`

GetUseSTSParamsOk returns a tuple with the UseSTSParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSTSParams

`func (o *AwsAuthenticationMethodsParams) SetUseSTSParams(v AwsUseSTSParams)`

SetUseSTSParams sets UseSTSParams field to given value.

### HasUseSTSParams

`func (o *AwsAuthenticationMethodsParams) HasUseSTSParams() bool`

HasUseSTSParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


