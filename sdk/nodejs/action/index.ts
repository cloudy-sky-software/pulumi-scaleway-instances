// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ListServerActionsArgs, ListServerActionsResult, ListServerActionsOutputArgs } from "./listServerActions";
export const listServerActions: typeof import("./listServerActions").listServerActions = null as any;
export const listServerActionsOutput: typeof import("./listServerActions").listServerActionsOutput = null as any;
utilities.lazyLoad(exports, ["listServerActions","listServerActionsOutput"], () => require("./listServerActions"));

export { ServerActionArgs } from "./serverAction";
export type ServerAction = import("./serverAction").ServerAction;
export const ServerAction: typeof import("./serverAction").ServerAction = null as any;
utilities.lazyLoad(exports, ["ServerAction"], () => require("./serverAction"));


// Export enums:
export * from "../types/enums/action";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway-instances:action:ServerAction":
                return new ServerAction(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway-instances", "action", _module)
