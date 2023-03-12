using System;

namespace Clrs.ch2;

public class DefaultMergeAlgorithm<T> : IMergeAlgorithm<T>
{
    private readonly Comparer<T> _comparer;

    public DefaultMergeAlgorithm()
    {
        _comparer = Comparer<T>.Default;
    }

    public void Merge(Span<T> subArray, ReadOnlySpan<T> leftPart, ReadOnlySpan<T> rightPart)
    {
        var left = new T[leftPart.Length];
        var right = new T[rightPart.Length];

        leftPart.CopyTo(left);
        rightPart.CopyTo(right);

        int leftIdx = 0;
        int rightIdx = 0;

        for (int i = 0; i < subArray.Length; i++)
        {
            if (leftIdx == left.Length)
                subArray[i] = right[rightIdx++];
            else if (rightIdx == right.Length)
                subArray[i] = left[leftIdx++];
            else if (_comparer.Compare(left[leftIdx], right[rightIdx]) is -1 or 0)
                subArray[i] = left[leftIdx++];
            else
                subArray[i] = right[rightIdx++];
        }
    }
}