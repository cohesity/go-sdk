/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the IpmiFruInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpmiFruInfo{}

// IpmiFruInfo Specifies the fru for the ipmi.
type IpmiFruInfo struct {
	// Specifies the IPMI Field Replaceable Unit(FRU) ID for given node.
	Id NullableInt32 `json:"Id,omitempty"`
	// Specifies the board manufacturer for given node.
	BoardMfg NullableString `json:"boardMfg,omitempty"`
	// Specifies the board manufacturing date for given node.
	BoardMfgDate NullableString `json:"boardMfgDate,omitempty"`
	// Specifies the board part number for given node.
	BoardPN NullableString `json:"boardPN,omitempty"`
	// Specifies the board product name for given node.
	BoardProduct NullableString `json:"boardProduct,omitempty"`
	// Specifies the board serial number for given node.
	BoardSerial NullableString `json:"boardSerial,omitempty"`
	// Specifies the information about chassis extras provided for given node.
	ChassisExtra NullableString `json:"chassisExtra,omitempty"`
	// Specifies the chassis part number for given node.
	ChassisPN NullableString `json:"chassisPN,omitempty"`
	// Specifies the chassis serial number for given node.
	ChassisSerial NullableString `json:"chassisSerial,omitempty"`
	// Specifies the type of chassis for given node.
	ChassisType NullableString `json:"chassisType,omitempty"`
	// Specifies the product manufacturer for given node.
	ProductMfg NullableString `json:"productMfg,omitempty"`
	// Specifies the product name for given node.
	ProductName NullableString `json:"productName,omitempty"`
	// Specifies the product part number for given node.
	ProductPN NullableString `json:"productPN,omitempty"`
	// Specifies the product serial number for given node.
	ProductSerial NullableString `json:"productSerial,omitempty"`
	// Specifies the product version for given node.
	ProductVersion NullableString `json:"productVersion,omitempty"`
}

// NewIpmiFruInfo instantiates a new IpmiFruInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmiFruInfo() *IpmiFruInfo {
	this := IpmiFruInfo{}
	return &this
}

// NewIpmiFruInfoWithDefaults instantiates a new IpmiFruInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmiFruInfoWithDefaults() *IpmiFruInfo {
	this := IpmiFruInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IpmiFruInfo) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IpmiFruInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IpmiFruInfo) UnsetId() {
	o.Id.Unset()
}

// GetBoardMfg returns the BoardMfg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetBoardMfg() string {
	if o == nil || IsNil(o.BoardMfg.Get()) {
		var ret string
		return ret
	}
	return *o.BoardMfg.Get()
}

// GetBoardMfgOk returns a tuple with the BoardMfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetBoardMfgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardMfg.Get(), o.BoardMfg.IsSet()
}

// HasBoardMfg returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasBoardMfg() bool {
	if o != nil && o.BoardMfg.IsSet() {
		return true
	}

	return false
}

// SetBoardMfg gets a reference to the given NullableString and assigns it to the BoardMfg field.
func (o *IpmiFruInfo) SetBoardMfg(v string) {
	o.BoardMfg.Set(&v)
}
// SetBoardMfgNil sets the value for BoardMfg to be an explicit nil
func (o *IpmiFruInfo) SetBoardMfgNil() {
	o.BoardMfg.Set(nil)
}

// UnsetBoardMfg ensures that no value is present for BoardMfg, not even an explicit nil
func (o *IpmiFruInfo) UnsetBoardMfg() {
	o.BoardMfg.Unset()
}

// GetBoardMfgDate returns the BoardMfgDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetBoardMfgDate() string {
	if o == nil || IsNil(o.BoardMfgDate.Get()) {
		var ret string
		return ret
	}
	return *o.BoardMfgDate.Get()
}

// GetBoardMfgDateOk returns a tuple with the BoardMfgDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetBoardMfgDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardMfgDate.Get(), o.BoardMfgDate.IsSet()
}

// HasBoardMfgDate returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasBoardMfgDate() bool {
	if o != nil && o.BoardMfgDate.IsSet() {
		return true
	}

	return false
}

// SetBoardMfgDate gets a reference to the given NullableString and assigns it to the BoardMfgDate field.
func (o *IpmiFruInfo) SetBoardMfgDate(v string) {
	o.BoardMfgDate.Set(&v)
}
// SetBoardMfgDateNil sets the value for BoardMfgDate to be an explicit nil
func (o *IpmiFruInfo) SetBoardMfgDateNil() {
	o.BoardMfgDate.Set(nil)
}

// UnsetBoardMfgDate ensures that no value is present for BoardMfgDate, not even an explicit nil
func (o *IpmiFruInfo) UnsetBoardMfgDate() {
	o.BoardMfgDate.Unset()
}

// GetBoardPN returns the BoardPN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetBoardPN() string {
	if o == nil || IsNil(o.BoardPN.Get()) {
		var ret string
		return ret
	}
	return *o.BoardPN.Get()
}

// GetBoardPNOk returns a tuple with the BoardPN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetBoardPNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardPN.Get(), o.BoardPN.IsSet()
}

// HasBoardPN returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasBoardPN() bool {
	if o != nil && o.BoardPN.IsSet() {
		return true
	}

	return false
}

// SetBoardPN gets a reference to the given NullableString and assigns it to the BoardPN field.
func (o *IpmiFruInfo) SetBoardPN(v string) {
	o.BoardPN.Set(&v)
}
// SetBoardPNNil sets the value for BoardPN to be an explicit nil
func (o *IpmiFruInfo) SetBoardPNNil() {
	o.BoardPN.Set(nil)
}

// UnsetBoardPN ensures that no value is present for BoardPN, not even an explicit nil
func (o *IpmiFruInfo) UnsetBoardPN() {
	o.BoardPN.Unset()
}

// GetBoardProduct returns the BoardProduct field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetBoardProduct() string {
	if o == nil || IsNil(o.BoardProduct.Get()) {
		var ret string
		return ret
	}
	return *o.BoardProduct.Get()
}

// GetBoardProductOk returns a tuple with the BoardProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetBoardProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardProduct.Get(), o.BoardProduct.IsSet()
}

// HasBoardProduct returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasBoardProduct() bool {
	if o != nil && o.BoardProduct.IsSet() {
		return true
	}

	return false
}

// SetBoardProduct gets a reference to the given NullableString and assigns it to the BoardProduct field.
func (o *IpmiFruInfo) SetBoardProduct(v string) {
	o.BoardProduct.Set(&v)
}
// SetBoardProductNil sets the value for BoardProduct to be an explicit nil
func (o *IpmiFruInfo) SetBoardProductNil() {
	o.BoardProduct.Set(nil)
}

// UnsetBoardProduct ensures that no value is present for BoardProduct, not even an explicit nil
func (o *IpmiFruInfo) UnsetBoardProduct() {
	o.BoardProduct.Unset()
}

// GetBoardSerial returns the BoardSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetBoardSerial() string {
	if o == nil || IsNil(o.BoardSerial.Get()) {
		var ret string
		return ret
	}
	return *o.BoardSerial.Get()
}

// GetBoardSerialOk returns a tuple with the BoardSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetBoardSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardSerial.Get(), o.BoardSerial.IsSet()
}

// HasBoardSerial returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasBoardSerial() bool {
	if o != nil && o.BoardSerial.IsSet() {
		return true
	}

	return false
}

// SetBoardSerial gets a reference to the given NullableString and assigns it to the BoardSerial field.
func (o *IpmiFruInfo) SetBoardSerial(v string) {
	o.BoardSerial.Set(&v)
}
// SetBoardSerialNil sets the value for BoardSerial to be an explicit nil
func (o *IpmiFruInfo) SetBoardSerialNil() {
	o.BoardSerial.Set(nil)
}

// UnsetBoardSerial ensures that no value is present for BoardSerial, not even an explicit nil
func (o *IpmiFruInfo) UnsetBoardSerial() {
	o.BoardSerial.Unset()
}

// GetChassisExtra returns the ChassisExtra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetChassisExtra() string {
	if o == nil || IsNil(o.ChassisExtra.Get()) {
		var ret string
		return ret
	}
	return *o.ChassisExtra.Get()
}

// GetChassisExtraOk returns a tuple with the ChassisExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetChassisExtraOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChassisExtra.Get(), o.ChassisExtra.IsSet()
}

// HasChassisExtra returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasChassisExtra() bool {
	if o != nil && o.ChassisExtra.IsSet() {
		return true
	}

	return false
}

// SetChassisExtra gets a reference to the given NullableString and assigns it to the ChassisExtra field.
func (o *IpmiFruInfo) SetChassisExtra(v string) {
	o.ChassisExtra.Set(&v)
}
// SetChassisExtraNil sets the value for ChassisExtra to be an explicit nil
func (o *IpmiFruInfo) SetChassisExtraNil() {
	o.ChassisExtra.Set(nil)
}

// UnsetChassisExtra ensures that no value is present for ChassisExtra, not even an explicit nil
func (o *IpmiFruInfo) UnsetChassisExtra() {
	o.ChassisExtra.Unset()
}

// GetChassisPN returns the ChassisPN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetChassisPN() string {
	if o == nil || IsNil(o.ChassisPN.Get()) {
		var ret string
		return ret
	}
	return *o.ChassisPN.Get()
}

// GetChassisPNOk returns a tuple with the ChassisPN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetChassisPNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChassisPN.Get(), o.ChassisPN.IsSet()
}

// HasChassisPN returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasChassisPN() bool {
	if o != nil && o.ChassisPN.IsSet() {
		return true
	}

	return false
}

// SetChassisPN gets a reference to the given NullableString and assigns it to the ChassisPN field.
func (o *IpmiFruInfo) SetChassisPN(v string) {
	o.ChassisPN.Set(&v)
}
// SetChassisPNNil sets the value for ChassisPN to be an explicit nil
func (o *IpmiFruInfo) SetChassisPNNil() {
	o.ChassisPN.Set(nil)
}

// UnsetChassisPN ensures that no value is present for ChassisPN, not even an explicit nil
func (o *IpmiFruInfo) UnsetChassisPN() {
	o.ChassisPN.Unset()
}

// GetChassisSerial returns the ChassisSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetChassisSerial() string {
	if o == nil || IsNil(o.ChassisSerial.Get()) {
		var ret string
		return ret
	}
	return *o.ChassisSerial.Get()
}

// GetChassisSerialOk returns a tuple with the ChassisSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetChassisSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChassisSerial.Get(), o.ChassisSerial.IsSet()
}

// HasChassisSerial returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasChassisSerial() bool {
	if o != nil && o.ChassisSerial.IsSet() {
		return true
	}

	return false
}

// SetChassisSerial gets a reference to the given NullableString and assigns it to the ChassisSerial field.
func (o *IpmiFruInfo) SetChassisSerial(v string) {
	o.ChassisSerial.Set(&v)
}
// SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil
func (o *IpmiFruInfo) SetChassisSerialNil() {
	o.ChassisSerial.Set(nil)
}

// UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
func (o *IpmiFruInfo) UnsetChassisSerial() {
	o.ChassisSerial.Unset()
}

// GetChassisType returns the ChassisType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetChassisType() string {
	if o == nil || IsNil(o.ChassisType.Get()) {
		var ret string
		return ret
	}
	return *o.ChassisType.Get()
}

// GetChassisTypeOk returns a tuple with the ChassisType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetChassisTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChassisType.Get(), o.ChassisType.IsSet()
}

// HasChassisType returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasChassisType() bool {
	if o != nil && o.ChassisType.IsSet() {
		return true
	}

	return false
}

// SetChassisType gets a reference to the given NullableString and assigns it to the ChassisType field.
func (o *IpmiFruInfo) SetChassisType(v string) {
	o.ChassisType.Set(&v)
}
// SetChassisTypeNil sets the value for ChassisType to be an explicit nil
func (o *IpmiFruInfo) SetChassisTypeNil() {
	o.ChassisType.Set(nil)
}

// UnsetChassisType ensures that no value is present for ChassisType, not even an explicit nil
func (o *IpmiFruInfo) UnsetChassisType() {
	o.ChassisType.Unset()
}

// GetProductMfg returns the ProductMfg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetProductMfg() string {
	if o == nil || IsNil(o.ProductMfg.Get()) {
		var ret string
		return ret
	}
	return *o.ProductMfg.Get()
}

// GetProductMfgOk returns a tuple with the ProductMfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetProductMfgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductMfg.Get(), o.ProductMfg.IsSet()
}

// HasProductMfg returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasProductMfg() bool {
	if o != nil && o.ProductMfg.IsSet() {
		return true
	}

	return false
}

// SetProductMfg gets a reference to the given NullableString and assigns it to the ProductMfg field.
func (o *IpmiFruInfo) SetProductMfg(v string) {
	o.ProductMfg.Set(&v)
}
// SetProductMfgNil sets the value for ProductMfg to be an explicit nil
func (o *IpmiFruInfo) SetProductMfgNil() {
	o.ProductMfg.Set(nil)
}

// UnsetProductMfg ensures that no value is present for ProductMfg, not even an explicit nil
func (o *IpmiFruInfo) UnsetProductMfg() {
	o.ProductMfg.Unset()
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *IpmiFruInfo) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *IpmiFruInfo) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *IpmiFruInfo) UnsetProductName() {
	o.ProductName.Unset()
}

// GetProductPN returns the ProductPN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetProductPN() string {
	if o == nil || IsNil(o.ProductPN.Get()) {
		var ret string
		return ret
	}
	return *o.ProductPN.Get()
}

// GetProductPNOk returns a tuple with the ProductPN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetProductPNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductPN.Get(), o.ProductPN.IsSet()
}

// HasProductPN returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasProductPN() bool {
	if o != nil && o.ProductPN.IsSet() {
		return true
	}

	return false
}

// SetProductPN gets a reference to the given NullableString and assigns it to the ProductPN field.
func (o *IpmiFruInfo) SetProductPN(v string) {
	o.ProductPN.Set(&v)
}
// SetProductPNNil sets the value for ProductPN to be an explicit nil
func (o *IpmiFruInfo) SetProductPNNil() {
	o.ProductPN.Set(nil)
}

// UnsetProductPN ensures that no value is present for ProductPN, not even an explicit nil
func (o *IpmiFruInfo) UnsetProductPN() {
	o.ProductPN.Unset()
}

// GetProductSerial returns the ProductSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetProductSerial() string {
	if o == nil || IsNil(o.ProductSerial.Get()) {
		var ret string
		return ret
	}
	return *o.ProductSerial.Get()
}

// GetProductSerialOk returns a tuple with the ProductSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetProductSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductSerial.Get(), o.ProductSerial.IsSet()
}

// HasProductSerial returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasProductSerial() bool {
	if o != nil && o.ProductSerial.IsSet() {
		return true
	}

	return false
}

// SetProductSerial gets a reference to the given NullableString and assigns it to the ProductSerial field.
func (o *IpmiFruInfo) SetProductSerial(v string) {
	o.ProductSerial.Set(&v)
}
// SetProductSerialNil sets the value for ProductSerial to be an explicit nil
func (o *IpmiFruInfo) SetProductSerialNil() {
	o.ProductSerial.Set(nil)
}

// UnsetProductSerial ensures that no value is present for ProductSerial, not even an explicit nil
func (o *IpmiFruInfo) UnsetProductSerial() {
	o.ProductSerial.Unset()
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiFruInfo) GetProductVersion() string {
	if o == nil || IsNil(o.ProductVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProductVersion.Get()
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiFruInfo) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductVersion.Get(), o.ProductVersion.IsSet()
}

// HasProductVersion returns a boolean if a field has been set.
func (o *IpmiFruInfo) HasProductVersion() bool {
	if o != nil && o.ProductVersion.IsSet() {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given NullableString and assigns it to the ProductVersion field.
func (o *IpmiFruInfo) SetProductVersion(v string) {
	o.ProductVersion.Set(&v)
}
// SetProductVersionNil sets the value for ProductVersion to be an explicit nil
func (o *IpmiFruInfo) SetProductVersionNil() {
	o.ProductVersion.Set(nil)
}

// UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
func (o *IpmiFruInfo) UnsetProductVersion() {
	o.ProductVersion.Unset()
}

func (o IpmiFruInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpmiFruInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.BoardMfg.IsSet() {
		toSerialize["boardMfg"] = o.BoardMfg.Get()
	}
	if o.BoardMfgDate.IsSet() {
		toSerialize["boardMfgDate"] = o.BoardMfgDate.Get()
	}
	if o.BoardPN.IsSet() {
		toSerialize["boardPN"] = o.BoardPN.Get()
	}
	if o.BoardProduct.IsSet() {
		toSerialize["boardProduct"] = o.BoardProduct.Get()
	}
	if o.BoardSerial.IsSet() {
		toSerialize["boardSerial"] = o.BoardSerial.Get()
	}
	if o.ChassisExtra.IsSet() {
		toSerialize["chassisExtra"] = o.ChassisExtra.Get()
	}
	if o.ChassisPN.IsSet() {
		toSerialize["chassisPN"] = o.ChassisPN.Get()
	}
	if o.ChassisSerial.IsSet() {
		toSerialize["chassisSerial"] = o.ChassisSerial.Get()
	}
	if o.ChassisType.IsSet() {
		toSerialize["chassisType"] = o.ChassisType.Get()
	}
	if o.ProductMfg.IsSet() {
		toSerialize["productMfg"] = o.ProductMfg.Get()
	}
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.ProductPN.IsSet() {
		toSerialize["productPN"] = o.ProductPN.Get()
	}
	if o.ProductSerial.IsSet() {
		toSerialize["productSerial"] = o.ProductSerial.Get()
	}
	if o.ProductVersion.IsSet() {
		toSerialize["productVersion"] = o.ProductVersion.Get()
	}
	return toSerialize, nil
}

type NullableIpmiFruInfo struct {
	value *IpmiFruInfo
	isSet bool
}

func (v NullableIpmiFruInfo) Get() *IpmiFruInfo {
	return v.value
}

func (v *NullableIpmiFruInfo) Set(val *IpmiFruInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmiFruInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmiFruInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmiFruInfo(val *IpmiFruInfo) *NullableIpmiFruInfo {
	return &NullableIpmiFruInfo{value: val, isSet: true}
}

func (v NullableIpmiFruInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmiFruInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


