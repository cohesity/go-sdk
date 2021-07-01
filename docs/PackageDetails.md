# PackageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowntimeRequired** | Pointer to **NullableBool** | Specifies whether or not downtime is required to update to this package. | [optional] 
**InstalledOnNodes** | Pointer to **[]int64** | Specifies the list of IDs of nodes on the cluster where this package is currently installed. | [optional] 
**PackageName** | Pointer to **NullableString** | Specifies the name of the current package. | [optional] 
**ReleaseDate** | Pointer to **NullableString** | Specifies the release date of this package. | [optional] 

## Methods

### NewPackageDetails

`func NewPackageDetails() *PackageDetails`

NewPackageDetails instantiates a new PackageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDetailsWithDefaults

`func NewPackageDetailsWithDefaults() *PackageDetails`

NewPackageDetailsWithDefaults instantiates a new PackageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDowntimeRequired

`func (o *PackageDetails) GetDowntimeRequired() bool`

GetDowntimeRequired returns the DowntimeRequired field if non-nil, zero value otherwise.

### GetDowntimeRequiredOk

`func (o *PackageDetails) GetDowntimeRequiredOk() (*bool, bool)`

GetDowntimeRequiredOk returns a tuple with the DowntimeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntimeRequired

`func (o *PackageDetails) SetDowntimeRequired(v bool)`

SetDowntimeRequired sets DowntimeRequired field to given value.

### HasDowntimeRequired

`func (o *PackageDetails) HasDowntimeRequired() bool`

HasDowntimeRequired returns a boolean if a field has been set.

### SetDowntimeRequiredNil

`func (o *PackageDetails) SetDowntimeRequiredNil(b bool)`

 SetDowntimeRequiredNil sets the value for DowntimeRequired to be an explicit nil

### UnsetDowntimeRequired
`func (o *PackageDetails) UnsetDowntimeRequired()`

UnsetDowntimeRequired ensures that no value is present for DowntimeRequired, not even an explicit nil
### GetInstalledOnNodes

`func (o *PackageDetails) GetInstalledOnNodes() []int64`

GetInstalledOnNodes returns the InstalledOnNodes field if non-nil, zero value otherwise.

### GetInstalledOnNodesOk

`func (o *PackageDetails) GetInstalledOnNodesOk() (*[]int64, bool)`

GetInstalledOnNodesOk returns a tuple with the InstalledOnNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledOnNodes

`func (o *PackageDetails) SetInstalledOnNodes(v []int64)`

SetInstalledOnNodes sets InstalledOnNodes field to given value.

### HasInstalledOnNodes

`func (o *PackageDetails) HasInstalledOnNodes() bool`

HasInstalledOnNodes returns a boolean if a field has been set.

### SetInstalledOnNodesNil

`func (o *PackageDetails) SetInstalledOnNodesNil(b bool)`

 SetInstalledOnNodesNil sets the value for InstalledOnNodes to be an explicit nil

### UnsetInstalledOnNodes
`func (o *PackageDetails) UnsetInstalledOnNodes()`

UnsetInstalledOnNodes ensures that no value is present for InstalledOnNodes, not even an explicit nil
### GetPackageName

`func (o *PackageDetails) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *PackageDetails) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *PackageDetails) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *PackageDetails) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### SetPackageNameNil

`func (o *PackageDetails) SetPackageNameNil(b bool)`

 SetPackageNameNil sets the value for PackageName to be an explicit nil

### UnsetPackageName
`func (o *PackageDetails) UnsetPackageName()`

UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
### GetReleaseDate

`func (o *PackageDetails) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PackageDetails) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PackageDetails) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PackageDetails) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *PackageDetails) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *PackageDetails) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


