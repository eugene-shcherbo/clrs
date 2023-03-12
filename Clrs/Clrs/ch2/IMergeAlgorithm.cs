namespace Clrs.ch2;

public interface IMergeAlgorithm<T>
{   
    void Merge(Span<T> subArray, ReadOnlySpan<T> leftPart, ReadOnlySpan<T> rightPart);
}