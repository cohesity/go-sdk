# McmAgentImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **NullableString** | Specifies the type of the agent platform. | [optional] 
**PlatformSubTypes** | Pointer to [**[]PlatformSubTypeAgentImageInfo**](PlatformSubTypeAgentImageInfo.md) | Specifies the agent information of platform subtypes. | [optional] 
**DownloadURL** | Pointer to **NullableString** | Specifies the URL for agent downlaod. | [optional] 

## Methods

### NewMcmAgentImage

`func NewMcmAgentImage() *McmAgentImage`

NewMcmAgentImage instantiates a new McmAgentImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmAgentImageWithDefaults

`func NewMcmAgentImageWithDefaults() *McmAgentImage`

NewMcmAgentImageWithDefaults instantiates a new McmAgentImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *McmAgentImage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *McmAgentImage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *McmAgentImage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *McmAgentImage) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *McmAgentImage) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *McmAgentImage) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPlatformSubTypes

`func (o *McmAgentImage) GetPlatformSubTypes() []PlatformSubTypeAgentImageInfo`

GetPlatformSubTypes returns the PlatformSubTypes field if non-nil, zero value otherwise.

### GetPlatformSubTypesOk

`func (o *McmAgentImage) GetPlatformSubTypesOk() (*[]PlatformSubTypeAgentImageInfo, bool)`

GetPlatformSubTypesOk returns a tuple with the PlatformSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSubTypes

`func (o *McmAgentImage) SetPlatformSubTypes(v []PlatformSubTypeAgentImageInfo)`

SetPlatformSubTypes sets PlatformSubTypes field to given value.

### HasPlatformSubTypes

`func (o *McmAgentImage) HasPlatformSubTypes() bool`

HasPlatformSubTypes returns a boolean if a field has been set.

### GetDownloadURL

`func (o *McmAgentImage) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *McmAgentImage) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *McmAgentImage) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.

### HasDownloadURL

`func (o *McmAgentImage) HasDownloadURL() bool`

HasDownloadURL returns a boolean if a field has been set.

### SetDownloadURLNil

`func (o *McmAgentImage) SetDownloadURLNil(b bool)`

 SetDownloadURLNil sets the value for DownloadURL to be an explicit nil

### UnsetDownloadURL
`func (o *McmAgentImage) UnsetDownloadURL()`

UnsetDownloadURL ensures that no value is present for DownloadURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


