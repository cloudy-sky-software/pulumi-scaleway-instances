// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


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

export const SecurityGroupRuleAction = {
    Accept: "accept",
    Drop: "drop",
} as const;

export type SecurityGroupRuleAction = (typeof SecurityGroupRuleAction)[keyof typeof SecurityGroupRuleAction];

export const SecurityGroupRuleDirection = {
    Inbound: "inbound",
    Outbound: "outbound",
} as const;

export type SecurityGroupRuleDirection = (typeof SecurityGroupRuleDirection)[keyof typeof SecurityGroupRuleDirection];

export const SecurityGroupRuleProtocol = {
    Tcp: "TCP",
    Udp: "UDP",
    Icmp: "ICMP",
    Any: "ANY",
} as const;

export type SecurityGroupRuleProtocol = (typeof SecurityGroupRuleProtocol)[keyof typeof SecurityGroupRuleProtocol];
