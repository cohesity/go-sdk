# CommonGcpExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**ClientEmailAddress** | **NullableString** | Specifies the client email address of the external target. | 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies the client private key of the external target. | [optional] 
**ProjectId** | **NullableString** | Specifies the project Id of the external target. | 

## Methods

### NewCommonGcpExternalTargetParams

`func NewCommonGcpExternalTargetParams(bucketName NullableString, clientEmailAddress NullableString, projectId NullableString, ) *CommonGcpExternalTargetParams`

NewCommonGcpExternalTargetParams instantiates a new CommonGcpExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonGcpExternalTargetParamsWithDefaults

`func NewCommonGcpExternalTargetParamsWithDefaults() *CommonGcpExternalTargetParams`

NewCommonGcpExternalTargetParamsWithDefaults instantiates a new CommonGcpExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CommonGcpExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonGcpExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonGcpExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonGcpExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonGcpExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetClientEmailAddress

`func (o *CommonGcpExternalTargetParams) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *CommonGcpExternalTargetParams) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *CommonGcpExternalTargetParams) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.


### SetClientEmailAddressNil

`func (o *CommonGcpExternalTargetParams) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *CommonGcpExternalTargetParams) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *CommonGcpExternalTargetParams) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *CommonGcpExternalTargetParams) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *CommonGcpExternalTargetParams) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *CommonGcpExternalTargetParams) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *CommonGcpExternalTargetParams) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *CommonGcpExternalTargetParams) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetProjectId

`func (o *CommonGcpExternalTargetParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CommonGcpExternalTargetParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CommonGcpExternalTargetParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *CommonGcpExternalTargetParams) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *CommonGcpExternalTargetParams) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


