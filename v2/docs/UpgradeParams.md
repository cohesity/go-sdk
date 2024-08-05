# UpgradeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortOnPreChecksFailure** | Pointer to **bool** | Specifies if healthchecks failure will cause upgrade to be aborted. By default we abort upgrade if there are healthchecks failures .Cluster will stop the upgrade.and present the failures which need to be resolved before proceeding with upgrade. If set to false upgrade will not be aborted on healthchecks failure.  | [optional] [default to true]
**AutoAgentUpgrade** | Pointer to **bool** | Upgrade Cohesity agents on servers of registered sources.  | [optional] 
**IgnoreSwIncompatibility** | Pointer to **bool** | If set to true, software incomaptibility checks are ignored. Applicable for operations: * &#x60;DownloadAndUpgradeWithPatch&#x60; * &#x60;DownloadAndUpgrade&#x60; * &#x60;Upgrade&#x60; * &#x60;UpgradeAndPatch&#x60;  | [optional] [default to false]
**Md5Sum** | Pointer to **string** | md5Sum of the upgrade package. Applicable for operations: * &#x60;DownloadAndUpgradeWithPatch&#x60; * &#x60;DownloadAndUpgrade&#x60; * &#x60;Upgrade&#x60; * &#x60;UpgradeAndPatch&#x60;  | [optional] 
**PackageUrl** | Pointer to [**ArtifactUrl**](ArtifactUrl.md) |  | [optional] 
**RunUpgradeInParallel** | Pointer to **bool** | If set to true, upgrade will run in parallel on all nodes. Applicable for operations: * &#x60;DownloadAndUpgradeWithPatch&#x60; * &#x60;DownloadAndUpgrade&#x60; * &#x60;Upgrade&#x60; * &#x60;UpgradeAndPatch&#x60;  | [optional] [default to false]
**VersionName** | Pointer to **string** | Version name of the package if the package is already downloaded. Example: 6.3.1h_release-20210714_0fad884e. Applicable for operations: * &#x60;Upgrade&#x60; * &#x60;UpgradeAndPatch&#x60;  | [optional] 

## Methods

### NewUpgradeParams

`func NewUpgradeParams() *UpgradeParams`

NewUpgradeParams instantiates a new UpgradeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeParamsWithDefaults

`func NewUpgradeParamsWithDefaults() *UpgradeParams`

NewUpgradeParamsWithDefaults instantiates a new UpgradeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortOnPreChecksFailure

`func (o *UpgradeParams) GetAbortOnPreChecksFailure() bool`

GetAbortOnPreChecksFailure returns the AbortOnPreChecksFailure field if non-nil, zero value otherwise.

### GetAbortOnPreChecksFailureOk

`func (o *UpgradeParams) GetAbortOnPreChecksFailureOk() (*bool, bool)`

GetAbortOnPreChecksFailureOk returns a tuple with the AbortOnPreChecksFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortOnPreChecksFailure

`func (o *UpgradeParams) SetAbortOnPreChecksFailure(v bool)`

SetAbortOnPreChecksFailure sets AbortOnPreChecksFailure field to given value.

### HasAbortOnPreChecksFailure

`func (o *UpgradeParams) HasAbortOnPreChecksFailure() bool`

HasAbortOnPreChecksFailure returns a boolean if a field has been set.

### GetAutoAgentUpgrade

`func (o *UpgradeParams) GetAutoAgentUpgrade() bool`

GetAutoAgentUpgrade returns the AutoAgentUpgrade field if non-nil, zero value otherwise.

### GetAutoAgentUpgradeOk

`func (o *UpgradeParams) GetAutoAgentUpgradeOk() (*bool, bool)`

GetAutoAgentUpgradeOk returns a tuple with the AutoAgentUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAgentUpgrade

`func (o *UpgradeParams) SetAutoAgentUpgrade(v bool)`

SetAutoAgentUpgrade sets AutoAgentUpgrade field to given value.

### HasAutoAgentUpgrade

`func (o *UpgradeParams) HasAutoAgentUpgrade() bool`

HasAutoAgentUpgrade returns a boolean if a field has been set.

### GetIgnoreSwIncompatibility

`func (o *UpgradeParams) GetIgnoreSwIncompatibility() bool`

GetIgnoreSwIncompatibility returns the IgnoreSwIncompatibility field if non-nil, zero value otherwise.

### GetIgnoreSwIncompatibilityOk

`func (o *UpgradeParams) GetIgnoreSwIncompatibilityOk() (*bool, bool)`

GetIgnoreSwIncompatibilityOk returns a tuple with the IgnoreSwIncompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSwIncompatibility

`func (o *UpgradeParams) SetIgnoreSwIncompatibility(v bool)`

SetIgnoreSwIncompatibility sets IgnoreSwIncompatibility field to given value.

### HasIgnoreSwIncompatibility

`func (o *UpgradeParams) HasIgnoreSwIncompatibility() bool`

HasIgnoreSwIncompatibility returns a boolean if a field has been set.

### GetMd5Sum

`func (o *UpgradeParams) GetMd5Sum() string`

GetMd5Sum returns the Md5Sum field if non-nil, zero value otherwise.

### GetMd5SumOk

`func (o *UpgradeParams) GetMd5SumOk() (*string, bool)`

GetMd5SumOk returns a tuple with the Md5Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Sum

`func (o *UpgradeParams) SetMd5Sum(v string)`

SetMd5Sum sets Md5Sum field to given value.

### HasMd5Sum

`func (o *UpgradeParams) HasMd5Sum() bool`

HasMd5Sum returns a boolean if a field has been set.

### GetPackageUrl

`func (o *UpgradeParams) GetPackageUrl() ArtifactUrl`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *UpgradeParams) GetPackageUrlOk() (*ArtifactUrl, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *UpgradeParams) SetPackageUrl(v ArtifactUrl)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *UpgradeParams) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetRunUpgradeInParallel

`func (o *UpgradeParams) GetRunUpgradeInParallel() bool`

GetRunUpgradeInParallel returns the RunUpgradeInParallel field if non-nil, zero value otherwise.

### GetRunUpgradeInParallelOk

`func (o *UpgradeParams) GetRunUpgradeInParallelOk() (*bool, bool)`

GetRunUpgradeInParallelOk returns a tuple with the RunUpgradeInParallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunUpgradeInParallel

`func (o *UpgradeParams) SetRunUpgradeInParallel(v bool)`

SetRunUpgradeInParallel sets RunUpgradeInParallel field to given value.

### HasRunUpgradeInParallel

`func (o *UpgradeParams) HasRunUpgradeInParallel() bool`

HasRunUpgradeInParallel returns a boolean if a field has been set.

### GetVersionName

`func (o *UpgradeParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *UpgradeParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *UpgradeParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *UpgradeParams) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


