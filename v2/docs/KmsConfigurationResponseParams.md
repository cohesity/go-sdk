# KmsConfigurationResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKmsParams** | Pointer to [**AwsKmsConfigurationResponse**](AwsKmsConfigurationResponse.md) |  | [optional] 
**ExternalTargetIds** | Pointer to **[]int64** | Ids of external targets used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to an external target, it cannot be changed. | [optional] 
**KmipKmsParams** | Pointer to [**KmipKmsConfigurationResponse**](KmipKmsConfigurationResponse.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name of the KMS. | [optional] 
**OwnershipContext** | Pointer to **NullableString** | Describes the consumption of the KMS key whether it is used for local or FortKnox. | [optional] 
**StorageDomainIds** | Pointer to **[]int64** | Ids of storage domains used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to a storage domain, it cannot be changed. | [optional] 
**Type** | Pointer to **NullableString** | Type of KMS. &#39;InternalKms&#39; indicates the internal cluster KMS. &#39;AwsKms&#39; indicates AWS KMS. &#39;KmipKms&#39; indicates any KMIP compliant KMS. | [optional] 
**UsageType** | Pointer to **NullableString** | Specifies the usage type of the kms config. &#39;kArchival&#39; indicates this is used for regular archival. &#39;kRpaasArchival&#39; indicates this is used for RPaaS only. | [optional] 

## Methods

### NewKmsConfigurationResponseParams

`func NewKmsConfigurationResponseParams() *KmsConfigurationResponseParams`

NewKmsConfigurationResponseParams instantiates a new KmsConfigurationResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationResponseParamsWithDefaults

`func NewKmsConfigurationResponseParamsWithDefaults() *KmsConfigurationResponseParams`

NewKmsConfigurationResponseParamsWithDefaults instantiates a new KmsConfigurationResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKmsParams

`func (o *KmsConfigurationResponseParams) GetAwsKmsParams() AwsKmsConfigurationResponse`

GetAwsKmsParams returns the AwsKmsParams field if non-nil, zero value otherwise.

### GetAwsKmsParamsOk

`func (o *KmsConfigurationResponseParams) GetAwsKmsParamsOk() (*AwsKmsConfigurationResponse, bool)`

GetAwsKmsParamsOk returns a tuple with the AwsKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsParams

`func (o *KmsConfigurationResponseParams) SetAwsKmsParams(v AwsKmsConfigurationResponse)`

SetAwsKmsParams sets AwsKmsParams field to given value.

### HasAwsKmsParams

`func (o *KmsConfigurationResponseParams) HasAwsKmsParams() bool`

HasAwsKmsParams returns a boolean if a field has been set.

### GetExternalTargetIds

`func (o *KmsConfigurationResponseParams) GetExternalTargetIds() []int64`

GetExternalTargetIds returns the ExternalTargetIds field if non-nil, zero value otherwise.

### GetExternalTargetIdsOk

`func (o *KmsConfigurationResponseParams) GetExternalTargetIdsOk() (*[]int64, bool)`

GetExternalTargetIdsOk returns a tuple with the ExternalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetIds

`func (o *KmsConfigurationResponseParams) SetExternalTargetIds(v []int64)`

SetExternalTargetIds sets ExternalTargetIds field to given value.

### HasExternalTargetIds

`func (o *KmsConfigurationResponseParams) HasExternalTargetIds() bool`

HasExternalTargetIds returns a boolean if a field has been set.

### SetExternalTargetIdsNil

`func (o *KmsConfigurationResponseParams) SetExternalTargetIdsNil(b bool)`

 SetExternalTargetIdsNil sets the value for ExternalTargetIds to be an explicit nil

### UnsetExternalTargetIds
`func (o *KmsConfigurationResponseParams) UnsetExternalTargetIds()`

UnsetExternalTargetIds ensures that no value is present for ExternalTargetIds, not even an explicit nil
### GetKmipKmsParams

`func (o *KmsConfigurationResponseParams) GetKmipKmsParams() KmipKmsConfigurationResponse`

GetKmipKmsParams returns the KmipKmsParams field if non-nil, zero value otherwise.

### GetKmipKmsParamsOk

`func (o *KmsConfigurationResponseParams) GetKmipKmsParamsOk() (*KmipKmsConfigurationResponse, bool)`

GetKmipKmsParamsOk returns a tuple with the KmipKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipKmsParams

`func (o *KmsConfigurationResponseParams) SetKmipKmsParams(v KmipKmsConfigurationResponse)`

SetKmipKmsParams sets KmipKmsParams field to given value.

### HasKmipKmsParams

`func (o *KmsConfigurationResponseParams) HasKmipKmsParams() bool`

HasKmipKmsParams returns a boolean if a field has been set.

### GetName

`func (o *KmsConfigurationResponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsConfigurationResponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsConfigurationResponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsConfigurationResponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KmsConfigurationResponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KmsConfigurationResponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnershipContext

`func (o *KmsConfigurationResponseParams) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *KmsConfigurationResponseParams) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *KmsConfigurationResponseParams) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *KmsConfigurationResponseParams) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *KmsConfigurationResponseParams) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *KmsConfigurationResponseParams) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetStorageDomainIds

`func (o *KmsConfigurationResponseParams) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *KmsConfigurationResponseParams) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *KmsConfigurationResponseParams) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *KmsConfigurationResponseParams) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *KmsConfigurationResponseParams) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *KmsConfigurationResponseParams) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetType

`func (o *KmsConfigurationResponseParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsConfigurationResponseParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsConfigurationResponseParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsConfigurationResponseParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *KmsConfigurationResponseParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KmsConfigurationResponseParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsageType

`func (o *KmsConfigurationResponseParams) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *KmsConfigurationResponseParams) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *KmsConfigurationResponseParams) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *KmsConfigurationResponseParams) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *KmsConfigurationResponseParams) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *KmsConfigurationResponseParams) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


