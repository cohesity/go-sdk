# MsGroupParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxRestoreParams** | Pointer to [**MsGroupParamMailboxRestoreParams**](MsGroupParamMailboxRestoreParams.md) |  | [optional] 
**MailboxRestoreType** | Pointer to **NullableString** | Specifies whether mailbox restore is full or granular. | [optional] 
**RecoverEntireGroup** | Pointer to **NullableBool** | Specifies if the entire Group (mailbox + site) is to be restored. | [optional] 
**RecoverObject** | [**MsGroupParamRecoverObject**](MsGroupParamRecoverObject.md) |  | 
**SiteRestoreParams** | Pointer to [**[]OneDriveParam**](OneDriveParam.md) | Specifies the parameters to recover a MSGroup site document. | [optional] 
**SiteRestoreType** | Pointer to **NullableString** | Specifies whether site restore is full or granular. | [optional] 

## Methods

### NewMsGroupParam

`func NewMsGroupParam(recoverObject MsGroupParamRecoverObject, ) *MsGroupParam`

NewMsGroupParam instantiates a new MsGroupParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsGroupParamWithDefaults

`func NewMsGroupParamWithDefaults() *MsGroupParam`

NewMsGroupParamWithDefaults instantiates a new MsGroupParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxRestoreParams

`func (o *MsGroupParam) GetMailboxRestoreParams() MsGroupParamMailboxRestoreParams`

GetMailboxRestoreParams returns the MailboxRestoreParams field if non-nil, zero value otherwise.

### GetMailboxRestoreParamsOk

`func (o *MsGroupParam) GetMailboxRestoreParamsOk() (*MsGroupParamMailboxRestoreParams, bool)`

GetMailboxRestoreParamsOk returns a tuple with the MailboxRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxRestoreParams

`func (o *MsGroupParam) SetMailboxRestoreParams(v MsGroupParamMailboxRestoreParams)`

SetMailboxRestoreParams sets MailboxRestoreParams field to given value.

### HasMailboxRestoreParams

`func (o *MsGroupParam) HasMailboxRestoreParams() bool`

HasMailboxRestoreParams returns a boolean if a field has been set.

### GetMailboxRestoreType

`func (o *MsGroupParam) GetMailboxRestoreType() string`

GetMailboxRestoreType returns the MailboxRestoreType field if non-nil, zero value otherwise.

### GetMailboxRestoreTypeOk

`func (o *MsGroupParam) GetMailboxRestoreTypeOk() (*string, bool)`

GetMailboxRestoreTypeOk returns a tuple with the MailboxRestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxRestoreType

`func (o *MsGroupParam) SetMailboxRestoreType(v string)`

SetMailboxRestoreType sets MailboxRestoreType field to given value.

### HasMailboxRestoreType

`func (o *MsGroupParam) HasMailboxRestoreType() bool`

HasMailboxRestoreType returns a boolean if a field has been set.

### SetMailboxRestoreTypeNil

`func (o *MsGroupParam) SetMailboxRestoreTypeNil(b bool)`

 SetMailboxRestoreTypeNil sets the value for MailboxRestoreType to be an explicit nil

### UnsetMailboxRestoreType
`func (o *MsGroupParam) UnsetMailboxRestoreType()`

UnsetMailboxRestoreType ensures that no value is present for MailboxRestoreType, not even an explicit nil
### GetRecoverEntireGroup

`func (o *MsGroupParam) GetRecoverEntireGroup() bool`

GetRecoverEntireGroup returns the RecoverEntireGroup field if non-nil, zero value otherwise.

### GetRecoverEntireGroupOk

`func (o *MsGroupParam) GetRecoverEntireGroupOk() (*bool, bool)`

GetRecoverEntireGroupOk returns a tuple with the RecoverEntireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireGroup

`func (o *MsGroupParam) SetRecoverEntireGroup(v bool)`

SetRecoverEntireGroup sets RecoverEntireGroup field to given value.

### HasRecoverEntireGroup

`func (o *MsGroupParam) HasRecoverEntireGroup() bool`

HasRecoverEntireGroup returns a boolean if a field has been set.

### SetRecoverEntireGroupNil

`func (o *MsGroupParam) SetRecoverEntireGroupNil(b bool)`

 SetRecoverEntireGroupNil sets the value for RecoverEntireGroup to be an explicit nil

### UnsetRecoverEntireGroup
`func (o *MsGroupParam) UnsetRecoverEntireGroup()`

UnsetRecoverEntireGroup ensures that no value is present for RecoverEntireGroup, not even an explicit nil
### GetRecoverObject

`func (o *MsGroupParam) GetRecoverObject() MsGroupParamRecoverObject`

GetRecoverObject returns the RecoverObject field if non-nil, zero value otherwise.

### GetRecoverObjectOk

`func (o *MsGroupParam) GetRecoverObjectOk() (*MsGroupParamRecoverObject, bool)`

GetRecoverObjectOk returns a tuple with the RecoverObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverObject

`func (o *MsGroupParam) SetRecoverObject(v MsGroupParamRecoverObject)`

SetRecoverObject sets RecoverObject field to given value.


### GetSiteRestoreParams

`func (o *MsGroupParam) GetSiteRestoreParams() []OneDriveParam`

GetSiteRestoreParams returns the SiteRestoreParams field if non-nil, zero value otherwise.

### GetSiteRestoreParamsOk

`func (o *MsGroupParam) GetSiteRestoreParamsOk() (*[]OneDriveParam, bool)`

GetSiteRestoreParamsOk returns a tuple with the SiteRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRestoreParams

`func (o *MsGroupParam) SetSiteRestoreParams(v []OneDriveParam)`

SetSiteRestoreParams sets SiteRestoreParams field to given value.

### HasSiteRestoreParams

`func (o *MsGroupParam) HasSiteRestoreParams() bool`

HasSiteRestoreParams returns a boolean if a field has been set.

### SetSiteRestoreParamsNil

`func (o *MsGroupParam) SetSiteRestoreParamsNil(b bool)`

 SetSiteRestoreParamsNil sets the value for SiteRestoreParams to be an explicit nil

### UnsetSiteRestoreParams
`func (o *MsGroupParam) UnsetSiteRestoreParams()`

UnsetSiteRestoreParams ensures that no value is present for SiteRestoreParams, not even an explicit nil
### GetSiteRestoreType

`func (o *MsGroupParam) GetSiteRestoreType() string`

GetSiteRestoreType returns the SiteRestoreType field if non-nil, zero value otherwise.

### GetSiteRestoreTypeOk

`func (o *MsGroupParam) GetSiteRestoreTypeOk() (*string, bool)`

GetSiteRestoreTypeOk returns a tuple with the SiteRestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteRestoreType

`func (o *MsGroupParam) SetSiteRestoreType(v string)`

SetSiteRestoreType sets SiteRestoreType field to given value.

### HasSiteRestoreType

`func (o *MsGroupParam) HasSiteRestoreType() bool`

HasSiteRestoreType returns a boolean if a field has been set.

### SetSiteRestoreTypeNil

`func (o *MsGroupParam) SetSiteRestoreTypeNil(b bool)`

 SetSiteRestoreTypeNil sets the value for SiteRestoreType to be an explicit nil

### UnsetSiteRestoreType
`func (o *MsGroupParam) UnsetSiteRestoreType()`

UnsetSiteRestoreType ensures that no value is present for SiteRestoreType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


