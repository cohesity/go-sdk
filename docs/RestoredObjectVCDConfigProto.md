# RestoredObjectVCDConfigProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVapp** | Pointer to **NullableBool** | Whether the restored object is a VApp. | [optional] 
**IsVappTemplate** | Pointer to **NullableBool** | Whether the restored object is a VApp template. | [optional] 
**RestoredVappInfo** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoredVappTemplateInfo** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**VappEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**VcenterConnectorParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**VdcEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoredObjectVCDConfigProto

`func NewRestoredObjectVCDConfigProto() *RestoredObjectVCDConfigProto`

NewRestoredObjectVCDConfigProto instantiates a new RestoredObjectVCDConfigProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoredObjectVCDConfigProtoWithDefaults

`func NewRestoredObjectVCDConfigProtoWithDefaults() *RestoredObjectVCDConfigProto`

NewRestoredObjectVCDConfigProtoWithDefaults instantiates a new RestoredObjectVCDConfigProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVapp

`func (o *RestoredObjectVCDConfigProto) GetIsVapp() bool`

GetIsVapp returns the IsVapp field if non-nil, zero value otherwise.

### GetIsVappOk

`func (o *RestoredObjectVCDConfigProto) GetIsVappOk() (*bool, bool)`

GetIsVappOk returns a tuple with the IsVapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVapp

`func (o *RestoredObjectVCDConfigProto) SetIsVapp(v bool)`

SetIsVapp sets IsVapp field to given value.

### HasIsVapp

`func (o *RestoredObjectVCDConfigProto) HasIsVapp() bool`

HasIsVapp returns a boolean if a field has been set.

### SetIsVappNil

`func (o *RestoredObjectVCDConfigProto) SetIsVappNil(b bool)`

 SetIsVappNil sets the value for IsVapp to be an explicit nil

### UnsetIsVapp
`func (o *RestoredObjectVCDConfigProto) UnsetIsVapp()`

UnsetIsVapp ensures that no value is present for IsVapp, not even an explicit nil
### GetIsVappTemplate

`func (o *RestoredObjectVCDConfigProto) GetIsVappTemplate() bool`

GetIsVappTemplate returns the IsVappTemplate field if non-nil, zero value otherwise.

### GetIsVappTemplateOk

`func (o *RestoredObjectVCDConfigProto) GetIsVappTemplateOk() (*bool, bool)`

GetIsVappTemplateOk returns a tuple with the IsVappTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVappTemplate

`func (o *RestoredObjectVCDConfigProto) SetIsVappTemplate(v bool)`

SetIsVappTemplate sets IsVappTemplate field to given value.

### HasIsVappTemplate

`func (o *RestoredObjectVCDConfigProto) HasIsVappTemplate() bool`

HasIsVappTemplate returns a boolean if a field has been set.

### SetIsVappTemplateNil

`func (o *RestoredObjectVCDConfigProto) SetIsVappTemplateNil(b bool)`

 SetIsVappTemplateNil sets the value for IsVappTemplate to be an explicit nil

### UnsetIsVappTemplate
`func (o *RestoredObjectVCDConfigProto) UnsetIsVappTemplate()`

UnsetIsVappTemplate ensures that no value is present for IsVappTemplate, not even an explicit nil
### GetRestoredVappInfo

`func (o *RestoredObjectVCDConfigProto) GetRestoredVappInfo() EntityProto`

GetRestoredVappInfo returns the RestoredVappInfo field if non-nil, zero value otherwise.

### GetRestoredVappInfoOk

`func (o *RestoredObjectVCDConfigProto) GetRestoredVappInfoOk() (*EntityProto, bool)`

GetRestoredVappInfoOk returns a tuple with the RestoredVappInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredVappInfo

`func (o *RestoredObjectVCDConfigProto) SetRestoredVappInfo(v EntityProto)`

SetRestoredVappInfo sets RestoredVappInfo field to given value.

### HasRestoredVappInfo

`func (o *RestoredObjectVCDConfigProto) HasRestoredVappInfo() bool`

HasRestoredVappInfo returns a boolean if a field has been set.

### GetRestoredVappTemplateInfo

`func (o *RestoredObjectVCDConfigProto) GetRestoredVappTemplateInfo() EntityProto`

GetRestoredVappTemplateInfo returns the RestoredVappTemplateInfo field if non-nil, zero value otherwise.

### GetRestoredVappTemplateInfoOk

`func (o *RestoredObjectVCDConfigProto) GetRestoredVappTemplateInfoOk() (*EntityProto, bool)`

GetRestoredVappTemplateInfoOk returns a tuple with the RestoredVappTemplateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredVappTemplateInfo

`func (o *RestoredObjectVCDConfigProto) SetRestoredVappTemplateInfo(v EntityProto)`

SetRestoredVappTemplateInfo sets RestoredVappTemplateInfo field to given value.

### HasRestoredVappTemplateInfo

`func (o *RestoredObjectVCDConfigProto) HasRestoredVappTemplateInfo() bool`

HasRestoredVappTemplateInfo returns a boolean if a field has been set.

### GetVappEntity

`func (o *RestoredObjectVCDConfigProto) GetVappEntity() EntityProto`

GetVappEntity returns the VappEntity field if non-nil, zero value otherwise.

### GetVappEntityOk

`func (o *RestoredObjectVCDConfigProto) GetVappEntityOk() (*EntityProto, bool)`

GetVappEntityOk returns a tuple with the VappEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVappEntity

`func (o *RestoredObjectVCDConfigProto) SetVappEntity(v EntityProto)`

SetVappEntity sets VappEntity field to given value.

### HasVappEntity

`func (o *RestoredObjectVCDConfigProto) HasVappEntity() bool`

HasVappEntity returns a boolean if a field has been set.

### GetVcenterConnectorParams

`func (o *RestoredObjectVCDConfigProto) GetVcenterConnectorParams() ConnectorParams`

GetVcenterConnectorParams returns the VcenterConnectorParams field if non-nil, zero value otherwise.

### GetVcenterConnectorParamsOk

`func (o *RestoredObjectVCDConfigProto) GetVcenterConnectorParamsOk() (*ConnectorParams, bool)`

GetVcenterConnectorParamsOk returns a tuple with the VcenterConnectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterConnectorParams

`func (o *RestoredObjectVCDConfigProto) SetVcenterConnectorParams(v ConnectorParams)`

SetVcenterConnectorParams sets VcenterConnectorParams field to given value.

### HasVcenterConnectorParams

`func (o *RestoredObjectVCDConfigProto) HasVcenterConnectorParams() bool`

HasVcenterConnectorParams returns a boolean if a field has been set.

### GetVdcEntity

`func (o *RestoredObjectVCDConfigProto) GetVdcEntity() EntityProto`

GetVdcEntity returns the VdcEntity field if non-nil, zero value otherwise.

### GetVdcEntityOk

`func (o *RestoredObjectVCDConfigProto) GetVdcEntityOk() (*EntityProto, bool)`

GetVdcEntityOk returns a tuple with the VdcEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcEntity

`func (o *RestoredObjectVCDConfigProto) SetVdcEntity(v EntityProto)`

SetVdcEntity sets VdcEntity field to given value.

### HasVdcEntity

`func (o *RestoredObjectVCDConfigProto) HasVdcEntity() bool`

HasVdcEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


