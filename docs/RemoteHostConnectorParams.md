# RemoteHostConnectorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**HostAddress** | Pointer to **NullableString** | Address (i.e., IP, hostname or FQDN) of the host to connect to. Magneto will connect using ssh or equivalent to the host. | [optional] 
**HostType** | Pointer to **NullableInt32** | Type of host to connect to. | [optional] 

## Methods

### NewRemoteHostConnectorParams

`func NewRemoteHostConnectorParams() *RemoteHostConnectorParams`

NewRemoteHostConnectorParams instantiates a new RemoteHostConnectorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteHostConnectorParamsWithDefaults

`func NewRemoteHostConnectorParamsWithDefaults() *RemoteHostConnectorParams`

NewRemoteHostConnectorParamsWithDefaults instantiates a new RemoteHostConnectorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *RemoteHostConnectorParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RemoteHostConnectorParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RemoteHostConnectorParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RemoteHostConnectorParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHostAddress

`func (o *RemoteHostConnectorParams) GetHostAddress() string`

GetHostAddress returns the HostAddress field if non-nil, zero value otherwise.

### GetHostAddressOk

`func (o *RemoteHostConnectorParams) GetHostAddressOk() (*string, bool)`

GetHostAddressOk returns a tuple with the HostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAddress

`func (o *RemoteHostConnectorParams) SetHostAddress(v string)`

SetHostAddress sets HostAddress field to given value.

### HasHostAddress

`func (o *RemoteHostConnectorParams) HasHostAddress() bool`

HasHostAddress returns a boolean if a field has been set.

### SetHostAddressNil

`func (o *RemoteHostConnectorParams) SetHostAddressNil(b bool)`

 SetHostAddressNil sets the value for HostAddress to be an explicit nil

### UnsetHostAddress
`func (o *RemoteHostConnectorParams) UnsetHostAddress()`

UnsetHostAddress ensures that no value is present for HostAddress, not even an explicit nil
### GetHostType

`func (o *RemoteHostConnectorParams) GetHostType() int32`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *RemoteHostConnectorParams) GetHostTypeOk() (*int32, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *RemoteHostConnectorParams) SetHostType(v int32)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *RemoteHostConnectorParams) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *RemoteHostConnectorParams) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *RemoteHostConnectorParams) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


