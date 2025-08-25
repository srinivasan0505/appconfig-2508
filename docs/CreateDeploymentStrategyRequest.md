

# CreateDeploymentStrategyRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | A name for the deployment strategy. |  |
|**description** | **String** | A description of the deployment strategy. |  [optional] |
|**deploymentDurationInMinutes** | **Integer** | Total amount of time for a deployment to last. |  |
|**finalBakeTimeInMinutes** | **Integer** | Specifies the amount of time AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete. If an alarm is triggered during this time, AppConfig rolls back the deployment. You must configure permissions for AppConfig to roll back based on CloudWatch alarms. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/getting-started-with-appconfig-cloudwatch-alarms-permissions.html\&quot;&gt;Configuring permissions for rollback based on Amazon CloudWatch alarms&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;. |  [optional] |
|**growthFactor** | **Float** | The percentage of targets to receive a deployed configuration during each interval. |  |
|**growthType** | [**GrowthTypeEnum**](#GrowthTypeEnum) | &lt;p&gt;The algorithm used to define how percentage grows over time. AppConfig supports the following growth types:&lt;/p&gt; &lt;p&gt; &lt;b&gt;Linear&lt;/b&gt;: For this type, AppConfig processes the deployment by dividing the total number of targets by the value specified for &lt;code&gt;Step percentage&lt;/code&gt;. For example, a linear deployment that uses a &lt;code&gt;Step percentage&lt;/code&gt; of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.&lt;/p&gt; &lt;p&gt; &lt;b&gt;Exponential&lt;/b&gt;: For this type, AppConfig processes the deployment exponentially using the following formula: &lt;code&gt;G*(2^N)&lt;/code&gt;. In this formula, &lt;code&gt;G&lt;/code&gt; is the growth factor specified by the user and &lt;code&gt;N&lt;/code&gt; is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:&lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^0)&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^1)&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^2)&lt;/code&gt; &lt;/p&gt; &lt;p&gt;Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.&lt;/p&gt; |  [optional] |
|**replicateTo** | [**ReplicateToEnum**](#ReplicateToEnum) | Save the deployment strategy to a Systems Manager (SSM) document. |  [optional] |
|**tags** | **Map&lt;String, String&gt;** | Metadata to assign to the deployment strategy. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |  [optional] |



## Enum: GrowthTypeEnum

| Name | Value |
|---- | -----|
| LINEAR | &quot;LINEAR&quot; |
| EXPONENTIAL | &quot;EXPONENTIAL&quot; |



## Enum: ReplicateToEnum

| Name | Value |
|---- | -----|
| NONE | &quot;NONE&quot; |
| SSM_DOCUMENT | &quot;SSM_DOCUMENT&quot; |



