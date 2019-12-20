package azurerm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/provider"
)

func Provider() terraform.ResourceProvider {
	dataSources := map[string]*schema.Resource{
		"azurerm_app_service_plan":                          dataSourceAppServicePlan(),
		"azurerm_app_service_certificate":                   dataSourceAppServiceCertificate(),
		"azurerm_app_service":                               dataSourceArmAppService(),
		"azurerm_app_service_certificate_order":             dataSourceArmAppServiceCertificateOrder(),
		"azurerm_application_security_group":                dataSourceArmApplicationSecurityGroup(),
		"azurerm_builtin_role_definition":                   dataSourceArmBuiltInRoleDefinition(),
		"azurerm_client_config":                             dataSourceArmClientConfig(),
		"azurerm_express_route_circuit":                     dataSourceArmExpressRouteCircuit(),
		"azurerm_firewall":                                  dataSourceArmFirewall(),
		"azurerm_image":                                     dataSourceArmImage(),
		"azurerm_maps_account":                              dataSourceArmMapsAccount(),
		"azurerm_key_vault_access_policy":                   dataSourceArmKeyVaultAccessPolicy(),
		"azurerm_key_vault_key":                             dataSourceArmKeyVaultKey(),
		"azurerm_key_vault_secret":                          dataSourceArmKeyVaultSecret(),
		"azurerm_key_vault":                                 dataSourceArmKeyVault(),
		"azurerm_lb":                                        dataSourceArmLoadBalancer(),
		"azurerm_lb_backend_address_pool":                   dataSourceArmLoadBalancerBackendAddressPool(),
		"azurerm_monitor_action_group":                      dataSourceArmMonitorActionGroup(),
		"azurerm_monitor_diagnostic_categories":             dataSourceArmMonitorDiagnosticCategories(),
		"azurerm_monitor_log_profile":                       dataSourceArmMonitorLogProfile(),
		"azurerm_mssql_elasticpool":                         dataSourceArmMsSqlElasticpool(),
		"azurerm_nat_gateway":                               dataSourceArmNatGateway(),
		"azurerm_netapp_account":                            dataSourceArmNetAppAccount(),
		"azurerm_netapp_pool":                               dataSourceArmNetAppPool(),
		"azurerm_netapp_volume":                             dataSourceArmNetAppVolume(),
		"azurerm_network_ddos_protection_plan":              dataSourceNetworkDDoSProtectionPlan(),
		"azurerm_network_interface":                         dataSourceArmNetworkInterface(),
		"azurerm_network_security_group":                    dataSourceArmNetworkSecurityGroup(),
		"azurerm_network_watcher":                           dataSourceArmNetworkWatcher(),
		"azurerm_notification_hub_namespace":                dataSourceNotificationHubNamespace(),
		"azurerm_notification_hub":                          dataSourceNotificationHub(),
		"azurerm_platform_image":                            dataSourceArmPlatformImage(),
		"azurerm_policy_definition":                         dataSourceArmPolicyDefinition(),
		"azurerm_postgresql_server":                         dataSourcePostgreSqlServer(),
		"azurerm_private_link_endpoint_connection":          dataSourceArmPrivateLinkEndpointConnection(),
		"azurerm_private_endpoint_connection":               dataSourceArmPrivateEndpointConnection(),
		"azurerm_private_link_service":                      dataSourceArmPrivateLinkService(),
		"azurerm_private_link_service_endpoint_connections": dataSourceArmPrivateLinkServiceEndpointConnections(),
		"azurerm_proximity_placement_group":                 dataSourceArmProximityPlacementGroup(),
		"azurerm_public_ip":                                 dataSourceArmPublicIP(),
		"azurerm_public_ips":                                dataSourceArmPublicIPs(),
		"azurerm_public_ip_prefix":                          dataSourceArmPublicIpPrefix(),
		"azurerm_recovery_services_vault":                   dataSourceArmRecoveryServicesVault(),
		"azurerm_recovery_services_protection_policy_vm":    dataSourceArmRecoveryServicesProtectionPolicyVm(),
		"azurerm_redis_cache":                               dataSourceArmRedisCache(),
		"azurerm_resources":                                 dataSourceArmResources(),
		"azurerm_resource_group":                            dataSourceArmResourceGroup(),
		"azurerm_route_table":                               dataSourceArmRouteTable(),
		"azurerm_scheduler_job_collection":                  dataSourceArmSchedulerJobCollection(),
		"azurerm_servicebus_namespace":                      dataSourceArmServiceBusNamespace(),
		"azurerm_servicebus_namespace_authorization_rule":   dataSourceArmServiceBusNamespaceAuthorizationRule(),
		"azurerm_sql_server":                                dataSourceSqlServer(),
		"azurerm_sql_database":                              dataSourceSqlDatabase(),
		"azurerm_stream_analytics_job":                      dataSourceArmStreamAnalyticsJob(),
		"azurerm_storage_account_blob_container_sas":        dataSourceArmStorageAccountBlobContainerSharedAccessSignature(),
		"azurerm_storage_account_sas":                       dataSourceArmStorageAccountSharedAccessSignature(),
		"azurerm_storage_account":                           dataSourceArmStorageAccount(),
		"azurerm_storage_management_policy":                 dataSourceArmStorageManagementPolicy(),
		"azurerm_subnet":                                    dataSourceArmSubnet(),
		"azurerm_subscription":                              dataSourceArmSubscription(),
		"azurerm_subscriptions":                             dataSourceArmSubscriptions(),
		"azurerm_traffic_manager_geographical_location":     dataSourceArmTrafficManagerGeographicalLocation(),
		"azurerm_user_assigned_identity":                    dataSourceArmUserAssignedIdentity(),
		"azurerm_virtual_hub":                               dataSourceArmVirtualHub(),
		"azurerm_virtual_machine":                           dataSourceArmVirtualMachine(),
		"azurerm_virtual_network_gateway":                   dataSourceArmVirtualNetworkGateway(),
		"azurerm_virtual_network_gateway_connection":        dataSourceArmVirtualNetworkGatewayConnection(),
		"azurerm_virtual_network":                           dataSourceArmVirtualNetwork(),
	}

	resources := map[string]*schema.Resource{
		"azurerm_advanced_threat_protection":                            resourceArmAdvancedThreatProtection(),
		"azurerm_app_service_active_slot":                               resourceArmAppServiceActiveSlot(),
		"azurerm_app_service_certificate":                               resourceArmAppServiceCertificate(),
		"azurerm_app_service_certificate_order":                         resourceArmAppServiceCertificateOrder(),
		"azurerm_app_service_custom_hostname_binding":                   resourceArmAppServiceCustomHostnameBinding(),
		"azurerm_app_service_plan":                                      resourceArmAppServicePlan(),
		"azurerm_app_service_slot":                                      resourceArmAppServiceSlot(),
		"azurerm_app_service_source_control_token":                      resourceArmAppServiceSourceControlToken(),
		"azurerm_app_service_virtual_network_swift_connection":          resourceArmAppServiceVirtualNetworkSwiftConnection(),
		"azurerm_app_service":                                           resourceArmAppService(),
		"azurerm_application_gateway":                                   resourceArmApplicationGateway(),
		"azurerm_application_security_group":                            resourceArmApplicationSecurityGroup(),
		"azurerm_autoscale_setting":                                     resourceArmAutoScaleSetting(),
		"azurerm_backup_container_storage_account":                      resourceArmBackupProtectionContainerStorageAccount(),
		"azurerm_backup_policy_file_share":                              resourceArmBackupProtectionPolicyFileShare(),
		"azurerm_backup_protected_file_share":                           resourceArmBackupProtectedFileShare(),
		"azurerm_backup_protected_vm":                                   resourceArmRecoveryServicesBackupProtectedVM(),
		"azurerm_backup_policy_vm":                                      resourceArmBackupProtectionPolicyVM(),
		"azurerm_bastion_host":                                          resourceArmBastionHost(),
		"azurerm_connection_monitor":                                    resourceArmConnectionMonitor(),
		"azurerm_dashboard":                                             resourceArmDashboard(),
		"azurerm_ddos_protection_plan":                                  resourceArmDDoSProtectionPlan(),
		"azurerm_express_route_circuit_authorization":                   resourceArmExpressRouteCircuitAuthorization(),
		"azurerm_express_route_circuit_peering":                         resourceArmExpressRouteCircuitPeering(),
		"azurerm_express_route_circuit":                                 resourceArmExpressRouteCircuit(),
		"azurerm_firewall_application_rule_collection":                  resourceArmFirewallApplicationRuleCollection(),
		"azurerm_firewall_nat_rule_collection":                          resourceArmFirewallNatRuleCollection(),
		"azurerm_firewall_network_rule_collection":                      resourceArmFirewallNetworkRuleCollection(),
		"azurerm_firewall":                                              resourceArmFirewall(),
		"azurerm_function_app":                                          resourceArmFunctionApp(),
		"azurerm_image":                                                 resourceArmImage(),
		"azurerm_key_vault_access_policy":                               resourceArmKeyVaultAccessPolicy(),
		"azurerm_key_vault_certificate":                                 resourceArmKeyVaultCertificate(),
		"azurerm_key_vault_key":                                         resourceArmKeyVaultKey(),
		"azurerm_key_vault_secret":                                      resourceArmKeyVaultSecret(),
		"azurerm_key_vault":                                             resourceArmKeyVault(),
		"azurerm_lb_backend_address_pool":                               resourceArmLoadBalancerBackendAddressPool(),
		"azurerm_lb_nat_pool":                                           resourceArmLoadBalancerNatPool(),
		"azurerm_lb_nat_rule":                                           resourceArmLoadBalancerNatRule(),
		"azurerm_lb_probe":                                              resourceArmLoadBalancerProbe(),
		"azurerm_lb_outbound_rule":                                      resourceArmLoadBalancerOutboundRule(),
		"azurerm_lb_rule":                                               resourceArmLoadBalancerRule(),
		"azurerm_lb":                                                    resourceArmLoadBalancer(),
		"azurerm_local_network_gateway":                                 resourceArmLocalNetworkGateway(),
		"azurerm_management_lock":                                       resourceArmManagementLock(),
		"azurerm_maps_account":                                          resourceArmMapsAccount(),
		"azurerm_mariadb_configuration":                                 resourceArmMariaDbConfiguration(),
		"azurerm_mariadb_database":                                      resourceArmMariaDbDatabase(),
		"azurerm_mariadb_firewall_rule":                                 resourceArmMariaDBFirewallRule(),
		"azurerm_mariadb_server":                                        resourceArmMariaDbServer(),
		"azurerm_mariadb_virtual_network_rule":                          resourceArmMariaDbVirtualNetworkRule(),
		"azurerm_marketplace_agreement":                                 resourceArmMarketplaceAgreement(),
		"azurerm_media_services_account":                                resourceArmMediaServicesAccount(),
		"azurerm_metric_alertrule":                                      resourceArmMetricAlertRule(),
		"azurerm_monitor_autoscale_setting":                             resourceArmMonitorAutoScaleSetting(),
		"azurerm_monitor_action_group":                                  resourceArmMonitorActionGroup(),
		"azurerm_monitor_activity_log_alert":                            resourceArmMonitorActivityLogAlert(),
		"azurerm_monitor_diagnostic_setting":                            resourceArmMonitorDiagnosticSetting(),
		"azurerm_monitor_log_profile":                                   resourceArmMonitorLogProfile(),
		"azurerm_monitor_metric_alert":                                  resourceArmMonitorMetricAlert(),
		"azurerm_monitor_metric_alertrule":                              resourceArmMonitorMetricAlertRule(),
		"azurerm_mssql_elasticpool":                                     resourceArmMsSqlElasticPool(),
		"azurerm_mssql_database_vulnerability_assessment_rule_baseline": resourceArmMssqlDatabaseVulnerabilityAssessmentRuleBaseline(),
		"azurerm_mssql_server_security_alert_policy":                    resourceArmMssqlServerSecurityAlertPolicy(),
		"azurerm_mssql_server_vulnerability_assessment":                 resourceArmMssqlServerVulnerabilityAssessment(),
		"azurerm_mysql_configuration":                                   resourceArmMySQLConfiguration(),
		"azurerm_mysql_database":                                        resourceArmMySqlDatabase(),
		"azurerm_mysql_firewall_rule":                                   resourceArmMySqlFirewallRule(),
		"azurerm_mysql_server":                                          resourceArmMySqlServer(),
		"azurerm_mysql_virtual_network_rule":                            resourceArmMySqlVirtualNetworkRule(),
		"azurerm_nat_gateway":                                           resourceArmNatGateway(),
		"azurerm_network_connection_monitor":                            resourceArmNetworkConnectionMonitor(),
		"azurerm_network_ddos_protection_plan":                          resourceArmNetworkDDoSProtectionPlan(),
		"azurerm_network_interface":                                     resourceArmNetworkInterface(),
		"azurerm_network_interface_application_gateway_backend_address_pool_association": resourceArmNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(),
		"azurerm_network_interface_application_security_group_association":               resourceArmNetworkInterfaceApplicationSecurityGroupAssociation(),
		"azurerm_network_interface_backend_address_pool_association":                     resourceArmNetworkInterfaceBackendAddressPoolAssociation(),
		"azurerm_network_interface_nat_rule_association":                                 resourceArmNetworkInterfaceNatRuleAssociation(),
		"azurerm_network_packet_capture":                                                 resourceArmNetworkPacketCapture(),
		"azurerm_network_profile":                                                        resourceArmNetworkProfile(),
		"azurerm_network_security_group":                                                 resourceArmNetworkSecurityGroup(),
		"azurerm_network_security_rule":                                                  resourceArmNetworkSecurityRule(),
		"azurerm_network_watcher_flow_log":                                               resourceArmNetworkWatcherFlowLog(),
		"azurerm_network_watcher":                                                        resourceArmNetworkWatcher(),
		"azurerm_netapp_account":                                                         resourceArmNetAppAccount(),
		"azurerm_netapp_pool":                                                            resourceArmNetAppPool(),
		"azurerm_netapp_volume":                                                          resourceArmNetAppVolume(),
		"azurerm_notification_hub_authorization_rule":                                    resourceArmNotificationHubAuthorizationRule(),
		"azurerm_notification_hub_namespace":                                             resourceArmNotificationHubNamespace(),
		"azurerm_notification_hub":                                                       resourceArmNotificationHub(),
		"azurerm_packet_capture":                                                         resourceArmPacketCapture(),
		"azurerm_policy_assignment":                                                      resourceArmPolicyAssignment(),
		"azurerm_policy_definition":                                                      resourceArmPolicyDefinition(),
		"azurerm_policy_set_definition":                                                  resourceArmPolicySetDefinition(),
		"azurerm_point_to_site_vpn_gateway":                                              resourceArmPointToSiteVPNGateway(),
		"azurerm_postgresql_configuration":                                               resourceArmPostgreSQLConfiguration(),
		"azurerm_postgresql_database":                                                    resourceArmPostgreSQLDatabase(),
		"azurerm_postgresql_firewall_rule":                                               resourceArmPostgreSQLFirewallRule(),
		"azurerm_postgresql_server":                                                      resourceArmPostgreSQLServer(),
		"azurerm_postgresql_virtual_network_rule":                                        resourceArmPostgreSQLVirtualNetworkRule(),
		"azurerm_private_dns_zone":                                                       resourceArmPrivateDnsZone(),
		"azurerm_private_dns_a_record":                                                   resourceArmPrivateDnsARecord(),
		"azurerm_private_dns_aaaa_record":                                                resourceArmPrivateDnsAaaaRecord(),
		"azurerm_private_dns_cname_record":                                               resourceArmPrivateDnsCNameRecord(),
		"azurerm_private_dns_mx_record":                                                  resourceArmPrivateDnsMxRecord(),
		"azurerm_private_dns_ptr_record":                                                 resourceArmPrivateDnsPtrRecord(),
		"azurerm_private_dns_srv_record":                                                 resourceArmPrivateDnsSrvRecord(),
		"azurerm_private_dns_zone_virtual_network_link":                                  resourceArmPrivateDnsZoneVirtualNetworkLink(),
		"azurerm_private_link_endpoint":                                                  resourceArmPrivateLinkEndpoint(),
		"azurerm_private_endpoint":                                                       resourceArmPrivateEndpoint(),
		"azurerm_private_link_service":                                                   resourceArmPrivateLinkService(),
		"azurerm_proximity_placement_group":                                              resourceArmProximityPlacementGroup(),
		"azurerm_public_ip":                                                              resourceArmPublicIp(),
		"azurerm_public_ip_prefix":                                                       resourceArmPublicIpPrefix(),
		"azurerm_recovery_network_mapping":                                               resourceArmRecoveryServicesNetworkMapping(),
		"azurerm_recovery_replicated_vm":                                                 resourceArmRecoveryServicesReplicatedVm(),
		"azurerm_recovery_services_fabric":                                               resourceArmRecoveryServicesFabric(),
		"azurerm_recovery_services_protected_vm":                                         resourceArmRecoveryServicesProtectedVm(),
		"azurerm_recovery_services_protection_container":                                 resourceArmRecoveryServicesProtectionContainer(),
		"azurerm_recovery_services_protection_container_mapping":                         resourceArmRecoveryServicesProtectionContainerMapping(),
		"azurerm_recovery_services_protection_policy_vm":                                 resourceArmRecoveryServicesProtectionPolicyVm(),
		"azurerm_recovery_services_replication_policy":                                   resourceArmRecoveryServicesReplicationPolicy(),
		"azurerm_recovery_services_vault":                                                resourceArmRecoveryServicesVault(),
		"azurerm_redis_cache":                                                            resourceArmRedisCache(),
		"azurerm_redis_firewall_rule":                                                    resourceArmRedisFirewallRule(),
		"azurerm_relay_hybrid_connection":                                                resourceArmHybridConnection(),
		"azurerm_relay_namespace":                                                        resourceArmRelayNamespace(),
		"azurerm_resource_group":                                                         resourceArmResourceGroup(),
		"azurerm_route_table":                                                            resourceArmRouteTable(),
		"azurerm_route":                                                                  resourceArmRoute(),
		"azurerm_scheduler_job_collection":                                               resourceArmSchedulerJobCollection(),
		"azurerm_scheduler_job":                                                          resourceArmSchedulerJob(),
		"azurerm_search_service":                                                         resourceArmSearchService(),
		"azurerm_security_center_contact":                                                resourceArmSecurityCenterContact(),
		"azurerm_security_center_subscription_pricing":                                   resourceArmSecurityCenterSubscriptionPricing(),
		"azurerm_security_center_workspace":                                              resourceArmSecurityCenterWorkspace(),
		"azurerm_service_fabric_cluster":                                                 resourceArmServiceFabricCluster(),
		"azurerm_servicebus_namespace_authorization_rule":                                resourceArmServiceBusNamespaceAuthorizationRule(),
		"azurerm_servicebus_namespace":                                                   resourceArmServiceBusNamespace(),
		"azurerm_servicebus_queue_authorization_rule":                                    resourceArmServiceBusQueueAuthorizationRule(),
		"azurerm_servicebus_queue":                                                       resourceArmServiceBusQueue(),
		"azurerm_servicebus_subscription_rule":                                           resourceArmServiceBusSubscriptionRule(),
		"azurerm_servicebus_subscription":                                                resourceArmServiceBusSubscription(),
		"azurerm_servicebus_topic_authorization_rule":                                    resourceArmServiceBusTopicAuthorizationRule(),
		"azurerm_servicebus_topic":                                                       resourceArmServiceBusTopic(),
		"azurerm_shared_image_gallery":                                                   resourceArmSharedImageGallery(),
		"azurerm_shared_image_version":                                                   resourceArmSharedImageVersion(),
		"azurerm_shared_image":                                                           resourceArmSharedImage(),
		"azurerm_signalr_service":                                                        resourceArmSignalRService(),
		"azurerm_site_recovery_fabric":                                                   resourceArmSiteRecoveryFabric(),
		"azurerm_site_recovery_network_mapping":                                          resourceArmSiteRecoveryNetworkMapping(),
		"azurerm_site_recovery_protection_container":                                     resourceArmSiteRecoveryProtectionContainer(),
		"azurerm_site_recovery_protection_container_mapping":                             resourceArmSiteRecoveryProtectionContainerMapping(),
		"azurerm_site_recovery_replicated_vm":                                            resourceArmSiteRecoveryReplicatedVM(),
		"azurerm_site_recovery_replication_policy":                                       resourceArmSiteRecoveryReplicationPolicy(),
		"azurerm_sql_active_directory_administrator":                                     resourceArmSqlAdministrator(),
		"azurerm_sql_database":                                                           resourceArmSqlDatabase(),
		"azurerm_sql_elasticpool":                                                        resourceArmSqlElasticPool(),
		"azurerm_sql_failover_group":                                                     resourceArmSqlFailoverGroup(),
		"azurerm_sql_firewall_rule":                                                      resourceArmSqlFirewallRule(),
		"azurerm_sql_server":                                                             resourceArmSqlServer(),
		"azurerm_sql_virtual_network_rule":                                               resourceArmSqlVirtualNetworkRule(),
		"azurerm_storage_account":                                                        resourceArmStorageAccount(),
		"azurerm_storage_account_network_rules":                                          resourceArmStorageAccountNetworkRules(),
		"azurerm_storage_blob":                                                           resourceArmStorageBlob(),
		"azurerm_storage_container":                                                      resourceArmStorageContainer(),
		"azurerm_storage_data_lake_gen2_filesystem":                                      resourceArmStorageDataLakeGen2FileSystem(),
		"azurerm_storage_management_policy":                                              resourceArmStorageManagementPolicy(),
		"azurerm_storage_queue":                                                          resourceArmStorageQueue(),
		"azurerm_storage_share":                                                          resourceArmStorageShare(),
		"azurerm_storage_share_directory":                                                resourceArmStorageShareDirectory(),
		"azurerm_storage_table":                                                          resourceArmStorageTable(),
		"azurerm_storage_table_entity":                                                   resourceArmStorageTableEntity(),
		"azurerm_stream_analytics_job":                                                   resourceArmStreamAnalyticsJob(),
		"azurerm_stream_analytics_function_javascript_udf":                               resourceArmStreamAnalyticsFunctionUDF(),
		"azurerm_stream_analytics_output_blob":                                           resourceArmStreamAnalyticsOutputBlob(),
		"azurerm_stream_analytics_output_mssql":                                          resourceArmStreamAnalyticsOutputSql(),
		"azurerm_stream_analytics_output_eventhub":                                       resourceArmStreamAnalyticsOutputEventHub(),
		"azurerm_stream_analytics_output_servicebus_queue":                               resourceArmStreamAnalyticsOutputServiceBusQueue(),
		"azurerm_stream_analytics_output_servicebus_topic":                               resourceArmStreamAnalyticsOutputServiceBusTopic(),
		"azurerm_stream_analytics_reference_input_blob":                                  resourceArmStreamAnalyticsReferenceInputBlob(),
		"azurerm_stream_analytics_stream_input_blob":                                     resourceArmStreamAnalyticsStreamInputBlob(),
		"azurerm_stream_analytics_stream_input_eventhub":                                 resourceArmStreamAnalyticsStreamInputEventHub(),
		"azurerm_stream_analytics_stream_input_iothub":                                   resourceArmStreamAnalyticsStreamInputIoTHub(),
		"azurerm_subnet_network_security_group_association":                              resourceArmSubnetNetworkSecurityGroupAssociation(),
		"azurerm_subnet_route_table_association":                                         resourceArmSubnetRouteTableAssociation(),
		"azurerm_subnet_nat_gateway_association":                                         resourceArmSubnetNatGatewayAssociation(),
		"azurerm_subnet":                                                                 resourceArmSubnet(),
		"azurerm_template_deployment":                                                    resourceArmTemplateDeployment(),
		"azurerm_traffic_manager_endpoint":                                               resourceArmTrafficManagerEndpoint(),
		"azurerm_traffic_manager_profile":                                                resourceArmTrafficManagerProfile(),
		"azurerm_user_assigned_identity":                                                 resourceArmUserAssignedIdentity(),
		"azurerm_virtual_hub":                                                            resourceArmVirtualHub(),
		"azurerm_virtual_network_gateway_connection":                                     resourceArmVirtualNetworkGatewayConnection(),
		"azurerm_virtual_network_gateway":                                                resourceArmVirtualNetworkGateway(),
		"azurerm_virtual_network_peering":                                                resourceArmVirtualNetworkPeering(),
		"azurerm_virtual_network":                                                        resourceArmVirtualNetwork(),
		"azurerm_virtual_wan":                                                            resourceArmVirtualWan(),
		"azurerm_vpn_gateway":                                                            resourceArmVPNGateway(),
		"azurerm_vpn_server_configuration":                                               resourceArmVPNServerConfiguration(),
		"azurerm_web_application_firewall_policy":                                        resourceArmWebApplicationFirewallPolicy(),
	}

	return provider.AzureProvider(dataSources, resources)
}
