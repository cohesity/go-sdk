# AzureCloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageAccessKey** | Pointer to **NullableString** | Specifies the access key to use when accessing a storage tier in a Azure cloud service. | [optional] 
**StorageAccountName** | Pointer to **NullableString** | Specifies the account name to use when accessing a storage tier in a Azure cloud service. | [optional] 
**TierType** | Pointer to **NullableString** | Specifies the storage class of Azure. AzureTierType specifies the storage class for Azure. &#39;kAzureTierHot&#39; indicates a tier type of Azure properties that is accessed frequently. &#39;kAzureTierCool&#39; indicates a tier type of Azure properties that is accessed less frequently, and stored for at least 30 days. &#39;kAzureTierArchive&#39; indicates a tier type of Azure properties that is accessed rarely and stored for at least 180 days. | [optional] 

## Methods

### NewAzureCloudCredentials

`func NewAzureCloudCredentials() *AzureCloudCredentials`

NewAzureCloudCredentials instantiates a new AzureCloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudCredentialsWithDefaults

`func NewAzureCloudCredentialsWithDefaults() *AzureCloudCredentials`

NewAzureCloudCredentialsWithDefaults instantiates a new AzureCloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageAccessKey

`func (o *AzureCloudCredentials) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *AzureCloudCredentials) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *AzureCloudCredentials) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *AzureCloudCredentials) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *AzureCloudCredentials) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *AzureCloudCredentials) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetStorageAccountName

`func (o *AzureCloudCredentials) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureCloudCredentials) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureCloudCredentials) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *AzureCloudCredentials) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### SetStorageAccountNameNil

`func (o *AzureCloudCredentials) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *AzureCloudCredentials) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetTierType

`func (o *AzureCloudCredentials) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *AzureCloudCredentials) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *AzureCloudCredentials) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *AzureCloudCredentials) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### SetTierTypeNil

`func (o *AzureCloudCredentials) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *AzureCloudCredentials) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


