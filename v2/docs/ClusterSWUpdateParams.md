# ClusterSWUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssessSoftwareUpdateParams** | Pointer to [**AssessSoftwareUpdateParams**](AssessSoftwareUpdateParams.md) |  | [optional] 
**OperationType** | **string** | The operation type. * &#x60;DownloadUpgradePackage&#x60; - Operation to download upgrade package. * &#x60;DownloadPatchPackage&#x60; - Operation to download patch package. * &#x60;DownloadUpgradeAndPatchPackages&#x60; - Operation to download upgrade    and patch packages. * &#x60;DownloadAndUpgrade&#x60; - Operation to download package and    and then upgrade the cluster. * &#x60;DownloadAndApplyPatch&#x60; - Operation to download package and    and then apply the patch. * &#x60;DownloadAndUpgradeWithPatch&#x60; - Operation to download upgrade   and patch packages, and then, upgrade the cluster and immediately   patch it * &#x60;Upgrade&#x60; - Operation to upgrade the software on the cluster. * &#x60;ApplyPatch&#x60; - Operation to apply the patch. * &#x60;RevertPatch&#x60; - Operation to revert the patch. * &#x60;UpgradeAndPatch&#x60; - Operation to upgrade the software on the   cluster and apply a patch. * &#x60;AssessSoftwareUpdate&#x60; - Operation to perform checks to assess   the state of cluster pre/post software update (upgrade/patch). * &#x60;AbortApplyPatch&#x60; - Operation to abort the patch. * &#x60;AbortUpgrade&#x60; - Operation to abort the upgrade.  | 
**PatchParams** | Pointer to [**PatchParams**](PatchParams.md) |  | [optional] 
**UpgradeParams** | Pointer to [**UpgradeParams**](UpgradeParams.md) |  | [optional] 

## Methods

### NewClusterSWUpdateParams

`func NewClusterSWUpdateParams(operationType string, ) *ClusterSWUpdateParams`

NewClusterSWUpdateParams instantiates a new ClusterSWUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSWUpdateParamsWithDefaults

`func NewClusterSWUpdateParamsWithDefaults() *ClusterSWUpdateParams`

NewClusterSWUpdateParamsWithDefaults instantiates a new ClusterSWUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessSoftwareUpdateParams

`func (o *ClusterSWUpdateParams) GetAssessSoftwareUpdateParams() AssessSoftwareUpdateParams`

GetAssessSoftwareUpdateParams returns the AssessSoftwareUpdateParams field if non-nil, zero value otherwise.

### GetAssessSoftwareUpdateParamsOk

`func (o *ClusterSWUpdateParams) GetAssessSoftwareUpdateParamsOk() (*AssessSoftwareUpdateParams, bool)`

GetAssessSoftwareUpdateParamsOk returns a tuple with the AssessSoftwareUpdateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessSoftwareUpdateParams

`func (o *ClusterSWUpdateParams) SetAssessSoftwareUpdateParams(v AssessSoftwareUpdateParams)`

SetAssessSoftwareUpdateParams sets AssessSoftwareUpdateParams field to given value.

### HasAssessSoftwareUpdateParams

`func (o *ClusterSWUpdateParams) HasAssessSoftwareUpdateParams() bool`

HasAssessSoftwareUpdateParams returns a boolean if a field has been set.

### GetOperationType

`func (o *ClusterSWUpdateParams) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ClusterSWUpdateParams) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ClusterSWUpdateParams) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetPatchParams

`func (o *ClusterSWUpdateParams) GetPatchParams() PatchParams`

GetPatchParams returns the PatchParams field if non-nil, zero value otherwise.

### GetPatchParamsOk

`func (o *ClusterSWUpdateParams) GetPatchParamsOk() (*PatchParams, bool)`

GetPatchParamsOk returns a tuple with the PatchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchParams

`func (o *ClusterSWUpdateParams) SetPatchParams(v PatchParams)`

SetPatchParams sets PatchParams field to given value.

### HasPatchParams

`func (o *ClusterSWUpdateParams) HasPatchParams() bool`

HasPatchParams returns a boolean if a field has been set.

### GetUpgradeParams

`func (o *ClusterSWUpdateParams) GetUpgradeParams() UpgradeParams`

GetUpgradeParams returns the UpgradeParams field if non-nil, zero value otherwise.

### GetUpgradeParamsOk

`func (o *ClusterSWUpdateParams) GetUpgradeParamsOk() (*UpgradeParams, bool)`

GetUpgradeParamsOk returns a tuple with the UpgradeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeParams

`func (o *ClusterSWUpdateParams) SetUpgradeParams(v UpgradeParams)`

SetUpgradeParams sets UpgradeParams field to given value.

### HasUpgradeParams

`func (o *ClusterSWUpdateParams) HasUpgradeParams() bool`

HasUpgradeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


