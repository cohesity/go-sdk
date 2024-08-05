# ClusterRolesAuthMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustedProfileId** | **NullableString** | Specifies the trusted profile ID for fetching roles. This ID will be used during multiple API calls made to the external identity management service. | 

## Methods

### NewClusterRolesAuthMetadata

`func NewClusterRolesAuthMetadata(trustedProfileId NullableString, ) *ClusterRolesAuthMetadata`

NewClusterRolesAuthMetadata instantiates a new ClusterRolesAuthMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRolesAuthMetadataWithDefaults

`func NewClusterRolesAuthMetadataWithDefaults() *ClusterRolesAuthMetadata`

NewClusterRolesAuthMetadataWithDefaults instantiates a new ClusterRolesAuthMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrustedProfileId

`func (o *ClusterRolesAuthMetadata) GetTrustedProfileId() string`

GetTrustedProfileId returns the TrustedProfileId field if non-nil, zero value otherwise.

### GetTrustedProfileIdOk

`func (o *ClusterRolesAuthMetadata) GetTrustedProfileIdOk() (*string, bool)`

GetTrustedProfileIdOk returns a tuple with the TrustedProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProfileId

`func (o *ClusterRolesAuthMetadata) SetTrustedProfileId(v string)`

SetTrustedProfileId sets TrustedProfileId field to given value.


### SetTrustedProfileIdNil

`func (o *ClusterRolesAuthMetadata) SetTrustedProfileIdNil(b bool)`

 SetTrustedProfileIdNil sets the value for TrustedProfileId to be an explicit nil

### UnsetTrustedProfileId
`func (o *ClusterRolesAuthMetadata) UnsetTrustedProfileId()`

UnsetTrustedProfileId ensures that no value is present for TrustedProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


