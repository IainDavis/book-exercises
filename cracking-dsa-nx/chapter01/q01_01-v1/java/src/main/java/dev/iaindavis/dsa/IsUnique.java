package dev.iaindavis.dsa;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

public class IsUnique {
    public static boolean NestedLoops(String maybeUnique) {
        for(int base = 0; base < maybeUnique.length() - 1;  base++) {
            for(int compare = base + 1; compare < maybeUnique.length(); compare++) {
                if(maybeUnique.codePointAt(base) == maybeUnique.codePointAt(compare)) {
                    return false;
                }
            }
        }

        return true;
    }

    public static boolean withStreamsDistinct(String maybeUnique) {
        return maybeUnique.codePoints().distinct().count() == maybeUnique.codePointCount(0, maybeUnique.length());
    }

    public static boolean withSetStreamsToArray(String maybeUnique) {
        Set<Integer> seen = new HashSet<Integer>();
        for(Integer cp : maybeUnique.codePoints().toArray()) {
            if(seen.contains(cp)) return false;
            seen.add(cp);
        }
        return true;
    }

    public static boolean streamToSet(String maybeUnique) {
        int stringLength = maybeUnique.length();
        return maybeUnique.codePoints().boxed().collect(Collectors.toSet()).size() == stringLength;
    }
}