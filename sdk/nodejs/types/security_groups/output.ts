// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export interface ScalewayInstanceV1GetSecurityGroupResponse {
    security_group?: outputs.security_groups.ScalewayInstanceV1SecurityGroup;
}
/**
 * scalewayInstanceV1GetSecurityGroupResponseProvideDefaults sets the appropriate defaults for ScalewayInstanceV1GetSecurityGroupResponse
 */
export function scalewayInstanceV1GetSecurityGroupResponseProvideDefaults(val: ScalewayInstanceV1GetSecurityGroupResponse): ScalewayInstanceV1GetSecurityGroupResponse {
    return {
        ...val,
        security_group: (val.security_group ? outputs.security_groups.scalewayInstanceV1SecurityGroupProvideDefaults(val.security_group) : undefined),
    };
}

export interface ScalewayInstanceV1ListSecurityGroupsResponse {
    security_groups?: outputs.security_groups.ScalewayInstanceV1SecurityGroup[];
    total_count?: number;
}

export interface ScalewayInstanceV1SecurityGroup {
    /**
     * The security group creation date (RFC 3339 format)
     */
    creation_date?: string;
    /**
     * The security groups description
     */
    description?: string;
    /**
     * True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a ticket if you need to make it configurable.
     */
    enable_default_security?: boolean;
    id?: string;
    /**
     * The default inbound policy
     */
    inbound_default_policy?: enums.security_groups.ScalewayInstanceV1SecurityGroupInboundDefaultPolicy;
    /**
     * The security group modification date (RFC 3339 format)
     */
    modification_date?: string;
    /**
     * The security groups name
     */
    name: string;
    /**
     * The security groups organization ID
     */
    organization?: string;
    /**
     * True if it is your default security group for this organization ID
     */
    organization_default?: boolean;
    /**
     * The default outbound policy
     */
    outbound_default_policy?: enums.security_groups.ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy;
    /**
     * The security group project ID
     */
    project: string;
    /**
     * True if it is your default security group for this project ID
     */
    project_default?: boolean;
    /**
     * List of servers attached to this security group
     */
    servers?: outputs.security_groups.ScalewayInstanceV1ServerSummary[];
    /**
     * Security group state
     */
    state?: enums.security_groups.ScalewayInstanceV1SecurityGroupState;
    /**
     * True if the security group is stateful
     */
    stateful?: boolean;
    /**
     * The security group tags
     */
    tags?: string[];
    /**
     * The zone in which is the security group
     */
    zone?: string;
}
/**
 * scalewayInstanceV1SecurityGroupProvideDefaults sets the appropriate defaults for ScalewayInstanceV1SecurityGroup
 */
export function scalewayInstanceV1SecurityGroupProvideDefaults(val: ScalewayInstanceV1SecurityGroup): ScalewayInstanceV1SecurityGroup {
    return {
        ...val,
        inbound_default_policy: (val.inbound_default_policy) ?? "accept",
        outbound_default_policy: (val.outbound_default_policy) ?? "accept",
        state: (val.state) ?? "available",
    };
}

export interface ScalewayInstanceV1ServerSummary {
    id?: string;
    name?: string;
}
