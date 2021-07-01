# UpgradeClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetSwVersion** | **NullableString** | Specifies the target software version. If specified, all Nodes on the Cluster will be searched to see if they have had the specified software package uploaded to them. If the specified package is found, then it will be used for the upgrade. | 

## Methods

### NewUpgradeClusterParameters

`func NewUpgradeClusterParameters(targetSwVersion NullableString, ) *UpgradeClusterParameters`

NewUpgradeClusterParameters instantiates a new UpgradeClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeClusterParametersWithDefaults

`func NewUpgradeClusterParametersWithDefaults() *UpgradeClusterParameters`

NewUpgradeClusterParametersWithDefaults instantiates a new UpgradeClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetSwVersion

`func (o *UpgradeClusterParameters) GetTargetSwVersion() string`

GetTargetSwVersion returns the TargetSwVersion field if non-nil, zero value otherwise.

### GetTargetSwVersionOk

`func (o *UpgradeClusterParameters) GetTargetSwVersionOk() (*string, bool)`

GetTargetSwVersionOk returns a tuple with the TargetSwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSwVersion

`func (o *UpgradeClusterParameters) SetTargetSwVersion(v string)`

SetTargetSwVersion sets TargetSwVersion field to given value.


### SetTargetSwVersionNil

`func (o *UpgradeClusterParameters) SetTargetSwVersionNil(b bool)`

 SetTargetSwVersionNil sets the value for TargetSwVersion to be an explicit nil

### UnsetTargetSwVersion
`func (o *UpgradeClusterParameters) UnsetTargetSwVersion()`

UnsetTargetSwVersion ensures that no value is present for TargetSwVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


