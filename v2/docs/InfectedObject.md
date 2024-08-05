# InfectedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the infected object. | 
**ObjectName** | **NullableString** | Specifies the key name of the infected object. | 
**VersionId** | Pointer to **NullableString** | Specifies the version id of the infected object. | [optional] 

## Methods

### NewInfectedObject

`func NewInfectedObject(bucketName NullableString, objectName NullableString, ) *InfectedObject`

NewInfectedObject instantiates a new InfectedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedObjectWithDefaults

`func NewInfectedObjectWithDefaults() *InfectedObject`

NewInfectedObjectWithDefaults instantiates a new InfectedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *InfectedObject) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *InfectedObject) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *InfectedObject) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *InfectedObject) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *InfectedObject) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetObjectName

`func (o *InfectedObject) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *InfectedObject) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *InfectedObject) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *InfectedObject) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *InfectedObject) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetVersionId

`func (o *InfectedObject) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *InfectedObject) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *InfectedObject) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *InfectedObject) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### SetVersionIdNil

`func (o *InfectedObject) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *InfectedObject) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


