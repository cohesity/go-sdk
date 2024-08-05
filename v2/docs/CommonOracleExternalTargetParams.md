# CommonOracleExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **NullableString** | Specifies the access key id of the external target. | 
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**Region** | **NullableString** | Specifies the region of the external target. | 
**StorageAccessKey** | Pointer to **NullableString** | Specifies the storage access key of the external target. | [optional] 
**Tenancy** | **NullableString** | Specifies the tenancy of the external target. | 

## Methods

### NewCommonOracleExternalTargetParams

`func NewCommonOracleExternalTargetParams(accessKeyId NullableString, bucketName NullableString, region NullableString, tenancy NullableString, ) *CommonOracleExternalTargetParams`

NewCommonOracleExternalTargetParams instantiates a new CommonOracleExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonOracleExternalTargetParamsWithDefaults

`func NewCommonOracleExternalTargetParamsWithDefaults() *CommonOracleExternalTargetParams`

NewCommonOracleExternalTargetParamsWithDefaults instantiates a new CommonOracleExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *CommonOracleExternalTargetParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CommonOracleExternalTargetParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CommonOracleExternalTargetParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *CommonOracleExternalTargetParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *CommonOracleExternalTargetParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetBucketName

`func (o *CommonOracleExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonOracleExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonOracleExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonOracleExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonOracleExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *CommonOracleExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonOracleExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonOracleExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *CommonOracleExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonOracleExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageAccessKey

`func (o *CommonOracleExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *CommonOracleExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *CommonOracleExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *CommonOracleExternalTargetParams) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *CommonOracleExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *CommonOracleExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetTenancy

`func (o *CommonOracleExternalTargetParams) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *CommonOracleExternalTargetParams) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *CommonOracleExternalTargetParams) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.


### SetTenancyNil

`func (o *CommonOracleExternalTargetParams) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *CommonOracleExternalTargetParams) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


