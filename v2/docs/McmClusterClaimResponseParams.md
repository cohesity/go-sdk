# McmClusterClaimResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCaChain** | Pointer to **NullableString** | Specifies the CA chain that is used to sign the cluster certificate. | [optional] 
**ClusterCertificate** | Pointer to **NullableString** | Specifies the cluster certificate. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name. | [optional] 
**ClusterPrivateKey** | Pointer to **NullableString** | Specifies the cluster private key. | [optional] 
**HeliosCertificate** | Pointer to **NullableString** | Specifies the Helios certificate that can be used to authenticate API calls made from Helios to cluster. | [optional] 
**Passphrase** | Pointer to **NullableString** | Specifies the passphrase (if used) to encrypt the cluster private key. | [optional] 
**SfAccountId** | Pointer to **NullableString** | Specifies the Salesforce account id used to claim the cluster. | [optional] 

## Methods

### NewMcmClusterClaimResponseParams

`func NewMcmClusterClaimResponseParams() *McmClusterClaimResponseParams`

NewMcmClusterClaimResponseParams instantiates a new McmClusterClaimResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmClusterClaimResponseParamsWithDefaults

`func NewMcmClusterClaimResponseParamsWithDefaults() *McmClusterClaimResponseParams`

NewMcmClusterClaimResponseParamsWithDefaults instantiates a new McmClusterClaimResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCaChain

`func (o *McmClusterClaimResponseParams) GetClusterCaChain() string`

GetClusterCaChain returns the ClusterCaChain field if non-nil, zero value otherwise.

### GetClusterCaChainOk

`func (o *McmClusterClaimResponseParams) GetClusterCaChainOk() (*string, bool)`

GetClusterCaChainOk returns a tuple with the ClusterCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCaChain

`func (o *McmClusterClaimResponseParams) SetClusterCaChain(v string)`

SetClusterCaChain sets ClusterCaChain field to given value.

### HasClusterCaChain

`func (o *McmClusterClaimResponseParams) HasClusterCaChain() bool`

HasClusterCaChain returns a boolean if a field has been set.

### SetClusterCaChainNil

`func (o *McmClusterClaimResponseParams) SetClusterCaChainNil(b bool)`

 SetClusterCaChainNil sets the value for ClusterCaChain to be an explicit nil

### UnsetClusterCaChain
`func (o *McmClusterClaimResponseParams) UnsetClusterCaChain()`

UnsetClusterCaChain ensures that no value is present for ClusterCaChain, not even an explicit nil
### GetClusterCertificate

`func (o *McmClusterClaimResponseParams) GetClusterCertificate() string`

GetClusterCertificate returns the ClusterCertificate field if non-nil, zero value otherwise.

### GetClusterCertificateOk

`func (o *McmClusterClaimResponseParams) GetClusterCertificateOk() (*string, bool)`

GetClusterCertificateOk returns a tuple with the ClusterCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCertificate

`func (o *McmClusterClaimResponseParams) SetClusterCertificate(v string)`

SetClusterCertificate sets ClusterCertificate field to given value.

### HasClusterCertificate

`func (o *McmClusterClaimResponseParams) HasClusterCertificate() bool`

HasClusterCertificate returns a boolean if a field has been set.

### SetClusterCertificateNil

`func (o *McmClusterClaimResponseParams) SetClusterCertificateNil(b bool)`

 SetClusterCertificateNil sets the value for ClusterCertificate to be an explicit nil

### UnsetClusterCertificate
`func (o *McmClusterClaimResponseParams) UnsetClusterCertificate()`

UnsetClusterCertificate ensures that no value is present for ClusterCertificate, not even an explicit nil
### GetClusterId

`func (o *McmClusterClaimResponseParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmClusterClaimResponseParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmClusterClaimResponseParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmClusterClaimResponseParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmClusterClaimResponseParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmClusterClaimResponseParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmClusterClaimResponseParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmClusterClaimResponseParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmClusterClaimResponseParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmClusterClaimResponseParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmClusterClaimResponseParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmClusterClaimResponseParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *McmClusterClaimResponseParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *McmClusterClaimResponseParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *McmClusterClaimResponseParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *McmClusterClaimResponseParams) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *McmClusterClaimResponseParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *McmClusterClaimResponseParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetClusterPrivateKey

`func (o *McmClusterClaimResponseParams) GetClusterPrivateKey() string`

GetClusterPrivateKey returns the ClusterPrivateKey field if non-nil, zero value otherwise.

### GetClusterPrivateKeyOk

`func (o *McmClusterClaimResponseParams) GetClusterPrivateKeyOk() (*string, bool)`

GetClusterPrivateKeyOk returns a tuple with the ClusterPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPrivateKey

`func (o *McmClusterClaimResponseParams) SetClusterPrivateKey(v string)`

SetClusterPrivateKey sets ClusterPrivateKey field to given value.

### HasClusterPrivateKey

`func (o *McmClusterClaimResponseParams) HasClusterPrivateKey() bool`

HasClusterPrivateKey returns a boolean if a field has been set.

### SetClusterPrivateKeyNil

`func (o *McmClusterClaimResponseParams) SetClusterPrivateKeyNil(b bool)`

 SetClusterPrivateKeyNil sets the value for ClusterPrivateKey to be an explicit nil

### UnsetClusterPrivateKey
`func (o *McmClusterClaimResponseParams) UnsetClusterPrivateKey()`

UnsetClusterPrivateKey ensures that no value is present for ClusterPrivateKey, not even an explicit nil
### GetHeliosCertificate

`func (o *McmClusterClaimResponseParams) GetHeliosCertificate() string`

GetHeliosCertificate returns the HeliosCertificate field if non-nil, zero value otherwise.

### GetHeliosCertificateOk

`func (o *McmClusterClaimResponseParams) GetHeliosCertificateOk() (*string, bool)`

GetHeliosCertificateOk returns a tuple with the HeliosCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosCertificate

`func (o *McmClusterClaimResponseParams) SetHeliosCertificate(v string)`

SetHeliosCertificate sets HeliosCertificate field to given value.

### HasHeliosCertificate

`func (o *McmClusterClaimResponseParams) HasHeliosCertificate() bool`

HasHeliosCertificate returns a boolean if a field has been set.

### SetHeliosCertificateNil

`func (o *McmClusterClaimResponseParams) SetHeliosCertificateNil(b bool)`

 SetHeliosCertificateNil sets the value for HeliosCertificate to be an explicit nil

### UnsetHeliosCertificate
`func (o *McmClusterClaimResponseParams) UnsetHeliosCertificate()`

UnsetHeliosCertificate ensures that no value is present for HeliosCertificate, not even an explicit nil
### GetPassphrase

`func (o *McmClusterClaimResponseParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *McmClusterClaimResponseParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *McmClusterClaimResponseParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *McmClusterClaimResponseParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *McmClusterClaimResponseParams) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *McmClusterClaimResponseParams) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetSfAccountId

`func (o *McmClusterClaimResponseParams) GetSfAccountId() string`

GetSfAccountId returns the SfAccountId field if non-nil, zero value otherwise.

### GetSfAccountIdOk

`func (o *McmClusterClaimResponseParams) GetSfAccountIdOk() (*string, bool)`

GetSfAccountIdOk returns a tuple with the SfAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccountId

`func (o *McmClusterClaimResponseParams) SetSfAccountId(v string)`

SetSfAccountId sets SfAccountId field to given value.

### HasSfAccountId

`func (o *McmClusterClaimResponseParams) HasSfAccountId() bool`

HasSfAccountId returns a boolean if a field has been set.

### SetSfAccountIdNil

`func (o *McmClusterClaimResponseParams) SetSfAccountIdNil(b bool)`

 SetSfAccountIdNil sets the value for SfAccountId to be an explicit nil

### UnsetSfAccountId
`func (o *McmClusterClaimResponseParams) UnsetSfAccountId()`

UnsetSfAccountId ensures that no value is present for SfAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


