# SourcesForSid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionSources** | Pointer to [**[]ProtectionSource**](ProtectionSource.md) | Array of Protection Sources.  Specifies the Protection Source objects that the specified principal has permissions to access. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the security identifier (SID) of the principal. | [optional] 
**Views** | Pointer to [**[]View**](View.md) | Array of View Names.  Specifies the names of the Views that the specified principal has permissions to access. | [optional] 

## Methods

### NewSourcesForSid

`func NewSourcesForSid() *SourcesForSid`

NewSourcesForSid instantiates a new SourcesForSid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesForSidWithDefaults

`func NewSourcesForSidWithDefaults() *SourcesForSid`

NewSourcesForSidWithDefaults instantiates a new SourcesForSid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionSources

`func (o *SourcesForSid) GetProtectionSources() []ProtectionSource`

GetProtectionSources returns the ProtectionSources field if non-nil, zero value otherwise.

### GetProtectionSourcesOk

`func (o *SourcesForSid) GetProtectionSourcesOk() (*[]ProtectionSource, bool)`

GetProtectionSourcesOk returns a tuple with the ProtectionSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSources

`func (o *SourcesForSid) SetProtectionSources(v []ProtectionSource)`

SetProtectionSources sets ProtectionSources field to given value.

### HasProtectionSources

`func (o *SourcesForSid) HasProtectionSources() bool`

HasProtectionSources returns a boolean if a field has been set.

### SetProtectionSourcesNil

`func (o *SourcesForSid) SetProtectionSourcesNil(b bool)`

 SetProtectionSourcesNil sets the value for ProtectionSources to be an explicit nil

### UnsetProtectionSources
`func (o *SourcesForSid) UnsetProtectionSources()`

UnsetProtectionSources ensures that no value is present for ProtectionSources, not even an explicit nil
### GetSid

`func (o *SourcesForSid) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SourcesForSid) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SourcesForSid) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SourcesForSid) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SourcesForSid) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SourcesForSid) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetViews

`func (o *SourcesForSid) GetViews() []View`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *SourcesForSid) GetViewsOk() (*[]View, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *SourcesForSid) SetViews(v []View)`

SetViews sets Views field to given value.

### HasViews

`func (o *SourcesForSid) HasViews() bool`

HasViews returns a boolean if a field has been set.

### SetViewsNil

`func (o *SourcesForSid) SetViewsNil(b bool)`

 SetViewsNil sets the value for Views to be an explicit nil

### UnsetViews
`func (o *SourcesForSid) UnsetViews()`

UnsetViews ensures that no value is present for Views, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


