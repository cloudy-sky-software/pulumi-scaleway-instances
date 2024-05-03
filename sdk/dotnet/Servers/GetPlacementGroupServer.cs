// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers
{
    public static class GetPlacementGroupServer
    {
        public static Task<GetPlacementGroupServerResult> InvokeAsync(GetPlacementGroupServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlacementGroupServerResult>("scaleway-instances:servers:getPlacementGroupServer", args ?? new GetPlacementGroupServerArgs(), options.WithDefaults());

        public static Output<GetPlacementGroupServerResult> Invoke(GetPlacementGroupServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlacementGroupServerResult>("scaleway-instances:servers:getPlacementGroupServer", args ?? new GetPlacementGroupServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlacementGroupServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the placement group
        /// </summary>
        [Input("placementGroupId", required: true)]
        public string PlacementGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetPlacementGroupServerArgs()
        {
        }
        public static new GetPlacementGroupServerArgs Empty => new GetPlacementGroupServerArgs();
    }

    public sealed class GetPlacementGroupServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the placement group
        /// </summary>
        [Input("placementGroupId", required: true)]
        public Input<string> PlacementGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetPlacementGroupServerInvokeArgs()
        {
        }
        public static new GetPlacementGroupServerInvokeArgs Empty => new GetPlacementGroupServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlacementGroupServerResult
    {
        public readonly Outputs.ScalewayInstanceV1GetPlacementGroupServersResponse Items;

        [OutputConstructor]
        private GetPlacementGroupServerResult(Outputs.ScalewayInstanceV1GetPlacementGroupServersResponse items)
        {
            Items = items;
        }
    }
}
