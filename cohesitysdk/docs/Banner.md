# Banner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannerId** | Pointer to **NullableString** | Specifies a banner_id which can uniquely identify a banner. This may be the cluster_id, or the tenant_id, or the group_id, or the user SID etc. If this field is nil, the it is assumed to be the cluster_id. The content is stored against this &#39;row&#39; in Scribe. | [optional] 
**Content** | Pointer to **NullableString** | Specifies the content of the banner. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | createdTimeMsecs field is deprecated. Timestamp at which banner was created. | [optional] 
**Description** | Pointer to **NullableString** | description field is deprecated. Specifies the description of this banner. | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | lastUpdatedTimeMsecs field is deprecated. Timestamp at which banner was last updated. | [optional] 

## Methods

### NewBanner

`func NewBanner() *Banner`

NewBanner instantiates a new Banner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannerWithDefaults

`func NewBannerWithDefaults() *Banner`

NewBannerWithDefaults instantiates a new Banner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerId

`func (o *Banner) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *Banner) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *Banner) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *Banner) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *Banner) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *Banner) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetContent

`func (o *Banner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Banner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Banner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Banner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *Banner) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *Banner) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *Banner) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *Banner) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *Banner) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *Banner) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *Banner) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *Banner) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDescription

`func (o *Banner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Banner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Banner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Banner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Banner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Banner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *Banner) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *Banner) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *Banner) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *Banner) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *Banner) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *Banner) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


