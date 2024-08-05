# ClusterPackageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatiblePackages** | Pointer to **[]string** | Array of versionName values, representing compatible packages that are available on system.  | [optional] 
**Components** | Pointer to [**[]PackageComponent**](PackageComponent.md) | List of package componenets. Aplicable for one helios package  | [optional] 
**FileSizeBytes** | Pointer to **int64** | Size of file in bytes | [optional] 
**FixedIssues** | Pointer to [**[]ClusterPackageFixedIssue**](ClusterPackageFixedIssue.md) | List of issues fixed in a package. | [optional] 
**IsDowntimeRequired** | Pointer to **NullableBool** | Indicates whether package need downtime during installation | [optional] [default to false]
**Md5Checksum** | Pointer to **string** | MD5 Checksum | [optional] 
**NodeIds** | Pointer to **[]int64** | Node IDs where package is available | [optional] 
**PackageSubType** | Pointer to **string** | Sub-type of package - Security Patch or Product Patch | [optional] 
**PackageType** | Pointer to **string** | Type of the package - Upgrade or Patch | [optional] 
**ReleaseDate** | Pointer to **NullableTime** | Release date of the package. | [optional] 
**ReleaseVersion** | Pointer to **string** | Release version of the package. Examples: For upgrade package: &#39;6.6.0d_u6&#39;, &#39;7.0.&#39; For patch package - &#39;6.8.1-p1s1&#39;  | [optional] 
**Sha256Checksum** | Pointer to **string** | SHA256 Checksum | [optional] 
**Status** | Pointer to [**ClusterPackageStatus**](ClusterPackageStatus.md) |  | [optional] 
**VersionName** | Pointer to **NullableString** | Name of the package version. Example: &#39;6.6.0d_u6_release-20210714_0fad884e&#39;,   &#39;7.0.1_release-20230623_ddbb8c79&#39; for upgrade packages, &#39;6.8.1-p1s1-2023Jun26-221b8a5c&#39; for patch packages  | [optional] 

## Methods

### NewClusterPackageParams

`func NewClusterPackageParams() *ClusterPackageParams`

NewClusterPackageParams instantiates a new ClusterPackageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPackageParamsWithDefaults

`func NewClusterPackageParamsWithDefaults() *ClusterPackageParams`

NewClusterPackageParamsWithDefaults instantiates a new ClusterPackageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatiblePackages

`func (o *ClusterPackageParams) GetCompatiblePackages() []string`

GetCompatiblePackages returns the CompatiblePackages field if non-nil, zero value otherwise.

### GetCompatiblePackagesOk

`func (o *ClusterPackageParams) GetCompatiblePackagesOk() (*[]string, bool)`

GetCompatiblePackagesOk returns a tuple with the CompatiblePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatiblePackages

`func (o *ClusterPackageParams) SetCompatiblePackages(v []string)`

SetCompatiblePackages sets CompatiblePackages field to given value.

### HasCompatiblePackages

`func (o *ClusterPackageParams) HasCompatiblePackages() bool`

HasCompatiblePackages returns a boolean if a field has been set.

### GetComponents

`func (o *ClusterPackageParams) GetComponents() []PackageComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ClusterPackageParams) GetComponentsOk() (*[]PackageComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ClusterPackageParams) SetComponents(v []PackageComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ClusterPackageParams) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *ClusterPackageParams) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *ClusterPackageParams) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *ClusterPackageParams) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *ClusterPackageParams) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetFixedIssues

`func (o *ClusterPackageParams) GetFixedIssues() []ClusterPackageFixedIssue`

GetFixedIssues returns the FixedIssues field if non-nil, zero value otherwise.

### GetFixedIssuesOk

`func (o *ClusterPackageParams) GetFixedIssuesOk() (*[]ClusterPackageFixedIssue, bool)`

GetFixedIssuesOk returns a tuple with the FixedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIssues

`func (o *ClusterPackageParams) SetFixedIssues(v []ClusterPackageFixedIssue)`

SetFixedIssues sets FixedIssues field to given value.

### HasFixedIssues

`func (o *ClusterPackageParams) HasFixedIssues() bool`

HasFixedIssues returns a boolean if a field has been set.

### GetIsDowntimeRequired

`func (o *ClusterPackageParams) GetIsDowntimeRequired() bool`

GetIsDowntimeRequired returns the IsDowntimeRequired field if non-nil, zero value otherwise.

### GetIsDowntimeRequiredOk

`func (o *ClusterPackageParams) GetIsDowntimeRequiredOk() (*bool, bool)`

GetIsDowntimeRequiredOk returns a tuple with the IsDowntimeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDowntimeRequired

`func (o *ClusterPackageParams) SetIsDowntimeRequired(v bool)`

SetIsDowntimeRequired sets IsDowntimeRequired field to given value.

### HasIsDowntimeRequired

`func (o *ClusterPackageParams) HasIsDowntimeRequired() bool`

HasIsDowntimeRequired returns a boolean if a field has been set.

### SetIsDowntimeRequiredNil

`func (o *ClusterPackageParams) SetIsDowntimeRequiredNil(b bool)`

 SetIsDowntimeRequiredNil sets the value for IsDowntimeRequired to be an explicit nil

### UnsetIsDowntimeRequired
`func (o *ClusterPackageParams) UnsetIsDowntimeRequired()`

UnsetIsDowntimeRequired ensures that no value is present for IsDowntimeRequired, not even an explicit nil
### GetMd5Checksum

`func (o *ClusterPackageParams) GetMd5Checksum() string`

GetMd5Checksum returns the Md5Checksum field if non-nil, zero value otherwise.

### GetMd5ChecksumOk

`func (o *ClusterPackageParams) GetMd5ChecksumOk() (*string, bool)`

GetMd5ChecksumOk returns a tuple with the Md5Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Checksum

`func (o *ClusterPackageParams) SetMd5Checksum(v string)`

SetMd5Checksum sets Md5Checksum field to given value.

### HasMd5Checksum

`func (o *ClusterPackageParams) HasMd5Checksum() bool`

HasMd5Checksum returns a boolean if a field has been set.

### GetNodeIds

`func (o *ClusterPackageParams) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ClusterPackageParams) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ClusterPackageParams) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *ClusterPackageParams) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetPackageSubType

`func (o *ClusterPackageParams) GetPackageSubType() string`

GetPackageSubType returns the PackageSubType field if non-nil, zero value otherwise.

### GetPackageSubTypeOk

`func (o *ClusterPackageParams) GetPackageSubTypeOk() (*string, bool)`

GetPackageSubTypeOk returns a tuple with the PackageSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSubType

`func (o *ClusterPackageParams) SetPackageSubType(v string)`

SetPackageSubType sets PackageSubType field to given value.

### HasPackageSubType

`func (o *ClusterPackageParams) HasPackageSubType() bool`

HasPackageSubType returns a boolean if a field has been set.

### GetPackageType

`func (o *ClusterPackageParams) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ClusterPackageParams) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ClusterPackageParams) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ClusterPackageParams) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ClusterPackageParams) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ClusterPackageParams) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ClusterPackageParams) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ClusterPackageParams) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ClusterPackageParams) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ClusterPackageParams) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetReleaseVersion

`func (o *ClusterPackageParams) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *ClusterPackageParams) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *ClusterPackageParams) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *ClusterPackageParams) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.

### GetSha256Checksum

`func (o *ClusterPackageParams) GetSha256Checksum() string`

GetSha256Checksum returns the Sha256Checksum field if non-nil, zero value otherwise.

### GetSha256ChecksumOk

`func (o *ClusterPackageParams) GetSha256ChecksumOk() (*string, bool)`

GetSha256ChecksumOk returns a tuple with the Sha256Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Checksum

`func (o *ClusterPackageParams) SetSha256Checksum(v string)`

SetSha256Checksum sets Sha256Checksum field to given value.

### HasSha256Checksum

`func (o *ClusterPackageParams) HasSha256Checksum() bool`

HasSha256Checksum returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterPackageParams) GetStatus() ClusterPackageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterPackageParams) GetStatusOk() (*ClusterPackageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterPackageParams) SetStatus(v ClusterPackageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterPackageParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersionName

`func (o *ClusterPackageParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ClusterPackageParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ClusterPackageParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *ClusterPackageParams) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### SetVersionNameNil

`func (o *ClusterPackageParams) SetVersionNameNil(b bool)`

 SetVersionNameNil sets the value for VersionName to be an explicit nil

### UnsetVersionName
`func (o *ClusterPackageParams) UnsetVersionName()`

UnsetVersionName ensures that no value is present for VersionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


