# McmCohesionClaimResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceId** | Pointer to **NullableString** | Unique id of the cohesion appliance with AWS | [optional] 
**ApplianceName** | Pointer to **NullableString** | Specifies the name of the cohesion appliance. | [optional] 
**CaChain** | Pointer to **NullableString** | Specifies the CA chain that is used to sign the Cohesion Appliance certificate. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the Cohesion Appliance certificate. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id of the appliance. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id of the appliance. | [optional] 
**HeliosCertificate** | Pointer to **NullableString** | Specifies the Helios certificate that can be used to authenticate api calls made from Helios to Cohesion Appliance | [optional] 
**Passphrase** | Pointer to **NullableString** | Specifies the passphrase (if used) to encrypt the Cohesion Appliance private key. | [optional] 
**PrivateKey** | Pointer to **NullableString** | Specifies the Cohesion Appliance private key. | [optional] 

## Methods

### NewMcmCohesionClaimResponseParams

`func NewMcmCohesionClaimResponseParams() *McmCohesionClaimResponseParams`

NewMcmCohesionClaimResponseParams instantiates a new McmCohesionClaimResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmCohesionClaimResponseParamsWithDefaults

`func NewMcmCohesionClaimResponseParamsWithDefaults() *McmCohesionClaimResponseParams`

NewMcmCohesionClaimResponseParamsWithDefaults instantiates a new McmCohesionClaimResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceId

`func (o *McmCohesionClaimResponseParams) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *McmCohesionClaimResponseParams) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *McmCohesionClaimResponseParams) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *McmCohesionClaimResponseParams) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### SetApplianceIdNil

`func (o *McmCohesionClaimResponseParams) SetApplianceIdNil(b bool)`

 SetApplianceIdNil sets the value for ApplianceId to be an explicit nil

### UnsetApplianceId
`func (o *McmCohesionClaimResponseParams) UnsetApplianceId()`

UnsetApplianceId ensures that no value is present for ApplianceId, not even an explicit nil
### GetApplianceName

`func (o *McmCohesionClaimResponseParams) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *McmCohesionClaimResponseParams) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *McmCohesionClaimResponseParams) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.

### HasApplianceName

`func (o *McmCohesionClaimResponseParams) HasApplianceName() bool`

HasApplianceName returns a boolean if a field has been set.

### SetApplianceNameNil

`func (o *McmCohesionClaimResponseParams) SetApplianceNameNil(b bool)`

 SetApplianceNameNil sets the value for ApplianceName to be an explicit nil

### UnsetApplianceName
`func (o *McmCohesionClaimResponseParams) UnsetApplianceName()`

UnsetApplianceName ensures that no value is present for ApplianceName, not even an explicit nil
### GetCaChain

`func (o *McmCohesionClaimResponseParams) GetCaChain() string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *McmCohesionClaimResponseParams) GetCaChainOk() (*string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *McmCohesionClaimResponseParams) SetCaChain(v string)`

SetCaChain sets CaChain field to given value.

### HasCaChain

`func (o *McmCohesionClaimResponseParams) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.

### SetCaChainNil

`func (o *McmCohesionClaimResponseParams) SetCaChainNil(b bool)`

 SetCaChainNil sets the value for CaChain to be an explicit nil

### UnsetCaChain
`func (o *McmCohesionClaimResponseParams) UnsetCaChain()`

UnsetCaChain ensures that no value is present for CaChain, not even an explicit nil
### GetCertificate

`func (o *McmCohesionClaimResponseParams) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *McmCohesionClaimResponseParams) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *McmCohesionClaimResponseParams) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *McmCohesionClaimResponseParams) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *McmCohesionClaimResponseParams) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *McmCohesionClaimResponseParams) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetClusterId

`func (o *McmCohesionClaimResponseParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmCohesionClaimResponseParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmCohesionClaimResponseParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmCohesionClaimResponseParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmCohesionClaimResponseParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmCohesionClaimResponseParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmCohesionClaimResponseParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmCohesionClaimResponseParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmCohesionClaimResponseParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmCohesionClaimResponseParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmCohesionClaimResponseParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmCohesionClaimResponseParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetHeliosCertificate

`func (o *McmCohesionClaimResponseParams) GetHeliosCertificate() string`

GetHeliosCertificate returns the HeliosCertificate field if non-nil, zero value otherwise.

### GetHeliosCertificateOk

`func (o *McmCohesionClaimResponseParams) GetHeliosCertificateOk() (*string, bool)`

GetHeliosCertificateOk returns a tuple with the HeliosCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosCertificate

`func (o *McmCohesionClaimResponseParams) SetHeliosCertificate(v string)`

SetHeliosCertificate sets HeliosCertificate field to given value.

### HasHeliosCertificate

`func (o *McmCohesionClaimResponseParams) HasHeliosCertificate() bool`

HasHeliosCertificate returns a boolean if a field has been set.

### SetHeliosCertificateNil

`func (o *McmCohesionClaimResponseParams) SetHeliosCertificateNil(b bool)`

 SetHeliosCertificateNil sets the value for HeliosCertificate to be an explicit nil

### UnsetHeliosCertificate
`func (o *McmCohesionClaimResponseParams) UnsetHeliosCertificate()`

UnsetHeliosCertificate ensures that no value is present for HeliosCertificate, not even an explicit nil
### GetPassphrase

`func (o *McmCohesionClaimResponseParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *McmCohesionClaimResponseParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *McmCohesionClaimResponseParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *McmCohesionClaimResponseParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *McmCohesionClaimResponseParams) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *McmCohesionClaimResponseParams) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetPrivateKey

`func (o *McmCohesionClaimResponseParams) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *McmCohesionClaimResponseParams) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *McmCohesionClaimResponseParams) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *McmCohesionClaimResponseParams) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *McmCohesionClaimResponseParams) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *McmCohesionClaimResponseParams) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


