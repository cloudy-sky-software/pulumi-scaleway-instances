// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers
{
    public static class GetServer
    {
        public static Task<Outputs.ScalewayInstanceV1GetServerResponse> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1GetServerResponse>("scaleway-instances:servers:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1GetServerResponse> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1GetServerResponse>("scaleway-instances:servers:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the server you want to get
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetServerArgs()
        {
        }
        public static new GetServerArgs Empty => new GetServerArgs();
    }

    public sealed class GetServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the server you want to get
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
        public static new GetServerInvokeArgs Empty => new GetServerInvokeArgs();
    }
}
