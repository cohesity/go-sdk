# RegisterSwiftParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the tenant Id who&#39;s Swift service will be registered. | 
**KeystoneCredentials** | Pointer to [**KeystoneCredentials**](KeystoneCredentials.md) |  | [optional] 

## Methods

### NewRegisterSwiftParams

`func NewRegisterSwiftParams(tenantId NullableString, ) *RegisterSwiftParams`

NewRegisterSwiftParams instantiates a new RegisterSwiftParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSwiftParamsWithDefaults

`func NewRegisterSwiftParamsWithDefaults() *RegisterSwiftParams`

NewRegisterSwiftParamsWithDefaults instantiates a new RegisterSwiftParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RegisterSwiftParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RegisterSwiftParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RegisterSwiftParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *RegisterSwiftParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RegisterSwiftParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetKeystoneCredentials

`func (o *RegisterSwiftParams) GetKeystoneCredentials() KeystoneCredentials`

GetKeystoneCredentials returns the KeystoneCredentials field if non-nil, zero value otherwise.

### GetKeystoneCredentialsOk

`func (o *RegisterSwiftParams) GetKeystoneCredentialsOk() (*KeystoneCredentials, bool)`

GetKeystoneCredentialsOk returns a tuple with the KeystoneCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneCredentials

`func (o *RegisterSwiftParams) SetKeystoneCredentials(v KeystoneCredentials)`

SetKeystoneCredentials sets KeystoneCredentials field to given value.

### HasKeystoneCredentials

`func (o *RegisterSwiftParams) HasKeystoneCredentials() bool`

HasKeystoneCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


