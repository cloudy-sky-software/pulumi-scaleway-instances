// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPrivateNIC(args: GetPrivateNICArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateNICResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:private_nics:getPrivateNIC", {
        "id": args.id,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

export interface GetPrivateNICArgs {
    id: string;
    serverId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface GetPrivateNICResult {
    readonly items: outputs.private_nics.ScalewayInstanceV1GetPrivateNICResponse;
}
export function getPrivateNICOutput(args: GetPrivateNICOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateNICResult> {
    return pulumi.output(args).apply((a: any) => getPrivateNIC(a, opts))
}

export interface GetPrivateNICOutputArgs {
    id: pulumi.Input<string>;
    serverId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
