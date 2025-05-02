# TatOpeExStatus

TatOpeExStatus defines the observed state of TatOpeEx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**state** | **str** | INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \&quot;make\&quot; to regenerate code after modifying this file | [optional] 

## Example

```python
from tatope.models.tat_ope_ex_status import TatOpeExStatus

# TODO update the JSON string below
json = "{}"
# create an instance of TatOpeExStatus from a JSON string
tat_ope_ex_status_instance = TatOpeExStatus.from_json(json)
# print the JSON string representation of the object
print(TatOpeExStatus.to_json())

# convert the object into a dict
tat_ope_ex_status_dict = tat_ope_ex_status_instance.to_dict()
# create an instance of TatOpeExStatus from a dict
tat_ope_ex_status_from_dict = TatOpeExStatus.from_dict(tat_ope_ex_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


