# PatchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortOnPreChecksFailure** | Pointer to **bool** | Specifies if healthchecks failure will cause apply patch to be aborted. By default we patch if there are healthchecks failures. Cluster will stop the apply patch and present the failures which need to be resolved before proceeding with apply patch. If set to false, apply patchwill not be aborted on healthchecks failure.  | [optional] [default to true]
**ApplyPatchInParallel** | Pointer to **bool** | If set to true, patch will be applied in parallel on all nodes. Applicable for operations: * &#x60;ApplyPatch&#x60; * &#x60;DownloadAndApplyPatch&#x60;  | [optional] [default to false]
**NodeIds** | Pointer to **[]int64** | Node IDs where patch has to be applied.  If unspecified, patch will be applied on all nodes.  | [optional] 
**PackageUrl** | Pointer to [**ArtifactUrl**](ArtifactUrl.md) |  | [optional] 
**VersionName** | Pointer to **string** | Version name of the package if the package is already downloaded. Example: 7.0.1-p1-2023Jul04-cc6d7c5f Applicable for operations: * &#x60;ApplyPatch&#x60; * &#x60;RevertPatch&#x60; * &#x60;UpgradeAndPatch&#x60;  | [optional] 

## Methods

### NewPatchParams

`func NewPatchParams() *PatchParams`

NewPatchParams instantiates a new PatchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchParamsWithDefaults

`func NewPatchParamsWithDefaults() *PatchParams`

NewPatchParamsWithDefaults instantiates a new PatchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortOnPreChecksFailure

`func (o *PatchParams) GetAbortOnPreChecksFailure() bool`

GetAbortOnPreChecksFailure returns the AbortOnPreChecksFailure field if non-nil, zero value otherwise.

### GetAbortOnPreChecksFailureOk

`func (o *PatchParams) GetAbortOnPreChecksFailureOk() (*bool, bool)`

GetAbortOnPreChecksFailureOk returns a tuple with the AbortOnPreChecksFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortOnPreChecksFailure

`func (o *PatchParams) SetAbortOnPreChecksFailure(v bool)`

SetAbortOnPreChecksFailure sets AbortOnPreChecksFailure field to given value.

### HasAbortOnPreChecksFailure

`func (o *PatchParams) HasAbortOnPreChecksFailure() bool`

HasAbortOnPreChecksFailure returns a boolean if a field has been set.

### GetApplyPatchInParallel

`func (o *PatchParams) GetApplyPatchInParallel() bool`

GetApplyPatchInParallel returns the ApplyPatchInParallel field if non-nil, zero value otherwise.

### GetApplyPatchInParallelOk

`func (o *PatchParams) GetApplyPatchInParallelOk() (*bool, bool)`

GetApplyPatchInParallelOk returns a tuple with the ApplyPatchInParallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPatchInParallel

`func (o *PatchParams) SetApplyPatchInParallel(v bool)`

SetApplyPatchInParallel sets ApplyPatchInParallel field to given value.

### HasApplyPatchInParallel

`func (o *PatchParams) HasApplyPatchInParallel() bool`

HasApplyPatchInParallel returns a boolean if a field has been set.

### GetNodeIds

`func (o *PatchParams) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *PatchParams) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *PatchParams) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *PatchParams) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetPackageUrl

`func (o *PatchParams) GetPackageUrl() ArtifactUrl`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *PatchParams) GetPackageUrlOk() (*ArtifactUrl, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *PatchParams) SetPackageUrl(v ArtifactUrl)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *PatchParams) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetVersionName

`func (o *PatchParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *PatchParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *PatchParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *PatchParams) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


