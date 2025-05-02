# V1TableRowCondition

TableRowCondition allows a row to be marked with additional information.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** | Human readable message indicating details about last transition. | [optional] 
**reason** | **str** | (brief) machine readable reason for the condition&#39;s last transition. | [optional] 
**status** | **str** | Status of the condition, one of True, False, Unknown. | [default to '']
**type** | **str** | Type of row condition. The only defined value is &#39;Completed&#39; indicating that the object this row represents has reached a completed state and may be given less visual priority than other rows. Clients are not required to honor any conditions but should be consistent where possible about handling the conditions. | [default to '']

## Example

```python
from tatope.models.v1_table_row_condition import V1TableRowCondition

# TODO update the JSON string below
json = "{}"
# create an instance of V1TableRowCondition from a JSON string
v1_table_row_condition_instance = V1TableRowCondition.from_json(json)
# print the JSON string representation of the object
print(V1TableRowCondition.to_json())

# convert the object into a dict
v1_table_row_condition_dict = v1_table_row_condition_instance.to_dict()
# create an instance of V1TableRowCondition from a dict
v1_table_row_condition_from_dict = V1TableRowCondition.from_dict(v1_table_row_condition_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


