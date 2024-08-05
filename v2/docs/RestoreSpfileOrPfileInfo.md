# RestoreSpfileOrPfileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to **NullableString** | Specifies the location where spfile/file will be restored. If this is empty and shouldRestoreSpfileOrPfile is true we restore at default location: $ORACLE_HOME/dbs | [optional] 
**ShouldRestoreSpfileOrPfile** | Pointer to **NullableBool** | Specifies whether to restore spfile/pfile or skip it. | [optional] 

## Methods

### NewRestoreSpfileOrPfileInfo

`func NewRestoreSpfileOrPfileInfo() *RestoreSpfileOrPfileInfo`

NewRestoreSpfileOrPfileInfo instantiates a new RestoreSpfileOrPfileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSpfileOrPfileInfoWithDefaults

`func NewRestoreSpfileOrPfileInfoWithDefaults() *RestoreSpfileOrPfileInfo`

NewRestoreSpfileOrPfileInfoWithDefaults instantiates a new RestoreSpfileOrPfileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *RestoreSpfileOrPfileInfo) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *RestoreSpfileOrPfileInfo) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *RestoreSpfileOrPfileInfo) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *RestoreSpfileOrPfileInfo) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### SetFileLocationNil

`func (o *RestoreSpfileOrPfileInfo) SetFileLocationNil(b bool)`

 SetFileLocationNil sets the value for FileLocation to be an explicit nil

### UnsetFileLocation
`func (o *RestoreSpfileOrPfileInfo) UnsetFileLocation()`

UnsetFileLocation ensures that no value is present for FileLocation, not even an explicit nil
### GetShouldRestoreSpfileOrPfile

`func (o *RestoreSpfileOrPfileInfo) GetShouldRestoreSpfileOrPfile() bool`

GetShouldRestoreSpfileOrPfile returns the ShouldRestoreSpfileOrPfile field if non-nil, zero value otherwise.

### GetShouldRestoreSpfileOrPfileOk

`func (o *RestoreSpfileOrPfileInfo) GetShouldRestoreSpfileOrPfileOk() (*bool, bool)`

GetShouldRestoreSpfileOrPfileOk returns a tuple with the ShouldRestoreSpfileOrPfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRestoreSpfileOrPfile

`func (o *RestoreSpfileOrPfileInfo) SetShouldRestoreSpfileOrPfile(v bool)`

SetShouldRestoreSpfileOrPfile sets ShouldRestoreSpfileOrPfile field to given value.

### HasShouldRestoreSpfileOrPfile

`func (o *RestoreSpfileOrPfileInfo) HasShouldRestoreSpfileOrPfile() bool`

HasShouldRestoreSpfileOrPfile returns a boolean if a field has been set.

### SetShouldRestoreSpfileOrPfileNil

`func (o *RestoreSpfileOrPfileInfo) SetShouldRestoreSpfileOrPfileNil(b bool)`

 SetShouldRestoreSpfileOrPfileNil sets the value for ShouldRestoreSpfileOrPfile to be an explicit nil

### UnsetShouldRestoreSpfileOrPfile
`func (o *RestoreSpfileOrPfileInfo) UnsetShouldRestoreSpfileOrPfile()`

UnsetShouldRestoreSpfileOrPfile ensures that no value is present for ShouldRestoreSpfileOrPfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


