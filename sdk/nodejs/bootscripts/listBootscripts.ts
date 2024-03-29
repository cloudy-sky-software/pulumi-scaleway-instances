// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listBootscripts(args: ListBootscriptsArgs, opts?: pulumi.InvokeOptions): Promise<ListBootscriptsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:bootscripts:listBootscripts", {
        "zone": args.zone,
    }, opts);
}

export interface ListBootscriptsArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListBootscriptsResult {
    readonly items: outputs.bootscripts.ScalewayInstanceV1ListBootscriptsResponse;
}
export function listBootscriptsOutput(args: ListBootscriptsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListBootscriptsResult> {
    return pulumi.output(args).apply((a: any) => listBootscripts(a, opts))
}

export interface ListBootscriptsOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
