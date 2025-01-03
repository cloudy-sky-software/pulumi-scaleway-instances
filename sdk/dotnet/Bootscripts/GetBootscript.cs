// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Bootscripts
{
    public static class GetBootscript
    {
        public static Task<Outputs.ScalewayInstanceV1GetBootscriptResponse> InvokeAsync(GetBootscriptArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1GetBootscriptResponse>("scaleway-instances:bootscripts:getBootscript", args ?? new GetBootscriptArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1GetBootscriptResponse> Invoke(GetBootscriptInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1GetBootscriptResponse>("scaleway-instances:bootscripts:getBootscript", args ?? new GetBootscriptInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1GetBootscriptResponse> Invoke(GetBootscriptInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1GetBootscriptResponse>("scaleway-instances:bootscripts:getBootscript", args ?? new GetBootscriptInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBootscriptArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetBootscriptArgs()
        {
        }
        public static new GetBootscriptArgs Empty => new GetBootscriptArgs();
    }

    public sealed class GetBootscriptInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetBootscriptInvokeArgs()
        {
        }
        public static new GetBootscriptInvokeArgs Empty => new GetBootscriptInvokeArgs();
    }
}
