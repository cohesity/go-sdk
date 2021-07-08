# SharePointRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreToOriginalSite** | Pointer to **NullableBool** | Specifies whether the objects are to be restored to the original drive. | [optional] 
**SiteOwnerList** | Pointer to [**[]SiteOwner**](SiteOwner.md) | Specifies the list of SharePoint Sites whose Document Repositories are being restored. | [optional] 
**TargetDocumentLibraryName** | Pointer to **NullableString** | Specifies the target document library name within the alternate site. | [optional] 
**TargetDocumentLibraryPrefix** | Pointer to **NullableString** | Specifies a custom prefix for the document libraries when being restored to the original or an alternate site. | [optional] 
**TargetSite** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 

## Methods

### NewSharePointRestoreParameters

`func NewSharePointRestoreParameters() *SharePointRestoreParameters`

NewSharePointRestoreParameters instantiates a new SharePointRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePointRestoreParametersWithDefaults

`func NewSharePointRestoreParametersWithDefaults() *SharePointRestoreParameters`

NewSharePointRestoreParametersWithDefaults instantiates a new SharePointRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreToOriginalSite

`func (o *SharePointRestoreParameters) GetRestoreToOriginalSite() bool`

GetRestoreToOriginalSite returns the RestoreToOriginalSite field if non-nil, zero value otherwise.

### GetRestoreToOriginalSiteOk

`func (o *SharePointRestoreParameters) GetRestoreToOriginalSiteOk() (*bool, bool)`

GetRestoreToOriginalSiteOk returns a tuple with the RestoreToOriginalSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginalSite

`func (o *SharePointRestoreParameters) SetRestoreToOriginalSite(v bool)`

SetRestoreToOriginalSite sets RestoreToOriginalSite field to given value.

### HasRestoreToOriginalSite

`func (o *SharePointRestoreParameters) HasRestoreToOriginalSite() bool`

HasRestoreToOriginalSite returns a boolean if a field has been set.

### SetRestoreToOriginalSiteNil

`func (o *SharePointRestoreParameters) SetRestoreToOriginalSiteNil(b bool)`

 SetRestoreToOriginalSiteNil sets the value for RestoreToOriginalSite to be an explicit nil

### UnsetRestoreToOriginalSite
`func (o *SharePointRestoreParameters) UnsetRestoreToOriginalSite()`

UnsetRestoreToOriginalSite ensures that no value is present for RestoreToOriginalSite, not even an explicit nil
### GetSiteOwnerList

`func (o *SharePointRestoreParameters) GetSiteOwnerList() []SiteOwner`

GetSiteOwnerList returns the SiteOwnerList field if non-nil, zero value otherwise.

### GetSiteOwnerListOk

`func (o *SharePointRestoreParameters) GetSiteOwnerListOk() (*[]SiteOwner, bool)`

GetSiteOwnerListOk returns a tuple with the SiteOwnerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwnerList

`func (o *SharePointRestoreParameters) SetSiteOwnerList(v []SiteOwner)`

SetSiteOwnerList sets SiteOwnerList field to given value.

### HasSiteOwnerList

`func (o *SharePointRestoreParameters) HasSiteOwnerList() bool`

HasSiteOwnerList returns a boolean if a field has been set.

### SetSiteOwnerListNil

`func (o *SharePointRestoreParameters) SetSiteOwnerListNil(b bool)`

 SetSiteOwnerListNil sets the value for SiteOwnerList to be an explicit nil

### UnsetSiteOwnerList
`func (o *SharePointRestoreParameters) UnsetSiteOwnerList()`

UnsetSiteOwnerList ensures that no value is present for SiteOwnerList, not even an explicit nil
### GetTargetDocumentLibraryName

`func (o *SharePointRestoreParameters) GetTargetDocumentLibraryName() string`

GetTargetDocumentLibraryName returns the TargetDocumentLibraryName field if non-nil, zero value otherwise.

### GetTargetDocumentLibraryNameOk

`func (o *SharePointRestoreParameters) GetTargetDocumentLibraryNameOk() (*string, bool)`

GetTargetDocumentLibraryNameOk returns a tuple with the TargetDocumentLibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocumentLibraryName

`func (o *SharePointRestoreParameters) SetTargetDocumentLibraryName(v string)`

SetTargetDocumentLibraryName sets TargetDocumentLibraryName field to given value.

### HasTargetDocumentLibraryName

`func (o *SharePointRestoreParameters) HasTargetDocumentLibraryName() bool`

HasTargetDocumentLibraryName returns a boolean if a field has been set.

### SetTargetDocumentLibraryNameNil

`func (o *SharePointRestoreParameters) SetTargetDocumentLibraryNameNil(b bool)`

 SetTargetDocumentLibraryNameNil sets the value for TargetDocumentLibraryName to be an explicit nil

### UnsetTargetDocumentLibraryName
`func (o *SharePointRestoreParameters) UnsetTargetDocumentLibraryName()`

UnsetTargetDocumentLibraryName ensures that no value is present for TargetDocumentLibraryName, not even an explicit nil
### GetTargetDocumentLibraryPrefix

`func (o *SharePointRestoreParameters) GetTargetDocumentLibraryPrefix() string`

GetTargetDocumentLibraryPrefix returns the TargetDocumentLibraryPrefix field if non-nil, zero value otherwise.

### GetTargetDocumentLibraryPrefixOk

`func (o *SharePointRestoreParameters) GetTargetDocumentLibraryPrefixOk() (*string, bool)`

GetTargetDocumentLibraryPrefixOk returns a tuple with the TargetDocumentLibraryPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocumentLibraryPrefix

`func (o *SharePointRestoreParameters) SetTargetDocumentLibraryPrefix(v string)`

SetTargetDocumentLibraryPrefix sets TargetDocumentLibraryPrefix field to given value.

### HasTargetDocumentLibraryPrefix

`func (o *SharePointRestoreParameters) HasTargetDocumentLibraryPrefix() bool`

HasTargetDocumentLibraryPrefix returns a boolean if a field has been set.

### SetTargetDocumentLibraryPrefixNil

`func (o *SharePointRestoreParameters) SetTargetDocumentLibraryPrefixNil(b bool)`

 SetTargetDocumentLibraryPrefixNil sets the value for TargetDocumentLibraryPrefix to be an explicit nil

### UnsetTargetDocumentLibraryPrefix
`func (o *SharePointRestoreParameters) UnsetTargetDocumentLibraryPrefix()`

UnsetTargetDocumentLibraryPrefix ensures that no value is present for TargetDocumentLibraryPrefix, not even an explicit nil
### GetTargetSite

`func (o *SharePointRestoreParameters) GetTargetSite() ProtectionSource`

GetTargetSite returns the TargetSite field if non-nil, zero value otherwise.

### GetTargetSiteOk

`func (o *SharePointRestoreParameters) GetTargetSiteOk() (*ProtectionSource, bool)`

GetTargetSiteOk returns a tuple with the TargetSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSite

`func (o *SharePointRestoreParameters) SetTargetSite(v ProtectionSource)`

SetTargetSite sets TargetSite field to given value.

### HasTargetSite

`func (o *SharePointRestoreParameters) HasTargetSite() bool`

HasTargetSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


