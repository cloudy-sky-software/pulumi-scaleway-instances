// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.UserData
{
    public static class GetServerUserData
    {
        public static Task<Outputs.ScalewayStdFile> InvokeAsync(GetServerUserDataArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayStdFile>("scaleway-instances:user_data:getServerUserData", args ?? new GetServerUserDataArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayStdFile> Invoke(GetServerUserDataInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayStdFile>("scaleway-instances:user_data:getServerUserData", args ?? new GetServerUserDataInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerUserDataArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key of the user data to get
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// UUID of the server
        /// </summary>
        [Input("serverId", required: true)]
        public string ServerId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetServerUserDataArgs()
        {
        }
        public static new GetServerUserDataArgs Empty => new GetServerUserDataArgs();
    }

    public sealed class GetServerUserDataInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key of the user data to get
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// UUID of the server
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetServerUserDataInvokeArgs()
        {
        }
        public static new GetServerUserDataInvokeArgs Empty => new GetServerUserDataInvokeArgs();
    }
}
