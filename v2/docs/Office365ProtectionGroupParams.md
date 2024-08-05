# Office365ProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]Office365ProtectionGroupObjectParams**](Office365ProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**OneDriveProtectionTypeParams** | Pointer to [**Office365OneDriveProtectionGroupParams**](Office365OneDriveProtectionGroupParams.md) |  | [optional] 
**OutlookProtectionTypeParams** | Pointer to [**Office365OutlookProtectionGroupParams**](Office365OutlookProtectionGroupParams.md) |  | [optional] 
**ProtectionTypes** | **[]string** | Specifies the Office 365 Protection Group types. | 
**PublicFoldersProtectionTypeParams** | Pointer to [**Office365PublicFoldersProtectionGroupParams**](Office365PublicFoldersProtectionGroupParams.md) |  | [optional] 
**SharePointProtectionTypeParams** | Pointer to [**Office365SharePointProtectionGroupParams**](Office365SharePointProtectionGroupParams.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewOffice365ProtectionGroupParams

`func NewOffice365ProtectionGroupParams(objects []Office365ProtectionGroupObjectParams, protectionTypes []string, ) *Office365ProtectionGroupParams`

NewOffice365ProtectionGroupParams instantiates a new Office365ProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ProtectionGroupParamsWithDefaults

`func NewOffice365ProtectionGroupParamsWithDefaults() *Office365ProtectionGroupParams`

NewOffice365ProtectionGroupParamsWithDefaults instantiates a new Office365ProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *Office365ProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *Office365ProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *Office365ProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *Office365ProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *Office365ProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *Office365ProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetIndexingPolicy

`func (o *Office365ProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *Office365ProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *Office365ProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *Office365ProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *Office365ProtectionGroupParams) GetObjects() []Office365ProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Office365ProtectionGroupParams) GetObjectsOk() (*[]Office365ProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Office365ProtectionGroupParams) SetObjects(v []Office365ProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetOneDriveProtectionTypeParams

`func (o *Office365ProtectionGroupParams) GetOneDriveProtectionTypeParams() Office365OneDriveProtectionGroupParams`

GetOneDriveProtectionTypeParams returns the OneDriveProtectionTypeParams field if non-nil, zero value otherwise.

### GetOneDriveProtectionTypeParamsOk

`func (o *Office365ProtectionGroupParams) GetOneDriveProtectionTypeParamsOk() (*Office365OneDriveProtectionGroupParams, bool)`

GetOneDriveProtectionTypeParamsOk returns a tuple with the OneDriveProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveProtectionTypeParams

`func (o *Office365ProtectionGroupParams) SetOneDriveProtectionTypeParams(v Office365OneDriveProtectionGroupParams)`

SetOneDriveProtectionTypeParams sets OneDriveProtectionTypeParams field to given value.

### HasOneDriveProtectionTypeParams

`func (o *Office365ProtectionGroupParams) HasOneDriveProtectionTypeParams() bool`

HasOneDriveProtectionTypeParams returns a boolean if a field has been set.

### GetOutlookProtectionTypeParams

`func (o *Office365ProtectionGroupParams) GetOutlookProtectionTypeParams() Office365OutlookProtectionGroupParams`

GetOutlookProtectionTypeParams returns the OutlookProtectionTypeParams field if non-nil, zero value otherwise.

### GetOutlookProtectionTypeParamsOk

`func (o *Office365ProtectionGroupParams) GetOutlookProtectionTypeParamsOk() (*Office365OutlookProtectionGroupParams, bool)`

GetOutlookProtectionTypeParamsOk returns a tuple with the OutlookProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookProtectionTypeParams

`func (o *Office365ProtectionGroupParams) SetOutlookProtectionTypeParams(v Office365OutlookProtectionGroupParams)`

SetOutlookProtectionTypeParams sets OutlookProtectionTypeParams field to given value.

### HasOutlookProtectionTypeParams

`func (o *Office365ProtectionGroupParams) HasOutlookProtectionTypeParams() bool`

HasOutlookProtectionTypeParams returns a boolean if a field has been set.

### GetProtectionTypes

`func (o *Office365ProtectionGroupParams) GetProtectionTypes() []string`

GetProtectionTypes returns the ProtectionTypes field if non-nil, zero value otherwise.

### GetProtectionTypesOk

`func (o *Office365ProtectionGroupParams) GetProtectionTypesOk() (*[]string, bool)`

GetProtectionTypesOk returns a tuple with the ProtectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionTypes

`func (o *Office365ProtectionGroupParams) SetProtectionTypes(v []string)`

SetProtectionTypes sets ProtectionTypes field to given value.


### GetPublicFoldersProtectionTypeParams

`func (o *Office365ProtectionGroupParams) GetPublicFoldersProtectionTypeParams() Office365PublicFoldersProtectionGroupParams`

GetPublicFoldersProtectionTypeParams returns the PublicFoldersProtectionTypeParams field if non-nil, zero value otherwise.

### GetPublicFoldersProtectionTypeParamsOk

`func (o *Office365ProtectionGroupParams) GetPublicFoldersProtectionTypeParamsOk() (*Office365PublicFoldersProtectionGroupParams, bool)`

GetPublicFoldersProtectionTypeParamsOk returns a tuple with the PublicFoldersProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFoldersProtectionTypeParams

`func (o *Office365ProtectionGroupParams) SetPublicFoldersProtectionTypeParams(v Office365PublicFoldersProtectionGroupParams)`

SetPublicFoldersProtectionTypeParams sets PublicFoldersProtectionTypeParams field to given value.

### HasPublicFoldersProtectionTypeParams

`func (o *Office365ProtectionGroupParams) HasPublicFoldersProtectionTypeParams() bool`

HasPublicFoldersProtectionTypeParams returns a boolean if a field has been set.

### GetSharePointProtectionTypeParams

`func (o *Office365ProtectionGroupParams) GetSharePointProtectionTypeParams() Office365SharePointProtectionGroupParams`

GetSharePointProtectionTypeParams returns the SharePointProtectionTypeParams field if non-nil, zero value otherwise.

### GetSharePointProtectionTypeParamsOk

`func (o *Office365ProtectionGroupParams) GetSharePointProtectionTypeParamsOk() (*Office365SharePointProtectionGroupParams, bool)`

GetSharePointProtectionTypeParamsOk returns a tuple with the SharePointProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePointProtectionTypeParams

`func (o *Office365ProtectionGroupParams) SetSharePointProtectionTypeParams(v Office365SharePointProtectionGroupParams)`

SetSharePointProtectionTypeParams sets SharePointProtectionTypeParams field to given value.

### HasSharePointProtectionTypeParams

`func (o *Office365ProtectionGroupParams) HasSharePointProtectionTypeParams() bool`

HasSharePointProtectionTypeParams returns a boolean if a field has been set.

### GetSourceId

`func (o *Office365ProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Office365ProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Office365ProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Office365ProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Office365ProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Office365ProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Office365ProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Office365ProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Office365ProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Office365ProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Office365ProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Office365ProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


