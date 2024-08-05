# KmsConfiguration

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
**Id** | Pointer to **NullableInt64** | Id of KMS. | [optional] [readonly] 
**State** | Pointer to **NullableString** | Specifies the state of KMS. &#39;Active&#39; indicates that KMS is reachable from cluster. &#39;InActive&#39; indicates that KMS is not reachable from cluster. &#39;MarkedForRemoval&#39; indicates that KMS is marked for removal and the removal process is in progress. | [optional] [readonly] 

## Methods

### NewKmsConfiguration

`func NewKmsConfiguration() *KmsConfiguration`

NewKmsConfiguration instantiates a new KmsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationWithDefaults

`func NewKmsConfigurationWithDefaults() *KmsConfiguration`

NewKmsConfigurationWithDefaults instantiates a new KmsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKmsParams

`func (o *KmsConfiguration) GetAwsKmsParams() AwsKmsConfigurationResponse`

GetAwsKmsParams returns the AwsKmsParams field if non-nil, zero value otherwise.

### GetAwsKmsParamsOk

`func (o *KmsConfiguration) GetAwsKmsParamsOk() (*AwsKmsConfigurationResponse, bool)`

GetAwsKmsParamsOk returns a tuple with the AwsKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsParams

`func (o *KmsConfiguration) SetAwsKmsParams(v AwsKmsConfigurationResponse)`

SetAwsKmsParams sets AwsKmsParams field to given value.

### HasAwsKmsParams

`func (o *KmsConfiguration) HasAwsKmsParams() bool`

HasAwsKmsParams returns a boolean if a field has been set.

### GetExternalTargetIds

`func (o *KmsConfiguration) GetExternalTargetIds() []int64`

GetExternalTargetIds returns the ExternalTargetIds field if non-nil, zero value otherwise.

### GetExternalTargetIdsOk

`func (o *KmsConfiguration) GetExternalTargetIdsOk() (*[]int64, bool)`

GetExternalTargetIdsOk returns a tuple with the ExternalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetIds

`func (o *KmsConfiguration) SetExternalTargetIds(v []int64)`

SetExternalTargetIds sets ExternalTargetIds field to given value.

### HasExternalTargetIds

`func (o *KmsConfiguration) HasExternalTargetIds() bool`

HasExternalTargetIds returns a boolean if a field has been set.

### SetExternalTargetIdsNil

`func (o *KmsConfiguration) SetExternalTargetIdsNil(b bool)`

 SetExternalTargetIdsNil sets the value for ExternalTargetIds to be an explicit nil

### UnsetExternalTargetIds
`func (o *KmsConfiguration) UnsetExternalTargetIds()`

UnsetExternalTargetIds ensures that no value is present for ExternalTargetIds, not even an explicit nil
### GetKmipKmsParams

`func (o *KmsConfiguration) GetKmipKmsParams() KmipKmsConfigurationResponse`

GetKmipKmsParams returns the KmipKmsParams field if non-nil, zero value otherwise.

### GetKmipKmsParamsOk

`func (o *KmsConfiguration) GetKmipKmsParamsOk() (*KmipKmsConfigurationResponse, bool)`

GetKmipKmsParamsOk returns a tuple with the KmipKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipKmsParams

`func (o *KmsConfiguration) SetKmipKmsParams(v KmipKmsConfigurationResponse)`

SetKmipKmsParams sets KmipKmsParams field to given value.

### HasKmipKmsParams

`func (o *KmsConfiguration) HasKmipKmsParams() bool`

HasKmipKmsParams returns a boolean if a field has been set.

### GetName

`func (o *KmsConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KmsConfiguration) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KmsConfiguration) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnershipContext

`func (o *KmsConfiguration) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *KmsConfiguration) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *KmsConfiguration) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *KmsConfiguration) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *KmsConfiguration) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *KmsConfiguration) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetStorageDomainIds

`func (o *KmsConfiguration) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *KmsConfiguration) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *KmsConfiguration) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *KmsConfiguration) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *KmsConfiguration) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *KmsConfiguration) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetType

`func (o *KmsConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *KmsConfiguration) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KmsConfiguration) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsageType

`func (o *KmsConfiguration) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *KmsConfiguration) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *KmsConfiguration) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *KmsConfiguration) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *KmsConfiguration) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *KmsConfiguration) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetId

`func (o *KmsConfiguration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsConfiguration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsConfiguration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KmsConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KmsConfiguration) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KmsConfiguration) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetState

`func (o *KmsConfiguration) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KmsConfiguration) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KmsConfiguration) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KmsConfiguration) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *KmsConfiguration) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *KmsConfiguration) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


