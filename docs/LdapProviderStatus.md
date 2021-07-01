# LdapProviderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusMessage** | Pointer to **NullableString** | Specifies the connection status message of an LDAP provider. | [optional] 

## Methods

### NewLdapProviderStatus

`func NewLdapProviderStatus() *LdapProviderStatus`

NewLdapProviderStatus instantiates a new LdapProviderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapProviderStatusWithDefaults

`func NewLdapProviderStatusWithDefaults() *LdapProviderStatus`

NewLdapProviderStatusWithDefaults instantiates a new LdapProviderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusMessage

`func (o *LdapProviderStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *LdapProviderStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *LdapProviderStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *LdapProviderStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *LdapProviderStatus) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *LdapProviderStatus) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


