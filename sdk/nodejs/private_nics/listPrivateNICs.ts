// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listPrivateNICs(args: ListPrivateNICsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.private_nics.ScalewayInstanceV1ListPrivateNICsResponse> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:private_nics:listPrivateNICs", {
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

export interface ListPrivateNICsArgs {
    serverId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function listPrivateNICsOutput(args: ListPrivateNICsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.private_nics.ScalewayInstanceV1ListPrivateNICsResponse> {
    return pulumi.output(args).apply((a: any) => listPrivateNICs(a, opts))
}

export interface ListPrivateNICsOutputArgs {
    serverId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
