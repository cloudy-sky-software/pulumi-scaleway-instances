// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.Export
{
    /// <summary>
    /// The task status
    /// </summary>
    [EnumType]
    public readonly struct ExportSnapshotScalewayInstanceV1TaskStatus : IEquatable<ExportSnapshotScalewayInstanceV1TaskStatus>
    {
        private readonly string _value;

        private ExportSnapshotScalewayInstanceV1TaskStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ExportSnapshotScalewayInstanceV1TaskStatus Pending { get; } = new ExportSnapshotScalewayInstanceV1TaskStatus("pending");
        public static ExportSnapshotScalewayInstanceV1TaskStatus Started { get; } = new ExportSnapshotScalewayInstanceV1TaskStatus("started");
        public static ExportSnapshotScalewayInstanceV1TaskStatus Success { get; } = new ExportSnapshotScalewayInstanceV1TaskStatus("success");
        public static ExportSnapshotScalewayInstanceV1TaskStatus Failure { get; } = new ExportSnapshotScalewayInstanceV1TaskStatus("failure");
        public static ExportSnapshotScalewayInstanceV1TaskStatus Retry { get; } = new ExportSnapshotScalewayInstanceV1TaskStatus("retry");

        public static bool operator ==(ExportSnapshotScalewayInstanceV1TaskStatus left, ExportSnapshotScalewayInstanceV1TaskStatus right) => left.Equals(right);
        public static bool operator !=(ExportSnapshotScalewayInstanceV1TaskStatus left, ExportSnapshotScalewayInstanceV1TaskStatus right) => !left.Equals(right);

        public static explicit operator string(ExportSnapshotScalewayInstanceV1TaskStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExportSnapshotScalewayInstanceV1TaskStatus other && Equals(other);
        public bool Equals(ExportSnapshotScalewayInstanceV1TaskStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
