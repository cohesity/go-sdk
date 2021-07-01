# DownloadCftResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **[]int32** | Specifies the content of the file. | [optional] 
**FileName** | Pointer to **NullableString** | Specifies the content of the CFT. in: body Specifies the file name of the cloud formation template. | [optional] 

## Methods

### NewDownloadCftResponse

`func NewDownloadCftResponse() *DownloadCftResponse`

NewDownloadCftResponse instantiates a new DownloadCftResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadCftResponseWithDefaults

`func NewDownloadCftResponseWithDefaults() *DownloadCftResponse`

NewDownloadCftResponseWithDefaults instantiates a new DownloadCftResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DownloadCftResponse) GetContent() []int32`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DownloadCftResponse) GetContentOk() (*[]int32, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DownloadCftResponse) SetContent(v []int32)`

SetContent sets Content field to given value.

### HasContent

`func (o *DownloadCftResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DownloadCftResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DownloadCftResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFileName

`func (o *DownloadCftResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DownloadCftResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DownloadCftResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DownloadCftResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *DownloadCftResponse) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *DownloadCftResponse) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


