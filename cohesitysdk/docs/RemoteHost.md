# RemoteHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** | Specifies the address (IP, hostname or FQDN) of the remote host that will run the script. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the OS type of the remote host that will run the script. Currently only &#39;kLinux&#39; is supported. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 

## Methods

### NewRemoteHost

`func NewRemoteHost() *RemoteHost`

NewRemoteHost instantiates a new RemoteHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteHostWithDefaults

`func NewRemoteHostWithDefaults() *RemoteHost`

NewRemoteHostWithDefaults instantiates a new RemoteHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RemoteHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RemoteHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RemoteHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RemoteHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *RemoteHost) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *RemoteHost) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetType

`func (o *RemoteHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemoteHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemoteHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RemoteHost) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RemoteHost) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RemoteHost) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


