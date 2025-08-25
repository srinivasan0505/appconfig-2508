

# UpdateConfigurationProfileRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | The name of the configuration profile. |  [optional] |
|**description** | **String** | A description of the configuration profile. |  [optional] |
|**retrievalRoleArn** | **String** | The ARN of an IAM role with permission to access the configuration at the specified &lt;code&gt;LocationUri&lt;/code&gt;. |  [optional] |
|**validators** | [**List&lt;Validator&gt;**](Validator.md) | A list of methods for validating the configuration. |  [optional] |



