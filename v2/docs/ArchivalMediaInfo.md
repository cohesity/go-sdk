# ArchivalMediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Specifies a unique identifier for the media. | [optional] [readonly] 
**IsOnline** | Pointer to **NullableBool** | Specifies a flag that indicates if the media is online or offline. Offline media must be manually loaded into the media library before a recovery can occur. | [optional] 
**Location** | Pointer to **NullableString** | Specifies the location of the offline media as recorded by the backup administrator using media management software. | [optional] [readonly] 

## Methods

### NewArchivalMediaInfo

`func NewArchivalMediaInfo() *ArchivalMediaInfo`

NewArchivalMediaInfo instantiates a new ArchivalMediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalMediaInfoWithDefaults

`func NewArchivalMediaInfoWithDefaults() *ArchivalMediaInfo`

NewArchivalMediaInfoWithDefaults instantiates a new ArchivalMediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ArchivalMediaInfo) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ArchivalMediaInfo) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ArchivalMediaInfo) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ArchivalMediaInfo) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ArchivalMediaInfo) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ArchivalMediaInfo) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetIsOnline

`func (o *ArchivalMediaInfo) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *ArchivalMediaInfo) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *ArchivalMediaInfo) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *ArchivalMediaInfo) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *ArchivalMediaInfo) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *ArchivalMediaInfo) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetLocation

`func (o *ArchivalMediaInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ArchivalMediaInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ArchivalMediaInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ArchivalMediaInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ArchivalMediaInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ArchivalMediaInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


