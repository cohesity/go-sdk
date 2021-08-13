# ObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmwareParams** | Pointer to [**VmwareObjectEntityParams**](VmwareObjectEntityParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectParams**](IsilonObjectParams.md) | Specifies the parameters for Isilon object. | [optional] 
**NetappParams** | Pointer to [**NetappObjectParams**](NetappObjectParams.md) | Specifies the parameters for NetApp object. | [optional] 
**GenericNasParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for GenericNas object. | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectParams**](FlashbladeObjectParams.md) | Specifies the parameters for Flashblade object. | [optional] 
**ElastifileParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for Elastifile object. | [optional] 
**GpfsParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for GPFS object. | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectEntityParams**](MssqlObjectEntityParams.md) | Specifies the parameters for Msssql object. | [optional] 
**OracleParams** | Pointer to [**OracleObjectEntityParams**](OracleObjectEntityParams.md) | Specifies the parameters for Oracle object. | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectEntityParams**](PhysicalObjectEntityParams.md) | Specifies the parameters for Physical object. | [optional] 
**SharepointParams** | Pointer to [**SharepointObjectEntityParams**](SharepointObjectEntityParams.md) | Specifies the parameters for Sharepoint object. | [optional] 

## Methods

### NewObjectAllOf

`func NewObjectAllOf() *ObjectAllOf`

NewObjectAllOf instantiates a new ObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectAllOfWithDefaults

`func NewObjectAllOfWithDefaults() *ObjectAllOf`

NewObjectAllOfWithDefaults instantiates a new ObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmwareParams

`func (o *ObjectAllOf) GetVmwareParams() VmwareObjectEntityParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ObjectAllOf) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ObjectAllOf) SetVmwareParams(v VmwareObjectEntityParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ObjectAllOf) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ObjectAllOf) GetIsilonParams() IsilonObjectParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ObjectAllOf) GetIsilonParamsOk() (*IsilonObjectParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ObjectAllOf) SetIsilonParams(v IsilonObjectParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ObjectAllOf) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ObjectAllOf) GetNetappParams() NetappObjectParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ObjectAllOf) GetNetappParamsOk() (*NetappObjectParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ObjectAllOf) SetNetappParams(v NetappObjectParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ObjectAllOf) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ObjectAllOf) GetGenericNasParams() CommonNasObjectParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ObjectAllOf) GetGenericNasParamsOk() (*CommonNasObjectParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ObjectAllOf) SetGenericNasParams(v CommonNasObjectParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ObjectAllOf) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ObjectAllOf) GetFlashbladeParams() FlashbladeObjectParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ObjectAllOf) GetFlashbladeParamsOk() (*FlashbladeObjectParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ObjectAllOf) SetFlashbladeParams(v FlashbladeObjectParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ObjectAllOf) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *ObjectAllOf) GetElastifileParams() CommonNasObjectParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ObjectAllOf) GetElastifileParamsOk() (*CommonNasObjectParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ObjectAllOf) SetElastifileParams(v CommonNasObjectParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ObjectAllOf) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ObjectAllOf) GetGpfsParams() CommonNasObjectParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ObjectAllOf) GetGpfsParamsOk() (*CommonNasObjectParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ObjectAllOf) SetGpfsParams(v CommonNasObjectParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ObjectAllOf) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ObjectAllOf) GetMssqlParams() MssqlObjectEntityParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ObjectAllOf) GetMssqlParamsOk() (*MssqlObjectEntityParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ObjectAllOf) SetMssqlParams(v MssqlObjectEntityParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ObjectAllOf) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ObjectAllOf) GetOracleParams() OracleObjectEntityParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ObjectAllOf) GetOracleParamsOk() (*OracleObjectEntityParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ObjectAllOf) SetOracleParams(v OracleObjectEntityParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ObjectAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *ObjectAllOf) GetPhysicalParams() PhysicalObjectEntityParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ObjectAllOf) GetPhysicalParamsOk() (*PhysicalObjectEntityParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ObjectAllOf) SetPhysicalParams(v PhysicalObjectEntityParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ObjectAllOf) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *ObjectAllOf) GetSharepointParams() SharepointObjectEntityParams`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *ObjectAllOf) GetSharepointParamsOk() (*SharepointObjectEntityParams, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *ObjectAllOf) SetSharepointParams(v SharepointObjectEntityParams)`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *ObjectAllOf) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


