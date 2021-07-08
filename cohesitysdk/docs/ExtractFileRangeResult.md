# ExtractFileRangeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]int32** | The actual data bytes. | [optional] 
**Eof** | Pointer to **NullableBool** | Will be true if start_offset &gt; file length or EOF is reached. This is an alternative to using file_length to determine when entire file is read. Used when fetching from a view. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**FileLength** | Pointer to **NullableInt64** | The total length of the file. This field would be set provided no error had occurred (indicated by error field above). | [optional] 
**StartOffset** | Pointer to **NullableInt64** | The offset from which data was read. | [optional] 

## Methods

### NewExtractFileRangeResult

`func NewExtractFileRangeResult() *ExtractFileRangeResult`

NewExtractFileRangeResult instantiates a new ExtractFileRangeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractFileRangeResultWithDefaults

`func NewExtractFileRangeResultWithDefaults() *ExtractFileRangeResult`

NewExtractFileRangeResultWithDefaults instantiates a new ExtractFileRangeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExtractFileRangeResult) GetData() []int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExtractFileRangeResult) GetDataOk() (*[]int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExtractFileRangeResult) SetData(v []int32)`

SetData sets Data field to given value.

### HasData

`func (o *ExtractFileRangeResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ExtractFileRangeResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ExtractFileRangeResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEof

`func (o *ExtractFileRangeResult) GetEof() bool`

GetEof returns the Eof field if non-nil, zero value otherwise.

### GetEofOk

`func (o *ExtractFileRangeResult) GetEofOk() (*bool, bool)`

GetEofOk returns a tuple with the Eof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEof

`func (o *ExtractFileRangeResult) SetEof(v bool)`

SetEof sets Eof field to given value.

### HasEof

`func (o *ExtractFileRangeResult) HasEof() bool`

HasEof returns a boolean if a field has been set.

### SetEofNil

`func (o *ExtractFileRangeResult) SetEofNil(b bool)`

 SetEofNil sets the value for Eof to be an explicit nil

### UnsetEof
`func (o *ExtractFileRangeResult) UnsetEof()`

UnsetEof ensures that no value is present for Eof, not even an explicit nil
### GetError

`func (o *ExtractFileRangeResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExtractFileRangeResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExtractFileRangeResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *ExtractFileRangeResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFileLength

`func (o *ExtractFileRangeResult) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ExtractFileRangeResult) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ExtractFileRangeResult) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ExtractFileRangeResult) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### SetFileLengthNil

`func (o *ExtractFileRangeResult) SetFileLengthNil(b bool)`

 SetFileLengthNil sets the value for FileLength to be an explicit nil

### UnsetFileLength
`func (o *ExtractFileRangeResult) UnsetFileLength()`

UnsetFileLength ensures that no value is present for FileLength, not even an explicit nil
### GetStartOffset

`func (o *ExtractFileRangeResult) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *ExtractFileRangeResult) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *ExtractFileRangeResult) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *ExtractFileRangeResult) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### SetStartOffsetNil

`func (o *ExtractFileRangeResult) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *ExtractFileRangeResult) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


