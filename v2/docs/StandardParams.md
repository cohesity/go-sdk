# StandardParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodType** | **NullableString** | Specifies the Authentication method(IamArn/IamRole) used by api | 
**IamRoleAwsCredentials** | Pointer to [**IamRoleAwsCredentials**](IamRoleAwsCredentials.md) |  | [optional] 
**IamUserAwsCredentials** | Pointer to [**IamUserAwsCredentials**](IamUserAwsCredentials.md) |  | [optional] 

## Methods

### NewStandardParams

`func NewStandardParams(authMethodType NullableString, ) *StandardParams`

NewStandardParams instantiates a new StandardParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardParamsWithDefaults

`func NewStandardParamsWithDefaults() *StandardParams`

NewStandardParamsWithDefaults instantiates a new StandardParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodType

`func (o *StandardParams) GetAuthMethodType() string`

GetAuthMethodType returns the AuthMethodType field if non-nil, zero value otherwise.

### GetAuthMethodTypeOk

`func (o *StandardParams) GetAuthMethodTypeOk() (*string, bool)`

GetAuthMethodTypeOk returns a tuple with the AuthMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodType

`func (o *StandardParams) SetAuthMethodType(v string)`

SetAuthMethodType sets AuthMethodType field to given value.


### SetAuthMethodTypeNil

`func (o *StandardParams) SetAuthMethodTypeNil(b bool)`

 SetAuthMethodTypeNil sets the value for AuthMethodType to be an explicit nil

### UnsetAuthMethodType
`func (o *StandardParams) UnsetAuthMethodType()`

UnsetAuthMethodType ensures that no value is present for AuthMethodType, not even an explicit nil
### GetIamRoleAwsCredentials

`func (o *StandardParams) GetIamRoleAwsCredentials() IamRoleAwsCredentials`

GetIamRoleAwsCredentials returns the IamRoleAwsCredentials field if non-nil, zero value otherwise.

### GetIamRoleAwsCredentialsOk

`func (o *StandardParams) GetIamRoleAwsCredentialsOk() (*IamRoleAwsCredentials, bool)`

GetIamRoleAwsCredentialsOk returns a tuple with the IamRoleAwsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleAwsCredentials

`func (o *StandardParams) SetIamRoleAwsCredentials(v IamRoleAwsCredentials)`

SetIamRoleAwsCredentials sets IamRoleAwsCredentials field to given value.

### HasIamRoleAwsCredentials

`func (o *StandardParams) HasIamRoleAwsCredentials() bool`

HasIamRoleAwsCredentials returns a boolean if a field has been set.

### GetIamUserAwsCredentials

`func (o *StandardParams) GetIamUserAwsCredentials() IamUserAwsCredentials`

GetIamUserAwsCredentials returns the IamUserAwsCredentials field if non-nil, zero value otherwise.

### GetIamUserAwsCredentialsOk

`func (o *StandardParams) GetIamUserAwsCredentialsOk() (*IamUserAwsCredentials, bool)`

GetIamUserAwsCredentialsOk returns a tuple with the IamUserAwsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamUserAwsCredentials

`func (o *StandardParams) SetIamUserAwsCredentials(v IamUserAwsCredentials)`

SetIamUserAwsCredentials sets IamUserAwsCredentials field to given value.

### HasIamUserAwsCredentials

`func (o *StandardParams) HasIamUserAwsCredentials() bool`

HasIamUserAwsCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


