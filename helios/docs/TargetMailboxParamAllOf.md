# TargetMailboxParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetFolderPath** | **NullableString** | Specifies the path to the target folder. | 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies the id of the domain for alternate domain recovery. | [optional] 

## Methods

### NewTargetMailboxParamAllOf

`func NewTargetMailboxParamAllOf(targetFolderPath NullableString, ) *TargetMailboxParamAllOf`

NewTargetMailboxParamAllOf instantiates a new TargetMailboxParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetMailboxParamAllOfWithDefaults

`func NewTargetMailboxParamAllOfWithDefaults() *TargetMailboxParamAllOf`

NewTargetMailboxParamAllOfWithDefaults instantiates a new TargetMailboxParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetFolderPath

`func (o *TargetMailboxParamAllOf) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *TargetMailboxParamAllOf) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *TargetMailboxParamAllOf) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.


### SetTargetFolderPathNil

`func (o *TargetMailboxParamAllOf) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *TargetMailboxParamAllOf) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetParentSourceId

`func (o *TargetMailboxParamAllOf) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *TargetMailboxParamAllOf) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *TargetMailboxParamAllOf) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *TargetMailboxParamAllOf) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *TargetMailboxParamAllOf) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *TargetMailboxParamAllOf) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


