// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.User_data.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ListServerUserDataResponse
    {
        public readonly ImmutableArray<string> User_data;

        [OutputConstructor]
        private ScalewayInstanceV1ListServerUserDataResponse(ImmutableArray<string> user_data)
        {
            User_data = user_data;
        }
    }
}
