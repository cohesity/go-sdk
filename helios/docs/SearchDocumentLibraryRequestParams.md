# SearchDocumentLibraryRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | Pointer to **NullableString** | Specifies the search string to filter the files/folders. User can specify a wildcard character &#39;*&#39; as a suffix to a string where all item names are matched with the prefix string. | [optional] 
**IncludeFolders** | Pointer to **NullableBool** | Specifies whether to include folders in the response. Default is true. | [optional] [default to true]
**IncludeFiles** | Pointer to **NullableBool** | Specifies whether to include files in the response. Default is true. | [optional] [default to true]
**CategoryTypes** | Pointer to **[]string** | Specifies a list of document library types. Only items within the given types will be returned. | [optional] 
**CreationStartTimeSecs** | Pointer to **NullableInt64** | Specifies the start time in Unix timestamp epoch in seconds when the file/folder is created. | [optional] 
**CreationEndTimeSecs** | Pointer to **NullableInt64** | Specifies the end time in Unix timestamp epoch in seconds when the file/folder is created. | [optional] 
**SizeBytesLowerLimit** | Pointer to **NullableInt64** | Specifies the minimum size of the file in bytes. | [optional] 
**SizeBytesUpperLimit** | Pointer to **NullableInt64** | Specifies the maximum size of the file in bytes. | [optional] 
**OwnerNames** | Pointer to **[]string** | Specifies the list of owner names to filter on owner of the file/folder. | [optional] 
**O365Params** | Pointer to [**O365SearchSharepointRequestParams**](O365SearchSharepointRequestParams.md) |  | [optional] 

## Methods

### NewSearchDocumentLibraryRequestParams

`func NewSearchDocumentLibraryRequestParams() *SearchDocumentLibraryRequestParams`

NewSearchDocumentLibraryRequestParams instantiates a new SearchDocumentLibraryRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDocumentLibraryRequestParamsWithDefaults

`func NewSearchDocumentLibraryRequestParamsWithDefaults() *SearchDocumentLibraryRequestParams`

NewSearchDocumentLibraryRequestParamsWithDefaults instantiates a new SearchDocumentLibraryRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *SearchDocumentLibraryRequestParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *SearchDocumentLibraryRequestParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *SearchDocumentLibraryRequestParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.

### HasSearchString

`func (o *SearchDocumentLibraryRequestParams) HasSearchString() bool`

HasSearchString returns a boolean if a field has been set.

### SetSearchStringNil

`func (o *SearchDocumentLibraryRequestParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *SearchDocumentLibraryRequestParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetIncludeFolders

`func (o *SearchDocumentLibraryRequestParams) GetIncludeFolders() bool`

GetIncludeFolders returns the IncludeFolders field if non-nil, zero value otherwise.

### GetIncludeFoldersOk

`func (o *SearchDocumentLibraryRequestParams) GetIncludeFoldersOk() (*bool, bool)`

GetIncludeFoldersOk returns a tuple with the IncludeFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFolders

`func (o *SearchDocumentLibraryRequestParams) SetIncludeFolders(v bool)`

SetIncludeFolders sets IncludeFolders field to given value.

### HasIncludeFolders

`func (o *SearchDocumentLibraryRequestParams) HasIncludeFolders() bool`

HasIncludeFolders returns a boolean if a field has been set.

### SetIncludeFoldersNil

`func (o *SearchDocumentLibraryRequestParams) SetIncludeFoldersNil(b bool)`

 SetIncludeFoldersNil sets the value for IncludeFolders to be an explicit nil

### UnsetIncludeFolders
`func (o *SearchDocumentLibraryRequestParams) UnsetIncludeFolders()`

UnsetIncludeFolders ensures that no value is present for IncludeFolders, not even an explicit nil
### GetIncludeFiles

`func (o *SearchDocumentLibraryRequestParams) GetIncludeFiles() bool`

GetIncludeFiles returns the IncludeFiles field if non-nil, zero value otherwise.

### GetIncludeFilesOk

`func (o *SearchDocumentLibraryRequestParams) GetIncludeFilesOk() (*bool, bool)`

GetIncludeFilesOk returns a tuple with the IncludeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFiles

`func (o *SearchDocumentLibraryRequestParams) SetIncludeFiles(v bool)`

SetIncludeFiles sets IncludeFiles field to given value.

### HasIncludeFiles

`func (o *SearchDocumentLibraryRequestParams) HasIncludeFiles() bool`

HasIncludeFiles returns a boolean if a field has been set.

### SetIncludeFilesNil

`func (o *SearchDocumentLibraryRequestParams) SetIncludeFilesNil(b bool)`

 SetIncludeFilesNil sets the value for IncludeFiles to be an explicit nil

### UnsetIncludeFiles
`func (o *SearchDocumentLibraryRequestParams) UnsetIncludeFiles()`

UnsetIncludeFiles ensures that no value is present for IncludeFiles, not even an explicit nil
### GetCategoryTypes

`func (o *SearchDocumentLibraryRequestParams) GetCategoryTypes() []string`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *SearchDocumentLibraryRequestParams) GetCategoryTypesOk() (*[]string, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *SearchDocumentLibraryRequestParams) SetCategoryTypes(v []string)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *SearchDocumentLibraryRequestParams) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### SetCategoryTypesNil

`func (o *SearchDocumentLibraryRequestParams) SetCategoryTypesNil(b bool)`

 SetCategoryTypesNil sets the value for CategoryTypes to be an explicit nil

### UnsetCategoryTypes
`func (o *SearchDocumentLibraryRequestParams) UnsetCategoryTypes()`

UnsetCategoryTypes ensures that no value is present for CategoryTypes, not even an explicit nil
### GetCreationStartTimeSecs

`func (o *SearchDocumentLibraryRequestParams) GetCreationStartTimeSecs() int64`

GetCreationStartTimeSecs returns the CreationStartTimeSecs field if non-nil, zero value otherwise.

### GetCreationStartTimeSecsOk

`func (o *SearchDocumentLibraryRequestParams) GetCreationStartTimeSecsOk() (*int64, bool)`

GetCreationStartTimeSecsOk returns a tuple with the CreationStartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationStartTimeSecs

`func (o *SearchDocumentLibraryRequestParams) SetCreationStartTimeSecs(v int64)`

SetCreationStartTimeSecs sets CreationStartTimeSecs field to given value.

### HasCreationStartTimeSecs

`func (o *SearchDocumentLibraryRequestParams) HasCreationStartTimeSecs() bool`

HasCreationStartTimeSecs returns a boolean if a field has been set.

### SetCreationStartTimeSecsNil

`func (o *SearchDocumentLibraryRequestParams) SetCreationStartTimeSecsNil(b bool)`

 SetCreationStartTimeSecsNil sets the value for CreationStartTimeSecs to be an explicit nil

### UnsetCreationStartTimeSecs
`func (o *SearchDocumentLibraryRequestParams) UnsetCreationStartTimeSecs()`

UnsetCreationStartTimeSecs ensures that no value is present for CreationStartTimeSecs, not even an explicit nil
### GetCreationEndTimeSecs

`func (o *SearchDocumentLibraryRequestParams) GetCreationEndTimeSecs() int64`

GetCreationEndTimeSecs returns the CreationEndTimeSecs field if non-nil, zero value otherwise.

### GetCreationEndTimeSecsOk

`func (o *SearchDocumentLibraryRequestParams) GetCreationEndTimeSecsOk() (*int64, bool)`

GetCreationEndTimeSecsOk returns a tuple with the CreationEndTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationEndTimeSecs

`func (o *SearchDocumentLibraryRequestParams) SetCreationEndTimeSecs(v int64)`

SetCreationEndTimeSecs sets CreationEndTimeSecs field to given value.

### HasCreationEndTimeSecs

`func (o *SearchDocumentLibraryRequestParams) HasCreationEndTimeSecs() bool`

HasCreationEndTimeSecs returns a boolean if a field has been set.

### SetCreationEndTimeSecsNil

`func (o *SearchDocumentLibraryRequestParams) SetCreationEndTimeSecsNil(b bool)`

 SetCreationEndTimeSecsNil sets the value for CreationEndTimeSecs to be an explicit nil

### UnsetCreationEndTimeSecs
`func (o *SearchDocumentLibraryRequestParams) UnsetCreationEndTimeSecs()`

UnsetCreationEndTimeSecs ensures that no value is present for CreationEndTimeSecs, not even an explicit nil
### GetSizeBytesLowerLimit

`func (o *SearchDocumentLibraryRequestParams) GetSizeBytesLowerLimit() int64`

GetSizeBytesLowerLimit returns the SizeBytesLowerLimit field if non-nil, zero value otherwise.

### GetSizeBytesLowerLimitOk

`func (o *SearchDocumentLibraryRequestParams) GetSizeBytesLowerLimitOk() (*int64, bool)`

GetSizeBytesLowerLimitOk returns a tuple with the SizeBytesLowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytesLowerLimit

`func (o *SearchDocumentLibraryRequestParams) SetSizeBytesLowerLimit(v int64)`

SetSizeBytesLowerLimit sets SizeBytesLowerLimit field to given value.

### HasSizeBytesLowerLimit

`func (o *SearchDocumentLibraryRequestParams) HasSizeBytesLowerLimit() bool`

HasSizeBytesLowerLimit returns a boolean if a field has been set.

### SetSizeBytesLowerLimitNil

`func (o *SearchDocumentLibraryRequestParams) SetSizeBytesLowerLimitNil(b bool)`

 SetSizeBytesLowerLimitNil sets the value for SizeBytesLowerLimit to be an explicit nil

### UnsetSizeBytesLowerLimit
`func (o *SearchDocumentLibraryRequestParams) UnsetSizeBytesLowerLimit()`

UnsetSizeBytesLowerLimit ensures that no value is present for SizeBytesLowerLimit, not even an explicit nil
### GetSizeBytesUpperLimit

`func (o *SearchDocumentLibraryRequestParams) GetSizeBytesUpperLimit() int64`

GetSizeBytesUpperLimit returns the SizeBytesUpperLimit field if non-nil, zero value otherwise.

### GetSizeBytesUpperLimitOk

`func (o *SearchDocumentLibraryRequestParams) GetSizeBytesUpperLimitOk() (*int64, bool)`

GetSizeBytesUpperLimitOk returns a tuple with the SizeBytesUpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytesUpperLimit

`func (o *SearchDocumentLibraryRequestParams) SetSizeBytesUpperLimit(v int64)`

SetSizeBytesUpperLimit sets SizeBytesUpperLimit field to given value.

### HasSizeBytesUpperLimit

`func (o *SearchDocumentLibraryRequestParams) HasSizeBytesUpperLimit() bool`

HasSizeBytesUpperLimit returns a boolean if a field has been set.

### SetSizeBytesUpperLimitNil

`func (o *SearchDocumentLibraryRequestParams) SetSizeBytesUpperLimitNil(b bool)`

 SetSizeBytesUpperLimitNil sets the value for SizeBytesUpperLimit to be an explicit nil

### UnsetSizeBytesUpperLimit
`func (o *SearchDocumentLibraryRequestParams) UnsetSizeBytesUpperLimit()`

UnsetSizeBytesUpperLimit ensures that no value is present for SizeBytesUpperLimit, not even an explicit nil
### GetOwnerNames

`func (o *SearchDocumentLibraryRequestParams) GetOwnerNames() []string`

GetOwnerNames returns the OwnerNames field if non-nil, zero value otherwise.

### GetOwnerNamesOk

`func (o *SearchDocumentLibraryRequestParams) GetOwnerNamesOk() (*[]string, bool)`

GetOwnerNamesOk returns a tuple with the OwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNames

`func (o *SearchDocumentLibraryRequestParams) SetOwnerNames(v []string)`

SetOwnerNames sets OwnerNames field to given value.

### HasOwnerNames

`func (o *SearchDocumentLibraryRequestParams) HasOwnerNames() bool`

HasOwnerNames returns a boolean if a field has been set.

### SetOwnerNamesNil

`func (o *SearchDocumentLibraryRequestParams) SetOwnerNamesNil(b bool)`

 SetOwnerNamesNil sets the value for OwnerNames to be an explicit nil

### UnsetOwnerNames
`func (o *SearchDocumentLibraryRequestParams) UnsetOwnerNames()`

UnsetOwnerNames ensures that no value is present for OwnerNames, not even an explicit nil
### GetO365Params

`func (o *SearchDocumentLibraryRequestParams) GetO365Params() O365SearchSharepointRequestParams`

GetO365Params returns the O365Params field if non-nil, zero value otherwise.

### GetO365ParamsOk

`func (o *SearchDocumentLibraryRequestParams) GetO365ParamsOk() (*O365SearchSharepointRequestParams, bool)`

GetO365ParamsOk returns a tuple with the O365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO365Params

`func (o *SearchDocumentLibraryRequestParams) SetO365Params(v O365SearchSharepointRequestParams)`

SetO365Params sets O365Params field to given value.

### HasO365Params

`func (o *SearchDocumentLibraryRequestParams) HasO365Params() bool`

HasO365Params returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


