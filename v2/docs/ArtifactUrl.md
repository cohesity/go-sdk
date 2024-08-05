# ArtifactUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaders** | Pointer to [**[]AuthHeader**](AuthHeader.md) | HTTP headers to be included in download requests for the purposes of authenticating the client, in case the package is hosted in secure file server or artifactory.  | [optional] 
**Url** | Pointer to **string** | The URL where the artifact can be downloaded from.  | [optional] 

## Methods

### NewArtifactUrl

`func NewArtifactUrl() *ArtifactUrl`

NewArtifactUrl instantiates a new ArtifactUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactUrlWithDefaults

`func NewArtifactUrlWithDefaults() *ArtifactUrl`

NewArtifactUrlWithDefaults instantiates a new ArtifactUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaders

`func (o *ArtifactUrl) GetAuthHeaders() []AuthHeader`

GetAuthHeaders returns the AuthHeaders field if non-nil, zero value otherwise.

### GetAuthHeadersOk

`func (o *ArtifactUrl) GetAuthHeadersOk() (*[]AuthHeader, bool)`

GetAuthHeadersOk returns a tuple with the AuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaders

`func (o *ArtifactUrl) SetAuthHeaders(v []AuthHeader)`

SetAuthHeaders sets AuthHeaders field to given value.

### HasAuthHeaders

`func (o *ArtifactUrl) HasAuthHeaders() bool`

HasAuthHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *ArtifactUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArtifactUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArtifactUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ArtifactUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


