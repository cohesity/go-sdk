# HdfsBrowseRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentPath** | Pointer to **NullableString** | Specifies the path whose contents are to be returned. The last token in the path can be a regex. In this case the regex is applied on the contents of the path upto the second-last token and the matching contents are returned. | [optional] 

## Methods

### NewHdfsBrowseRequestParams

`func NewHdfsBrowseRequestParams() *HdfsBrowseRequestParams`

NewHdfsBrowseRequestParams instantiates a new HdfsBrowseRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsBrowseRequestParamsWithDefaults

`func NewHdfsBrowseRequestParamsWithDefaults() *HdfsBrowseRequestParams`

NewHdfsBrowseRequestParamsWithDefaults instantiates a new HdfsBrowseRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentPath

`func (o *HdfsBrowseRequestParams) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *HdfsBrowseRequestParams) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *HdfsBrowseRequestParams) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *HdfsBrowseRequestParams) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.

### SetParentPathNil

`func (o *HdfsBrowseRequestParams) SetParentPathNil(b bool)`

 SetParentPathNil sets the value for ParentPath to be an explicit nil

### UnsetParentPath
`func (o *HdfsBrowseRequestParams) UnsetParentPath()`

UnsetParentPath ensures that no value is present for ParentPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


