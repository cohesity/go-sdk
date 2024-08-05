# PhysicalFileBackupPathParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedPaths** | Pointer to **[]string** | Specifies a set of paths nested under the include path which should be excluded from the Protection Group. | [optional] 
**IncludedPath** | **string** | Specifies a path to be included on the source. All paths under this path will be included unless they are specifically mentioned in excluded paths. | 
**SkipNestedVolumes** | Pointer to **NullableBool** | Specifies whether to skip any nested volumes (both local and network) that are mounted under include path. Applicable only for windows sources. | [optional] 

## Methods

### NewPhysicalFileBackupPathParams

`func NewPhysicalFileBackupPathParams(includedPath string, ) *PhysicalFileBackupPathParams`

NewPhysicalFileBackupPathParams instantiates a new PhysicalFileBackupPathParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalFileBackupPathParamsWithDefaults

`func NewPhysicalFileBackupPathParamsWithDefaults() *PhysicalFileBackupPathParams`

NewPhysicalFileBackupPathParamsWithDefaults instantiates a new PhysicalFileBackupPathParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedPaths

`func (o *PhysicalFileBackupPathParams) GetExcludedPaths() []string`

GetExcludedPaths returns the ExcludedPaths field if non-nil, zero value otherwise.

### GetExcludedPathsOk

`func (o *PhysicalFileBackupPathParams) GetExcludedPathsOk() (*[]string, bool)`

GetExcludedPathsOk returns a tuple with the ExcludedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPaths

`func (o *PhysicalFileBackupPathParams) SetExcludedPaths(v []string)`

SetExcludedPaths sets ExcludedPaths field to given value.

### HasExcludedPaths

`func (o *PhysicalFileBackupPathParams) HasExcludedPaths() bool`

HasExcludedPaths returns a boolean if a field has been set.

### SetExcludedPathsNil

`func (o *PhysicalFileBackupPathParams) SetExcludedPathsNil(b bool)`

 SetExcludedPathsNil sets the value for ExcludedPaths to be an explicit nil

### UnsetExcludedPaths
`func (o *PhysicalFileBackupPathParams) UnsetExcludedPaths()`

UnsetExcludedPaths ensures that no value is present for ExcludedPaths, not even an explicit nil
### GetIncludedPath

`func (o *PhysicalFileBackupPathParams) GetIncludedPath() string`

GetIncludedPath returns the IncludedPath field if non-nil, zero value otherwise.

### GetIncludedPathOk

`func (o *PhysicalFileBackupPathParams) GetIncludedPathOk() (*string, bool)`

GetIncludedPathOk returns a tuple with the IncludedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedPath

`func (o *PhysicalFileBackupPathParams) SetIncludedPath(v string)`

SetIncludedPath sets IncludedPath field to given value.


### GetSkipNestedVolumes

`func (o *PhysicalFileBackupPathParams) GetSkipNestedVolumes() bool`

GetSkipNestedVolumes returns the SkipNestedVolumes field if non-nil, zero value otherwise.

### GetSkipNestedVolumesOk

`func (o *PhysicalFileBackupPathParams) GetSkipNestedVolumesOk() (*bool, bool)`

GetSkipNestedVolumesOk returns a tuple with the SkipNestedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNestedVolumes

`func (o *PhysicalFileBackupPathParams) SetSkipNestedVolumes(v bool)`

SetSkipNestedVolumes sets SkipNestedVolumes field to given value.

### HasSkipNestedVolumes

`func (o *PhysicalFileBackupPathParams) HasSkipNestedVolumes() bool`

HasSkipNestedVolumes returns a boolean if a field has been set.

### SetSkipNestedVolumesNil

`func (o *PhysicalFileBackupPathParams) SetSkipNestedVolumesNil(b bool)`

 SetSkipNestedVolumesNil sets the value for SkipNestedVolumes to be an explicit nil

### UnsetSkipNestedVolumes
`func (o *PhysicalFileBackupPathParams) UnsetSkipNestedVolumes()`

UnsetSkipNestedVolumes ensures that no value is present for SkipNestedVolumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


