# PatchUnpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestPath** | **NullableString** | Specifies the destination path. | 
**SourcePath** | **NullableString** | Specifies the source path. | 

## Methods

### NewPatchUnpackRequest

`func NewPatchUnpackRequest(destPath NullableString, sourcePath NullableString, ) *PatchUnpackRequest`

NewPatchUnpackRequest instantiates a new PatchUnpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUnpackRequestWithDefaults

`func NewPatchUnpackRequestWithDefaults() *PatchUnpackRequest`

NewPatchUnpackRequestWithDefaults instantiates a new PatchUnpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestPath

`func (o *PatchUnpackRequest) GetDestPath() string`

GetDestPath returns the DestPath field if non-nil, zero value otherwise.

### GetDestPathOk

`func (o *PatchUnpackRequest) GetDestPathOk() (*string, bool)`

GetDestPathOk returns a tuple with the DestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPath

`func (o *PatchUnpackRequest) SetDestPath(v string)`

SetDestPath sets DestPath field to given value.


### SetDestPathNil

`func (o *PatchUnpackRequest) SetDestPathNil(b bool)`

 SetDestPathNil sets the value for DestPath to be an explicit nil

### UnsetDestPath
`func (o *PatchUnpackRequest) UnsetDestPath()`

UnsetDestPath ensures that no value is present for DestPath, not even an explicit nil
### GetSourcePath

`func (o *PatchUnpackRequest) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *PatchUnpackRequest) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *PatchUnpackRequest) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.


### SetSourcePathNil

`func (o *PatchUnpackRequest) SetSourcePathNil(b bool)`

 SetSourcePathNil sets the value for SourcePath to be an explicit nil

### UnsetSourcePath
`func (o *PatchUnpackRequest) UnsetSourcePath()`

UnsetSourcePath ensures that no value is present for SourcePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


