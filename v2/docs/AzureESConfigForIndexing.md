# AzureESConfigForIndexing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **NullableString** | Client Id for the KeyVault. | 
**EsDomain** | **NullableString** | Fully qualified ES domain name. | 
**SecretName** | **NullableString** | Name of the secret corresponding to tenant&#39;s ES creds. | 
**VaultURL** | **NullableString** | URL of the KeyVault where ES creds will be stored. | 

## Methods

### NewAzureESConfigForIndexing

`func NewAzureESConfigForIndexing(clientId NullableString, esDomain NullableString, secretName NullableString, vaultURL NullableString, ) *AzureESConfigForIndexing`

NewAzureESConfigForIndexing instantiates a new AzureESConfigForIndexing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureESConfigForIndexingWithDefaults

`func NewAzureESConfigForIndexingWithDefaults() *AzureESConfigForIndexing`

NewAzureESConfigForIndexingWithDefaults instantiates a new AzureESConfigForIndexing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureESConfigForIndexing) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureESConfigForIndexing) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureESConfigForIndexing) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### SetClientIdNil

`func (o *AzureESConfigForIndexing) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *AzureESConfigForIndexing) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetEsDomain

`func (o *AzureESConfigForIndexing) GetEsDomain() string`

GetEsDomain returns the EsDomain field if non-nil, zero value otherwise.

### GetEsDomainOk

`func (o *AzureESConfigForIndexing) GetEsDomainOk() (*string, bool)`

GetEsDomainOk returns a tuple with the EsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsDomain

`func (o *AzureESConfigForIndexing) SetEsDomain(v string)`

SetEsDomain sets EsDomain field to given value.


### SetEsDomainNil

`func (o *AzureESConfigForIndexing) SetEsDomainNil(b bool)`

 SetEsDomainNil sets the value for EsDomain to be an explicit nil

### UnsetEsDomain
`func (o *AzureESConfigForIndexing) UnsetEsDomain()`

UnsetEsDomain ensures that no value is present for EsDomain, not even an explicit nil
### GetSecretName

`func (o *AzureESConfigForIndexing) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *AzureESConfigForIndexing) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *AzureESConfigForIndexing) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### SetSecretNameNil

`func (o *AzureESConfigForIndexing) SetSecretNameNil(b bool)`

 SetSecretNameNil sets the value for SecretName to be an explicit nil

### UnsetSecretName
`func (o *AzureESConfigForIndexing) UnsetSecretName()`

UnsetSecretName ensures that no value is present for SecretName, not even an explicit nil
### GetVaultURL

`func (o *AzureESConfigForIndexing) GetVaultURL() string`

GetVaultURL returns the VaultURL field if non-nil, zero value otherwise.

### GetVaultURLOk

`func (o *AzureESConfigForIndexing) GetVaultURLOk() (*string, bool)`

GetVaultURLOk returns a tuple with the VaultURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultURL

`func (o *AzureESConfigForIndexing) SetVaultURL(v string)`

SetVaultURL sets VaultURL field to given value.


### SetVaultURLNil

`func (o *AzureESConfigForIndexing) SetVaultURLNil(b bool)`

 SetVaultURLNil sets the value for VaultURL to be an explicit nil

### UnsetVaultURL
`func (o *AzureESConfigForIndexing) UnsetVaultURL()`

UnsetVaultURL ensures that no value is present for VaultURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


