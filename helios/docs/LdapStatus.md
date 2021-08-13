# LdapStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **NullableString** | Specifies the connection status. | [optional] 

## Methods

### NewLdapStatus

`func NewLdapStatus() *LdapStatus`

NewLdapStatus instantiates a new LdapStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapStatusWithDefaults

`func NewLdapStatusWithDefaults() *LdapStatus`

NewLdapStatusWithDefaults instantiates a new LdapStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *LdapStatus) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *LdapStatus) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *LdapStatus) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *LdapStatus) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### SetConnectionStatusNil

`func (o *LdapStatus) SetConnectionStatusNil(b bool)`

 SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil

### UnsetConnectionStatus
`func (o *LdapStatus) UnsetConnectionStatus()`

UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


