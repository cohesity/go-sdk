# FileStubbingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoOrphanDataCleanup** | Pointer to **NullableBool** | Specifies whether to remove the orphan data from the target if the symlink is removed from the source. | [optional] [default to true]
**DowntieringFileAge** | Pointer to [**DowntieringFileAgePolicy**](DowntieringFileAgePolicy.md) |  | [optional] 
**SkipBackSymlink** | Pointer to **NullableBool** | Specifies whether to create a symlink for the migrated data from source to target. | [optional] [default to true]

## Methods

### NewFileStubbingParams

`func NewFileStubbingParams() *FileStubbingParams`

NewFileStubbingParams instantiates a new FileStubbingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStubbingParamsWithDefaults

`func NewFileStubbingParamsWithDefaults() *FileStubbingParams`

NewFileStubbingParamsWithDefaults instantiates a new FileStubbingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoOrphanDataCleanup

`func (o *FileStubbingParams) GetAutoOrphanDataCleanup() bool`

GetAutoOrphanDataCleanup returns the AutoOrphanDataCleanup field if non-nil, zero value otherwise.

### GetAutoOrphanDataCleanupOk

`func (o *FileStubbingParams) GetAutoOrphanDataCleanupOk() (*bool, bool)`

GetAutoOrphanDataCleanupOk returns a tuple with the AutoOrphanDataCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOrphanDataCleanup

`func (o *FileStubbingParams) SetAutoOrphanDataCleanup(v bool)`

SetAutoOrphanDataCleanup sets AutoOrphanDataCleanup field to given value.

### HasAutoOrphanDataCleanup

`func (o *FileStubbingParams) HasAutoOrphanDataCleanup() bool`

HasAutoOrphanDataCleanup returns a boolean if a field has been set.

### SetAutoOrphanDataCleanupNil

`func (o *FileStubbingParams) SetAutoOrphanDataCleanupNil(b bool)`

 SetAutoOrphanDataCleanupNil sets the value for AutoOrphanDataCleanup to be an explicit nil

### UnsetAutoOrphanDataCleanup
`func (o *FileStubbingParams) UnsetAutoOrphanDataCleanup()`

UnsetAutoOrphanDataCleanup ensures that no value is present for AutoOrphanDataCleanup, not even an explicit nil
### GetDowntieringFileAge

`func (o *FileStubbingParams) GetDowntieringFileAge() DowntieringFileAgePolicy`

GetDowntieringFileAge returns the DowntieringFileAge field if non-nil, zero value otherwise.

### GetDowntieringFileAgeOk

`func (o *FileStubbingParams) GetDowntieringFileAgeOk() (*DowntieringFileAgePolicy, bool)`

GetDowntieringFileAgeOk returns a tuple with the DowntieringFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieringFileAge

`func (o *FileStubbingParams) SetDowntieringFileAge(v DowntieringFileAgePolicy)`

SetDowntieringFileAge sets DowntieringFileAge field to given value.

### HasDowntieringFileAge

`func (o *FileStubbingParams) HasDowntieringFileAge() bool`

HasDowntieringFileAge returns a boolean if a field has been set.

### GetSkipBackSymlink

`func (o *FileStubbingParams) GetSkipBackSymlink() bool`

GetSkipBackSymlink returns the SkipBackSymlink field if non-nil, zero value otherwise.

### GetSkipBackSymlinkOk

`func (o *FileStubbingParams) GetSkipBackSymlinkOk() (*bool, bool)`

GetSkipBackSymlinkOk returns a tuple with the SkipBackSymlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBackSymlink

`func (o *FileStubbingParams) SetSkipBackSymlink(v bool)`

SetSkipBackSymlink sets SkipBackSymlink field to given value.

### HasSkipBackSymlink

`func (o *FileStubbingParams) HasSkipBackSymlink() bool`

HasSkipBackSymlink returns a boolean if a field has been set.

### SetSkipBackSymlinkNil

`func (o *FileStubbingParams) SetSkipBackSymlinkNil(b bool)`

 SetSkipBackSymlinkNil sets the value for SkipBackSymlink to be an explicit nil

### UnsetSkipBackSymlink
`func (o *FileStubbingParams) UnsetSkipBackSymlink()`

UnsetSkipBackSymlink ensures that no value is present for SkipBackSymlink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


