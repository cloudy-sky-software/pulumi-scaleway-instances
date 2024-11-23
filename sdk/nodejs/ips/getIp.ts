// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getIp(args: GetIpArgs, opts?: pulumi.InvokeOptions): Promise<outputs.ips.ScalewayInstanceV1GetIpResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:ips:getIp", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetIpArgs {
    /**
     * The IP ID or address to get
     */
    id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getIpOutput(args: GetIpOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.ips.ScalewayInstanceV1GetIpResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:ips:getIp", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetIpOutputArgs {
    /**
     * The IP ID or address to get
     */
    id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
