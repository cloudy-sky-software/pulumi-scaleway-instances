// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.PrivateNics
{
    /// <summary>
    /// The private NIC state
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1PrivateNICState : IEquatable<ScalewayInstanceV1PrivateNICState>
    {
        private readonly string _value;

        private ScalewayInstanceV1PrivateNICState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1PrivateNICState Available { get; } = new ScalewayInstanceV1PrivateNICState("available");
        public static ScalewayInstanceV1PrivateNICState Syncing { get; } = new ScalewayInstanceV1PrivateNICState("syncing");
        public static ScalewayInstanceV1PrivateNICState SyncingError { get; } = new ScalewayInstanceV1PrivateNICState("syncing_error");

        public static bool operator ==(ScalewayInstanceV1PrivateNICState left, ScalewayInstanceV1PrivateNICState right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1PrivateNICState left, ScalewayInstanceV1PrivateNICState right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1PrivateNICState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1PrivateNICState other && Equals(other);
        public bool Equals(ScalewayInstanceV1PrivateNICState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
