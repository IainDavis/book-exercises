package dev.iaindavis.dsa;

import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.api.DynamicTest;
import org.junit.jupiter.api.TestFactory;

import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Stream;

public class IsUniqueTest {

    private static final Map<String, Function<String, Boolean>> implementations = Map.of(
        "Brute force implementation with nested loops", IsUnique::NestedLoops,
        "Using streams with distinct and count", IsUnique::withStreamsDistinct,
        "Using streams to array and set comparison", IsUnique::withSetStreamsToArray
    );

    private static final List<Map.Entry<String, Boolean>> testCases = List.of(
        Map.entry("abc", true),
        Map.entry("aabc", false),
        Map.entry("🙉💩", true),
        Map.entry("🙉🙉", false)
    );

    @TestFactory
    Stream<DynamicTest> tests() {
        return implementations.entrySet().stream()
            .flatMap(entry -> {
                String description = entry.getKey();
                Function<String, Boolean> impl = entry.getValue();
                return testCases.stream()
                    .map(testCase -> DynamicTest.dynamicTest(
                        String.format("%s: input='%s' → expected=%s", description, testCase.getKey(), testCase.getValue()),
                        () -> assertEquals(testCase.getValue(), impl.apply(testCase.getKey()))
                    ));
            });
    }
}
