/**
 * MCP Server function for No description available
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Get_ExtensionassociationsMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_ExtensionassociationsHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("resource_identifier")) {
            queryParams.add("resource_identifier=" + args.get("resource_identifier"));
        }
        if (args.containsKey("extension_identifier")) {
            queryParams.add("extension_identifier=" + args.get("extension_identifier"));
        }
        if (args.containsKey("next_token")) {
            queryParams.add("next_token=" + args.get("next_token"));
        }
        if (args.containsKey("MaxResults")) {
            queryParams.add("MaxResults=" + args.get("MaxResults"));
        }
        if (args.containsKey("NextToken")) {
            queryParams.add("NextToken=" + args.get("NextToken"));
        }
        if (args.containsKey("extension_version_number")) {
            queryParams.add("extension_version_number=" + args.get("extension_version_number"));
        }
        if (args.containsKey("max_results")) {
            queryParams.add("max_results=" + args.get("max_results"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_extensionassociations" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createGet_ExtensionassociationsTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> resource_identifierProperty = new HashMap<>();
        resource_identifierProperty.put("type", "string");
        resource_identifierProperty.put("required", false);
        resource_identifierProperty.put("description", "The ARN of an application, configuration profile, or environment.");
        properties.put("resource_identifier", resource_identifierProperty);
        Map<String, Object> extension_identifierProperty = new HashMap<>();
        extension_identifierProperty.put("type", "string");
        extension_identifierProperty.put("required", false);
        extension_identifierProperty.put("description", "The name, the ID, or the Amazon Resource Name (ARN) of the extension.");
        properties.put("extension_identifier", extension_identifierProperty);
        Map<String, Object> next_tokenProperty = new HashMap<>();
        next_tokenProperty.put("type", "string");
        next_tokenProperty.put("required", false);
        next_tokenProperty.put("description", "A token to start the list. Use this token to get the next set of results or pass null to get the first set of results.");
        properties.put("next_token", next_tokenProperty);
        Map<String, Object> MaxResultsProperty = new HashMap<>();
        MaxResultsProperty.put("type", "string");
        MaxResultsProperty.put("required", false);
        MaxResultsProperty.put("description", "Pagination limit");
        properties.put("MaxResults", MaxResultsProperty);
        Map<String, Object> NextTokenProperty = new HashMap<>();
        NextTokenProperty.put("type", "string");
        NextTokenProperty.put("required", false);
        NextTokenProperty.put("description", "Pagination token");
        properties.put("NextToken", NextTokenProperty);
        Map<String, Object> extension_version_numberProperty = new HashMap<>();
        extension_version_numberProperty.put("type", "string");
        extension_version_numberProperty.put("required", false);
        extension_version_numberProperty.put("description", "The version number for the extension defined in the association.");
        properties.put("extension_version_number", extension_version_numberProperty);
        Map<String, Object> max_resultsProperty = new HashMap<>();
        max_resultsProperty.put("type", "string");
        max_resultsProperty.put("required", false);
        max_resultsProperty.put("description", "The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.");
        properties.put("max_results", max_resultsProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_extensionassociations",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_ExtensionassociationsHandler(config));
    }
    
}