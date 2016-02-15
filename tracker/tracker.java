// LargestTracker.java
import java.util.List;

public class LargestTracker {
    /**
     * guarantees the creation of a single instance across the virtual machine.
     * assumed to be called very frequently.
     * 
     * @return an instance of largesttracker
     */
    static LargestTracker getInstance();

    /**
     * Returns a list in O(n log m) time OR BETTER where n is the number of entries 
     * added to LargestTracker and m is numberOfTopLargestElements. Duplicates are allowed
     * 
     * @param numberOfTopLargestElements
     *            the number of top-most-elements to return
     * @return the top-most-elements in the tracker sorted in ascending order
     */
    List<Integer> getNLargest(int numberOfTopLargestElements);

    /**
     * Adds an entry to the tracker. This method must operate in O(log n) time
     * OR BETTER
     * @param anEntry
     *            the entry to add to the tracker. Entries need not be unique.
     */
    void add(int anEntry);

    /**
     * Removes all the entries from the tracker. This should return in constant
     * time.
     */
    void clear();
}
// end LargestTracker.java
