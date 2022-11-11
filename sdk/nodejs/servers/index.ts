// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPlacementGroupServersArgs, GetPlacementGroupServersResult, GetPlacementGroupServersOutputArgs } from "./getPlacementGroupServers";
export const getPlacementGroupServers: typeof import("./getPlacementGroupServers").getPlacementGroupServers = null as any;
export const getPlacementGroupServersOutput: typeof import("./getPlacementGroupServers").getPlacementGroupServersOutput = null as any;
utilities.lazyLoad(exports, ["getPlacementGroupServers","getPlacementGroupServersOutput"], () => require("./getPlacementGroupServers"));

export { GetServerArgs, GetServerResult, GetServerOutputArgs } from "./getServer";
export const getServer: typeof import("./getServer").getServer = null as any;
export const getServerOutput: typeof import("./getServer").getServerOutput = null as any;
utilities.lazyLoad(exports, ["getServer","getServerOutput"], () => require("./getServer"));

export { ListServersArgs, ListServersResult, ListServersOutputArgs } from "./listServers";
export const listServers: typeof import("./listServers").listServers = null as any;
export const listServersOutput: typeof import("./listServers").listServersOutput = null as any;
utilities.lazyLoad(exports, ["listServers","listServersOutput"], () => require("./listServers"));

export { ListServersTypesArgs, ListServersTypesResult, ListServersTypesOutputArgs } from "./listServersTypes";
export const listServersTypes: typeof import("./listServersTypes").listServersTypes = null as any;
export const listServersTypesOutput: typeof import("./listServersTypes").listServersTypesOutput = null as any;
utilities.lazyLoad(exports, ["listServersTypes","listServersTypesOutput"], () => require("./listServersTypes"));

export { PlacementGroupServersArgs } from "./placementGroupServers";
export type PlacementGroupServers = import("./placementGroupServers").PlacementGroupServers;
export const PlacementGroupServers: typeof import("./placementGroupServers").PlacementGroupServers = null as any;
utilities.lazyLoad(exports, ["PlacementGroupServers"], () => require("./placementGroupServers"));

export { ServerArgs } from "./server";
export type Server = import("./server").Server;
export const Server: typeof import("./server").Server = null as any;
utilities.lazyLoad(exports, ["Server"], () => require("./server"));


// Export enums:
export * from "../types/enums/servers";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway-instances:servers:PlacementGroupServers":
                return new PlacementGroupServers(name, <any>undefined, { urn })
            case "scaleway-instances:servers:Server":
                return new Server(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway-instances", "servers", _module)