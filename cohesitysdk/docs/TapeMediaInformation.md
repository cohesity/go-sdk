# TapeMediaInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Specifies a unique identifier for the media. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the offline media as recorded by the backup administrator using media management software. | [optional] 
**Online** | Pointer to **NullableBool** | Specifies a flag that indicates if the media is online or offline. Offline media must be manually loaded into the media library before a recovery can occur. | [optional] 

## Methods

### NewTapeMediaInformation

`func NewTapeMediaInformation() *TapeMediaInformation`

NewTapeMediaInformation instantiates a new TapeMediaInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTapeMediaInformationWithDefaults

`func NewTapeMediaInformationWithDefaults() *TapeMediaInformation`

NewTapeMediaInformationWithDefaults instantiates a new TapeMediaInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *TapeMediaInformation) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *TapeMediaInformation) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *TapeMediaInformation) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *TapeMediaInformation) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *TapeMediaInformation) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *TapeMediaInformation) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetLocation

`func (o *TapeMediaInformation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TapeMediaInformation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TapeMediaInformation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TapeMediaInformation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *TapeMediaInformation) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *TapeMediaInformation) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOnline

`func (o *TapeMediaInformation) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *TapeMediaInformation) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *TapeMediaInformation) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *TapeMediaInformation) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### SetOnlineNil

`func (o *TapeMediaInformation) SetOnlineNil(b bool)`

 SetOnlineNil sets the value for Online to be an explicit nil

### UnsetOnline
`func (o *TapeMediaInformation) UnsetOnline()`

UnsetOnline ensures that no value is present for Online, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


