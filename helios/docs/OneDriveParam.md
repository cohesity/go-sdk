# OneDriveParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the OneDrive id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the OneDrive name. | [optional] 
**RecoverEntireDrive** | Pointer to **NullableBool** | Specifies whether to recover the whole OneDrive. This is set to false when excluding recovering specific drive items. | [optional] 
**RecoverItems** | Pointer to [**[]OneDriveItem**](OneDriveItem.md) | Specifies a list of OneDrive items to recover. | [optional] 

## Methods

### NewOneDriveParam

`func NewOneDriveParam() *OneDriveParam`

NewOneDriveParam instantiates a new OneDriveParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveParamWithDefaults

`func NewOneDriveParamWithDefaults() *OneDriveParam`

NewOneDriveParamWithDefaults instantiates a new OneDriveParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OneDriveParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneDriveParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneDriveParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OneDriveParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OneDriveParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OneDriveParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *OneDriveParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OneDriveParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OneDriveParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OneDriveParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OneDriveParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OneDriveParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRecoverEntireDrive

`func (o *OneDriveParam) GetRecoverEntireDrive() bool`

GetRecoverEntireDrive returns the RecoverEntireDrive field if non-nil, zero value otherwise.

### GetRecoverEntireDriveOk

`func (o *OneDriveParam) GetRecoverEntireDriveOk() (*bool, bool)`

GetRecoverEntireDriveOk returns a tuple with the RecoverEntireDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireDrive

`func (o *OneDriveParam) SetRecoverEntireDrive(v bool)`

SetRecoverEntireDrive sets RecoverEntireDrive field to given value.

### HasRecoverEntireDrive

`func (o *OneDriveParam) HasRecoverEntireDrive() bool`

HasRecoverEntireDrive returns a boolean if a field has been set.

### SetRecoverEntireDriveNil

`func (o *OneDriveParam) SetRecoverEntireDriveNil(b bool)`

 SetRecoverEntireDriveNil sets the value for RecoverEntireDrive to be an explicit nil

### UnsetRecoverEntireDrive
`func (o *OneDriveParam) UnsetRecoverEntireDrive()`

UnsetRecoverEntireDrive ensures that no value is present for RecoverEntireDrive, not even an explicit nil
### GetRecoverItems

`func (o *OneDriveParam) GetRecoverItems() []OneDriveItem`

GetRecoverItems returns the RecoverItems field if non-nil, zero value otherwise.

### GetRecoverItemsOk

`func (o *OneDriveParam) GetRecoverItemsOk() (*[]OneDriveItem, bool)`

GetRecoverItemsOk returns a tuple with the RecoverItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverItems

`func (o *OneDriveParam) SetRecoverItems(v []OneDriveItem)`

SetRecoverItems sets RecoverItems field to given value.

### HasRecoverItems

`func (o *OneDriveParam) HasRecoverItems() bool`

HasRecoverItems returns a boolean if a field has been set.

### SetRecoverItemsNil

`func (o *OneDriveParam) SetRecoverItemsNil(b bool)`

 SetRecoverItemsNil sets the value for RecoverItems to be an explicit nil

### UnsetRecoverItems
`func (o *OneDriveParam) UnsetRecoverItems()`

UnsetRecoverItems ensures that no value is present for RecoverItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


