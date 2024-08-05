# CommonAuthenticationMethodParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **NullableString** | Specifies the AWS External Target Authentication type. | 

## Methods

### NewCommonAuthenticationMethodParams

`func NewCommonAuthenticationMethodParams(authenticationType NullableString, ) *CommonAuthenticationMethodParams`

NewCommonAuthenticationMethodParams instantiates a new CommonAuthenticationMethodParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAuthenticationMethodParamsWithDefaults

`func NewCommonAuthenticationMethodParamsWithDefaults() *CommonAuthenticationMethodParams`

NewCommonAuthenticationMethodParamsWithDefaults instantiates a new CommonAuthenticationMethodParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *CommonAuthenticationMethodParams) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CommonAuthenticationMethodParams) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CommonAuthenticationMethodParams) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### SetAuthenticationTypeNil

`func (o *CommonAuthenticationMethodParams) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *CommonAuthenticationMethodParams) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


