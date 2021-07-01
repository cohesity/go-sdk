# DownloadPackageParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **NullableString** | Specifies a URL from which the package can be downloaded to the Cluster. | 

## Methods

### NewDownloadPackageParameters

`func NewDownloadPackageParameters(url NullableString, ) *DownloadPackageParameters`

NewDownloadPackageParameters instantiates a new DownloadPackageParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadPackageParametersWithDefaults

`func NewDownloadPackageParametersWithDefaults() *DownloadPackageParameters`

NewDownloadPackageParametersWithDefaults instantiates a new DownloadPackageParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DownloadPackageParameters) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DownloadPackageParameters) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DownloadPackageParameters) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *DownloadPackageParameters) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *DownloadPackageParameters) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


