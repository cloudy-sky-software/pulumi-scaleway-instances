// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listServers(args: ListServersArgs, opts?: pulumi.InvokeOptions): Promise<ListServersResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway-instances:servers:listServers", {
        "zone": args.zone,
    }, opts);
}

export interface ListServersArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListServersResult {
    readonly items: outputs.servers.ScalewayInstanceV1ListServersResponse;
}

export function listServersOutput(args: ListServersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListServersResult> {
    return pulumi.output(args).apply(a => listServers(a, opts))
}

export interface ListServersOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
