# ObjectSnapshotType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **NullableString** | This captures any additional message about the snapshot itself, e.g. if the app-consistent snapshot had to fallback to crash consistent, this will contain that. | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewObjectSnapshotType

`func NewObjectSnapshotType() *ObjectSnapshotType`

NewObjectSnapshotType instantiates a new ObjectSnapshotType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotTypeWithDefaults

`func NewObjectSnapshotTypeWithDefaults() *ObjectSnapshotType`

NewObjectSnapshotTypeWithDefaults instantiates a new ObjectSnapshotType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ObjectSnapshotType) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ObjectSnapshotType) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ObjectSnapshotType) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ObjectSnapshotType) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *ObjectSnapshotType) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *ObjectSnapshotType) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetType

`func (o *ObjectSnapshotType) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectSnapshotType) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectSnapshotType) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ObjectSnapshotType) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ObjectSnapshotType) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ObjectSnapshotType) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


