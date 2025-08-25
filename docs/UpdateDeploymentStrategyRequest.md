

# UpdateDeploymentStrategyRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**description** | **String** | A description of the deployment strategy. |  [optional] |
|**deploymentDurationInMinutes** | **Integer** | Total amount of time for a deployment to last. |  [optional] |
|**finalBakeTimeInMinutes** | **Integer** | The amount of time that AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic rollback. |  [optional] |
|**growthFactor** | **Float** | The percentage of targets to receive a deployed configuration during each interval. |  [optional] |
|**growthType** | [**GrowthTypeEnum**](#GrowthTypeEnum) | &lt;p&gt;The algorithm used to define how percentage grows over time. AppConfig supports the following growth types:&lt;/p&gt; &lt;p&gt; &lt;b&gt;Linear&lt;/b&gt;: For this type, AppConfig processes the deployment by increments of the growth factor evenly distributed over the deployment time. For example, a linear deployment that uses a growth factor of 20 initially makes the configuration available to 20 percent of the targets. After 1/5th of the deployment time has passed, the system updates the percentage to 40 percent. This continues until 100% of the targets are set to receive the deployed configuration.&lt;/p&gt; &lt;p&gt; &lt;b&gt;Exponential&lt;/b&gt;: For this type, AppConfig processes the deployment exponentially using the following formula: &lt;code&gt;G*(2^N)&lt;/code&gt;. In this formula, &lt;code&gt;G&lt;/code&gt; is the growth factor specified by the user and &lt;code&gt;N&lt;/code&gt; is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:&lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^0)&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^1)&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;2*(2^2)&lt;/code&gt; &lt;/p&gt; &lt;p&gt;Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.&lt;/p&gt; |  [optional] |



## Enum: GrowthTypeEnum

| Name | Value |
|---- | -----|
| LINEAR | &quot;LINEAR&quot; |
| EXPONENTIAL | &quot;EXPONENTIAL&quot; |



