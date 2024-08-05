# ClusterDocsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | Pointer to **NullableString** | Specifies the purpose for having external hyperlink to documentation. | [optional] 
**Url** | Pointer to **NullableString** | Specifies the URL to access the endpoint for the given documentation purpose. | [optional] 

## Methods

### NewClusterDocsMetadata

`func NewClusterDocsMetadata() *ClusterDocsMetadata`

NewClusterDocsMetadata instantiates a new ClusterDocsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDocsMetadataWithDefaults

`func NewClusterDocsMetadataWithDefaults() *ClusterDocsMetadata`

NewClusterDocsMetadataWithDefaults instantiates a new ClusterDocsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *ClusterDocsMetadata) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ClusterDocsMetadata) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ClusterDocsMetadata) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ClusterDocsMetadata) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *ClusterDocsMetadata) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *ClusterDocsMetadata) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetUrl

`func (o *ClusterDocsMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterDocsMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterDocsMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClusterDocsMetadata) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ClusterDocsMetadata) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ClusterDocsMetadata) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


