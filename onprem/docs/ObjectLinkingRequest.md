# ObjectLinkingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectMap** | Pointer to [**[]ReplicaFailoverObject**](ReplicaFailoverObject.md) | Specifies the objectMap that will be used to create linking between given replicated source entity and newly restored entity on erplication cluster. | [optional] 

## Methods

### NewObjectLinkingRequest

`func NewObjectLinkingRequest() *ObjectLinkingRequest`

NewObjectLinkingRequest instantiates a new ObjectLinkingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectLinkingRequestWithDefaults

`func NewObjectLinkingRequestWithDefaults() *ObjectLinkingRequest`

NewObjectLinkingRequestWithDefaults instantiates a new ObjectLinkingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectMap

`func (o *ObjectLinkingRequest) GetObjectMap() []ReplicaFailoverObject`

GetObjectMap returns the ObjectMap field if non-nil, zero value otherwise.

### GetObjectMapOk

`func (o *ObjectLinkingRequest) GetObjectMapOk() (*[]ReplicaFailoverObject, bool)`

GetObjectMapOk returns a tuple with the ObjectMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMap

`func (o *ObjectLinkingRequest) SetObjectMap(v []ReplicaFailoverObject)`

SetObjectMap sets ObjectMap field to given value.

### HasObjectMap

`func (o *ObjectLinkingRequest) HasObjectMap() bool`

HasObjectMap returns a boolean if a field has been set.

### SetObjectMapNil

`func (o *ObjectLinkingRequest) SetObjectMapNil(b bool)`

 SetObjectMapNil sets the value for ObjectMap to be an explicit nil

### UnsetObjectMap
`func (o *ObjectLinkingRequest) UnsetObjectMap()`

UnsetObjectMap ensures that no value is present for ObjectMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


