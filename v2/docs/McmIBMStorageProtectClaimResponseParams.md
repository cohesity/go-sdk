# McmIBMStorageProtectClaimResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCaChain** | Pointer to **NullableString** | Specifies the CA chain that is used to sign the cluster certificate. | [optional] 
**ClusterCertificate** | Pointer to **NullableString** | Specifies the Cluster certificate. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster ID. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation ID. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name. | [optional] 
**ClusterPrivateKey** | Pointer to **NullableString** | Specifies the Cluster private key. | [optional] 
**HeliosCertificate** | Pointer to **NullableString** | Specifies the Helios certificate that can be used to authenticate api calls made from Helios to cluster. | [optional] 
**Passphrase** | Pointer to **NullableString** | Specifies the passphrase (if used) to encrypt the cluster private key. | [optional] 
**SfAccountId** | Pointer to **NullableString** | Specifies the Salesforce account ID used to claim the cluster. | [optional] 

## Methods

### NewMcmIBMStorageProtectClaimResponseParams

`func NewMcmIBMStorageProtectClaimResponseParams() *McmIBMStorageProtectClaimResponseParams`

NewMcmIBMStorageProtectClaimResponseParams instantiates a new McmIBMStorageProtectClaimResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmIBMStorageProtectClaimResponseParamsWithDefaults

`func NewMcmIBMStorageProtectClaimResponseParamsWithDefaults() *McmIBMStorageProtectClaimResponseParams`

NewMcmIBMStorageProtectClaimResponseParamsWithDefaults instantiates a new McmIBMStorageProtectClaimResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCaChain

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterCaChain() string`

GetClusterCaChain returns the ClusterCaChain field if non-nil, zero value otherwise.

### GetClusterCaChainOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterCaChainOk() (*string, bool)`

GetClusterCaChainOk returns a tuple with the ClusterCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCaChain

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterCaChain(v string)`

SetClusterCaChain sets ClusterCaChain field to given value.

### HasClusterCaChain

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterCaChain() bool`

HasClusterCaChain returns a boolean if a field has been set.

### SetClusterCaChainNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterCaChainNil(b bool)`

 SetClusterCaChainNil sets the value for ClusterCaChain to be an explicit nil

### UnsetClusterCaChain
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterCaChain()`

UnsetClusterCaChain ensures that no value is present for ClusterCaChain, not even an explicit nil
### GetClusterCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterCertificate() string`

GetClusterCertificate returns the ClusterCertificate field if non-nil, zero value otherwise.

### GetClusterCertificateOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterCertificateOk() (*string, bool)`

GetClusterCertificateOk returns a tuple with the ClusterCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterCertificate(v string)`

SetClusterCertificate sets ClusterCertificate field to given value.

### HasClusterCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterCertificate() bool`

HasClusterCertificate returns a boolean if a field has been set.

### SetClusterCertificateNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterCertificateNil(b bool)`

 SetClusterCertificateNil sets the value for ClusterCertificate to be an explicit nil

### UnsetClusterCertificate
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterCertificate()`

UnsetClusterCertificate ensures that no value is present for ClusterCertificate, not even an explicit nil
### GetClusterId

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetClusterPrivateKey

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterPrivateKey() string`

GetClusterPrivateKey returns the ClusterPrivateKey field if non-nil, zero value otherwise.

### GetClusterPrivateKeyOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetClusterPrivateKeyOk() (*string, bool)`

GetClusterPrivateKeyOk returns a tuple with the ClusterPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPrivateKey

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterPrivateKey(v string)`

SetClusterPrivateKey sets ClusterPrivateKey field to given value.

### HasClusterPrivateKey

`func (o *McmIBMStorageProtectClaimResponseParams) HasClusterPrivateKey() bool`

HasClusterPrivateKey returns a boolean if a field has been set.

### SetClusterPrivateKeyNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetClusterPrivateKeyNil(b bool)`

 SetClusterPrivateKeyNil sets the value for ClusterPrivateKey to be an explicit nil

### UnsetClusterPrivateKey
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetClusterPrivateKey()`

UnsetClusterPrivateKey ensures that no value is present for ClusterPrivateKey, not even an explicit nil
### GetHeliosCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) GetHeliosCertificate() string`

GetHeliosCertificate returns the HeliosCertificate field if non-nil, zero value otherwise.

### GetHeliosCertificateOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetHeliosCertificateOk() (*string, bool)`

GetHeliosCertificateOk returns a tuple with the HeliosCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) SetHeliosCertificate(v string)`

SetHeliosCertificate sets HeliosCertificate field to given value.

### HasHeliosCertificate

`func (o *McmIBMStorageProtectClaimResponseParams) HasHeliosCertificate() bool`

HasHeliosCertificate returns a boolean if a field has been set.

### SetHeliosCertificateNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetHeliosCertificateNil(b bool)`

 SetHeliosCertificateNil sets the value for HeliosCertificate to be an explicit nil

### UnsetHeliosCertificate
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetHeliosCertificate()`

UnsetHeliosCertificate ensures that no value is present for HeliosCertificate, not even an explicit nil
### GetPassphrase

`func (o *McmIBMStorageProtectClaimResponseParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *McmIBMStorageProtectClaimResponseParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *McmIBMStorageProtectClaimResponseParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetSfAccountId

`func (o *McmIBMStorageProtectClaimResponseParams) GetSfAccountId() string`

GetSfAccountId returns the SfAccountId field if non-nil, zero value otherwise.

### GetSfAccountIdOk

`func (o *McmIBMStorageProtectClaimResponseParams) GetSfAccountIdOk() (*string, bool)`

GetSfAccountIdOk returns a tuple with the SfAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccountId

`func (o *McmIBMStorageProtectClaimResponseParams) SetSfAccountId(v string)`

SetSfAccountId sets SfAccountId field to given value.

### HasSfAccountId

`func (o *McmIBMStorageProtectClaimResponseParams) HasSfAccountId() bool`

HasSfAccountId returns a boolean if a field has been set.

### SetSfAccountIdNil

`func (o *McmIBMStorageProtectClaimResponseParams) SetSfAccountIdNil(b bool)`

 SetSfAccountIdNil sets the value for SfAccountId to be an explicit nil

### UnsetSfAccountId
`func (o *McmIBMStorageProtectClaimResponseParams) UnsetSfAccountId()`

UnsetSfAccountId ensures that no value is present for SfAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


