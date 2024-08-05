# UdaIndexParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Whether the dynamic config support is enabled or not. | [optional] 
**Id** | Pointer to **NullableInt64** | Unique id for each Uda connector. | [optional] 
**Label** | Pointer to **NullableString** | Label to be shown On UI screen. | [optional] 
**SourceType** | Pointer to **NullableString** | Type of the uda connector. | [optional] 
**Syntax** | Pointer to **NullableInt32** | Syntax used for configuration file. | [optional] 
**UiFeatureFlag** | Pointer to **NullableString** | UI feature flag that controls the visibility of the connector. | [optional] 
**Version** | Pointer to **NullableInt64** | Version of the connector. | [optional] 

## Methods

### NewUdaIndexParams

`func NewUdaIndexParams() *UdaIndexParams`

NewUdaIndexParams instantiates a new UdaIndexParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaIndexParamsWithDefaults

`func NewUdaIndexParamsWithDefaults() *UdaIndexParams`

NewUdaIndexParamsWithDefaults instantiates a new UdaIndexParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UdaIndexParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UdaIndexParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UdaIndexParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UdaIndexParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UdaIndexParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UdaIndexParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetId

`func (o *UdaIndexParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UdaIndexParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UdaIndexParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UdaIndexParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UdaIndexParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UdaIndexParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *UdaIndexParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UdaIndexParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UdaIndexParams) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UdaIndexParams) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *UdaIndexParams) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *UdaIndexParams) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetSourceType

`func (o *UdaIndexParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UdaIndexParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UdaIndexParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UdaIndexParams) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *UdaIndexParams) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *UdaIndexParams) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSyntax

`func (o *UdaIndexParams) GetSyntax() int32`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *UdaIndexParams) GetSyntaxOk() (*int32, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *UdaIndexParams) SetSyntax(v int32)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *UdaIndexParams) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### SetSyntaxNil

`func (o *UdaIndexParams) SetSyntaxNil(b bool)`

 SetSyntaxNil sets the value for Syntax to be an explicit nil

### UnsetSyntax
`func (o *UdaIndexParams) UnsetSyntax()`

UnsetSyntax ensures that no value is present for Syntax, not even an explicit nil
### GetUiFeatureFlag

`func (o *UdaIndexParams) GetUiFeatureFlag() string`

GetUiFeatureFlag returns the UiFeatureFlag field if non-nil, zero value otherwise.

### GetUiFeatureFlagOk

`func (o *UdaIndexParams) GetUiFeatureFlagOk() (*string, bool)`

GetUiFeatureFlagOk returns a tuple with the UiFeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiFeatureFlag

`func (o *UdaIndexParams) SetUiFeatureFlag(v string)`

SetUiFeatureFlag sets UiFeatureFlag field to given value.

### HasUiFeatureFlag

`func (o *UdaIndexParams) HasUiFeatureFlag() bool`

HasUiFeatureFlag returns a boolean if a field has been set.

### SetUiFeatureFlagNil

`func (o *UdaIndexParams) SetUiFeatureFlagNil(b bool)`

 SetUiFeatureFlagNil sets the value for UiFeatureFlag to be an explicit nil

### UnsetUiFeatureFlag
`func (o *UdaIndexParams) UnsetUiFeatureFlag()`

UnsetUiFeatureFlag ensures that no value is present for UiFeatureFlag, not even an explicit nil
### GetVersion

`func (o *UdaIndexParams) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UdaIndexParams) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UdaIndexParams) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UdaIndexParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *UdaIndexParams) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *UdaIndexParams) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


