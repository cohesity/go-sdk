# KmsConfigurationCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTargetIds** | Pointer to **[]int64** | Ids of external targets used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to an external target, it cannot be changed. | [optional] 
**KmipKmsParams** | Pointer to [**KmipKmsConfiguration**](KmipKmsConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the KMS. | 
**StorageDomainIds** | Pointer to **[]int64** | Ids of storage domains used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to a storage domain, it cannot be changed. | [optional] 
**AwsKmsParams** | Pointer to [**AwsKmsConfiguration**](AwsKmsConfiguration.md) |  | [optional] 
**OwnershipContext** | Pointer to **NullableString** | Specifies the ownership context of the kms config. &#39;Local&#39; indicates this is used for regular archival. &#39;FortKnox&#39; indicates this is used for FortKnox only. | [optional] 
**Type** | **string** | Type of KMS. &#39;InternalKms&#39; indicates the internal cluster KMS. &#39;AwsKms&#39; indicates AWS KMS. &#39;KmipKms&#39; indicates any KMIP compliant KMS. | 
**UsageType** | Pointer to **NullableString** | Specifies the usage type of the kms config. &#39;kArchival&#39; indicates this is used for regular archival. &#39;kRpaasArchival&#39; indicates this is used for RPaaS only. | [optional] 

## Methods

### NewKmsConfigurationCreateParams

`func NewKmsConfigurationCreateParams(name string, type_ string, ) *KmsConfigurationCreateParams`

NewKmsConfigurationCreateParams instantiates a new KmsConfigurationCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationCreateParamsWithDefaults

`func NewKmsConfigurationCreateParamsWithDefaults() *KmsConfigurationCreateParams`

NewKmsConfigurationCreateParamsWithDefaults instantiates a new KmsConfigurationCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTargetIds

`func (o *KmsConfigurationCreateParams) GetExternalTargetIds() []int64`

GetExternalTargetIds returns the ExternalTargetIds field if non-nil, zero value otherwise.

### GetExternalTargetIdsOk

`func (o *KmsConfigurationCreateParams) GetExternalTargetIdsOk() (*[]int64, bool)`

GetExternalTargetIdsOk returns a tuple with the ExternalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetIds

`func (o *KmsConfigurationCreateParams) SetExternalTargetIds(v []int64)`

SetExternalTargetIds sets ExternalTargetIds field to given value.

### HasExternalTargetIds

`func (o *KmsConfigurationCreateParams) HasExternalTargetIds() bool`

HasExternalTargetIds returns a boolean if a field has been set.

### SetExternalTargetIdsNil

`func (o *KmsConfigurationCreateParams) SetExternalTargetIdsNil(b bool)`

 SetExternalTargetIdsNil sets the value for ExternalTargetIds to be an explicit nil

### UnsetExternalTargetIds
`func (o *KmsConfigurationCreateParams) UnsetExternalTargetIds()`

UnsetExternalTargetIds ensures that no value is present for ExternalTargetIds, not even an explicit nil
### GetKmipKmsParams

`func (o *KmsConfigurationCreateParams) GetKmipKmsParams() KmipKmsConfiguration`

GetKmipKmsParams returns the KmipKmsParams field if non-nil, zero value otherwise.

### GetKmipKmsParamsOk

`func (o *KmsConfigurationCreateParams) GetKmipKmsParamsOk() (*KmipKmsConfiguration, bool)`

GetKmipKmsParamsOk returns a tuple with the KmipKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipKmsParams

`func (o *KmsConfigurationCreateParams) SetKmipKmsParams(v KmipKmsConfiguration)`

SetKmipKmsParams sets KmipKmsParams field to given value.

### HasKmipKmsParams

`func (o *KmsConfigurationCreateParams) HasKmipKmsParams() bool`

HasKmipKmsParams returns a boolean if a field has been set.

### GetName

`func (o *KmsConfigurationCreateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsConfigurationCreateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsConfigurationCreateParams) SetName(v string)`

SetName sets Name field to given value.


### GetStorageDomainIds

`func (o *KmsConfigurationCreateParams) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *KmsConfigurationCreateParams) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *KmsConfigurationCreateParams) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *KmsConfigurationCreateParams) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *KmsConfigurationCreateParams) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *KmsConfigurationCreateParams) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetAwsKmsParams

`func (o *KmsConfigurationCreateParams) GetAwsKmsParams() AwsKmsConfiguration`

GetAwsKmsParams returns the AwsKmsParams field if non-nil, zero value otherwise.

### GetAwsKmsParamsOk

`func (o *KmsConfigurationCreateParams) GetAwsKmsParamsOk() (*AwsKmsConfiguration, bool)`

GetAwsKmsParamsOk returns a tuple with the AwsKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsParams

`func (o *KmsConfigurationCreateParams) SetAwsKmsParams(v AwsKmsConfiguration)`

SetAwsKmsParams sets AwsKmsParams field to given value.

### HasAwsKmsParams

`func (o *KmsConfigurationCreateParams) HasAwsKmsParams() bool`

HasAwsKmsParams returns a boolean if a field has been set.

### GetOwnershipContext

`func (o *KmsConfigurationCreateParams) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *KmsConfigurationCreateParams) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *KmsConfigurationCreateParams) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *KmsConfigurationCreateParams) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *KmsConfigurationCreateParams) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *KmsConfigurationCreateParams) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetType

`func (o *KmsConfigurationCreateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsConfigurationCreateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsConfigurationCreateParams) SetType(v string)`

SetType sets Type field to given value.


### GetUsageType

`func (o *KmsConfigurationCreateParams) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *KmsConfigurationCreateParams) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *KmsConfigurationCreateParams) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *KmsConfigurationCreateParams) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *KmsConfigurationCreateParams) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *KmsConfigurationCreateParams) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


