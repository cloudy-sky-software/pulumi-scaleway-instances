// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSecurityGroupRuleArgs, GetSecurityGroupRuleResult, GetSecurityGroupRuleOutputArgs } from "./getSecurityGroupRule";
export const getSecurityGroupRule: typeof import("./getSecurityGroupRule").getSecurityGroupRule = null as any;
export const getSecurityGroupRuleOutput: typeof import("./getSecurityGroupRule").getSecurityGroupRuleOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityGroupRule","getSecurityGroupRuleOutput"], () => require("./getSecurityGroupRule"));

export { ListDefaultSecurityGroupRulesArgs, ListDefaultSecurityGroupRulesResult, ListDefaultSecurityGroupRulesOutputArgs } from "./listDefaultSecurityGroupRules";
export const listDefaultSecurityGroupRules: typeof import("./listDefaultSecurityGroupRules").listDefaultSecurityGroupRules = null as any;
export const listDefaultSecurityGroupRulesOutput: typeof import("./listDefaultSecurityGroupRules").listDefaultSecurityGroupRulesOutput = null as any;
utilities.lazyLoad(exports, ["listDefaultSecurityGroupRules","listDefaultSecurityGroupRulesOutput"], () => require("./listDefaultSecurityGroupRules"));

export { ListSecurityGroupRulesArgs, ListSecurityGroupRulesResult, ListSecurityGroupRulesOutputArgs } from "./listSecurityGroupRules";
export const listSecurityGroupRules: typeof import("./listSecurityGroupRules").listSecurityGroupRules = null as any;
export const listSecurityGroupRulesOutput: typeof import("./listSecurityGroupRules").listSecurityGroupRulesOutput = null as any;
utilities.lazyLoad(exports, ["listSecurityGroupRules","listSecurityGroupRulesOutput"], () => require("./listSecurityGroupRules"));

export { SecurityGroupRuleArgs } from "./securityGroupRule";
export type SecurityGroupRule = import("./securityGroupRule").SecurityGroupRule;
export const SecurityGroupRule: typeof import("./securityGroupRule").SecurityGroupRule = null as any;
utilities.lazyLoad(exports, ["SecurityGroupRule"], () => require("./securityGroupRule"));


// Export enums:
export * from "../types/enums/rules";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway-instances:rules:SecurityGroupRule":
                return new SecurityGroupRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway-instances", "rules", _module)
