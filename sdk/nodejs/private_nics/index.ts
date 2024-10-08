// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPrivateNICArgs, GetPrivateNICOutputArgs } from "./getPrivateNIC";
export const getPrivateNIC: typeof import("./getPrivateNIC").getPrivateNIC = null as any;
export const getPrivateNICOutput: typeof import("./getPrivateNIC").getPrivateNICOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateNIC","getPrivateNICOutput"], () => require("./getPrivateNIC"));

export { ListPrivateNICsArgs, ListPrivateNICsOutputArgs } from "./listPrivateNICs";
export const listPrivateNICs: typeof import("./listPrivateNICs").listPrivateNICs = null as any;
export const listPrivateNICsOutput: typeof import("./listPrivateNICs").listPrivateNICsOutput = null as any;
utilities.lazyLoad(exports, ["listPrivateNICs","listPrivateNICsOutput"], () => require("./listPrivateNICs"));

export { PrivateNICArgs } from "./privateNIC";
export type PrivateNIC = import("./privateNIC").PrivateNIC;
export const PrivateNIC: typeof import("./privateNIC").PrivateNIC = null as any;
utilities.lazyLoad(exports, ["PrivateNIC"], () => require("./privateNIC"));


// Export enums:
export * from "../types/enums/private_nics";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway-instances:private_nics:PrivateNIC":
                return new PrivateNIC(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway-instances", "private_nics", _module)
