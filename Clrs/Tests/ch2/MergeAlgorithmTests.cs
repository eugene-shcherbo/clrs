using Clrs.ch2;
using NUnit.Framework;

namespace Tests.ch2;

[TestFixture]
public class SentitelMergeAlgorithmTests : MergeAlgorithmBaseTests
{
    protected override IMergeAlgorithm<int> CreateAlgorithm()
    {
        return new SentitelMergeAlgorithm<int>(int.MaxValue);
    }
}

[TestFixture]
public class DefaultMergeAlgorithmTests : MergeAlgorithmBaseTests
{
    protected override IMergeAlgorithm<int> CreateAlgorithm()
    {
        return new DefaultMergeAlgorithm<int>();
    }
}