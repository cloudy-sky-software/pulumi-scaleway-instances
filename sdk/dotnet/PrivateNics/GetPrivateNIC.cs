// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.PrivateNics
{
    public static class GetPrivateNIC
    {
        public static Task<GetPrivateNICResult> InvokeAsync(GetPrivateNICArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateNICResult>("scaleway-instances:private_nics:getPrivateNIC", args ?? new GetPrivateNICArgs(), options.WithDefaults());

        public static Output<GetPrivateNICResult> Invoke(GetPrivateNICInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateNICResult>("scaleway-instances:private_nics:getPrivateNIC", args ?? new GetPrivateNICInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateNICArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("server_id", required: true)]
        public string Server_id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetPrivateNICArgs()
        {
        }
        public static new GetPrivateNICArgs Empty => new GetPrivateNICArgs();
    }

    public sealed class GetPrivateNICInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("server_id", required: true)]
        public Input<string> Server_id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetPrivateNICInvokeArgs()
        {
        }
        public static new GetPrivateNICInvokeArgs Empty => new GetPrivateNICInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateNICResult
    {
        public readonly Outputs.ScalewayInstanceV1GetPrivateNICResponse Items;

        [OutputConstructor]
        private GetPrivateNICResult(Outputs.ScalewayInstanceV1GetPrivateNICResponse items)
        {
            Items = items;
        }
    }
}
