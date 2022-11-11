// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.SecurityGroups
{
    /// <summary>
    /// The default inbound policy
    /// </summary>
    [EnumType]
    public readonly struct InboundDefaultPolicy : IEquatable<InboundDefaultPolicy>
    {
        private readonly string _value;

        private InboundDefaultPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InboundDefaultPolicy Accept { get; } = new InboundDefaultPolicy("accept");
        public static InboundDefaultPolicy Drop { get; } = new InboundDefaultPolicy("drop");

        public static bool operator ==(InboundDefaultPolicy left, InboundDefaultPolicy right) => left.Equals(right);
        public static bool operator !=(InboundDefaultPolicy left, InboundDefaultPolicy right) => !left.Equals(right);

        public static explicit operator string(InboundDefaultPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InboundDefaultPolicy other && Equals(other);
        public bool Equals(InboundDefaultPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default outbound policy
    /// </summary>
    [EnumType]
    public readonly struct OutboundDefaultPolicy : IEquatable<OutboundDefaultPolicy>
    {
        private readonly string _value;

        private OutboundDefaultPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OutboundDefaultPolicy Accept { get; } = new OutboundDefaultPolicy("accept");
        public static OutboundDefaultPolicy Drop { get; } = new OutboundDefaultPolicy("drop");

        public static bool operator ==(OutboundDefaultPolicy left, OutboundDefaultPolicy right) => left.Equals(right);
        public static bool operator !=(OutboundDefaultPolicy left, OutboundDefaultPolicy right) => !left.Equals(right);

        public static explicit operator string(OutboundDefaultPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OutboundDefaultPolicy other && Equals(other);
        public bool Equals(OutboundDefaultPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default inbound policy
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupInboundDefaultPolicy : IEquatable<ScalewayInstanceV1SecurityGroupInboundDefaultPolicy>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupInboundDefaultPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupInboundDefaultPolicy Accept { get; } = new ScalewayInstanceV1SecurityGroupInboundDefaultPolicy("accept");
        public static ScalewayInstanceV1SecurityGroupInboundDefaultPolicy Drop { get; } = new ScalewayInstanceV1SecurityGroupInboundDefaultPolicy("drop");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupInboundDefaultPolicy left, ScalewayInstanceV1SecurityGroupInboundDefaultPolicy right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupInboundDefaultPolicy left, ScalewayInstanceV1SecurityGroupInboundDefaultPolicy right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupInboundDefaultPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupInboundDefaultPolicy other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupInboundDefaultPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default outbound policy
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy : IEquatable<ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy Accept { get; } = new ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy("accept");
        public static ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy Drop { get; } = new ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy("drop");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy left, ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy left, ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupOutboundDefaultPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Security group state
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupState : IEquatable<ScalewayInstanceV1SecurityGroupState>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupState Available { get; } = new ScalewayInstanceV1SecurityGroupState("available");
        public static ScalewayInstanceV1SecurityGroupState Syncing { get; } = new ScalewayInstanceV1SecurityGroupState("syncing");
        public static ScalewayInstanceV1SecurityGroupState SyncingError { get; } = new ScalewayInstanceV1SecurityGroupState("syncing_error");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupState left, ScalewayInstanceV1SecurityGroupState right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupState left, ScalewayInstanceV1SecurityGroupState right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupState other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Security group state
    /// </summary>
    [EnumType]
    public readonly struct State : IEquatable<State>
    {
        private readonly string _value;

        private State(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static State Available { get; } = new State("available");
        public static State Syncing { get; } = new State("syncing");
        public static State SyncingError { get; } = new State("syncing_error");

        public static bool operator ==(State left, State right) => left.Equals(right);
        public static bool operator !=(State left, State right) => !left.Equals(right);

        public static explicit operator string(State value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is State other && Equals(other);
        public bool Equals(State other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
