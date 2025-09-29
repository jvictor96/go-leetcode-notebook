import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

public class MaxFreq {
    public static void main(String[] args) {
        System.out.println(max_frequency_elements(Arrays.asList(1,2,3,2,1,2,1,2)));
    }

    static int max_frequency_elements(List<Integer> ints) {
        HashMap<Integer, Integer> freqs = new HashMap<>();
        int bestFreq = 0;
        int total = 0;
        for (Integer integer : ints) {
            freqs.put(integer, freqs.getOrDefault(integer, 0) + 1);
            if (freqs.get(integer) > bestFreq) {
                bestFreq = freqs.get(integer);
                total = bestFreq;
            } else if(freqs.get(integer) == bestFreq) {
                total += bestFreq;
            }
        }
        return total;
    }

    static int max_frequency_elements_by_stream(List<Integer> ints) {
    HashMap<Integer, Integer> freqs = new HashMap<>();
    for (Integer integer : ints) {
        freqs.put(integer, freqs.getOrDefault(integer, 0) + 1);
    }
    int bestFreq = freqs.values().stream().max(Integer::compare).orElse(0);
    return freqs.values().stream()
                 .filter(f -> f == bestFreq)
                 .mapToInt(f -> f)
                 .sum();
    }

}