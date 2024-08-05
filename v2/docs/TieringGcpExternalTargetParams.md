# TieringGcpExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**ClientEmailAddress** | **NullableString** | Specifies the client email address of the external target. | 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies the client private key of the external target. | [optional] 
**ProjectId** | **NullableString** | Specifies the project Id of the external target. | 
**StorageClass** | **NullableString** | Specifies the GCP External Target class. | 

## Methods

### NewTieringGcpExternalTargetParams

`func NewTieringGcpExternalTargetParams(bucketName NullableString, clientEmailAddress NullableString, projectId NullableString, storageClass NullableString, ) *TieringGcpExternalTargetParams`

NewTieringGcpExternalTargetParams instantiates a new TieringGcpExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringGcpExternalTargetParamsWithDefaults

`func NewTieringGcpExternalTargetParamsWithDefaults() *TieringGcpExternalTargetParams`

NewTieringGcpExternalTargetParamsWithDefaults instantiates a new TieringGcpExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *TieringGcpExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TieringGcpExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TieringGcpExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *TieringGcpExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *TieringGcpExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetClientEmailAddress

`func (o *TieringGcpExternalTargetParams) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *TieringGcpExternalTargetParams) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *TieringGcpExternalTargetParams) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.


### SetClientEmailAddressNil

`func (o *TieringGcpExternalTargetParams) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *TieringGcpExternalTargetParams) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *TieringGcpExternalTargetParams) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *TieringGcpExternalTargetParams) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *TieringGcpExternalTargetParams) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *TieringGcpExternalTargetParams) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *TieringGcpExternalTargetParams) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *TieringGcpExternalTargetParams) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetProjectId

`func (o *TieringGcpExternalTargetParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TieringGcpExternalTargetParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TieringGcpExternalTargetParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *TieringGcpExternalTargetParams) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *TieringGcpExternalTargetParams) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetStorageClass

`func (o *TieringGcpExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *TieringGcpExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *TieringGcpExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *TieringGcpExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *TieringGcpExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


