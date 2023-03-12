using System;
using Clrs.ch2;
using NUnit.Framework;

namespace Tests.ch2;

[TestFixture]
public abstract class MergeAlgorithmBaseTests
{
    [Test]
    public void Merge_SameParts()
    {
        var expected = new[] { 8, 9, 9, 10, 11, 12, 1, 0, 12, 10 };
        var actual = new[] { 10, 11, 12, 8, 9, 9, 1, 0, 12, 10 };
        var algo = CreateAlgorithm();

        algo.Merge(actual.AsSpan(0, 6), actual.AsSpan(0, 3), actual.AsSpan(3, 3));

        Assert.AreEqual(expected, actual);
    }

    [Test]
    public void Merge_LeftGreater()
    {
        var expected = new[] { 8, 9, 9, 10, 11, 12, 15, 1, 0, 12, 10 };
        var actual = new[] { 8, 9, 10, 11, 12, 9, 15, 1, 0, 12, 10 };
        var algo = CreateAlgorithm();

        algo.Merge(actual.AsSpan(0, 7), actual.AsSpan(0, 5), actual.AsSpan(5, 2));

        Assert.AreEqual(expected, actual);
    }
    
    [Test]
    public void Merge_RightGreater()
    {
        var expected = new[] { 8, 9, 9, 11, 12, 15, 16, 18, 1, 0, 12, 10 };
        var actual = new[] { 11, 12, 8, 9, 9, 15, 16, 18, 1, 0, 12, 10 };
        var algo = CreateAlgorithm();

        algo.Merge(actual.AsSpan(0, 8), actual.AsSpan(0, 2), actual.AsSpan(2, 6));

        Assert.AreEqual(expected, actual);
    }

    protected abstract IMergeAlgorithm<int> CreateAlgorithm();
}