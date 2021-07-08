# S3BucketConfigProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to [**ACLProto**](ACLProto.md) |  | [optional] 
**LifecycleConfig** | Pointer to **map[string]interface{}** | Protobuf that describes the lifecycle configuration that is used to manage the lifecycle of objects in a bucket. | [optional] 
**OwnerInfo** | Pointer to [**OwnerInfo**](OwnerInfo.md) |  | [optional] 
**ProtocolType** | Pointer to **NullableInt32** | Protocol type of this bucket. | [optional] 
**SwiftContainerTag** | Pointer to [**SwiftContainerTaggingProto**](SwiftContainerTaggingProto.md) |  | [optional] 
**TagMap** | Pointer to [**[]S3BucketConfigProtoTagMapEntry**](S3BucketConfigProtoTagMapEntry.md) | Tags (or labels) assigned to the bucket. Tags are set of &lt;key, value&gt; pairs. | [optional] 
**VersioningState** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewS3BucketConfigProto

`func NewS3BucketConfigProto() *S3BucketConfigProto`

NewS3BucketConfigProto instantiates a new S3BucketConfigProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketConfigProtoWithDefaults

`func NewS3BucketConfigProtoWithDefaults() *S3BucketConfigProto`

NewS3BucketConfigProtoWithDefaults instantiates a new S3BucketConfigProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *S3BucketConfigProto) GetAcl() ACLProto`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *S3BucketConfigProto) GetAclOk() (*ACLProto, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *S3BucketConfigProto) SetAcl(v ACLProto)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *S3BucketConfigProto) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetLifecycleConfig

`func (o *S3BucketConfigProto) GetLifecycleConfig() map[string]interface{}`

GetLifecycleConfig returns the LifecycleConfig field if non-nil, zero value otherwise.

### GetLifecycleConfigOk

`func (o *S3BucketConfigProto) GetLifecycleConfigOk() (*map[string]interface{}, bool)`

GetLifecycleConfigOk returns a tuple with the LifecycleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleConfig

`func (o *S3BucketConfigProto) SetLifecycleConfig(v map[string]interface{})`

SetLifecycleConfig sets LifecycleConfig field to given value.

### HasLifecycleConfig

`func (o *S3BucketConfigProto) HasLifecycleConfig() bool`

HasLifecycleConfig returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *S3BucketConfigProto) GetOwnerInfo() OwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *S3BucketConfigProto) GetOwnerInfoOk() (*OwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *S3BucketConfigProto) SetOwnerInfo(v OwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *S3BucketConfigProto) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetProtocolType

`func (o *S3BucketConfigProto) GetProtocolType() int32`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *S3BucketConfigProto) GetProtocolTypeOk() (*int32, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *S3BucketConfigProto) SetProtocolType(v int32)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *S3BucketConfigProto) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### SetProtocolTypeNil

`func (o *S3BucketConfigProto) SetProtocolTypeNil(b bool)`

 SetProtocolTypeNil sets the value for ProtocolType to be an explicit nil

### UnsetProtocolType
`func (o *S3BucketConfigProto) UnsetProtocolType()`

UnsetProtocolType ensures that no value is present for ProtocolType, not even an explicit nil
### GetSwiftContainerTag

`func (o *S3BucketConfigProto) GetSwiftContainerTag() SwiftContainerTaggingProto`

GetSwiftContainerTag returns the SwiftContainerTag field if non-nil, zero value otherwise.

### GetSwiftContainerTagOk

`func (o *S3BucketConfigProto) GetSwiftContainerTagOk() (*SwiftContainerTaggingProto, bool)`

GetSwiftContainerTagOk returns a tuple with the SwiftContainerTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftContainerTag

`func (o *S3BucketConfigProto) SetSwiftContainerTag(v SwiftContainerTaggingProto)`

SetSwiftContainerTag sets SwiftContainerTag field to given value.

### HasSwiftContainerTag

`func (o *S3BucketConfigProto) HasSwiftContainerTag() bool`

HasSwiftContainerTag returns a boolean if a field has been set.

### GetTagMap

`func (o *S3BucketConfigProto) GetTagMap() []S3BucketConfigProtoTagMapEntry`

GetTagMap returns the TagMap field if non-nil, zero value otherwise.

### GetTagMapOk

`func (o *S3BucketConfigProto) GetTagMapOk() (*[]S3BucketConfigProtoTagMapEntry, bool)`

GetTagMapOk returns a tuple with the TagMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagMap

`func (o *S3BucketConfigProto) SetTagMap(v []S3BucketConfigProtoTagMapEntry)`

SetTagMap sets TagMap field to given value.

### HasTagMap

`func (o *S3BucketConfigProto) HasTagMap() bool`

HasTagMap returns a boolean if a field has been set.

### SetTagMapNil

`func (o *S3BucketConfigProto) SetTagMapNil(b bool)`

 SetTagMapNil sets the value for TagMap to be an explicit nil

### UnsetTagMap
`func (o *S3BucketConfigProto) UnsetTagMap()`

UnsetTagMap ensures that no value is present for TagMap, not even an explicit nil
### GetVersioningState

`func (o *S3BucketConfigProto) GetVersioningState() int32`

GetVersioningState returns the VersioningState field if non-nil, zero value otherwise.

### GetVersioningStateOk

`func (o *S3BucketConfigProto) GetVersioningStateOk() (*int32, bool)`

GetVersioningStateOk returns a tuple with the VersioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioningState

`func (o *S3BucketConfigProto) SetVersioningState(v int32)`

SetVersioningState sets VersioningState field to given value.

### HasVersioningState

`func (o *S3BucketConfigProto) HasVersioningState() bool`

HasVersioningState returns a boolean if a field has been set.

### SetVersioningStateNil

`func (o *S3BucketConfigProto) SetVersioningStateNil(b bool)`

 SetVersioningStateNil sets the value for VersioningState to be an explicit nil

### UnsetVersioningState
`func (o *S3BucketConfigProto) UnsetVersioningState()`

UnsetVersioningState ensures that no value is present for VersioningState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


