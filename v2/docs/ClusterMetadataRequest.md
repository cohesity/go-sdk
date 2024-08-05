# ClusterMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethods** | [**ClusterAuthMethodsMetadata**](ClusterAuthMethodsMetadata.md) |  | 
**CustomProperties** | Pointer to [**[]ClusterCustomMetadata**](ClusterCustomMetadata.md) | Specifies the list of custom properties associated with the cluster. API caller can choose to set the following properties using provided key and value fields. The input values must always be in the string format and each key must be unique. | [optional] 
**Docs** | Pointer to [**[]ClusterDocsMetadata**](ClusterDocsMetadata.md) | Specifies the docs related metadata specific to the cluster. This metadata mainly consists of any external hyperlinks to service provider&#39;s documentation. | [optional] 
**Sla** | Pointer to [**ClusterSLAMetadata**](ClusterSLAMetadata.md) |  | [optional] 

## Methods

### NewClusterMetadataRequest

`func NewClusterMetadataRequest(authenticationMethods ClusterAuthMethodsMetadata, ) *ClusterMetadataRequest`

NewClusterMetadataRequest instantiates a new ClusterMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetadataRequestWithDefaults

`func NewClusterMetadataRequestWithDefaults() *ClusterMetadataRequest`

NewClusterMetadataRequestWithDefaults instantiates a new ClusterMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethods

`func (o *ClusterMetadataRequest) GetAuthenticationMethods() ClusterAuthMethodsMetadata`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *ClusterMetadataRequest) GetAuthenticationMethodsOk() (*ClusterAuthMethodsMetadata, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *ClusterMetadataRequest) SetAuthenticationMethods(v ClusterAuthMethodsMetadata)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.


### GetCustomProperties

`func (o *ClusterMetadataRequest) GetCustomProperties() []ClusterCustomMetadata`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ClusterMetadataRequest) GetCustomPropertiesOk() (*[]ClusterCustomMetadata, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ClusterMetadataRequest) SetCustomProperties(v []ClusterCustomMetadata)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ClusterMetadataRequest) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ClusterMetadataRequest) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ClusterMetadataRequest) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetDocs

`func (o *ClusterMetadataRequest) GetDocs() []ClusterDocsMetadata`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *ClusterMetadataRequest) GetDocsOk() (*[]ClusterDocsMetadata, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *ClusterMetadataRequest) SetDocs(v []ClusterDocsMetadata)`

SetDocs sets Docs field to given value.

### HasDocs

`func (o *ClusterMetadataRequest) HasDocs() bool`

HasDocs returns a boolean if a field has been set.

### SetDocsNil

`func (o *ClusterMetadataRequest) SetDocsNil(b bool)`

 SetDocsNil sets the value for Docs to be an explicit nil

### UnsetDocs
`func (o *ClusterMetadataRequest) UnsetDocs()`

UnsetDocs ensures that no value is present for Docs, not even an explicit nil
### GetSla

`func (o *ClusterMetadataRequest) GetSla() ClusterSLAMetadata`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ClusterMetadataRequest) GetSlaOk() (*ClusterSLAMetadata, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ClusterMetadataRequest) SetSla(v ClusterSLAMetadata)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ClusterMetadataRequest) HasSla() bool`

HasSla returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


