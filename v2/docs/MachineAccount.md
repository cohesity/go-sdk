# MachineAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsHostName** | Pointer to **NullableString** | Specifies the DNS host name of the machine account. | [optional] 
**Encryption** | Pointer to **[]string** | Specifies a list of encryption types apply to the machine account. | [optional] 
**Name** | **NullableString** | Specifies the machine account name. | 
**ServicePrincipalNames** | Pointer to **[]string** | Specifies the customized Service Principal Names of the Machine Account. Service Principal Name should be unique across the Active Directory forest. | [optional] 

## Methods

### NewMachineAccount

`func NewMachineAccount(name NullableString, ) *MachineAccount`

NewMachineAccount instantiates a new MachineAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountWithDefaults

`func NewMachineAccountWithDefaults() *MachineAccount`

NewMachineAccountWithDefaults instantiates a new MachineAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsHostName

`func (o *MachineAccount) GetDnsHostName() string`

GetDnsHostName returns the DnsHostName field if non-nil, zero value otherwise.

### GetDnsHostNameOk

`func (o *MachineAccount) GetDnsHostNameOk() (*string, bool)`

GetDnsHostNameOk returns a tuple with the DnsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHostName

`func (o *MachineAccount) SetDnsHostName(v string)`

SetDnsHostName sets DnsHostName field to given value.

### HasDnsHostName

`func (o *MachineAccount) HasDnsHostName() bool`

HasDnsHostName returns a boolean if a field has been set.

### SetDnsHostNameNil

`func (o *MachineAccount) SetDnsHostNameNil(b bool)`

 SetDnsHostNameNil sets the value for DnsHostName to be an explicit nil

### UnsetDnsHostName
`func (o *MachineAccount) UnsetDnsHostName()`

UnsetDnsHostName ensures that no value is present for DnsHostName, not even an explicit nil
### GetEncryption

`func (o *MachineAccount) GetEncryption() []string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *MachineAccount) GetEncryptionOk() (*[]string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *MachineAccount) SetEncryption(v []string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *MachineAccount) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### SetEncryptionNil

`func (o *MachineAccount) SetEncryptionNil(b bool)`

 SetEncryptionNil sets the value for Encryption to be an explicit nil

### UnsetEncryption
`func (o *MachineAccount) UnsetEncryption()`

UnsetEncryption ensures that no value is present for Encryption, not even an explicit nil
### GetName

`func (o *MachineAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineAccount) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *MachineAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServicePrincipalNames

`func (o *MachineAccount) GetServicePrincipalNames() []string`

GetServicePrincipalNames returns the ServicePrincipalNames field if non-nil, zero value otherwise.

### GetServicePrincipalNamesOk

`func (o *MachineAccount) GetServicePrincipalNamesOk() (*[]string, bool)`

GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalNames

`func (o *MachineAccount) SetServicePrincipalNames(v []string)`

SetServicePrincipalNames sets ServicePrincipalNames field to given value.

### HasServicePrincipalNames

`func (o *MachineAccount) HasServicePrincipalNames() bool`

HasServicePrincipalNames returns a boolean if a field has been set.

### SetServicePrincipalNamesNil

`func (o *MachineAccount) SetServicePrincipalNamesNil(b bool)`

 SetServicePrincipalNamesNil sets the value for ServicePrincipalNames to be an explicit nil

### UnsetServicePrincipalNames
`func (o *MachineAccount) UnsetServicePrincipalNames()`

UnsetServicePrincipalNames ensures that no value is present for ServicePrincipalNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


