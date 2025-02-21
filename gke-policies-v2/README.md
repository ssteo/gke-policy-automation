# GKE Policy Automation library

## Policy structure

Please refer to the [Policy Authoring Guide](./AUTHORING.md) for details about structure
of our policy files.

## Available Policies

<!-- BEGIN POLICY-DOC -->
|Group|Title|Description|File|
|-|-|-|-|
|Availability|Control Plane redundancy|GKE cluster should be regional for maximum availability of control plane during upgrades and zonal outages|[gke-policies-v2/policy/control_plane_redundancy.rego](../gke-policies-v2/policy/control_plane_redundancy.rego)|
|Availability|Multi-zone node pools|GKE node pools should be regional (multiple zones) for maximum nodes availability during zonal outages|[gke-policies-v2/policy/node_pool_multi_zone.rego](../gke-policies-v2/policy/node_pool_multi_zone.rego)|
|Availability|Use Node Auto-Repair|GKE node pools should have Node Auto-Repair enabled to configure Kubernetes Engine|[gke-policies-v2/policy/node_pool_autorepair.rego](../gke-policies-v2/policy/node_pool_autorepair.rego)|
|Maintenance|Cloud Monitoring and Logging|GKE cluster should use Cloud Logging and Monitoring|[gke-policies-v2/policy/monitoring_and_logging.rego](../gke-policies-v2/policy/monitoring_and_logging.rego)|
|Management|Enable binary authorization in the cluster|GKE cluster should enable for deploy-time security control that ensures only trusted container images are deployed to gain tighter control over your container environment.|[gke-policies-v2/policy/cluster_binary_authorization.rego](../gke-policies-v2/policy/cluster_binary_authorization.rego)|
|Management|GKE Autopilot mode|GKE Autopilot mode is the recommended way to operate a GKE cluster|[gke-policies-v2/policy/autopilot_cluster.rego](../gke-policies-v2/policy/autopilot_cluster.rego)|
|Management|GKE VPC-native cluster|GKE cluster nodepool should be VPC-native as per our best-practices|[gke-policies-v2/policy/vpc_native_cluster.rego](../gke-policies-v2/policy/vpc_native_cluster.rego)|
|Management|Receive updates about new GKE versions|GKE cluster should be proactively receive updates about GKE upgrades and GKE versions|[gke-policies-v2/policy/cluster_receive_updates.rego](../gke-policies-v2/policy/cluster_receive_updates.rego)|
|Management|Schedule maintenance windows and exclusions|GKE cluster should schedule maintenance windows and exclusions to upgrade predictability and to align updates with off-peak business hours.|[gke-policies-v2/policy/cluster_maintenance_window.rego](../gke-policies-v2/policy/cluster_maintenance_window.rego)|
|Management|Use Compute Engine persistent disk CSI driver|Automatic deployment and management of the Compute Engine persisten disk CSI driver. The driver provides support for features like customer managed encryption keys or volume snapshots.|[gke-policies-v2/policy/cluster_gce_csi_driver.rego](../gke-policies-v2/policy/cluster_gce_csi_driver.rego)|
|Management|Version skew between node pools and control plane|Difference between cluster control plane version and node pools version should be no more than 2 minor versions.|[gke-policies-v2/policy/node_pool_version_skew.rego](../gke-policies-v2/policy/node_pool_version_skew.rego)|
|Scalability|GKE L4 ILB Subsetting|GKE cluster should use GKE L4 ILB Subsetting if nodes > 250|[gke-policies-v2/policy/ilb_subsetting.rego](../gke-policies-v2/policy/ilb_subsetting.rego)|
|Scalability|GKE Nodes Limit|GKE Nodes Limit|[gke-policies-v2/scalability/limits_nodes.rego](../gke-policies-v2/scalability/limits_nodes.rego)|
|Scalability|GKE node local DNS cache|GKE cluster should use node local DNS cache|[gke-policies-v2/policy/node_local_dns_cache.rego](../gke-policies-v2/policy/node_local_dns_cache.rego)|
|Scalability|Number of HPAs in a cluster|The optimal number of Horizonal Pod Autoscalers in a cluster|[gke-policies-v2/scalability/limits_hpas.rego](../gke-policies-v2/scalability/limits_hpas.rego)|
|Scalability|Number of PODs in a cluster|The total number of PODs running in a cluster|[gke-policies-v2/scalability/limits_pods.rego](../gke-policies-v2/scalability/limits_pods.rego)|
|Scalability|Number of PODs per node|The total number of PODs running on a single node|[gke-policies-v2/scalability/limits_pods_per_node.rego](../gke-policies-v2/scalability/limits_pods_per_node.rego)|
|Scalability|Number of containers in a cluster|The total number of containers running in a cluster|[gke-policies-v2/scalability/limits_containers.rego](../gke-policies-v2/scalability/limits_containers.rego)|
|Scalability|Number of nodes in a nodepool zone|The total number of nodes running in a single node pool zone|[gke-policies-v2/scalability/limits_nodes_per_pool_zone.rego](../gke-policies-v2/scalability/limits_nodes_per_pool_zone.rego)|
|Scalability|Number of secrets with KMS encryption|The total number of secrets when KMS secret encryption is enabled|[gke-policies-v2/scalability/limit_secrets_encryption.rego](../gke-policies-v2/scalability/limit_secrets_encryption.rego)|
|Scalability|Number of services in a cluster|The total number of services running in a cluster|[gke-policies-v2/scalability/limit_services.rego](../gke-policies-v2/scalability/limit_services.rego)|
|Scalability|Number of services per namespace|The total number of services running in single namespace|[gke-policies-v2/scalability/limit_services_per_ns.rego](../gke-policies-v2/scalability/limit_services_per_ns.rego)|
|Scalability|Use node pool autoscaling|GKE node pools should have autoscaling configured to proper resize nodes according to traffic|[gke-policies-v2/policy/node_pool_autoscaling.rego](../gke-policies-v2/policy/node_pool_autoscaling.rego)|
|Security|Control Plane endpoint access|Control Plane endpoint access should be limited to authorized networks only|[gke-policies-v2/policy/control_plane_access.rego](../gke-policies-v2/policy/control_plane_access.rego)|
|Security|Control Plane endpoint visibility|Control Plane endpoint should be locked from external access|[gke-policies-v2/policy/control_plane_endpoint.rego](../gke-policies-v2/policy/control_plane_endpoint.rego)|
|Security|Enrollment in Release Channels|GKE cluster should be enrolled in release channels|[gke-policies-v2/policy/cluster_release_channels.rego](../gke-policies-v2/policy/cluster_release_channels.rego)|
|Security|Ensure that node pool locations within Node Auto-Provisioning are covering more than one zone (or not enforced at all)|Node Auto-Provisioning configuration should cover more than one zone|[gke-policies-v2/policy/nap_forbid_single_zone.rego](../gke-policies-v2/policy/nap_forbid_single_zone.rego)|
|Security|Ensure that nodes in Node Auto-Provisioning node pools will use Container-Optimized OS|Nodes in Node Auto-Provisioning should use Container-Optimized OS|[gke-policies-v2/policy/nap_use_cos.rego](../gke-policies-v2/policy/nap_use_cos.rego)|
|Security|Ensure that nodes in Node Auto-Provisioning node pools will use integrity monitoring|Nodes in Node Auto-Provisioning should use integrity monitoring|[gke-policies-v2/policy/nap_integrity_monitoring.rego](../gke-policies-v2/policy/nap_integrity_monitoring.rego)|
|Security|Forbid default Service Accounts in Node Auto-Provisioning|Node Auto-Provisioning configuration should not allow default Service Accounts|[gke-policies-v2/policy/nap_forbid_default_sa.rego](../gke-policies-v2/policy/nap_forbid_default_sa.rego)|
|Security|Forbid default compute SA on node_pool|GKE node pools should have a dedicated sa with a restricted set of permissions|[gke-policies-v2/policy/node_pool_forbid_default_sa.rego](../gke-policies-v2/policy/node_pool_forbid_default_sa.rego)|
|Security|GKE Network Policies engine|GKE cluster should have Network Policies or Dataplane V2 enabled|[gke-policies-v2/policy/network_policies.rego](../gke-policies-v2/policy/network_policies.rego)|
|Security|GKE RBAC authorization|GKE cluster should use RBAC instead of legacy ABAC authorization|[gke-policies-v2/policy/control_plane_disable_legacy_authorization.rego](../gke-policies-v2/policy/control_plane_disable_legacy_authorization.rego)|
|Security|GKE Shielded Nodes|GKE cluster should use shielded nodes|[gke-policies-v2/policy/shielded_nodes.rego](../gke-policies-v2/policy/shielded_nodes.rego)|
|Security|GKE Workload Identity|GKE cluster should have Workload Identity enabled|[gke-policies-v2/policy/workload_identity.rego](../gke-policies-v2/policy/workload_identity.rego)|
|Security|GKE private cluster|GKE cluster should be private to ensure network isolation|[gke-policies-v2/policy/private_cluster.rego](../gke-policies-v2/policy/private_cluster.rego)|
|Security|Integrity monitoring on the nodes|GKE node pools should have integrity monitoring feature enabled to detect changes in a VM boot measurments|[gke-policies-v2/policy/node_pool_integrity_monitoring.rego](../gke-policies-v2/policy/node_pool_integrity_monitoring.rego)|
|Security|Kubernetes secrets encryption|GKE cluster should use encryption for kubernetes application secrets|[gke-policies-v2/policy/secret_encryption.rego](../gke-policies-v2/policy/secret_encryption.rego)|
|Security|Secure boot on the nodes|Secure Boot helps ensure that the system only runs authentic software by verifying the digital signature of all boot components, and halting the boot process if signature verification fails|[gke-policies-v2/policy/node_pool_secure_boot.rego](../gke-policies-v2/policy/node_pool_secure_boot.rego)|
|Security|Use Container-Optimized OS|GKE node pools should use Container-Optimized OS which is maintained by Google and optimized for running Docker containers with security and efficiency.|[gke-policies-v2/policy/node_pool_use_cos.rego](../gke-policies-v2/policy/node_pool_use_cos.rego)|
|Security|Use Node Auto-Upgrade|GKE node pools should have Node Auto-Upgrade enabled to configure Kubernetes Engine|[gke-policies-v2/policy/node_pool_autoupgrade.rego](../gke-policies-v2/policy/node_pool_autoupgrade.rego)|
|Security|Use RBAC Google group|GKE cluster should have RBAC security Google group enabled|[gke-policies-v2/policy/node_rbac_security_group.rego](../gke-policies-v2/policy/node_rbac_security_group.rego)|
