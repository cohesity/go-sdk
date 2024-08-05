# ClusterAuthMethodsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTargetAuthentication** | Pointer to [**ExternalTargetAuthMetadata**](ExternalTargetAuthMetadata.md) |  | [optional] 
**FetchRolesAuthentication** | [**ClusterRolesAuthMetadata**](ClusterRolesAuthMetadata.md) |  | 
**KmsAuthentication** | Pointer to [**KmsAuthMetadata**](KmsAuthMetadata.md) |  | [optional] 

## Methods

### NewClusterAuthMethodsMetadata

`func NewClusterAuthMethodsMetadata(fetchRolesAuthentication ClusterRolesAuthMetadata, ) *ClusterAuthMethodsMetadata`

NewClusterAuthMethodsMetadata instantiates a new ClusterAuthMethodsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuthMethodsMetadataWithDefaults

`func NewClusterAuthMethodsMetadataWithDefaults() *ClusterAuthMethodsMetadata`

NewClusterAuthMethodsMetadataWithDefaults instantiates a new ClusterAuthMethodsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTargetAuthentication

`func (o *ClusterAuthMethodsMetadata) GetExternalTargetAuthentication() ExternalTargetAuthMetadata`

GetExternalTargetAuthentication returns the ExternalTargetAuthentication field if non-nil, zero value otherwise.

### GetExternalTargetAuthenticationOk

`func (o *ClusterAuthMethodsMetadata) GetExternalTargetAuthenticationOk() (*ExternalTargetAuthMetadata, bool)`

GetExternalTargetAuthenticationOk returns a tuple with the ExternalTargetAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetAuthentication

`func (o *ClusterAuthMethodsMetadata) SetExternalTargetAuthentication(v ExternalTargetAuthMetadata)`

SetExternalTargetAuthentication sets ExternalTargetAuthentication field to given value.

### HasExternalTargetAuthentication

`func (o *ClusterAuthMethodsMetadata) HasExternalTargetAuthentication() bool`

HasExternalTargetAuthentication returns a boolean if a field has been set.

### GetFetchRolesAuthentication

`func (o *ClusterAuthMethodsMetadata) GetFetchRolesAuthentication() ClusterRolesAuthMetadata`

GetFetchRolesAuthentication returns the FetchRolesAuthentication field if non-nil, zero value otherwise.

### GetFetchRolesAuthenticationOk

`func (o *ClusterAuthMethodsMetadata) GetFetchRolesAuthenticationOk() (*ClusterRolesAuthMetadata, bool)`

GetFetchRolesAuthenticationOk returns a tuple with the FetchRolesAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchRolesAuthentication

`func (o *ClusterAuthMethodsMetadata) SetFetchRolesAuthentication(v ClusterRolesAuthMetadata)`

SetFetchRolesAuthentication sets FetchRolesAuthentication field to given value.


### GetKmsAuthentication

`func (o *ClusterAuthMethodsMetadata) GetKmsAuthentication() KmsAuthMetadata`

GetKmsAuthentication returns the KmsAuthentication field if non-nil, zero value otherwise.

### GetKmsAuthenticationOk

`func (o *ClusterAuthMethodsMetadata) GetKmsAuthenticationOk() (*KmsAuthMetadata, bool)`

GetKmsAuthenticationOk returns a tuple with the KmsAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsAuthentication

`func (o *ClusterAuthMethodsMetadata) SetKmsAuthentication(v KmsAuthMetadata)`

SetKmsAuthentication sets KmsAuthentication field to given value.

### HasKmsAuthentication

`func (o *ClusterAuthMethodsMetadata) HasKmsAuthentication() bool`

HasKmsAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


