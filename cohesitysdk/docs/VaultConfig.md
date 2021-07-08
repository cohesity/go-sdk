# VaultConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amazon** | Pointer to [**AmazonCloudCredentials**](AmazonCloudCredentials.md) |  | [optional] 
**Azure** | Pointer to [**AzureCloudCredentials**](AzureCloudCredentials.md) |  | [optional] 
**BucketName** | Pointer to **NullableString** | Specifies the name of a storage location of the Vault, where objects are stored. For Google and AMS, this storage location is called a bucket. For Microsoft Azure, this storage location is called a container. For QStar and NAS, you do not specify a storage location. | [optional] 
**Google** | Pointer to [**GoogleCloudCredentials**](GoogleCloudCredentials.md) |  | [optional] 
**Nas** | Pointer to [**NasCredentials**](NasCredentials.md) |  | [optional] 
**Oracle** | Pointer to [**OracleCloudCredentials**](OracleCloudCredentials.md) |  | [optional] 
**Qstar** | Pointer to [**QStarServerCredentials**](QStarServerCredentials.md) |  | [optional] 

## Methods

### NewVaultConfig

`func NewVaultConfig() *VaultConfig`

NewVaultConfig instantiates a new VaultConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultConfigWithDefaults

`func NewVaultConfigWithDefaults() *VaultConfig`

NewVaultConfigWithDefaults instantiates a new VaultConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmazon

`func (o *VaultConfig) GetAmazon() AmazonCloudCredentials`

GetAmazon returns the Amazon field if non-nil, zero value otherwise.

### GetAmazonOk

`func (o *VaultConfig) GetAmazonOk() (*AmazonCloudCredentials, bool)`

GetAmazonOk returns a tuple with the Amazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazon

`func (o *VaultConfig) SetAmazon(v AmazonCloudCredentials)`

SetAmazon sets Amazon field to given value.

### HasAmazon

`func (o *VaultConfig) HasAmazon() bool`

HasAmazon returns a boolean if a field has been set.

### GetAzure

`func (o *VaultConfig) GetAzure() AzureCloudCredentials`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *VaultConfig) GetAzureOk() (*AzureCloudCredentials, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *VaultConfig) SetAzure(v AzureCloudCredentials)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *VaultConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetBucketName

`func (o *VaultConfig) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *VaultConfig) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *VaultConfig) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *VaultConfig) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### SetBucketNameNil

`func (o *VaultConfig) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *VaultConfig) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetGoogle

`func (o *VaultConfig) GetGoogle() GoogleCloudCredentials`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *VaultConfig) GetGoogleOk() (*GoogleCloudCredentials, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *VaultConfig) SetGoogle(v GoogleCloudCredentials)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *VaultConfig) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetNas

`func (o *VaultConfig) GetNas() NasCredentials`

GetNas returns the Nas field if non-nil, zero value otherwise.

### GetNasOk

`func (o *VaultConfig) GetNasOk() (*NasCredentials, bool)`

GetNasOk returns a tuple with the Nas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNas

`func (o *VaultConfig) SetNas(v NasCredentials)`

SetNas sets Nas field to given value.

### HasNas

`func (o *VaultConfig) HasNas() bool`

HasNas returns a boolean if a field has been set.

### GetOracle

`func (o *VaultConfig) GetOracle() OracleCloudCredentials`

GetOracle returns the Oracle field if non-nil, zero value otherwise.

### GetOracleOk

`func (o *VaultConfig) GetOracleOk() (*OracleCloudCredentials, bool)`

GetOracleOk returns a tuple with the Oracle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracle

`func (o *VaultConfig) SetOracle(v OracleCloudCredentials)`

SetOracle sets Oracle field to given value.

### HasOracle

`func (o *VaultConfig) HasOracle() bool`

HasOracle returns a boolean if a field has been set.

### GetQstar

`func (o *VaultConfig) GetQstar() QStarServerCredentials`

GetQstar returns the Qstar field if non-nil, zero value otherwise.

### GetQstarOk

`func (o *VaultConfig) GetQstarOk() (*QStarServerCredentials, bool)`

GetQstarOk returns a tuple with the Qstar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQstar

`func (o *VaultConfig) SetQstar(v QStarServerCredentials)`

SetQstar sets Qstar field to given value.

### HasQstar

`func (o *VaultConfig) HasQstar() bool`

HasQstar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


