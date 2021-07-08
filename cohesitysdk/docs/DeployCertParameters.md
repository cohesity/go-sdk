# DeployCertParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertFileName** | Pointer to **NullableString** | Specifies the filename of the certificate. | [optional] 
**HostsInfoList** | Pointer to [**[]HostInfo**](HostInfo.md) | Specifies the list of all hosts on which the certificate is to be deployed. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the host such as &#39;kSapHana&#39;, &#39;kSapOracle&#39;, etc. Specifies the host type of host for generating and deploying a Certificate. &#39;kOther&#39; indicates it is any of the other hosts. &#39;kSapOracle&#39; indicates it is a SAP Oracle host. &#39;kSapHana&#39; indicates it is a SAP HANA host. | [optional] 
**ValidDays** | Pointer to **NullableInt64** | Specifies the number of days after which the certificate will expire. The user has to input the number of days (from the current date) till when the certificate is valid. | [optional] 

## Methods

### NewDeployCertParameters

`func NewDeployCertParameters() *DeployCertParameters`

NewDeployCertParameters instantiates a new DeployCertParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployCertParametersWithDefaults

`func NewDeployCertParametersWithDefaults() *DeployCertParameters`

NewDeployCertParametersWithDefaults instantiates a new DeployCertParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertFileName

`func (o *DeployCertParameters) GetCertFileName() string`

GetCertFileName returns the CertFileName field if non-nil, zero value otherwise.

### GetCertFileNameOk

`func (o *DeployCertParameters) GetCertFileNameOk() (*string, bool)`

GetCertFileNameOk returns a tuple with the CertFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileName

`func (o *DeployCertParameters) SetCertFileName(v string)`

SetCertFileName sets CertFileName field to given value.

### HasCertFileName

`func (o *DeployCertParameters) HasCertFileName() bool`

HasCertFileName returns a boolean if a field has been set.

### SetCertFileNameNil

`func (o *DeployCertParameters) SetCertFileNameNil(b bool)`

 SetCertFileNameNil sets the value for CertFileName to be an explicit nil

### UnsetCertFileName
`func (o *DeployCertParameters) UnsetCertFileName()`

UnsetCertFileName ensures that no value is present for CertFileName, not even an explicit nil
### GetHostsInfoList

`func (o *DeployCertParameters) GetHostsInfoList() []HostInfo`

GetHostsInfoList returns the HostsInfoList field if non-nil, zero value otherwise.

### GetHostsInfoListOk

`func (o *DeployCertParameters) GetHostsInfoListOk() (*[]HostInfo, bool)`

GetHostsInfoListOk returns a tuple with the HostsInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsInfoList

`func (o *DeployCertParameters) SetHostsInfoList(v []HostInfo)`

SetHostsInfoList sets HostsInfoList field to given value.

### HasHostsInfoList

`func (o *DeployCertParameters) HasHostsInfoList() bool`

HasHostsInfoList returns a boolean if a field has been set.

### SetHostsInfoListNil

`func (o *DeployCertParameters) SetHostsInfoListNil(b bool)`

 SetHostsInfoListNil sets the value for HostsInfoList to be an explicit nil

### UnsetHostsInfoList
`func (o *DeployCertParameters) UnsetHostsInfoList()`

UnsetHostsInfoList ensures that no value is present for HostsInfoList, not even an explicit nil
### GetType

`func (o *DeployCertParameters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeployCertParameters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeployCertParameters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeployCertParameters) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeployCertParameters) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeployCertParameters) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValidDays

`func (o *DeployCertParameters) GetValidDays() int64`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *DeployCertParameters) GetValidDaysOk() (*int64, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *DeployCertParameters) SetValidDays(v int64)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *DeployCertParameters) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### SetValidDaysNil

`func (o *DeployCertParameters) SetValidDaysNil(b bool)`

 SetValidDaysNil sets the value for ValidDays to be an explicit nil

### UnsetValidDays
`func (o *DeployCertParameters) UnsetValidDays()`

UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


