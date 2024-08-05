# KmsConfigurationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTargetIds** | Pointer to **[]int64** | Ids of external targets used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to an external target, it cannot be changed. | [optional] 
**KmipKmsParams** | Pointer to [**KmipKmsConfiguration**](KmipKmsConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the KMS. | 
**StorageDomainIds** | Pointer to **[]int64** | Ids of storage domains used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to a storage domain, it cannot be changed. | [optional] 
**AwsKmsParams** | Pointer to [**AwsKmsConfigurationUpdateParams**](AwsKmsConfigurationUpdateParams.md) |  | [optional] 

## Methods

### NewKmsConfigurationUpdateParams

`func NewKmsConfigurationUpdateParams(name string, ) *KmsConfigurationUpdateParams`

NewKmsConfigurationUpdateParams instantiates a new KmsConfigurationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationUpdateParamsWithDefaults

`func NewKmsConfigurationUpdateParamsWithDefaults() *KmsConfigurationUpdateParams`

NewKmsConfigurationUpdateParamsWithDefaults instantiates a new KmsConfigurationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTargetIds

`func (o *KmsConfigurationUpdateParams) GetExternalTargetIds() []int64`

GetExternalTargetIds returns the ExternalTargetIds field if non-nil, zero value otherwise.

### GetExternalTargetIdsOk

`func (o *KmsConfigurationUpdateParams) GetExternalTargetIdsOk() (*[]int64, bool)`

GetExternalTargetIdsOk returns a tuple with the ExternalTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetIds

`func (o *KmsConfigurationUpdateParams) SetExternalTargetIds(v []int64)`

SetExternalTargetIds sets ExternalTargetIds field to given value.

### HasExternalTargetIds

`func (o *KmsConfigurationUpdateParams) HasExternalTargetIds() bool`

HasExternalTargetIds returns a boolean if a field has been set.

### SetExternalTargetIdsNil

`func (o *KmsConfigurationUpdateParams) SetExternalTargetIdsNil(b bool)`

 SetExternalTargetIdsNil sets the value for ExternalTargetIds to be an explicit nil

### UnsetExternalTargetIds
`func (o *KmsConfigurationUpdateParams) UnsetExternalTargetIds()`

UnsetExternalTargetIds ensures that no value is present for ExternalTargetIds, not even an explicit nil
### GetKmipKmsParams

`func (o *KmsConfigurationUpdateParams) GetKmipKmsParams() KmipKmsConfiguration`

GetKmipKmsParams returns the KmipKmsParams field if non-nil, zero value otherwise.

### GetKmipKmsParamsOk

`func (o *KmsConfigurationUpdateParams) GetKmipKmsParamsOk() (*KmipKmsConfiguration, bool)`

GetKmipKmsParamsOk returns a tuple with the KmipKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipKmsParams

`func (o *KmsConfigurationUpdateParams) SetKmipKmsParams(v KmipKmsConfiguration)`

SetKmipKmsParams sets KmipKmsParams field to given value.

### HasKmipKmsParams

`func (o *KmsConfigurationUpdateParams) HasKmipKmsParams() bool`

HasKmipKmsParams returns a boolean if a field has been set.

### GetName

`func (o *KmsConfigurationUpdateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsConfigurationUpdateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsConfigurationUpdateParams) SetName(v string)`

SetName sets Name field to given value.


### GetStorageDomainIds

`func (o *KmsConfigurationUpdateParams) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *KmsConfigurationUpdateParams) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *KmsConfigurationUpdateParams) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *KmsConfigurationUpdateParams) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *KmsConfigurationUpdateParams) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *KmsConfigurationUpdateParams) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetAwsKmsParams

`func (o *KmsConfigurationUpdateParams) GetAwsKmsParams() AwsKmsConfigurationUpdateParams`

GetAwsKmsParams returns the AwsKmsParams field if non-nil, zero value otherwise.

### GetAwsKmsParamsOk

`func (o *KmsConfigurationUpdateParams) GetAwsKmsParamsOk() (*AwsKmsConfigurationUpdateParams, bool)`

GetAwsKmsParamsOk returns a tuple with the AwsKmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKmsParams

`func (o *KmsConfigurationUpdateParams) SetAwsKmsParams(v AwsKmsConfigurationUpdateParams)`

SetAwsKmsParams sets AwsKmsParams field to given value.

### HasAwsKmsParams

`func (o *KmsConfigurationUpdateParams) HasAwsKmsParams() bool`

HasAwsKmsParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


