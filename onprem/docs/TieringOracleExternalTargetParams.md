# TieringOracleExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**Region** | **NullableString** | Specifies the region of the external target. | 
**AccessKeyId** | **NullableString** | Specifies the access key id of the external target. | 
**StorageAccessKey** | **NullableString** | Specifies the storage access key of the external target. | 
**Tenancy** | **NullableString** | Specifies the tenancy of the external target. | 
**StorageClass** | **NullableString** | Specifies the Oracle External Target class. | 

## Methods

### NewTieringOracleExternalTargetParams

`func NewTieringOracleExternalTargetParams(bucketName NullableString, region NullableString, accessKeyId NullableString, storageAccessKey NullableString, tenancy NullableString, storageClass NullableString, ) *TieringOracleExternalTargetParams`

NewTieringOracleExternalTargetParams instantiates a new TieringOracleExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringOracleExternalTargetParamsWithDefaults

`func NewTieringOracleExternalTargetParamsWithDefaults() *TieringOracleExternalTargetParams`

NewTieringOracleExternalTargetParamsWithDefaults instantiates a new TieringOracleExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *TieringOracleExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TieringOracleExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TieringOracleExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *TieringOracleExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *TieringOracleExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *TieringOracleExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TieringOracleExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TieringOracleExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *TieringOracleExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *TieringOracleExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetAccessKeyId

`func (o *TieringOracleExternalTargetParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *TieringOracleExternalTargetParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *TieringOracleExternalTargetParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *TieringOracleExternalTargetParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *TieringOracleExternalTargetParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetStorageAccessKey

`func (o *TieringOracleExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *TieringOracleExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *TieringOracleExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.


### SetStorageAccessKeyNil

`func (o *TieringOracleExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *TieringOracleExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetTenancy

`func (o *TieringOracleExternalTargetParams) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *TieringOracleExternalTargetParams) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *TieringOracleExternalTargetParams) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.


### SetTenancyNil

`func (o *TieringOracleExternalTargetParams) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *TieringOracleExternalTargetParams) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
### GetStorageClass

`func (o *TieringOracleExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *TieringOracleExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *TieringOracleExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *TieringOracleExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *TieringOracleExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


