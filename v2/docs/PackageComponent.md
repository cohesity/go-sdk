# PackageComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageName** | Pointer to **string** | Name of sub package | [optional] 
**Release** | Pointer to **string** | Release Version of sub package. | [optional] 
**Version** | Pointer to **string** | Version of sub package. | [optional] 

## Methods

### NewPackageComponent

`func NewPackageComponent() *PackageComponent`

NewPackageComponent instantiates a new PackageComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageComponentWithDefaults

`func NewPackageComponentWithDefaults() *PackageComponent`

NewPackageComponentWithDefaults instantiates a new PackageComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageName

`func (o *PackageComponent) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *PackageComponent) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *PackageComponent) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *PackageComponent) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetRelease

`func (o *PackageComponent) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PackageComponent) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PackageComponent) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PackageComponent) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetVersion

`func (o *PackageComponent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageComponent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageComponent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageComponent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


