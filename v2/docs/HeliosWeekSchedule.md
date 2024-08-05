# HeliosWeekSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **[]string** | Specifies a list of days of the week when to start Protection Group Runs. &lt;br&gt; Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: &lt;br&gt;  unit: &#39;Weeks&#39; &lt;br&gt;  dayOfWeek: [&#39;Monday&#39;,&#39;Tuesday&#39;] | [optional] 

## Methods

### NewHeliosWeekSchedule

`func NewHeliosWeekSchedule() *HeliosWeekSchedule`

NewHeliosWeekSchedule instantiates a new HeliosWeekSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosWeekScheduleWithDefaults

`func NewHeliosWeekScheduleWithDefaults() *HeliosWeekSchedule`

NewHeliosWeekScheduleWithDefaults instantiates a new HeliosWeekSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *HeliosWeekSchedule) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *HeliosWeekSchedule) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *HeliosWeekSchedule) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *HeliosWeekSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *HeliosWeekSchedule) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *HeliosWeekSchedule) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


