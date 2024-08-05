# IpmiFruInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Specifies the IPMI Field Replaceable Unit(FRU) ID for given node. | [optional] 
**BoardMfg** | Pointer to **NullableString** | Specifies the board manufacturer for given node. | [optional] 
**BoardMfgDate** | Pointer to **NullableString** | Specifies the board manufacturing date for given node. | [optional] 
**BoardPN** | Pointer to **NullableString** | Specifies the board part number for given node. | [optional] 
**BoardProduct** | Pointer to **NullableString** | Specifies the board product name for given node. | [optional] 
**BoardSerial** | Pointer to **NullableString** | Specifies the board serial number for given node. | [optional] 
**ChassisExtra** | Pointer to **NullableString** | Specifies the information about chassis extras provided for given node. | [optional] 
**ChassisPN** | Pointer to **NullableString** | Specifies the chassis part number for given node. | [optional] 
**ChassisSerial** | Pointer to **NullableString** | Specifies the chassis serial number for given node. | [optional] 
**ChassisType** | Pointer to **NullableString** | Specifies the type of chassis for given node. | [optional] 
**ProductMfg** | Pointer to **NullableString** | Specifies the product manufacturer for given node. | [optional] 
**ProductName** | Pointer to **NullableString** | Specifies the product name for given node. | [optional] 
**ProductPN** | Pointer to **NullableString** | Specifies the product part number for given node. | [optional] 
**ProductSerial** | Pointer to **NullableString** | Specifies the product serial number for given node. | [optional] 
**ProductVersion** | Pointer to **NullableString** | Specifies the product version for given node. | [optional] 

## Methods

### NewIpmiFruInfo

`func NewIpmiFruInfo() *IpmiFruInfo`

NewIpmiFruInfo instantiates a new IpmiFruInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiFruInfoWithDefaults

`func NewIpmiFruInfoWithDefaults() *IpmiFruInfo`

NewIpmiFruInfoWithDefaults instantiates a new IpmiFruInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpmiFruInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpmiFruInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpmiFruInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpmiFruInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IpmiFruInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IpmiFruInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetBoardMfg

`func (o *IpmiFruInfo) GetBoardMfg() string`

GetBoardMfg returns the BoardMfg field if non-nil, zero value otherwise.

### GetBoardMfgOk

`func (o *IpmiFruInfo) GetBoardMfgOk() (*string, bool)`

GetBoardMfgOk returns a tuple with the BoardMfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardMfg

`func (o *IpmiFruInfo) SetBoardMfg(v string)`

SetBoardMfg sets BoardMfg field to given value.

### HasBoardMfg

`func (o *IpmiFruInfo) HasBoardMfg() bool`

HasBoardMfg returns a boolean if a field has been set.

### SetBoardMfgNil

`func (o *IpmiFruInfo) SetBoardMfgNil(b bool)`

 SetBoardMfgNil sets the value for BoardMfg to be an explicit nil

### UnsetBoardMfg
`func (o *IpmiFruInfo) UnsetBoardMfg()`

UnsetBoardMfg ensures that no value is present for BoardMfg, not even an explicit nil
### GetBoardMfgDate

`func (o *IpmiFruInfo) GetBoardMfgDate() string`

GetBoardMfgDate returns the BoardMfgDate field if non-nil, zero value otherwise.

### GetBoardMfgDateOk

`func (o *IpmiFruInfo) GetBoardMfgDateOk() (*string, bool)`

GetBoardMfgDateOk returns a tuple with the BoardMfgDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardMfgDate

`func (o *IpmiFruInfo) SetBoardMfgDate(v string)`

SetBoardMfgDate sets BoardMfgDate field to given value.

### HasBoardMfgDate

`func (o *IpmiFruInfo) HasBoardMfgDate() bool`

HasBoardMfgDate returns a boolean if a field has been set.

### SetBoardMfgDateNil

`func (o *IpmiFruInfo) SetBoardMfgDateNil(b bool)`

 SetBoardMfgDateNil sets the value for BoardMfgDate to be an explicit nil

### UnsetBoardMfgDate
`func (o *IpmiFruInfo) UnsetBoardMfgDate()`

UnsetBoardMfgDate ensures that no value is present for BoardMfgDate, not even an explicit nil
### GetBoardPN

`func (o *IpmiFruInfo) GetBoardPN() string`

GetBoardPN returns the BoardPN field if non-nil, zero value otherwise.

### GetBoardPNOk

`func (o *IpmiFruInfo) GetBoardPNOk() (*string, bool)`

GetBoardPNOk returns a tuple with the BoardPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardPN

`func (o *IpmiFruInfo) SetBoardPN(v string)`

SetBoardPN sets BoardPN field to given value.

### HasBoardPN

`func (o *IpmiFruInfo) HasBoardPN() bool`

HasBoardPN returns a boolean if a field has been set.

### SetBoardPNNil

`func (o *IpmiFruInfo) SetBoardPNNil(b bool)`

 SetBoardPNNil sets the value for BoardPN to be an explicit nil

### UnsetBoardPN
`func (o *IpmiFruInfo) UnsetBoardPN()`

UnsetBoardPN ensures that no value is present for BoardPN, not even an explicit nil
### GetBoardProduct

`func (o *IpmiFruInfo) GetBoardProduct() string`

GetBoardProduct returns the BoardProduct field if non-nil, zero value otherwise.

### GetBoardProductOk

`func (o *IpmiFruInfo) GetBoardProductOk() (*string, bool)`

GetBoardProductOk returns a tuple with the BoardProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardProduct

`func (o *IpmiFruInfo) SetBoardProduct(v string)`

SetBoardProduct sets BoardProduct field to given value.

### HasBoardProduct

`func (o *IpmiFruInfo) HasBoardProduct() bool`

HasBoardProduct returns a boolean if a field has been set.

### SetBoardProductNil

`func (o *IpmiFruInfo) SetBoardProductNil(b bool)`

 SetBoardProductNil sets the value for BoardProduct to be an explicit nil

### UnsetBoardProduct
`func (o *IpmiFruInfo) UnsetBoardProduct()`

UnsetBoardProduct ensures that no value is present for BoardProduct, not even an explicit nil
### GetBoardSerial

`func (o *IpmiFruInfo) GetBoardSerial() string`

GetBoardSerial returns the BoardSerial field if non-nil, zero value otherwise.

### GetBoardSerialOk

`func (o *IpmiFruInfo) GetBoardSerialOk() (*string, bool)`

GetBoardSerialOk returns a tuple with the BoardSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardSerial

`func (o *IpmiFruInfo) SetBoardSerial(v string)`

SetBoardSerial sets BoardSerial field to given value.

### HasBoardSerial

`func (o *IpmiFruInfo) HasBoardSerial() bool`

HasBoardSerial returns a boolean if a field has been set.

### SetBoardSerialNil

`func (o *IpmiFruInfo) SetBoardSerialNil(b bool)`

 SetBoardSerialNil sets the value for BoardSerial to be an explicit nil

### UnsetBoardSerial
`func (o *IpmiFruInfo) UnsetBoardSerial()`

UnsetBoardSerial ensures that no value is present for BoardSerial, not even an explicit nil
### GetChassisExtra

`func (o *IpmiFruInfo) GetChassisExtra() string`

GetChassisExtra returns the ChassisExtra field if non-nil, zero value otherwise.

### GetChassisExtraOk

`func (o *IpmiFruInfo) GetChassisExtraOk() (*string, bool)`

GetChassisExtraOk returns a tuple with the ChassisExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisExtra

`func (o *IpmiFruInfo) SetChassisExtra(v string)`

SetChassisExtra sets ChassisExtra field to given value.

### HasChassisExtra

`func (o *IpmiFruInfo) HasChassisExtra() bool`

HasChassisExtra returns a boolean if a field has been set.

### SetChassisExtraNil

`func (o *IpmiFruInfo) SetChassisExtraNil(b bool)`

 SetChassisExtraNil sets the value for ChassisExtra to be an explicit nil

### UnsetChassisExtra
`func (o *IpmiFruInfo) UnsetChassisExtra()`

UnsetChassisExtra ensures that no value is present for ChassisExtra, not even an explicit nil
### GetChassisPN

`func (o *IpmiFruInfo) GetChassisPN() string`

GetChassisPN returns the ChassisPN field if non-nil, zero value otherwise.

### GetChassisPNOk

`func (o *IpmiFruInfo) GetChassisPNOk() (*string, bool)`

GetChassisPNOk returns a tuple with the ChassisPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisPN

`func (o *IpmiFruInfo) SetChassisPN(v string)`

SetChassisPN sets ChassisPN field to given value.

### HasChassisPN

`func (o *IpmiFruInfo) HasChassisPN() bool`

HasChassisPN returns a boolean if a field has been set.

### SetChassisPNNil

`func (o *IpmiFruInfo) SetChassisPNNil(b bool)`

 SetChassisPNNil sets the value for ChassisPN to be an explicit nil

### UnsetChassisPN
`func (o *IpmiFruInfo) UnsetChassisPN()`

UnsetChassisPN ensures that no value is present for ChassisPN, not even an explicit nil
### GetChassisSerial

`func (o *IpmiFruInfo) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *IpmiFruInfo) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *IpmiFruInfo) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *IpmiFruInfo) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *IpmiFruInfo) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *IpmiFruInfo) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetChassisType

`func (o *IpmiFruInfo) GetChassisType() string`

GetChassisType returns the ChassisType field if non-nil, zero value otherwise.

### GetChassisTypeOk

`func (o *IpmiFruInfo) GetChassisTypeOk() (*string, bool)`

GetChassisTypeOk returns a tuple with the ChassisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisType

`func (o *IpmiFruInfo) SetChassisType(v string)`

SetChassisType sets ChassisType field to given value.

### HasChassisType

`func (o *IpmiFruInfo) HasChassisType() bool`

HasChassisType returns a boolean if a field has been set.

### SetChassisTypeNil

`func (o *IpmiFruInfo) SetChassisTypeNil(b bool)`

 SetChassisTypeNil sets the value for ChassisType to be an explicit nil

### UnsetChassisType
`func (o *IpmiFruInfo) UnsetChassisType()`

UnsetChassisType ensures that no value is present for ChassisType, not even an explicit nil
### GetProductMfg

`func (o *IpmiFruInfo) GetProductMfg() string`

GetProductMfg returns the ProductMfg field if non-nil, zero value otherwise.

### GetProductMfgOk

`func (o *IpmiFruInfo) GetProductMfgOk() (*string, bool)`

GetProductMfgOk returns a tuple with the ProductMfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMfg

`func (o *IpmiFruInfo) SetProductMfg(v string)`

SetProductMfg sets ProductMfg field to given value.

### HasProductMfg

`func (o *IpmiFruInfo) HasProductMfg() bool`

HasProductMfg returns a boolean if a field has been set.

### SetProductMfgNil

`func (o *IpmiFruInfo) SetProductMfgNil(b bool)`

 SetProductMfgNil sets the value for ProductMfg to be an explicit nil

### UnsetProductMfg
`func (o *IpmiFruInfo) UnsetProductMfg()`

UnsetProductMfg ensures that no value is present for ProductMfg, not even an explicit nil
### GetProductName

`func (o *IpmiFruInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IpmiFruInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IpmiFruInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IpmiFruInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *IpmiFruInfo) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *IpmiFruInfo) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetProductPN

`func (o *IpmiFruInfo) GetProductPN() string`

GetProductPN returns the ProductPN field if non-nil, zero value otherwise.

### GetProductPNOk

`func (o *IpmiFruInfo) GetProductPNOk() (*string, bool)`

GetProductPNOk returns a tuple with the ProductPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPN

`func (o *IpmiFruInfo) SetProductPN(v string)`

SetProductPN sets ProductPN field to given value.

### HasProductPN

`func (o *IpmiFruInfo) HasProductPN() bool`

HasProductPN returns a boolean if a field has been set.

### SetProductPNNil

`func (o *IpmiFruInfo) SetProductPNNil(b bool)`

 SetProductPNNil sets the value for ProductPN to be an explicit nil

### UnsetProductPN
`func (o *IpmiFruInfo) UnsetProductPN()`

UnsetProductPN ensures that no value is present for ProductPN, not even an explicit nil
### GetProductSerial

`func (o *IpmiFruInfo) GetProductSerial() string`

GetProductSerial returns the ProductSerial field if non-nil, zero value otherwise.

### GetProductSerialOk

`func (o *IpmiFruInfo) GetProductSerialOk() (*string, bool)`

GetProductSerialOk returns a tuple with the ProductSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSerial

`func (o *IpmiFruInfo) SetProductSerial(v string)`

SetProductSerial sets ProductSerial field to given value.

### HasProductSerial

`func (o *IpmiFruInfo) HasProductSerial() bool`

HasProductSerial returns a boolean if a field has been set.

### SetProductSerialNil

`func (o *IpmiFruInfo) SetProductSerialNil(b bool)`

 SetProductSerialNil sets the value for ProductSerial to be an explicit nil

### UnsetProductSerial
`func (o *IpmiFruInfo) UnsetProductSerial()`

UnsetProductSerial ensures that no value is present for ProductSerial, not even an explicit nil
### GetProductVersion

`func (o *IpmiFruInfo) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *IpmiFruInfo) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *IpmiFruInfo) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *IpmiFruInfo) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### SetProductVersionNil

`func (o *IpmiFruInfo) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *IpmiFruInfo) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


