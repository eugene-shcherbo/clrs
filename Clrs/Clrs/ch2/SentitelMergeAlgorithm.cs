namespace Clrs.ch2;

public class SentitelMergeAlgorithm<T> : IMergeAlgorithm<T>
{
    private readonly T _sentitel;
    private readonly IComparer<T> _comparer;

    public SentitelMergeAlgorithm(T sentitel)
    {
        _sentitel = sentitel;
        _comparer = Comparer<T>.Default;
    }

    public void Merge(Span<T> subArray, ReadOnlySpan<T> leftPart, ReadOnlySpan<T> rightPart)
    {
        var left = new T[leftPart.Length + 1];
        var right = new T[rightPart.Length + 1];

        leftPart.CopyTo(left);
        rightPart.CopyTo(right);

        left[^1] = _sentitel;
        right[^1] = _sentitel;

        var leftIdx = 0;
        var rightIdx = 0;

        for (var i = 0; i < subArray.Length; i++)
        {
            if (_comparer.Compare(left[leftIdx], right[rightIdx]) is -1 or 0)
                subArray[i] = left[leftIdx++];
            else
                subArray[i] = right[rightIdx++];
        }
    }
}