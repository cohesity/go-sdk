# DestroyClonedEntityInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClonedEntity** | Pointer to [**DestroyClonedEntityInfoProtoClonedEntity**](DestroyClonedEntityInfoProtoClonedEntity.md) |  | [optional] 
**ClonedEntityStatus** | Pointer to **NullableInt32** |  | [optional] 
**DestroyClonedEntityState** | Pointer to **NullableInt32** | The state of the destroy/teardown of a cloned entity (i.e, VM).  The following two fields are set by the slave in order for the master to find status of the destroy operation. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**FullViewName** | Pointer to **NullableString** | The full external view name where cloned objects are placed. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this destroy cloned entity info pertains to. | [optional] 

## Methods

### NewDestroyClonedEntityInfoProto

`func NewDestroyClonedEntityInfoProto() *DestroyClonedEntityInfoProto`

NewDestroyClonedEntityInfoProto instantiates a new DestroyClonedEntityInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyClonedEntityInfoProtoWithDefaults

`func NewDestroyClonedEntityInfoProtoWithDefaults() *DestroyClonedEntityInfoProto`

NewDestroyClonedEntityInfoProtoWithDefaults instantiates a new DestroyClonedEntityInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClonedEntity

`func (o *DestroyClonedEntityInfoProto) GetClonedEntity() DestroyClonedEntityInfoProtoClonedEntity`

GetClonedEntity returns the ClonedEntity field if non-nil, zero value otherwise.

### GetClonedEntityOk

`func (o *DestroyClonedEntityInfoProto) GetClonedEntityOk() (*DestroyClonedEntityInfoProtoClonedEntity, bool)`

GetClonedEntityOk returns a tuple with the ClonedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedEntity

`func (o *DestroyClonedEntityInfoProto) SetClonedEntity(v DestroyClonedEntityInfoProtoClonedEntity)`

SetClonedEntity sets ClonedEntity field to given value.

### HasClonedEntity

`func (o *DestroyClonedEntityInfoProto) HasClonedEntity() bool`

HasClonedEntity returns a boolean if a field has been set.

### GetClonedEntityStatus

`func (o *DestroyClonedEntityInfoProto) GetClonedEntityStatus() int32`

GetClonedEntityStatus returns the ClonedEntityStatus field if non-nil, zero value otherwise.

### GetClonedEntityStatusOk

`func (o *DestroyClonedEntityInfoProto) GetClonedEntityStatusOk() (*int32, bool)`

GetClonedEntityStatusOk returns a tuple with the ClonedEntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedEntityStatus

`func (o *DestroyClonedEntityInfoProto) SetClonedEntityStatus(v int32)`

SetClonedEntityStatus sets ClonedEntityStatus field to given value.

### HasClonedEntityStatus

`func (o *DestroyClonedEntityInfoProto) HasClonedEntityStatus() bool`

HasClonedEntityStatus returns a boolean if a field has been set.

### SetClonedEntityStatusNil

`func (o *DestroyClonedEntityInfoProto) SetClonedEntityStatusNil(b bool)`

 SetClonedEntityStatusNil sets the value for ClonedEntityStatus to be an explicit nil

### UnsetClonedEntityStatus
`func (o *DestroyClonedEntityInfoProto) UnsetClonedEntityStatus()`

UnsetClonedEntityStatus ensures that no value is present for ClonedEntityStatus, not even an explicit nil
### GetDestroyClonedEntityState

`func (o *DestroyClonedEntityInfoProto) GetDestroyClonedEntityState() int32`

GetDestroyClonedEntityState returns the DestroyClonedEntityState field if non-nil, zero value otherwise.

### GetDestroyClonedEntityStateOk

`func (o *DestroyClonedEntityInfoProto) GetDestroyClonedEntityStateOk() (*int32, bool)`

GetDestroyClonedEntityStateOk returns a tuple with the DestroyClonedEntityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyClonedEntityState

`func (o *DestroyClonedEntityInfoProto) SetDestroyClonedEntityState(v int32)`

SetDestroyClonedEntityState sets DestroyClonedEntityState field to given value.

### HasDestroyClonedEntityState

`func (o *DestroyClonedEntityInfoProto) HasDestroyClonedEntityState() bool`

HasDestroyClonedEntityState returns a boolean if a field has been set.

### SetDestroyClonedEntityStateNil

`func (o *DestroyClonedEntityInfoProto) SetDestroyClonedEntityStateNil(b bool)`

 SetDestroyClonedEntityStateNil sets the value for DestroyClonedEntityState to be an explicit nil

### UnsetDestroyClonedEntityState
`func (o *DestroyClonedEntityInfoProto) UnsetDestroyClonedEntityState()`

UnsetDestroyClonedEntityState ensures that no value is present for DestroyClonedEntityState, not even an explicit nil
### GetError

`func (o *DestroyClonedEntityInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestroyClonedEntityInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestroyClonedEntityInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *DestroyClonedEntityInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFullViewName

`func (o *DestroyClonedEntityInfoProto) GetFullViewName() string`

GetFullViewName returns the FullViewName field if non-nil, zero value otherwise.

### GetFullViewNameOk

`func (o *DestroyClonedEntityInfoProto) GetFullViewNameOk() (*string, bool)`

GetFullViewNameOk returns a tuple with the FullViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullViewName

`func (o *DestroyClonedEntityInfoProto) SetFullViewName(v string)`

SetFullViewName sets FullViewName field to given value.

### HasFullViewName

`func (o *DestroyClonedEntityInfoProto) HasFullViewName() bool`

HasFullViewName returns a boolean if a field has been set.

### SetFullViewNameNil

`func (o *DestroyClonedEntityInfoProto) SetFullViewNameNil(b bool)`

 SetFullViewNameNil sets the value for FullViewName to be an explicit nil

### UnsetFullViewName
`func (o *DestroyClonedEntityInfoProto) UnsetFullViewName()`

UnsetFullViewName ensures that no value is present for FullViewName, not even an explicit nil
### GetType

`func (o *DestroyClonedEntityInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestroyClonedEntityInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestroyClonedEntityInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DestroyClonedEntityInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DestroyClonedEntityInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DestroyClonedEntityInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


