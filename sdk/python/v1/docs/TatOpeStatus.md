# TatOpeStatus

TatOpeStatus defines the observed state of TatOpe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**state** | **str** |  | [optional] 

## Example

```python
from tatope.models.tat_ope_status import TatOpeStatus

# TODO update the JSON string below
json = "{}"
# create an instance of TatOpeStatus from a JSON string
tat_ope_status_instance = TatOpeStatus.from_json(json)
# print the JSON string representation of the object
print(TatOpeStatus.to_json())

# convert the object into a dict
tat_ope_status_dict = tat_ope_status_instance.to_dict()
# create an instance of TatOpeStatus from a dict
tat_ope_status_from_dict = TatOpeStatus.from_dict(tat_ope_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


