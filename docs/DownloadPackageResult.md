# DownloadPackageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies a message describing the result of the request to download a package to a Cluster. | [optional] 

## Methods

### NewDownloadPackageResult

`func NewDownloadPackageResult() *DownloadPackageResult`

NewDownloadPackageResult instantiates a new DownloadPackageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadPackageResultWithDefaults

`func NewDownloadPackageResultWithDefaults() *DownloadPackageResult`

NewDownloadPackageResultWithDefaults instantiates a new DownloadPackageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DownloadPackageResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DownloadPackageResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DownloadPackageResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DownloadPackageResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DownloadPackageResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DownloadPackageResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


