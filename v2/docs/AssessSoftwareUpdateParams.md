# AssessSoftwareUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageType** | **NullableString** | Type of software package. | 
**Phase** | **NullableString** | Specifies the phase of software update. | 
**VersionName** | **string** | Version name of the package. | 

## Methods

### NewAssessSoftwareUpdateParams

`func NewAssessSoftwareUpdateParams(packageType NullableString, phase NullableString, versionName string, ) *AssessSoftwareUpdateParams`

NewAssessSoftwareUpdateParams instantiates a new AssessSoftwareUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssessSoftwareUpdateParamsWithDefaults

`func NewAssessSoftwareUpdateParamsWithDefaults() *AssessSoftwareUpdateParams`

NewAssessSoftwareUpdateParamsWithDefaults instantiates a new AssessSoftwareUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageType

`func (o *AssessSoftwareUpdateParams) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *AssessSoftwareUpdateParams) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *AssessSoftwareUpdateParams) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.


### SetPackageTypeNil

`func (o *AssessSoftwareUpdateParams) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *AssessSoftwareUpdateParams) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetPhase

`func (o *AssessSoftwareUpdateParams) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AssessSoftwareUpdateParams) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AssessSoftwareUpdateParams) SetPhase(v string)`

SetPhase sets Phase field to given value.


### SetPhaseNil

`func (o *AssessSoftwareUpdateParams) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *AssessSoftwareUpdateParams) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetVersionName

`func (o *AssessSoftwareUpdateParams) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *AssessSoftwareUpdateParams) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *AssessSoftwareUpdateParams) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


