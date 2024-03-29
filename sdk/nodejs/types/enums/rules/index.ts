// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Action = {
    Accept: "accept",
    Drop: "drop",
} as const;

export type Action = (typeof Action)[keyof typeof Action];

export const Direction = {
    Inbound: "inbound",
    Outbound: "outbound",
} as const;

export type Direction = (typeof Direction)[keyof typeof Direction];

export const Protocol = {
    Tcp: "TCP",
    Udp: "UDP",
    Icmp: "ICMP",
    Any: "ANY",
} as const;

export type Protocol = (typeof Protocol)[keyof typeof Protocol];

export const ScalewayInstanceV1SecurityGroupRuleAction = {
    Accept: "accept",
    Drop: "drop",
} as const;

export type ScalewayInstanceV1SecurityGroupRuleAction = (typeof ScalewayInstanceV1SecurityGroupRuleAction)[keyof typeof ScalewayInstanceV1SecurityGroupRuleAction];

export const ScalewayInstanceV1SecurityGroupRuleDirection = {
    Inbound: "inbound",
    Outbound: "outbound",
} as const;

export type ScalewayInstanceV1SecurityGroupRuleDirection = (typeof ScalewayInstanceV1SecurityGroupRuleDirection)[keyof typeof ScalewayInstanceV1SecurityGroupRuleDirection];

export const ScalewayInstanceV1SecurityGroupRuleProtocol = {
    Tcp: "TCP",
    Udp: "UDP",
    Icmp: "ICMP",
    Any: "ANY",
} as const;

export type ScalewayInstanceV1SecurityGroupRuleProtocol = (typeof ScalewayInstanceV1SecurityGroupRuleProtocol)[keyof typeof ScalewayInstanceV1SecurityGroupRuleProtocol];

export const ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction = {
    Accept: "accept",
    Drop: "drop",
} as const;

/**
 * Action to apply when the rule matches a packet
 */
export type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction = (typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction)[keyof typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction];

export const ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection = {
    Inbound: "inbound",
    Outbound: "outbound",
} as const;

/**
 * Direction the rule applies to
 */
export type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection = (typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection)[keyof typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection];

export const ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol = {
    Tcp: "TCP",
    Udp: "UDP",
    Icmp: "ICMP",
    Any: "ANY",
} as const;

/**
 * Protocol family this rule applies to
 */
export type ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol = (typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol)[keyof typeof ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol];
