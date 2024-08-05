# HeliosTagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** | Specifies category of tag applied to the object. | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of tag applied to the object. | [optional] 
**SubCategory** | Pointer to **NullableString** | Specifies subCategory of tag applied to the object. | [optional] 
**ThirdPartyName** | Pointer to **NullableString** | Specifies thirdPartyName of tag applied to the object. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type (ex custom, thirdparty, system) of tag applied to the object. | [optional] 
**UiColor** | Pointer to **NullableString** | Specifies the color of tag applied to the object. | [optional] 
**UpdatedTimeUsecs** | Pointer to **NullableInt64** | Specifies update time of tag applied to the object. | [optional] 
**Uuid** | **NullableString** | Specifies Uuid of tag applied to the object. | 

## Methods

### NewHeliosTagInfo

`func NewHeliosTagInfo(uuid NullableString, ) *HeliosTagInfo`

NewHeliosTagInfo instantiates a new HeliosTagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTagInfoWithDefaults

`func NewHeliosTagInfoWithDefaults() *HeliosTagInfo`

NewHeliosTagInfoWithDefaults instantiates a new HeliosTagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *HeliosTagInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HeliosTagInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HeliosTagInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HeliosTagInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *HeliosTagInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *HeliosTagInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *HeliosTagInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosTagInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosTagInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosTagInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosTagInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosTagInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubCategory

`func (o *HeliosTagInfo) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *HeliosTagInfo) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *HeliosTagInfo) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *HeliosTagInfo) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### SetSubCategoryNil

`func (o *HeliosTagInfo) SetSubCategoryNil(b bool)`

 SetSubCategoryNil sets the value for SubCategory to be an explicit nil

### UnsetSubCategory
`func (o *HeliosTagInfo) UnsetSubCategory()`

UnsetSubCategory ensures that no value is present for SubCategory, not even an explicit nil
### GetThirdPartyName

`func (o *HeliosTagInfo) GetThirdPartyName() string`

GetThirdPartyName returns the ThirdPartyName field if non-nil, zero value otherwise.

### GetThirdPartyNameOk

`func (o *HeliosTagInfo) GetThirdPartyNameOk() (*string, bool)`

GetThirdPartyNameOk returns a tuple with the ThirdPartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyName

`func (o *HeliosTagInfo) SetThirdPartyName(v string)`

SetThirdPartyName sets ThirdPartyName field to given value.

### HasThirdPartyName

`func (o *HeliosTagInfo) HasThirdPartyName() bool`

HasThirdPartyName returns a boolean if a field has been set.

### SetThirdPartyNameNil

`func (o *HeliosTagInfo) SetThirdPartyNameNil(b bool)`

 SetThirdPartyNameNil sets the value for ThirdPartyName to be an explicit nil

### UnsetThirdPartyName
`func (o *HeliosTagInfo) UnsetThirdPartyName()`

UnsetThirdPartyName ensures that no value is present for ThirdPartyName, not even an explicit nil
### GetType

`func (o *HeliosTagInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosTagInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosTagInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HeliosTagInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HeliosTagInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosTagInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUiColor

`func (o *HeliosTagInfo) GetUiColor() string`

GetUiColor returns the UiColor field if non-nil, zero value otherwise.

### GetUiColorOk

`func (o *HeliosTagInfo) GetUiColorOk() (*string, bool)`

GetUiColorOk returns a tuple with the UiColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiColor

`func (o *HeliosTagInfo) SetUiColor(v string)`

SetUiColor sets UiColor field to given value.

### HasUiColor

`func (o *HeliosTagInfo) HasUiColor() bool`

HasUiColor returns a boolean if a field has been set.

### SetUiColorNil

`func (o *HeliosTagInfo) SetUiColorNil(b bool)`

 SetUiColorNil sets the value for UiColor to be an explicit nil

### UnsetUiColor
`func (o *HeliosTagInfo) UnsetUiColor()`

UnsetUiColor ensures that no value is present for UiColor, not even an explicit nil
### GetUpdatedTimeUsecs

`func (o *HeliosTagInfo) GetUpdatedTimeUsecs() int64`

GetUpdatedTimeUsecs returns the UpdatedTimeUsecs field if non-nil, zero value otherwise.

### GetUpdatedTimeUsecsOk

`func (o *HeliosTagInfo) GetUpdatedTimeUsecsOk() (*int64, bool)`

GetUpdatedTimeUsecsOk returns a tuple with the UpdatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimeUsecs

`func (o *HeliosTagInfo) SetUpdatedTimeUsecs(v int64)`

SetUpdatedTimeUsecs sets UpdatedTimeUsecs field to given value.

### HasUpdatedTimeUsecs

`func (o *HeliosTagInfo) HasUpdatedTimeUsecs() bool`

HasUpdatedTimeUsecs returns a boolean if a field has been set.

### SetUpdatedTimeUsecsNil

`func (o *HeliosTagInfo) SetUpdatedTimeUsecsNil(b bool)`

 SetUpdatedTimeUsecsNil sets the value for UpdatedTimeUsecs to be an explicit nil

### UnsetUpdatedTimeUsecs
`func (o *HeliosTagInfo) UnsetUpdatedTimeUsecs()`

UnsetUpdatedTimeUsecs ensures that no value is present for UpdatedTimeUsecs, not even an explicit nil
### GetUuid

`func (o *HeliosTagInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HeliosTagInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HeliosTagInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### SetUuidNil

`func (o *HeliosTagInfo) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *HeliosTagInfo) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


