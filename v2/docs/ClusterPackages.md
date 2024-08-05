# ClusterPackages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Packages** | Pointer to [**[]ClusterPackageParams**](ClusterPackageParams.md) | List of cluster software packages. | [optional] 

## Methods

### NewClusterPackages

`func NewClusterPackages() *ClusterPackages`

NewClusterPackages instantiates a new ClusterPackages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPackagesWithDefaults

`func NewClusterPackagesWithDefaults() *ClusterPackages`

NewClusterPackagesWithDefaults instantiates a new ClusterPackages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackages

`func (o *ClusterPackages) GetPackages() []ClusterPackageParams`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ClusterPackages) GetPackagesOk() (*[]ClusterPackageParams, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ClusterPackages) SetPackages(v []ClusterPackageParams)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *ClusterPackages) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


