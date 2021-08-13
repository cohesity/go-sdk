# SnapshotTagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **NullableString** | Specifies Id of tag applied to the object. | 
**RunIds** | Pointer to **[]string** | Specifies runs the tags are applied to. | [optional] 

## Methods

### NewSnapshotTagInfo

`func NewSnapshotTagInfo(tagId NullableString, ) *SnapshotTagInfo`

NewSnapshotTagInfo instantiates a new SnapshotTagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotTagInfoWithDefaults

`func NewSnapshotTagInfoWithDefaults() *SnapshotTagInfo`

NewSnapshotTagInfoWithDefaults instantiates a new SnapshotTagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *SnapshotTagInfo) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *SnapshotTagInfo) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *SnapshotTagInfo) SetTagId(v string)`

SetTagId sets TagId field to given value.


### SetTagIdNil

`func (o *SnapshotTagInfo) SetTagIdNil(b bool)`

 SetTagIdNil sets the value for TagId to be an explicit nil

### UnsetTagId
`func (o *SnapshotTagInfo) UnsetTagId()`

UnsetTagId ensures that no value is present for TagId, not even an explicit nil
### GetRunIds

`func (o *SnapshotTagInfo) GetRunIds() []string`

GetRunIds returns the RunIds field if non-nil, zero value otherwise.

### GetRunIdsOk

`func (o *SnapshotTagInfo) GetRunIdsOk() (*[]string, bool)`

GetRunIdsOk returns a tuple with the RunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunIds

`func (o *SnapshotTagInfo) SetRunIds(v []string)`

SetRunIds sets RunIds field to given value.

### HasRunIds

`func (o *SnapshotTagInfo) HasRunIds() bool`

HasRunIds returns a boolean if a field has been set.

### SetRunIdsNil

`func (o *SnapshotTagInfo) SetRunIdsNil(b bool)`

 SetRunIdsNil sets the value for RunIds to be an explicit nil

### UnsetRunIds
`func (o *SnapshotTagInfo) UnsetRunIds()`

UnsetRunIds ensures that no value is present for RunIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


