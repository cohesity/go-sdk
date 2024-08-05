# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChainPem** | Pointer to **[]string** | CA chain | [optional] 
**CertPem** | Pointer to **[]string** | Certificate (pem) to be imported | [optional] 
**PrivateKey** | Pointer to **string** | Private key | [optional] 
**ProtoType** | Pointer to **string** |  | [optional] 
**ServiceTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCertificate

`func NewCertificate() *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChainPem

`func (o *Certificate) GetCaChainPem() []string`

GetCaChainPem returns the CaChainPem field if non-nil, zero value otherwise.

### GetCaChainPemOk

`func (o *Certificate) GetCaChainPemOk() (*[]string, bool)`

GetCaChainPemOk returns a tuple with the CaChainPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChainPem

`func (o *Certificate) SetCaChainPem(v []string)`

SetCaChainPem sets CaChainPem field to given value.

### HasCaChainPem

`func (o *Certificate) HasCaChainPem() bool`

HasCaChainPem returns a boolean if a field has been set.

### GetCertPem

`func (o *Certificate) GetCertPem() []string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *Certificate) GetCertPemOk() (*[]string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *Certificate) SetCertPem(v []string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *Certificate) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### GetPrivateKey

`func (o *Certificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Certificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Certificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *Certificate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetProtoType

`func (o *Certificate) GetProtoType() string`

GetProtoType returns the ProtoType field if non-nil, zero value otherwise.

### GetProtoTypeOk

`func (o *Certificate) GetProtoTypeOk() (*string, bool)`

GetProtoTypeOk returns a tuple with the ProtoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoType

`func (o *Certificate) SetProtoType(v string)`

SetProtoType sets ProtoType field to given value.

### HasProtoType

`func (o *Certificate) HasProtoType() bool`

HasProtoType returns a boolean if a field has been set.

### GetServiceTypes

`func (o *Certificate) GetServiceTypes() []string`

GetServiceTypes returns the ServiceTypes field if non-nil, zero value otherwise.

### GetServiceTypesOk

`func (o *Certificate) GetServiceTypesOk() (*[]string, bool)`

GetServiceTypesOk returns a tuple with the ServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypes

`func (o *Certificate) SetServiceTypes(v []string)`

SetServiceTypes sets ServiceTypes field to given value.

### HasServiceTypes

`func (o *Certificate) HasServiceTypes() bool`

HasServiceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


