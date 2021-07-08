# MapReduceInstanceWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogPath** | Pointer to **NullableString** | LogPath is the path of the log files for the MR instance run. | [optional] 
**MrInstance** | Pointer to [**MapReduceInstance**](MapReduceInstance.md) |  | [optional] 
**OutputFilePathList** | Pointer to **[]string** | OutputFilePathList is the list containing the output files path suffix that Yoda uses to build the full path of the MR instance run output files. | [optional] 

## Methods

### NewMapReduceInstanceWrapper

`func NewMapReduceInstanceWrapper() *MapReduceInstanceWrapper`

NewMapReduceInstanceWrapper instantiates a new MapReduceInstanceWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapReduceInstanceWrapperWithDefaults

`func NewMapReduceInstanceWrapperWithDefaults() *MapReduceInstanceWrapper`

NewMapReduceInstanceWrapperWithDefaults instantiates a new MapReduceInstanceWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogPath

`func (o *MapReduceInstanceWrapper) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *MapReduceInstanceWrapper) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *MapReduceInstanceWrapper) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *MapReduceInstanceWrapper) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### SetLogPathNil

`func (o *MapReduceInstanceWrapper) SetLogPathNil(b bool)`

 SetLogPathNil sets the value for LogPath to be an explicit nil

### UnsetLogPath
`func (o *MapReduceInstanceWrapper) UnsetLogPath()`

UnsetLogPath ensures that no value is present for LogPath, not even an explicit nil
### GetMrInstance

`func (o *MapReduceInstanceWrapper) GetMrInstance() MapReduceInstance`

GetMrInstance returns the MrInstance field if non-nil, zero value otherwise.

### GetMrInstanceOk

`func (o *MapReduceInstanceWrapper) GetMrInstanceOk() (*MapReduceInstance, bool)`

GetMrInstanceOk returns a tuple with the MrInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrInstance

`func (o *MapReduceInstanceWrapper) SetMrInstance(v MapReduceInstance)`

SetMrInstance sets MrInstance field to given value.

### HasMrInstance

`func (o *MapReduceInstanceWrapper) HasMrInstance() bool`

HasMrInstance returns a boolean if a field has been set.

### GetOutputFilePathList

`func (o *MapReduceInstanceWrapper) GetOutputFilePathList() []string`

GetOutputFilePathList returns the OutputFilePathList field if non-nil, zero value otherwise.

### GetOutputFilePathListOk

`func (o *MapReduceInstanceWrapper) GetOutputFilePathListOk() (*[]string, bool)`

GetOutputFilePathListOk returns a tuple with the OutputFilePathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFilePathList

`func (o *MapReduceInstanceWrapper) SetOutputFilePathList(v []string)`

SetOutputFilePathList sets OutputFilePathList field to given value.

### HasOutputFilePathList

`func (o *MapReduceInstanceWrapper) HasOutputFilePathList() bool`

HasOutputFilePathList returns a boolean if a field has been set.

### SetOutputFilePathListNil

`func (o *MapReduceInstanceWrapper) SetOutputFilePathListNil(b bool)`

 SetOutputFilePathListNil sets the value for OutputFilePathList to be an explicit nil

### UnsetOutputFilePathList
`func (o *MapReduceInstanceWrapper) UnsetOutputFilePathList()`

UnsetOutputFilePathList ensures that no value is present for OutputFilePathList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


