# Office365SharePointProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludePaths** | Pointer to **[]string** | Array of paths to be excluded from backup. Specifies list of doclib/directory paths which should be excluded when backing up Office 365 source. supported exclusion: - doclib exclusion: whole doclib is excluded from backup. sample: /Doclib1 - directory exclusion: specified path in doclib will be excluded from backup. sample: /Doclib1/folderA/forderB Doclibs can be specified by either a) Doclib name - eg, Documents. b) Drive id of doclib - b!ZMSl2JRm0UeXLHfHR1m-iuD10p0CIV9qSa6TtgM Regular expressions are not supported. If not specified, all the doclibs within sharepoint site will be protected. | [optional] 
**PreservationHoldLibraryParams** | Pointer to [**Office365PreservationHoldLibraryParams**](Office365PreservationHoldLibraryParams.md) |  | [optional] 

## Methods

### NewOffice365SharePointProtectionGroupParams

`func NewOffice365SharePointProtectionGroupParams() *Office365SharePointProtectionGroupParams`

NewOffice365SharePointProtectionGroupParams instantiates a new Office365SharePointProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365SharePointProtectionGroupParamsWithDefaults

`func NewOffice365SharePointProtectionGroupParamsWithDefaults() *Office365SharePointProtectionGroupParams`

NewOffice365SharePointProtectionGroupParamsWithDefaults instantiates a new Office365SharePointProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludePaths

`func (o *Office365SharePointProtectionGroupParams) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *Office365SharePointProtectionGroupParams) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *Office365SharePointProtectionGroupParams) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *Office365SharePointProtectionGroupParams) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### SetExcludePathsNil

`func (o *Office365SharePointProtectionGroupParams) SetExcludePathsNil(b bool)`

 SetExcludePathsNil sets the value for ExcludePaths to be an explicit nil

### UnsetExcludePaths
`func (o *Office365SharePointProtectionGroupParams) UnsetExcludePaths()`

UnsetExcludePaths ensures that no value is present for ExcludePaths, not even an explicit nil
### GetPreservationHoldLibraryParams

`func (o *Office365SharePointProtectionGroupParams) GetPreservationHoldLibraryParams() Office365PreservationHoldLibraryParams`

GetPreservationHoldLibraryParams returns the PreservationHoldLibraryParams field if non-nil, zero value otherwise.

### GetPreservationHoldLibraryParamsOk

`func (o *Office365SharePointProtectionGroupParams) GetPreservationHoldLibraryParamsOk() (*Office365PreservationHoldLibraryParams, bool)`

GetPreservationHoldLibraryParamsOk returns a tuple with the PreservationHoldLibraryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreservationHoldLibraryParams

`func (o *Office365SharePointProtectionGroupParams) SetPreservationHoldLibraryParams(v Office365PreservationHoldLibraryParams)`

SetPreservationHoldLibraryParams sets PreservationHoldLibraryParams field to given value.

### HasPreservationHoldLibraryParams

`func (o *Office365SharePointProtectionGroupParams) HasPreservationHoldLibraryParams() bool`

HasPreservationHoldLibraryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


