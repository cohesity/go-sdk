# SecurityConfigCertificateBasedAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdMapping** | Pointer to **NullableString** | Specifies the field to be used in AD user for authentication. | [optional] 
**CertificateMapping** | Pointer to **NullableString** | Specifies the field to be used in certificate for authentication. | [optional] 
**EnableMappingBasedAuthentication** | Pointer to **NullableBool** | If true, certfication based authentication is done via configured mapping. Else it will proceed based on legacy serial number match. | [optional] 

## Methods

### NewSecurityConfigCertificateBasedAuth

`func NewSecurityConfigCertificateBasedAuth() *SecurityConfigCertificateBasedAuth`

NewSecurityConfigCertificateBasedAuth instantiates a new SecurityConfigCertificateBasedAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigCertificateBasedAuthWithDefaults

`func NewSecurityConfigCertificateBasedAuthWithDefaults() *SecurityConfigCertificateBasedAuth`

NewSecurityConfigCertificateBasedAuthWithDefaults instantiates a new SecurityConfigCertificateBasedAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdMapping

`func (o *SecurityConfigCertificateBasedAuth) GetAdMapping() string`

GetAdMapping returns the AdMapping field if non-nil, zero value otherwise.

### GetAdMappingOk

`func (o *SecurityConfigCertificateBasedAuth) GetAdMappingOk() (*string, bool)`

GetAdMappingOk returns a tuple with the AdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdMapping

`func (o *SecurityConfigCertificateBasedAuth) SetAdMapping(v string)`

SetAdMapping sets AdMapping field to given value.

### HasAdMapping

`func (o *SecurityConfigCertificateBasedAuth) HasAdMapping() bool`

HasAdMapping returns a boolean if a field has been set.

### SetAdMappingNil

`func (o *SecurityConfigCertificateBasedAuth) SetAdMappingNil(b bool)`

 SetAdMappingNil sets the value for AdMapping to be an explicit nil

### UnsetAdMapping
`func (o *SecurityConfigCertificateBasedAuth) UnsetAdMapping()`

UnsetAdMapping ensures that no value is present for AdMapping, not even an explicit nil
### GetCertificateMapping

`func (o *SecurityConfigCertificateBasedAuth) GetCertificateMapping() string`

GetCertificateMapping returns the CertificateMapping field if non-nil, zero value otherwise.

### GetCertificateMappingOk

`func (o *SecurityConfigCertificateBasedAuth) GetCertificateMappingOk() (*string, bool)`

GetCertificateMappingOk returns a tuple with the CertificateMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateMapping

`func (o *SecurityConfigCertificateBasedAuth) SetCertificateMapping(v string)`

SetCertificateMapping sets CertificateMapping field to given value.

### HasCertificateMapping

`func (o *SecurityConfigCertificateBasedAuth) HasCertificateMapping() bool`

HasCertificateMapping returns a boolean if a field has been set.

### SetCertificateMappingNil

`func (o *SecurityConfigCertificateBasedAuth) SetCertificateMappingNil(b bool)`

 SetCertificateMappingNil sets the value for CertificateMapping to be an explicit nil

### UnsetCertificateMapping
`func (o *SecurityConfigCertificateBasedAuth) UnsetCertificateMapping()`

UnsetCertificateMapping ensures that no value is present for CertificateMapping, not even an explicit nil
### GetEnableMappingBasedAuthentication

`func (o *SecurityConfigCertificateBasedAuth) GetEnableMappingBasedAuthentication() bool`

GetEnableMappingBasedAuthentication returns the EnableMappingBasedAuthentication field if non-nil, zero value otherwise.

### GetEnableMappingBasedAuthenticationOk

`func (o *SecurityConfigCertificateBasedAuth) GetEnableMappingBasedAuthenticationOk() (*bool, bool)`

GetEnableMappingBasedAuthenticationOk returns a tuple with the EnableMappingBasedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMappingBasedAuthentication

`func (o *SecurityConfigCertificateBasedAuth) SetEnableMappingBasedAuthentication(v bool)`

SetEnableMappingBasedAuthentication sets EnableMappingBasedAuthentication field to given value.

### HasEnableMappingBasedAuthentication

`func (o *SecurityConfigCertificateBasedAuth) HasEnableMappingBasedAuthentication() bool`

HasEnableMappingBasedAuthentication returns a boolean if a field has been set.

### SetEnableMappingBasedAuthenticationNil

`func (o *SecurityConfigCertificateBasedAuth) SetEnableMappingBasedAuthenticationNil(b bool)`

 SetEnableMappingBasedAuthenticationNil sets the value for EnableMappingBasedAuthentication to be an explicit nil

### UnsetEnableMappingBasedAuthentication
`func (o *SecurityConfigCertificateBasedAuth) UnsetEnableMappingBasedAuthentication()`

UnsetEnableMappingBasedAuthentication ensures that no value is present for EnableMappingBasedAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


