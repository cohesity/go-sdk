# RestoreSiteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreToOriginal** | Pointer to **NullableBool** | Whether or not all sites are restored to original location. | [optional] 
**SiteOwnerVec** | Pointer to [**[]RestoreSiteParamsSiteOwner**](RestoreSiteParamsSiteOwner.md) | The list of sites whose drives are being restored. | [optional] 
**TargetDocLibName** | Pointer to **NullableString** | Incase of alternate restore of granular items within document repositiories of sites to another site, a doc lib name has to be specified by the caller. NOTE: It can be safely assumed that this field will only be present in case of granular items restore only. | [optional] 
**TargetDocLibPrefix** | Pointer to **NullableString** | If alternate site is provided, customer may want to provide a custom prefix to document libraries that we create. In any case we would also have to distinguish the newly created document library as the alternate site provided by the customer may as well turn out to be the original backup site. | [optional] 
**TargetSite** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreSiteParams

`func NewRestoreSiteParams() *RestoreSiteParams`

NewRestoreSiteParams instantiates a new RestoreSiteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSiteParamsWithDefaults

`func NewRestoreSiteParamsWithDefaults() *RestoreSiteParams`

NewRestoreSiteParamsWithDefaults instantiates a new RestoreSiteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreToOriginal

`func (o *RestoreSiteParams) GetRestoreToOriginal() bool`

GetRestoreToOriginal returns the RestoreToOriginal field if non-nil, zero value otherwise.

### GetRestoreToOriginalOk

`func (o *RestoreSiteParams) GetRestoreToOriginalOk() (*bool, bool)`

GetRestoreToOriginalOk returns a tuple with the RestoreToOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginal

`func (o *RestoreSiteParams) SetRestoreToOriginal(v bool)`

SetRestoreToOriginal sets RestoreToOriginal field to given value.

### HasRestoreToOriginal

`func (o *RestoreSiteParams) HasRestoreToOriginal() bool`

HasRestoreToOriginal returns a boolean if a field has been set.

### SetRestoreToOriginalNil

`func (o *RestoreSiteParams) SetRestoreToOriginalNil(b bool)`

 SetRestoreToOriginalNil sets the value for RestoreToOriginal to be an explicit nil

### UnsetRestoreToOriginal
`func (o *RestoreSiteParams) UnsetRestoreToOriginal()`

UnsetRestoreToOriginal ensures that no value is present for RestoreToOriginal, not even an explicit nil
### GetSiteOwnerVec

`func (o *RestoreSiteParams) GetSiteOwnerVec() []RestoreSiteParamsSiteOwner`

GetSiteOwnerVec returns the SiteOwnerVec field if non-nil, zero value otherwise.

### GetSiteOwnerVecOk

`func (o *RestoreSiteParams) GetSiteOwnerVecOk() (*[]RestoreSiteParamsSiteOwner, bool)`

GetSiteOwnerVecOk returns a tuple with the SiteOwnerVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteOwnerVec

`func (o *RestoreSiteParams) SetSiteOwnerVec(v []RestoreSiteParamsSiteOwner)`

SetSiteOwnerVec sets SiteOwnerVec field to given value.

### HasSiteOwnerVec

`func (o *RestoreSiteParams) HasSiteOwnerVec() bool`

HasSiteOwnerVec returns a boolean if a field has been set.

### SetSiteOwnerVecNil

`func (o *RestoreSiteParams) SetSiteOwnerVecNil(b bool)`

 SetSiteOwnerVecNil sets the value for SiteOwnerVec to be an explicit nil

### UnsetSiteOwnerVec
`func (o *RestoreSiteParams) UnsetSiteOwnerVec()`

UnsetSiteOwnerVec ensures that no value is present for SiteOwnerVec, not even an explicit nil
### GetTargetDocLibName

`func (o *RestoreSiteParams) GetTargetDocLibName() string`

GetTargetDocLibName returns the TargetDocLibName field if non-nil, zero value otherwise.

### GetTargetDocLibNameOk

`func (o *RestoreSiteParams) GetTargetDocLibNameOk() (*string, bool)`

GetTargetDocLibNameOk returns a tuple with the TargetDocLibName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocLibName

`func (o *RestoreSiteParams) SetTargetDocLibName(v string)`

SetTargetDocLibName sets TargetDocLibName field to given value.

### HasTargetDocLibName

`func (o *RestoreSiteParams) HasTargetDocLibName() bool`

HasTargetDocLibName returns a boolean if a field has been set.

### SetTargetDocLibNameNil

`func (o *RestoreSiteParams) SetTargetDocLibNameNil(b bool)`

 SetTargetDocLibNameNil sets the value for TargetDocLibName to be an explicit nil

### UnsetTargetDocLibName
`func (o *RestoreSiteParams) UnsetTargetDocLibName()`

UnsetTargetDocLibName ensures that no value is present for TargetDocLibName, not even an explicit nil
### GetTargetDocLibPrefix

`func (o *RestoreSiteParams) GetTargetDocLibPrefix() string`

GetTargetDocLibPrefix returns the TargetDocLibPrefix field if non-nil, zero value otherwise.

### GetTargetDocLibPrefixOk

`func (o *RestoreSiteParams) GetTargetDocLibPrefixOk() (*string, bool)`

GetTargetDocLibPrefixOk returns a tuple with the TargetDocLibPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocLibPrefix

`func (o *RestoreSiteParams) SetTargetDocLibPrefix(v string)`

SetTargetDocLibPrefix sets TargetDocLibPrefix field to given value.

### HasTargetDocLibPrefix

`func (o *RestoreSiteParams) HasTargetDocLibPrefix() bool`

HasTargetDocLibPrefix returns a boolean if a field has been set.

### SetTargetDocLibPrefixNil

`func (o *RestoreSiteParams) SetTargetDocLibPrefixNil(b bool)`

 SetTargetDocLibPrefixNil sets the value for TargetDocLibPrefix to be an explicit nil

### UnsetTargetDocLibPrefix
`func (o *RestoreSiteParams) UnsetTargetDocLibPrefix()`

UnsetTargetDocLibPrefix ensures that no value is present for TargetDocLibPrefix, not even an explicit nil
### GetTargetSite

`func (o *RestoreSiteParams) GetTargetSite() EntityProto`

GetTargetSite returns the TargetSite field if non-nil, zero value otherwise.

### GetTargetSiteOk

`func (o *RestoreSiteParams) GetTargetSiteOk() (*EntityProto, bool)`

GetTargetSiteOk returns a tuple with the TargetSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSite

`func (o *RestoreSiteParams) SetTargetSite(v EntityProto)`

SetTargetSite sets TargetSite field to given value.

### HasTargetSite

`func (o *RestoreSiteParams) HasTargetSite() bool`

HasTargetSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


