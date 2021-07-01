# OracleSpecialParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppParamsList** | Pointer to [**[]OracleAppParams**](OracleAppParams.md) | Array of application parameters i.e. database parameters for standalone/RAC and DG parameters for data guard.  Specifies the list of parameters required at app entity level. | [optional] 
**ApplicationEntityIds** | Pointer to **[]int64** | Array of Ids of Application Entities like Oracle instances, and databases that should be protected in a Protection Source.  Specifies the subset of application entities like Oracle instances, and databases to protect in a Protection Source of type kOracle&#39;. If not specified, all application entities on the Protection Source. | [optional] 
**PersistMountpoints** | Pointer to **NullableBool** | Specifies if the mountpoints for Oracle view for the current host are to be persisted. | [optional] 

## Methods

### NewOracleSpecialParameters

`func NewOracleSpecialParameters() *OracleSpecialParameters`

NewOracleSpecialParameters instantiates a new OracleSpecialParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleSpecialParametersWithDefaults

`func NewOracleSpecialParametersWithDefaults() *OracleSpecialParameters`

NewOracleSpecialParametersWithDefaults instantiates a new OracleSpecialParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppParamsList

`func (o *OracleSpecialParameters) GetAppParamsList() []OracleAppParams`

GetAppParamsList returns the AppParamsList field if non-nil, zero value otherwise.

### GetAppParamsListOk

`func (o *OracleSpecialParameters) GetAppParamsListOk() (*[]OracleAppParams, bool)`

GetAppParamsListOk returns a tuple with the AppParamsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppParamsList

`func (o *OracleSpecialParameters) SetAppParamsList(v []OracleAppParams)`

SetAppParamsList sets AppParamsList field to given value.

### HasAppParamsList

`func (o *OracleSpecialParameters) HasAppParamsList() bool`

HasAppParamsList returns a boolean if a field has been set.

### SetAppParamsListNil

`func (o *OracleSpecialParameters) SetAppParamsListNil(b bool)`

 SetAppParamsListNil sets the value for AppParamsList to be an explicit nil

### UnsetAppParamsList
`func (o *OracleSpecialParameters) UnsetAppParamsList()`

UnsetAppParamsList ensures that no value is present for AppParamsList, not even an explicit nil
### GetApplicationEntityIds

`func (o *OracleSpecialParameters) GetApplicationEntityIds() []int64`

GetApplicationEntityIds returns the ApplicationEntityIds field if non-nil, zero value otherwise.

### GetApplicationEntityIdsOk

`func (o *OracleSpecialParameters) GetApplicationEntityIdsOk() (*[]int64, bool)`

GetApplicationEntityIdsOk returns a tuple with the ApplicationEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEntityIds

`func (o *OracleSpecialParameters) SetApplicationEntityIds(v []int64)`

SetApplicationEntityIds sets ApplicationEntityIds field to given value.

### HasApplicationEntityIds

`func (o *OracleSpecialParameters) HasApplicationEntityIds() bool`

HasApplicationEntityIds returns a boolean if a field has been set.

### SetApplicationEntityIdsNil

`func (o *OracleSpecialParameters) SetApplicationEntityIdsNil(b bool)`

 SetApplicationEntityIdsNil sets the value for ApplicationEntityIds to be an explicit nil

### UnsetApplicationEntityIds
`func (o *OracleSpecialParameters) UnsetApplicationEntityIds()`

UnsetApplicationEntityIds ensures that no value is present for ApplicationEntityIds, not even an explicit nil
### GetPersistMountpoints

`func (o *OracleSpecialParameters) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleSpecialParameters) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleSpecialParameters) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleSpecialParameters) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleSpecialParameters) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleSpecialParameters) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


