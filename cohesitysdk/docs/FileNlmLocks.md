# FileNlmLocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to [**FileId**](FileId.md) |  | [optional] 
**NlmLocks** | Pointer to [**[]NlmLock**](NlmLock.md) | Specifies the list of NLM locks in a view. | [optional] 

## Methods

### NewFileNlmLocks

`func NewFileNlmLocks() *FileNlmLocks`

NewFileNlmLocks instantiates a new FileNlmLocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileNlmLocksWithDefaults

`func NewFileNlmLocksWithDefaults() *FileNlmLocks`

NewFileNlmLocksWithDefaults instantiates a new FileNlmLocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *FileNlmLocks) GetFileId() FileId`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileNlmLocks) GetFileIdOk() (*FileId, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileNlmLocks) SetFileId(v FileId)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileNlmLocks) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetNlmLocks

`func (o *FileNlmLocks) GetNlmLocks() []NlmLock`

GetNlmLocks returns the NlmLocks field if non-nil, zero value otherwise.

### GetNlmLocksOk

`func (o *FileNlmLocks) GetNlmLocksOk() (*[]NlmLock, bool)`

GetNlmLocksOk returns a tuple with the NlmLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNlmLocks

`func (o *FileNlmLocks) SetNlmLocks(v []NlmLock)`

SetNlmLocks sets NlmLocks field to given value.

### HasNlmLocks

`func (o *FileNlmLocks) HasNlmLocks() bool`

HasNlmLocks returns a boolean if a field has been set.

### SetNlmLocksNil

`func (o *FileNlmLocks) SetNlmLocksNil(b bool)`

 SetNlmLocksNil sets the value for NlmLocks to be an explicit nil

### UnsetNlmLocks
`func (o *FileNlmLocks) UnsetNlmLocks()`

UnsetNlmLocks ensures that no value is present for NlmLocks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


