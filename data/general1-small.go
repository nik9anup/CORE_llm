   
import (
    "math"
)

// Check if in given list of numbers, are any two numbers closer to each other than given threshold.
// >>> HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5)
// false
// >>> HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3)
// true


func HasCloseElements(numbers []float64, threshold float64) bool {

    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            var distance float64 = math.Abs(numbers[i] - numbers[j])
            if distance < threshold {
                return true
            }
        }
    }
    return false
}



   

// Input to this function is a string containing multiple groups of nested parentheses. Your goal is to
// separate those group into separate strings and return the list of those.
// Separate groups are balanced (each open brace is properly closed) and not nested within each other
// Ignore any spaces in the input string.
// >>> SeparateParenGroups('( ) (( )) (( )( ))')
// ['()', '(())', '(()())']


func SeparateParenGroups(paren_string string) []string {

    result := make([]string, 0)
    current_string := make([]rune, 0)
    current_depth := 0

    for _, c := range paren_string {
        if c == '(' {
            current_depth += 1
            current_string = append(current_string, c)
        }else if c== ')'{
            current_depth -= 1
            current_string = append(current_string, c)

            if current_depth == 0{
                result = append(result, string(current_string))
                current_string =  make([]rune, 0)
            }
        }

    }
    return result
}



   
import (
    "math"
)

// Given a positive floating point number, it can be decomposed into
// and integer part (largest integer smaller than given number) and decimals
// (leftover part always smaller than 1).
// 
// Return the decimal part of the number.
// >>> TruncateNumber(3.5)
// 0.5


func TruncateNumber(number float64) float64 {

    return math.Mod(number,1)
}



   

// You're given a list of deposit and withdrawal operations on a bank account that starts with
// zero balance. Your task is to detect if at any point the balance of account fallls below zero, and
// at that point function should return true. Otherwise it should return false.
// >>> BelowZero([1, 2, 3])
// false
// >>> BelowZero([1, 2, -4, 5])
// true


func BelowZero(operations []int) bool {

    balance := 0
    for _, op := range operations {
        balance += op
        if balance < 0 {
            return true
        }
    }
    return false
}



   
import (
    "math"
)

// For a given list of input numbers, calculate Mean Absolute Deviation
// around the mean of this dataset.
// Mean Absolute Deviation is the average absolute difference between each
// element and a centerpoint (mean in this case):
// MAD = average | x - x_mean |
// >>> MeanAbsoluteDeviation([1.0, 2.0, 3.0, 4.0])
// 1.0


func MeanAbsoluteDeviation(numbers []float64) float64 {

    sum := func(numbers []float64) float64 {
        sum := 0.0
        for _, num := range numbers {
            sum += num
        }
        return sum
    }

    mean := sum(numbers) / float64(len(numbers))
    numList := make([]float64, 0)
    for _, x := range numbers {
        numList = append(numList, math.Abs(x-mean))
    }
    return sum(numList) / float64(len(numbers))
}



   

// Insert a number 'delimeter' between every two consecutive elements of input list `numbers'
// >>> Intersperse([], 4)
// []
// >>> Intersperse([1, 2, 3], 4)
// [1, 4, 2, 4, 3]


func Intersperse(numbers []int, delimeter int) []int {

    result := make([]int, 0)
    if len(numbers) == 0 {
        return result
    }
    for i := 0; i < len(numbers)-1; i++ {
        n := numbers[i]
        result = append(result, n)
        result = append(result, delimeter)
    }
    result = append(result, numbers[len(numbers)-1])
    return result
}



   
import (
    "math"
    "strings"
)

// Input to this function is a string represented multiple groups for nested parentheses separated by spaces.
// For each of the group, output the deepest level of nesting of parentheses.
// E.g. (()()) has maximum two levels of nesting while ((())) has three.
// 
// >>> ParseNestedParens('(()()) ((())) () ((())()())')
// [2, 3, 1, 3]


func ParseNestedParens(paren_string string) []int {

    parse_paren_group := func(s string) int {
        depth := 0
        max_depth := 0
        for _, c := range s {
            if c == '(' {
                depth += 1
                max_depth = int(math.Max(float64(depth), float64(max_depth)))
            } else {
                depth -= 1
            }
        }
        return max_depth
    }
    result := make([]int, 0)
    for _, x := range strings.Split(paren_string, " ") {
        result = append(result, parse_paren_group(x))
    }
    return result

}



   
import (
    "strings"
)

// Filter an input list of strings only for ones that contain given substring
// >>> FilterBySubstring([], 'a')
// []
// >>> FilterBySubstring(['abc', 'bacd', 'cde', 'array'], 'a')
// ['abc', 'bacd', 'array']


func FilterBySubstring(stringList []string, substring string) []string {

    result := make([]string, 0)
    for _, x := range stringList {
        if strings.Index(x, substring) != -1 {
            result = append(result, x)
        }
    }
    return result
}



   

// For a given list of integers, return a tuple consisting of a sum and a product of all the integers in a list.
// Empty sum should be equal to 0 and empty product should be equal to 1.
// >>> SumProduct([])
// (0, 1)
// >>> SumProduct([1, 2, 3, 4])
// (10, 24)


func SumProduct(numbers []int) [2]int {

    sum_value := 0
    prod_value := 1

    for _, n := range numbers {
        sum_value += n
        prod_value *= n
    }
    return [2]int{sum_value, prod_value}
}



   
import (
    "math"
)

// From a given list of integers, generate a list of rolling maximum element found until given moment
// in the sequence.
// >>> RollingMax([1, 2, 3, 2, 3, 4, 2])
// [1, 2, 3, 3, 3, 4, 4]


func RollingMax(numbers []int) []int {

    running_max := math.MinInt32
    result := make([]int, 0)

    for _, n := range numbers {
        if running_max == math.MinInt32 {
            running_max = n
        } else {
            running_max = int(math.Max(float64(running_max), float64(n)))
        }
        result = append(result, running_max)
    }

    return result
}



   
import (
    "strings"
)

// Find the shortest palindrome that begins with a supplied string.
// Algorithm idea is simple:
// - Find the longest postfix of supplied string that is a palindrome.
// - Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
// >>> MakePalindrome('')
// ''
// >>> MakePalindrome('cat')
// 'catac'
// >>> MakePalindrome('cata')
// 'catac'


func MakePalindrome(str string) string {

    if strings.TrimSpace(str) == "" {
        return ""
    }
    beginning_of_suffix := 0
    runes := []rune(str)
    for !IsPalindrome(string(runes[beginning_of_suffix:])) {
        beginning_of_suffix += 1
    }
    result := make([]rune, 0)
    for i := len(str[:beginning_of_suffix]) - 1; i >= 0; i-- {
        result = append(result, runes[i])
    }
    return str + string(result)
}



   
import (
    "fmt"
)

// Input are two strings a and b consisting only of 1s and 0s.
// Perform binary XOR on these inputs and return result also as a string.
// >>> StringXor('010', '110')
// '100'


func StringXor(a string, b string) string {

    s2b := func(bs string) int32 {
        result := int32(0)
        runes := []rune(bs)
        for _, r := range runes {
            result = result << 1
            temp := r - rune('0')
            result += temp
        }
        return result
    }
    ab := s2b(a)
    bb := s2b(b)
    res := ab ^ bb
    sprint := fmt.Sprintf("%b", res)
    for i := 0; i < len(a)-len(sprint); i++ {
        sprint = "0" + sprint
    }
    return sprint
}



   

// Out of list of strings, return the Longest one. Return the first one in case of multiple
// strings of the same length. Return nil in case the input list is empty.
// >>> Longest([])
// nil
// >>> Longest(['a', 'b', 'c'])
// 'a'
// >>> Longest(['a', 'bb', 'ccc'])
// 'ccc'


func Longest(strings []string) interface{}{

    if strings == nil || len(strings) == 0 {
        return nil
    }
    maxlen := 0
    maxi := 0
    for i, s := range strings {
        if maxlen < len(s) {
            maxlen = len(s)
            maxi = i
        }
    }
    return strings[maxi]
}



   

// Return a greatest common divisor of two integers a and b
// >>> GreatestCommonDivisor(3, 5)
// 1
// >>> GreatestCommonDivisor(25, 15)
// 5


func GreatestCommonDivisor(a int,b int) int{

    if b < 2 {
		return b
	}
	var gcd int = 1
	for i := 2; i < b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}



   

// Return list of all prefixes from shortest to longest of the input string
// >>> AllPrefixes('abc')
// ['a', 'ab', 'abc']


func AllPrefixes(str string) []string{

    prefixes := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		prefixes = append(prefixes, str[:i+1])
	}
	return prefixes
}



   
import (
    "strconv"
)

// Return a string containing space-delimited numbers starting from 0 upto n inclusive.
// >>> StringSequence(0)
// '0'
// >>> StringSequence(5)
// '0 1 2 3 4 5'


func StringSequence(n int) string{

    var seq string
    for i := 0; i <= n; i++ {
        seq += strconv.Itoa(i)
        if i != n {
            seq += " "
        }
    }
    return seq
}



   
import (
    "strings"
)

// Given a string, find out how many distinct characters (regardless of case) does it consist of
// >>> CountDistinctCharacters('xyzXYZ')
// 3
// >>> CountDistinctCharacters('Jerry')
// 4


func CountDistinctCharacters(str string) int{

    lower := strings.ToLower(str)
	count := 0
	set := make(map[rune]bool)
	for _, i := range lower {
		if set[i] == true {
			continue
		} else {
			set[i] = true
			count++
		}
	}
	return count
}




   

// Input to this function is a string representing musical notes in a special ASCII format.
// Your task is to parse this string and return list of integers corresponding to how many beats does each
// not last.
// 
// Here is a legend:
// 'o' - whole note, lasts four beats
// 'o|' - half note, lasts two beats
// '.|' - quater note, lasts one beat
// 
// >>> ParseMusic('o o| .| o| o| .| .| .| .| o o')
// [4, 2, 1, 2, 2, 1, 1, 1, 1, 4, 4]


func ParseMusic(music_string string) []int{

    note_map := map[string]int{"o": 4, "o|": 2, ".|": 1}
	split := strings.Split(music_string, " ")
	result := make([]int, 0)
	for _, x := range split {
		if i, ok := note_map[x]; ok {
			result = append(result, i)
		}
	}
	return result
}




   


// Find how many times a given substring can be found in the original string. Count overlaping cases.
// >>> HowManyTimes('', 'a')
// 0
// >>> HowManyTimes('aaa', 'a')
// 3
// >>> HowManyTimes('aaaa', 'aa')
// 3


func HowManyTimes(str string,substring string) int{

    times := 0
	for i := 0; i < (len(str) - len(substring) + 1); i++ {
		if str[i:i+len(substring)] == substring {
			times += 1
		}
	}
	return times
}




   
import (
    "sort"
    "strings"
)
// Input is a space-delimited string of numberals from 'zero' to 'nine'.
// Valid choices are 'zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight' and 'nine'.
// Return the string with numbers sorted from smallest to largest
// >>> SortNumbers('three one five')
// 'one three five'


func SortNumbers(numbers string) string{

    valueMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	stringMap := make(map[int]string)
	for s, i := range valueMap {
		stringMap[i] = s
	}
	split := strings.Split(numbers, " ")
	temp := make([]int, 0)
	for _, s := range split {
		if i, ok := valueMap[s]; ok {
			temp = append(temp, i)
		}
	}
	sort.Ints(temp)
	result := make([]string, 0)
	for _, i := range temp {
		result = append(result, stringMap[i])
	}
	return strings.Join(result, " ")
}




   

// From a supplied list of numbers (of length at least two) select and return two that are the closest to each
// other and return them in order (smaller number, larger number).
// >>> FindClosestElements([1.0, 2.0, 3.0, 4.0, 5.0, 2.2])
// (2.0, 2.2)
// >>> FindClosestElements([1.0, 2.0, 3.0, 4.0, 5.0, 2.0])
// (2.0, 2.0)


func FindClosestElements(numbers []float64) [2]float64 {

    distance := math.MaxFloat64
	var closestPair [2]float64
	for idx, elem := range numbers {
		for idx2, elem2 := range numbers {
			if idx != idx2 {
				if distance == math.MinInt64 {
					distance = math.Abs(elem - elem2)
					float64s := []float64{elem, elem2}
					sort.Float64s(float64s)
					closestPair = [2]float64{float64s[0], float64s[1]}
				} else {
					newDistance := math.Abs(elem - elem2)
					if newDistance < distance{
						distance = newDistance
						float64s := []float64{elem, elem2}
						sort.Float64s(float64s)
						closestPair = [2]float64{float64s[0], float64s[1]}
					}
				}
			}
		}
	}
	return closestPair
}




   

// Given list of numbers (of at least two elements), apply a linear transform to that list,
// such that the smallest number will become 0 and the largest will become 1
// >>> RescaleToUnit([1.0, 2.0, 3.0, 4.0, 5.0])
// [0.0, 0.25, 0.5, 0.75, 1.0]


func RescaleToUnit(numbers []float64) []float64 {

    smallest := numbers[0]
	largest := smallest
	for _, n := range numbers {
		if smallest > n {
			smallest = n
		}
		if largest < n {
			largest = n
		}
	}
	if smallest == largest {
		return numbers
	}
	for i, n := range numbers {
		numbers[i] = (n - smallest) / (largest - smallest)
	}
	return numbers
}



   

// Filter given list of any values only for integers
// >>> FilterIntegers(['a', 3.14, 5])
// [5]
// >>> FilterIntegers([1, 2, 3, 'abc', {}, []])
// [1, 2, 3]


func FilterIntegers(values []interface{}) []int {

    result := make([]int, 0)
    for _, val := range values {
        switch i := val.(type) {
        case int:
            result = append(result, i)
        }
    }
    return result
}



   

// Return length of given string
// >>> Strlen('')
// 0
// >>> Strlen('abc')
// 3


func Strlen(str string) int {

    return len(str)
}



   

// For a given number n, find the largest number that divides n evenly, smaller than n
// >>> LargestDivisor(15)
// 5


func LargestDivisor(n int) int {

    for i := n - 1; i > 0; i-- {
		if n % i == 0 {
			return i
		}
	}
	return 0
}



   
import (
    "math"
)
// Return list of prime factors of given integer in the order from smallest to largest.
// Each of the factors should be listed number of times corresponding to how many times it appeares in factorization.
// Input number should be equal to the product of all factors
// >>> Factorize(8)
// [2, 2, 2]
// >>> Factorize(25)
// [5, 5]
// >>> Factorize(70)
// [2, 5, 7]


func Factorize(n int) []int {

    fact := make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))+1); {
		if n%i == 0 {
			fact = append(fact, i)
			n = n / i
		} else {
			i++
		}
	}
	if n > 1 {
		fact = append(fact, n)
	}
	return fact
}



   

// From a list of integers, remove all elements that occur more than once.
// Keep order of elements left the same as in the input.
// >>> RemoveDuplicates([1, 2, 3, 2, 4])
// [1, 3, 4]


func RemoveDuplicates(numbers []int) []int {

    c := make(map[int] int)
	for _, number := range numbers {
		if i, ok := c[number]; ok {
			c[number] = i + 1
		} else {
			c[number] = 1
		}
	}
	result := make([]int, 0)
	for _, number := range numbers {
		if c[number] <= 1 {
			result = append(result, number)
		}
	}
	return result
}



   
import (
    "strings"
)

// For a given string, flip lowercase characters to uppercase and uppercase to lowercase.
// >>> FlipCase('Hello')
// 'hELLO'


func FlipCase(str string) string {

    result := []rune{}
    for _, c := range str {
        if c >= 'A' && c <= 'Z' {
            result = append(result, 'a' + ((c - 'A' + 26) % 26))
        } else if c >= 'a' && c <= 'z' {
            result = append(result, 'A' + ((c - 'a' + 26) % 26))
        } else {
            result = append(result, c)
        }
    }
    return string(result)
}



   

// Concatenate list of strings into a single string
// >>> Concatenate([])
// ''
// >>> Concatenate(['a', 'b', 'c'])
// 'abc'


func Concatenate(strings []string) string {

    if len(strings) == 0 {
		return ""
	}
	return strings[0] + Concatenate(strings[1:])
}



   

// Filter an input list of strings only for ones that start with a given prefix.
// >>> FilterByPrefix([], 'a')
// []
// >>> FilterByPrefix(['abc', 'bcd', 'cde', 'array'], 'a')
// ['abc', 'array']


func FilterByPrefix(strings []string,prefix string) []string {

    if len(strings) == 0 {
        return []string{}
    }
    res := make([]string, 0, len(strings))
	for _, s := range strings {
		if s[:len(prefix)] == prefix {
			res = append(res, s)
		}
	}
	return res
}




   

// Return only positive numbers in the list.
// >>> GetPositive([-1, 2, -4, 5, 6])
// [2, 5, 6]
// >>> GetPositive([5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10])
// [5, 3, 2, 3, 9, 123, 1]


func GetPositive(l []int) []int {

    res := make([]int, 0)
    for _, x := range l {
        if x > 0 {
            res = append(res, x)
        }
    }
    return res
}




   

// Return true if a given number is prime, and false otherwise.
// >>> IsPrime(6)
// false
// >>> IsPrime(101)
// true
// >>> IsPrime(11)
// true
// >>> IsPrime(13441)
// true
// >>> IsPrime(61)
// true
// >>> IsPrime(4)
// false
// >>> IsPrime(1)
// false


func IsPrime(n int) bool {

    if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}




   
import (
    "math"
)

// xs are coefficients of a polynomial.
// FindZero find x such that Poly(x) = 0.
// FindZero returns only only zero point, even if there are many.
// Moreover, FindZero only takes list xs having even number of coefficients
// and largest non zero coefficient as it guarantees
// a solution.
// >>> round(FindZero([1, 2]), 2) # f(x) = 1 + 2x
// -0.5
// >>> round(FindZero([-6, 11, -6, 1]), 2) # (x - 1) * (x - 2) * (x - 3) = -6 + 11x - 6x^2 + x^3
// 1.0


func FindZero(xs []int) float64 {

    begin := -1.0
	end := 1.0
	for Poly(xs, begin)*Poly(xs, end) > 0 {
		begin *= 2
		end *= 2
	}
	for end-begin > 1e-10 {
		center := (begin + end) / 2
		if Poly(xs, center)*Poly(xs, begin) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return begin
}



   
import (
    "sort"
)
// This function takes a list l and returns a list l' such that
// l' is identical to l in the indicies that are not divisible by three, while its values at the indicies that are divisible by three are equal
// to the values of the corresponding indicies of l, but sorted.
// >>> SortThird([1, 2, 3])
// [1, 2, 3]
// >>> SortThird([5, 6, 3, 4, 8, 9, 2])
// [2, 6, 3, 4, 8, 9, 5]


func SortThird(l []int) []int {

    temp := make([]int, 0)
	for i := 0; i < len(l); i = i + 3 {
		temp = append(temp, l[i])
	}
	sort.Ints(temp)
	j := 0
	for i := 0; i < len(l); i = i + 3 {
		l[i] = temp[j]
		j++
	}
	return l
}



   
import (
    "sort"
)
// Return sorted Unique elements in a list
// >>> Unique([5, 3, 5, 2, 3, 3, 9, 0, 123])
// [0, 2, 3, 5, 9, 123]


func Unique(l []int) []int {

    set := make(map[int]interface{})
	for _, i := range l {
		set[i]=nil
	}
	l = make([]int,0)
	for i, _ := range set {
		l = append(l, i)
	}
	sort.Ints(l)
	return l
}



   

// Return maximum element in the list.
// >>> MaxElement([1, 2, 3])
// 3
// >>> MaxElement([5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10])
// 123


func MaxElement(l []int) int {

    max := l[0]
	for _, x := range l {
		if x > max {
			max = x
		}
	}
	return max
}



   
import (
	"strconv"
	"strings"
)
// Return the number of times the digit 7 appears in integers less than n which are divisible by 11 or 13.
// >>> FizzBuzz(50)
// 0
// >>> FizzBuzz(78)
// 2
// >>> FizzBuzz(79)
// 3


func FizzBuzz(n int) int {

    ns := make([]int, 0)
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, i)
		}
	}
	temp := make([]string, 0)
	for _, i := range ns {
		temp = append(temp, strconv.Itoa(i))
	}
	join := strings.Join(temp, "")
	ans := 0
	for _, c := range join {
		if c == '7' {
			ans++
		}
	}
	return ans
}



   
import (
	"sort"
)
// This function takes a list l and returns a list l' such that
// l' is identical to l in the odd indicies, while its values at the even indicies are equal
// to the values of the even indicies of l, but sorted.
// >>> SortEven([1, 2, 3])
// [1, 2, 3]
// >>> SortEven([5, 6, 3, 4])
// [3, 6, 5, 4]


func SortEven(l []int) []int {

    evens := make([]int, 0)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}
	sort.Ints(evens)
	j := 0
	for i := 0; i < len(l); i += 2 {
		l[i] = evens[j]
		j++
	}
	return l
}



   
import (
    "math"
    "strings"
    "time"
)

// returns encoded string by cycling groups of three characters.
// takes as input string encoded with EncodeCyclic function. Returns decoded string.


func DecodeCyclic(s string) string {

    return EncodeCyclic(EncodeCyclic(s))
}



   
import (
	"math"
)
// PrimeFib returns n-th number that is a Fibonacci number and it's also prime.
// >>> PrimeFib(1)
// 2
// >>> PrimeFib(2)
// 3
// >>> PrimeFib(3)
// 5
// >>> PrimeFib(4)
// 13
// >>> PrimeFib(5)
// 89


func PrimeFib(n int) int {

    isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		for i := 2; i < int(math.Min(math.Sqrt(float64(p))+1, float64(p-1))); i++ {
			if p%i == 0 {
				return false
			}
		}
		return true
	}
	f := []int{0, 1}
	for {
		f = append(f, f[len(f)-1]+f[len(f)-2])
		if isPrime(f[len(f)-1]) {
			n -= 1
		}
		if n == 0 {
			return f[len(f)-1]
		}
	}
}



   

// TriplesSumToZero takes a list of integers as an input.
// it returns true if there are three distinct elements in the list that
// sum to zero, and false otherwise.
// 
// >>> TriplesSumToZero([1, 3, 5, 0])
// false
// >>> TriplesSumToZero([1, 3, -2, 1])
// true
// >>> TriplesSumToZero([1, 2, 3, 7])
// false
// >>> TriplesSumToZero([2, 4, -5, 3, 9, 7])
// true
// >>> TriplesSumToZero([1])
// false


func TriplesSumToZero(l []int) bool {

    for i := 0; i < len(l) - 2; i++ {
		for j := i + 1; j < len(l) - 1; j++ {
			for k := j + 1; k < len(l); k++ {
				if l[i] + l[j] + l[k] == 0 {
					return true
				}
			}
		}
	}
	return false
}



   

// Imagine a road that's a perfectly straight infinitely long line.
// n cars are driving left to right;  simultaneously, a different set of n cars
// are driving right to left.   The two sets of cars start out being very far from
// each other.  All cars move in the same speed.  Two cars are said to collide
// when a car that's moving left to right hits a car that's moving right to left.
// However, the cars are infinitely sturdy and strong; as a result, they continue moving
// in their trajectory as if they did not collide.
// 
// This function outputs the number of such collisions.


func CarRaceCollision(n int) int {

	return n * n
}



   

// Return list with elements incremented by 1.
// >>> IncrList([1, 2, 3])
// [2, 3, 4]
// >>> IncrList([5, 3, 5, 2, 3, 3, 9, 0, 123])
// [6, 4, 6, 3, 4, 4, 10, 1, 124]


func IncrList(l []int) []int {

    n := len(l)
	for i := 0; i < n; i++ {
		l[i]++
	}
	return l
}



   

// PairsSumToZero takes a list of integers as an input.
// it returns true if there are two distinct elements in the list that
// sum to zero, and false otherwise.
// >>> PairsSumToZero([1, 3, 5, 0])
// false
// >>> PairsSumToZero([1, 3, -2, 1])
// false
// >>> PairsSumToZero([1, 2, 3, 7])
// false
// >>> PairsSumToZero([2, 4, -5, 3, 5, 7])
// true
// >>> PairsSumToZero([1])
// false


func PairsSumToZero(l []int) bool {

    seen := map[int]bool{}
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] + l[j] == 0 {
				if _, ok := seen[l[i]]; !ok {
					seen[l[i]] = true
					return true
				}
				if _, ok := seen[l[j]]; !ok {
					seen[l[j]] = true
					return true
				}
			}
		}
	}
	return false
}



   
import (
    "strconv"
)

// Change numerical base of input number x to base.
// return string representation after the conversion.
// base numbers are less than 10.
// >>> ChangeBase(8, 3)
// '22'
// >>> ChangeBase(8, 2)
// '1000'
// >>> ChangeBase(7, 2)
// '111'


func ChangeBase(x int, base int) string {

    if x >= base {
        return ChangeBase(x/base, base) + ChangeBase(x%base, base)
    }
    return strconv.Itoa(x)
}



   

// Given length of a side and high return area for a triangle.
// >>> TriangleArea(5, 3)
// 7.5


func TriangleArea(a float64, h float64) float64 {

    return a * h / 2
}



   

// The Fib4 number sequence is a sequence similar to the Fibbonacci sequnece that's defined as follows:
// Fib4(0) -> 0
// Fib4(1) -> 0
// Fib4(2) -> 2
// Fib4(3) -> 0
// Fib4(n) -> Fib4(n-1) + Fib4(n-2) + Fib4(n-3) + Fib4(n-4).
// Please write a function to efficiently compute the n-th element of the Fib4 number sequence.  Do not use recursion.
// >>> Fib4(5)
// 4
// >>> Fib4(6)
// 8
// >>> Fib4(7)
// 14


func Fib4(n int) int {

    switch n {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 2
	case 3:
		return 0
	default:
		return Fib4(n-1) + Fib4(n-2) + Fib4(n-3) + Fib4(n-4)
	}
}



   
import (
	"sort"
)
// Return Median of elements in the list l.
// >>> Median([3, 1, 2, 4, 5])
// 3
// >>> Median([-10, 4, 6, 1000, 10, 20])
// 15.0


func Median(l []int) float64 {

    sort.Ints(l)
	if len(l)%2==1{
		return float64(l[len(l)/2])
	}else{
		return float64(l[len(l)/2-1]+l[len(l)/2])/2.0
	}
}



   

// Checks if given string is a palindrome
// >>> IsPalindrome('')
// true
// >>> IsPalindrome('aba')
// true
// >>> IsPalindrome('aaaaa')
// true
// >>> IsPalindrome('zbcd')
// false


func IsPalindrome(text string) bool {

    runes := []rune(text)
    result := make([]rune, 0)
    for i := len(runes) - 1; i >= 0; i-- {
        result = append(result, runes[i])
    }
    return text == string(result)
}



   

// Return 2^n modulo p (be aware of numerics).
// >>> Modp(3, 5)
// 3
// >>> Modp(1101, 101)
// 2
// >>> Modp(0, 101)
// 1
// >>> Modp(3, 11)
// 8
// >>> Modp(100, 101)
// 1


func Modp(n int,p int) int {

    ret := 1
    for i:= 0; i < n; i++ {
		ret = (2 * ret) % p
	}
    return ret
}



   

// returns encoded string by shifting every character by 5 in the alphabet.
// takes as input string encoded with encode_shift function. Returns decoded string.


func DecodeShift(s string) string {

    runes := []rune(s)
    newRunes := make([]rune, 0)
    for _, ch := range runes {
        newRunes = append(newRunes, (ch-5-'a')%26+'a')
    }
    return string(runes)
}



   
import (
    "regexp"
)
// RemoveVowels is a function that takes string and returns string without vowels.
// >>> RemoveVowels('')
// ''
// >>> RemoveVowels("abcdef\nghijklm")
// 'bcdf\nghjklm'
// >>> RemoveVowels('abcdef')
// 'bcdf'
// >>> RemoveVowels('aaaaa')
// ''
// >>> RemoveVowels('aaBAA')
// 'B'
// >>> RemoveVowels('zbcd')
// 'zbcd'


func RemoveVowels(text string) string {
    
    var re = regexp.MustCompile("[aeiouAEIOU]")
	text = re.ReplaceAllString(text, "")
	return text
}



   

// Return true if all numbers in the list l are below threshold t.
// >>> BelowThreshold([1, 2, 4, 10], 100)
// true
// >>> BelowThreshold([1, 20, 4, 10], 5)
// false


func BelowThreshold(l []int,t int) bool {

    for _, n := range l {
		if n >= t {
			return false
		}
	}
	return true
}



   

// Add two numbers x and y
// >>> Add(2, 3)
// 5
// >>> Add(5, 7)
// 12


func Add(x int, y int) int {

    return x + y
}



   

// Check if two words have the same characters.
// >>> SameChars('eabcdzzzz', 'dddzzzzzzzddeddabc')
// true
// >>> SameChars('abcd', 'dddddddabc')
// true
// >>> SameChars('dddddddabc', 'abcd')
// true
// >>> SameChars('eabcd', 'dddddddabc')
// false
// >>> SameChars('abcd', 'dddddddabce')
// false
// >>> SameChars('eabcdzzzz', 'dddzzzzzzzddddabc')
// false


func SameChars(s0 string, s1 string) bool {

    set0 := make(map[int32]interface{})
	set1 := make(map[int32]interface{})
	for _, i := range s0 {
		set0[i] = nil
	}
	for _, i := range s1 {
		set1[i] = nil
	}
	for i, _ := range set0 {
		if _,ok:=set1[i];!ok{
			return false
		}
	}
	for i, _ := range set1 {
		if _,ok:=set0[i];!ok{
			return false
		}
	}
	return true
}



   

// Return n-th Fibonacci number.
// >>> Fib(10)
// 55
// >>> Fib(1)
// 1
// >>> Fib(8)
// 21


func Fib(n int) int {

    if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}



   

// brackets is a string of "<" and ">".
// return true if every opening bracket has a corresponding closing bracket.
// 
// >>> CorrectBracketing("<")
// false
// >>> CorrectBracketing("<>")
// true
// >>> CorrectBracketing("<<><>>")
// true
// >>> CorrectBracketing("><<>")
// false


func CorrectBracketing(brackets string) bool {

    l := len(brackets)
	count := 0
	for index := 0; index < l; index++ {
		if brackets[index] == '<' {
			count++
		} else if brackets[index] == '>' {
			count--
		}
		if count < 0 {
			return false
		}
	}
    if count == 0 {
        return true
    } else {
        return false
    }
}



   

// Return true is list elements are Monotonically increasing or decreasing.
// >>> Monotonic([1, 2, 4, 20])
// true
// >>> Monotonic([1, 20, 4, 10])
// false
// >>> Monotonic([4, 1, 0, -10])
// true


func Monotonic(l []int) bool {

    flag := true
	if len(l) > 1 {
		for i := 0; i < len(l)-1; i++ {
			if l[i] != l[i+1] {
				flag = l[i] > l[i+1]
				break
			}
		}
	} else {
		return false
	}
	for i := 0; i < len(l)-1; i++ {
		if flag != (l[i] >= l[i+1]) {
			return false
		}
	}
	return true
}



   
import (
    "sort"
)
// Return sorted unique Common elements for two lists.
// >>> Common([1, 4, 3, 34, 653, 2, 5], [5, 7, 1, 5, 9, 653, 121])
// [1, 5, 653]
// >>> Common([5, 3, 2, 8], [3, 2])
// [2, 3]


func Common(l1 []int,l2 []int) []int {

    m := make(map[int]bool)
	for _, e1 := range l1 {
		if m[e1] {
			continue
		}
		for _, e2 := range l2 {
			if e1 == e2 {
				m[e1] = true
				break
			}
		}
	}
	res := make([]int, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}



   

// Return the largest prime factor of n. Assume n > 1 and is not a prime.
// >>> LargestPrimeFactor(13195)
// 29
// >>> LargestPrimeFactor(2048)
// 2


func LargestPrimeFactor(n int) int {

    isPrime := func(n int) bool {
        for i := 2; i < int(math.Pow(float64(n), 0.5)+1); i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }

    largest := 1
    for j := 2; j < n + 1; j++ {
		if n % j == 0 && isPrime(j) {
			if j > largest {
				largest = j
			}
		}
	}
    return largest
}



   

// SumToN is a function that sums numbers from 1 to n.
// >>> SumToN(30)
// 465
// >>> SumToN(100)
// 5050
// >>> SumToN(5)
// 15
// >>> SumToN(10)
// 55
// >>> SumToN(1)
// 1


func SumToN(n int) int {

    if n <= 0 {
		return 0
	} else {
		return n + SumToN(n - 1)
	}
}



   
import (
    "strings"
)
// brackets is a string of "(" and ")".
// return true if every opening bracket has a corresponding closing bracket.
// 
// >>> CorrectBracketing("(")
// false
// >>> CorrectBracketing("()")
// true
// >>> CorrectBracketing("(()())")
// true
// >>> CorrectBracketing(")(()")
// false


func CorrectBracketing(brackets string) bool {

    brackets = strings.Replace(brackets, "(", " ( ", -1)
	brackets = strings.Replace(brackets, ")", ") ", -1)
	open := 0
	for _, b := range brackets {
		if b == '(' {
			open++
		} else if b == ')' {
			open--
		}
		if open < 0 {
			return false
		}
	}
	return open == 0
}



   

// xs represent coefficients of a polynomial.
// xs[0] + xs[1] * x + xs[2] * x^2 + ....
// Return Derivative of this polynomial in the same form.
// >>> Derivative([3, 1, 2, 4, 5])
// [1, 4, 12, 20]
// >>> Derivative([1, 2, 3])
// [2, 6]


func Derivative(xs []int) []int {

    l := len(xs)
	y := make([]int, l - 1)
	for i := 0; i < l - 1; i++ {
		y[i] = xs[i + 1] * (i + 1)
	}
	return y
}



   

// The FibFib number sequence is a sequence similar to the Fibbonacci sequnece that's defined as follows:
// Fibfib(0) == 0
// Fibfib(1) == 0
// Fibfib(2) == 1
// Fibfib(n) == Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3).
// Please write a function to efficiently compute the n-th element of the Fibfib number sequence.
// >>> Fibfib(1)
// 0
// >>> Fibfib(5)
// 4
// >>> Fibfib(8)
// 24


func Fibfib(n int) int {

    if n <= 0 {
		return 0
	}
    switch n {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	default:
		return Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3)
	}
}



   
import (
    "strings"
)
// Write a function VowelsCount which takes a string representing
// a word as input and returns the number of vowels in the string.
// Vowels in this case are 'a', 'e', 'i', 'o', 'u'. Here, 'y' is also a
// vowel, but only when it is at the end of the given word.
// 
// Example:
// >>> VowelsCount("abcde")
// 2
// >>> VowelsCount("ACEDY")
// 3


func VowelsCount(s string) int {

    s = strings.ToLower(s)
	vowels := map[int32]interface{}{'a': nil, 'e': nil, 'i': nil, 'o': nil, 'u': nil}
	count := 0
	for _, i := range s {
		if _, ok := vowels[i]; ok {
			count++
		}
	}
	if (s[len(s)-1]) == 'y' {
		count++
	}
	return count
}



   
import (
    "strconv"
)
// Circular shift the digits of the integer x, shift the digits right by shift
// and return the result as a string.
// If shift > number of digits, return digits reversed.
// >>> CircularShift(12, 1)
// "21"
// >>> CircularShift(12, 2)
// "12"


func CircularShift(x int, shift int) string {

    s := strconv.Itoa(x)
	if shift > len(s) {
		runes := make([]rune, 0)
		for i := len(s)-1; i >= 0; i-- {
			runes = append(runes, rune(s[i]))
		}
		return string(runes)
	}else{
		return s[len(s)-shift:]+s[:len(s)-shift]
	}
}



   

// Task
// Write a function that takes a string as input and returns the sum of the upper characters only'
// ASCII codes.
// 
// Examples:
// Digitsum("") => 0
// Digitsum("abAB") => 131
// Digitsum("abcCd") => 67
// Digitsum("helloE") => 69
// Digitsum("woArBld") => 131
// Digitsum("aAaaaXa") => 153


func Digitsum(x string) int {

    if len(x) == 0 {
		return 0
	}
	result := 0
	for _, i := range x {
		if 'A' <= i && i <= 'Z' {
			result += int(i)
		}
	}
	return result
}



   
import (
	"strconv"
	"strings"
)
// In this task, you will be given a string that represents a number of apples and oranges
// that are distributed in a basket of fruit this basket contains
// apples, oranges, and mango fruits. Given the string that represents the total number of
// the oranges and apples and an integer that represent the total number of the fruits
// in the basket return the number of the mango fruits in the basket.
// for examble:
// FruitDistribution("5 apples and 6 oranges", 19) ->19 - 5 - 6 = 8
// FruitDistribution("0 apples and 1 oranges",3) -> 3 - 0 - 1 = 2
// FruitDistribution("2 apples and 3 oranges", 100) -> 100 - 2 - 3 = 95
// FruitDistribution("100 apples and 1 oranges",120) -> 120 - 100 - 1 = 19


func FruitDistribution(s string,n int) int {

    split := strings.Split(s, " ")
	for _, i := range split {
		atoi, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		n = n - atoi
	}
	return n
}



   
import (
    "math"
)
// Given an array representing a branch of a tree that has non-negative integer nodes
// your task is to Pluck one of the nodes and return it.
// The Plucked node should be the node with the smallest even value.
// If multiple nodes with the same smallest even value are found return the node that has smallest index.
// 
// The Plucked node should be returned in a list, [ smalest_value, its index ],
// If there are no even values or the given array is empty, return [].
// 
// Example 1:
// Input: [4,2,3]
// Output: [2, 1]
// Explanation: 2 has the smallest even value, and 2 has the smallest index.
// 
// Example 2:
// Input: [1,2,3]
// Output: [2, 1]
// Explanation: 2 has the smallest even value, and 2 has the smallest index.
// 
// Example 3:
// Input: []
// Output: []
// 
// Example 4:
// Input: [5, 0, 3, 0, 4, 2]
// Output: [0, 1]
// Explanation: 0 is the smallest value, but  there are two zeros,
// so we will choose the first zero, which has the smallest index.
// 
// Constraints:
// * 1 <= nodes.length <= 10000
// * 0 <= node.value


func Pluck(arr []int) []int {

    result := make([]int, 0)
	if len(arr) == 0 {
		return result
	}
	evens := make([]int, 0)
	min := math.MaxInt64
	minIndex := 0
	for i, x := range arr {
		if x%2 == 0 {
			evens = append(evens, x)
			if x < min {
				min = x
				minIndex = i
			}
		}
	}
	if len(evens) == 0 {
		return result
	}
	result = []int{min, minIndex}
	return result
}



   

// You are given a non-empty list of positive integers. Return the greatest integer that is greater than
// zero, and has a frequency greater than or equal to the value of the integer itself.
// The frequency of an integer is the number of times it appears in the list.
// If no such a value exist, return -1.
// Examples:
// Search([4, 1, 2, 2, 3, 1]) == 2
// Search([1, 2, 2, 3, 3, 3, 4, 4, 4]) == 3
// Search([5, 5, 4, 4, 4]) == -1


func Search(lst []int) int {

    countMap := make(map[int]int)
	for _, i := range lst {
		if count, ok := countMap[i]; ok {
			countMap[i] = count + 1
		} else {
			countMap[i] = 1
		}
	}
	max := -1
	for i, count := range countMap {
		if count >= i && count > max {
			max = i
		}
	}
	return max
}



   
import (
    "sort"
)
// Given list of integers, return list in strange order.
// Strange sorting, is when you start with the minimum value,
// then maximum of the remaining integers, then minimum and so on.
// 
// Examples:
// StrangeSortList([1, 2, 3, 4]) == [1, 4, 2, 3]
// StrangeSortList([5, 5, 5, 5]) == [5, 5, 5, 5]
// StrangeSortList([]) == []


func StrangeSortList(lst []int) []int {

    sort.Ints(lst)
	result := make([]int, 0)
	for i := 0; i < len(lst)/2; i++ {
		result = append(result, lst[i])
		result = append(result, lst[len(lst)-i-1])
	}
	if len(lst)%2 != 0 {
		result = append(result, lst[len(lst)/2])
	}
	return result
}



   
import (
    "math"
)
// Given the lengths of the three sides of a triangle. Return the area of
// the triangle rounded to 2 decimal points if the three sides form a valid triangle.
// Otherwise return -1
// Three sides make a valid triangle when the sum of any two sides is greater
// than the third side.
// Example:
// TriangleArea(3, 4, 5) == 6.00
// TriangleArea(1, 2, 10) == -1


func TriangleArea(a float64, b float64, c float64) interface{} {

    if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}
	s := (a + b + c) / 2
	area := math.Pow(s * (s - a) * (s - b) * (s - c), 0.5)
	area = math.Round(area*100)/100
	return area
}



   

// Write a function that returns true if the object q will fly, and false otherwise.
// The object q will fly if it's balanced (it is a palindromic list) and the sum of its elements is less than or equal the maximum possible weight w.
// 
// Example:
// WillItFly([1, 2], 5) ➞ false
// 1+2 is less than the maximum possible weight, but it's unbalanced.
// 
// WillItFly([3, 2, 3], 1) ➞ false
// it's balanced, but 3+2+3 is more than the maximum possible weight.
// 
// WillItFly([3, 2, 3], 9) ➞ true
// 3+2+3 is less than the maximum possible weight, and it's balanced.
// 
// WillItFly([3], 5) ➞ true
// 3 is less than the maximum possible weight, and it's balanced.


func WillItFly(q []int,w int) bool {

    sum := 0
	for i := 0; i < len(q); i++ {
		sum += q[i]
	}
	if sum <= w && isPalindrome(q) {
		return true
	}
	return false
}

func isPalindrome(arr []int) bool {
	for i := 0; i < (len(arr) / 2); i++ {
		if arr[i] != arr[len(arr) - i - 1] {
			return false
		}
	}
	return true
}



   

// Given an array arr of integers, find the minimum number of elements that
// need to be changed to make the array palindromic. A palindromic array is an array that
// is read the same backwards and forwards. In one change, you can change one element to any other element.
// 
// For example:
// SmallestChange([1,2,3,5,4,7,9,6]) == 4
// SmallestChange([1, 2, 3, 4, 3, 2, 2]) == 1
// SmallestChange([1, 2, 3, 2, 1]) == 0


func SmallestChange(arr []int) int {

    count := 0
	for i := 0; i < len(arr) - 1; i++ {
        a := arr[len(arr) - i - 1]
		if arr[i] != a {
			arr[i] = a
            count++
		}
	}
	return count
}



   

// Write a function that accepts two lists of strings and returns the list that has
// total number of chars in the all strings of the list less than the other list.
// 
// if the two lists have the same number of chars, return the first list.
// 
// Examples
// TotalMatch([], []) ➞ []
// TotalMatch(['hi', 'admin'], ['hI', 'Hi']) ➞ ['hI', 'Hi']
// TotalMatch(['hi', 'admin'], ['hi', 'hi', 'admin', 'project']) ➞ ['hi', 'admin']
// TotalMatch(['hi', 'admin'], ['hI', 'hi', 'hi']) ➞ ['hI', 'hi', 'hi']
// TotalMatch(['4'], ['1', '2', '3', '4', '5']) ➞ ['4']


func TotalMatch(lst1 []string,lst2 []string) []string {

    var numchar1 = 0
	var numchar2 = 0
	for _, item := range lst1 {
		numchar1 += len(item)
	}
	for _, item := range lst2 {
		numchar2 += len(item)
	}
	if numchar1 <= numchar2 {
		return lst1
	} else {
		return lst2
	}
}



   

// Write a function that returns true if the given number is the multiplication of 3 prime numbers
// and false otherwise.
// Knowing that (a) is less then 100.
// Example:
// IsMultiplyPrime(30) == true
// 30 = 2 * 3 * 5


func IsMultiplyPrime(a int) bool {

    isPrime := func(n int) bool {
        for i := 2; i < int(math.Pow(float64(n), 0.5)+1); i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }
    for i := 2; i < 101; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j < 101; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k < 101; k++ {
				if !isPrime(k) {
					continue
				}
				if i*j*k == a {
					return true
				}
			}
		}
	}
	return false
}



   

// Your task is to write a function that returns true if a number x is a simple
// power of n and false in other cases.
// x is a simple power of n if n**int=x
// For example:
// IsSimplePower(1, 4) => true
// IsSimplePower(2, 2) => true
// IsSimplePower(8, 2) => true
// IsSimplePower(3, 2) => false
// IsSimplePower(3, 1) => false
// IsSimplePower(5, 3) => false


func IsSimplePower(x int,n int) bool {

    if x == 1 {
		return true
	}
	if n==1 {
		return false
	}
	if x % n != 0 {
		return false
	}
	return IsSimplePower(x / n, n)
}



   
import (
    "math"
)
// Write a function that takes an integer a and returns true
// if this ingeger is a cube of some integer number.
// Note: you may assume the input is always valid.
// Examples:
// Iscube(1) ==> true
// Iscube(2) ==> false
// Iscube(-1) ==> true
// Iscube(64) ==> true
// Iscube(0) ==> true
// Iscube(180) ==> false


func Iscube(a int) bool {

    abs := math.Abs(float64(a))
	return int(math.Pow(math.Round(math.Pow(abs, 1.0/3.0)), 3.0)) == int(abs)
}



   

// You have been tasked to write a function that receives
// a hexadecimal number as a string and counts the number of hexadecimal
// digits that are primes (prime number, or a prime, is a natural number
// greater than 1 that is not a product of two smaller natural numbers).
// Hexadecimal digits are 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F.
// Prime numbers are 2, 3, 5, 7, 11, 13, 17,...
// So you have to determine a number of the following digits: 2, 3, 5, 7,
// B (=decimal 11), D (=decimal 13).
// Note: you may assume the input is always correct or empty string,
// and symbols A,B,C,D,E,F are always uppercase.
// Examples:
// For num = "AB" the output should be 1.
// For num = "1077E" the output should be 2.
// For num = "ABED1A33" the output should be 4.
// For num = "123456789ABCDEF0" the output should be 6.
// For num = "2020" the output should be 2.


func HexKey(num string) int {

    primes := map[int32]interface{}{'2': nil, '3': nil, '5': nil, '7': nil, 'B': nil, 'D': nil}
	total := 0
	for _, c := range num {
		if _, ok := primes[c]; ok {
			total++
		}
	}
	return total
}



   
import (
	"fmt"
)
// You will be given a number in decimal form and your task is to convert it to
// binary format. The function should return a string, with each character representing a binary
// number. Each character in the string will be '0' or '1'.
// 
// There will be an extra couple of characters 'db' at the beginning and at the end of the string.
// The extra characters are there to help with the format.
// 
// Examples:
// DecimalToBinary(15)   # returns "db1111db"
// DecimalToBinary(32)   # returns "db100000db"


func DecimalToBinary(decimal int) string {

    return fmt.Sprintf("db%bdb", decimal)
}



   

// You are given a string s.
// Your task is to check if the string is happy or not.
// A string is happy if its length is at least 3 and every 3 consecutive letters are distinct
// For example:
// IsHappy(a) => false
// IsHappy(aa) => false
// IsHappy(abcd) => true
// IsHappy(aabb) => false
// IsHappy(adb) => true
// IsHappy(xyy) => false


func IsHappy(s string) bool {

    if len(s) < 3 {
        return false
    }
    for i := 0; i < len(s)-2; i++ {
        if s[i] == s[i+1] || s[i+1] == s[i+2] || s[i] == s[i+2] {
            return false
        }
    }
    return true
}



   

// It is the last week of the semester and the teacher has to give the grades
// to students. The teacher has been making her own algorithm for grading.
// The only problem is, she has lost the code she used for grading.
// She has given you a list of GPAs for some students and you have to write
// a function that can output a list of letter grades using the following table:
// GPA       |    Letter grade
// 4.0                A+
// > 3.7                A
// > 3.3                A-
// > 3.0                B+
// > 2.7                B
// > 2.3                B-
// > 2.0                C+
// > 1.7                C
// > 1.3                C-
// > 1.0                D+
// > 0.7                D
// > 0.0                D-
// 0.0                E
// 
// 
// Example:
// grade_equation([4.0, 3, 1.7, 2, 3.5]) ==> ["A+", "B", "C-", "C", "A-"]


func NumericalLetterGrade(grades []float64) []string {

letter_grade := make([]string, 0, len(grades))
    for _, gpa := range grades {
        switch {
        case gpa == 4.0:
            letter_grade = append(letter_grade, "A+")
        case gpa > 3.7:
            letter_grade = append(letter_grade, "A")
        case gpa > 3.3:
            letter_grade = append(letter_grade, "A-")
        case gpa > 3.0:
            letter_grade = append(letter_grade, "B+")
        case gpa > 2.7:
            letter_grade = append(letter_grade, "B")
        case gpa > 2.3:
            letter_grade = append(letter_grade, "B-")
        case gpa > 2.0:
            letter_grade = append(letter_grade, "C+")
        case gpa > 1.7:
            letter_grade = append(letter_grade, "C")
        case gpa > 1.3:
            letter_grade = append(letter_grade, "C-")
        case gpa > 1.0:
            letter_grade = append(letter_grade, "D+")
        case gpa > 0.7:
            letter_grade = append(letter_grade, "D")
        case gpa > 0.0:
            letter_grade = append(letter_grade, "D-")
        default:
            letter_grade = append(letter_grade, "E")
        }

    }
    return letter_grade
}



   

// Write a function that takes a string and returns true if the string
// length is a prime number or false otherwise
// Examples
// PrimeLength('Hello') == true
// PrimeLength('abcdcba') == true
// PrimeLength('kittens') == true
// PrimeLength('orange') == false


func PrimeLength(s string) bool {

    l := len(s)
    if l == 0 || l == 1 {
        return false
    }
    for i := 2; i < l; i++ {
        if l%i == 0 {
            return false
        }
    }
    return true
}



   
import (
    "math"
)

// Given a positive integer n, return the count of the numbers of n-digit
// positive integers that start or end with 1.


func StartsOneEnds(n int) int {

    if n == 1 {
        return 1
    }
    return 18 * int(math.Pow(10, float64(n-2)))
}



   
import (
    "fmt"
    "strconv"
)

// Given a positive integer N, return the total sum of its digits in binary.
// 
// Example
// For N = 1000, the sum of digits will be 1 the output should be "1".
// For N = 150, the sum of digits will be 6 the output should be "110".
// For N = 147, the sum of digits will be 12 the output should be "1100".
// 
// Variables:
// @N integer
// Constraints: 0 ≤ N ≤ 10000.
// Output:
// a string of binary number


func Solve(N int) string {

    sum := 0
    for _, c := range strconv.Itoa(N) {
        sum += int(c - '0')
    }
    return fmt.Sprintf("%b", sum)
}



   

// Given a non-empty list of integers lst. Add the even elements that are at odd indices..
// 
// Examples:
// Add([4, 2, 6, 7]) ==> 2


func Add(lst []int) int {

    sum := 0
    for i := 1; i < len(lst); i += 2 {
        if lst[i]%2 == 0 {
            sum += lst[i]
        }
    }
    return sum
}



   
import (
    "sort"
    "strings"
)

// Write a function that takes a string and returns an ordered version of it.
// Ordered version of string, is a string where all words (separated by space)
// are replaced by a new word where all the characters arranged in
// ascending order based on ascii value.
// Note: You should keep the order of words and blank spaces in the sentence.
// 
// For example:
// AntiShuffle('Hi') returns 'Hi'
// AntiShuffle('hello') returns 'ehllo'
// AntiShuffle('Hello World!!!') returns 'Hello !!!Wdlor'


func AntiShuffle(s string) string {

    strs := make([]string, 0)
    for _, i := range strings.Fields(s) {
        word := []rune(i)
        sort.Slice(word, func(i, j int) bool {
            return word[i] < word[j]
        })
        strs = append(strs, string(word))
    }
    return strings.Join(strs, " ")
}



   
import (
    "sort"
)

// You are given a 2 dimensional data, as a nested lists,
// which is similar to matrix, however, unlike matrices,
// each row may contain a different number of columns.
// Given lst, and integer x, find integers x in the list,
// and return list of tuples, [(x1, y1), (x2, y2) ...] such that
// each tuple is a coordinate - (row, columns), starting with 0.
// Sort coordinates initially by rows in ascending order.
// Also, sort coordinates of the row by columns in descending order.
// 
// Examples:
// GetRow([
// [1,2,3,4,5,6],
// [1,2,3,4,1,6],
// [1,2,3,4,5,1]
// ], 1) == [(0, 0), (1, 4), (1, 0), (2, 5), (2, 0)]
// GetRow([], 1) == []
// GetRow([[], [1], [1, 2, 3]], 3) == [(2, 2)]


func GetRow(lst [][]int, x int) [][2]int {

    coords := make([][2]int, 0)
    for i, row := range lst {
        for j, item := range row {
            if item == x {
                coords = append(coords, [2]int{i, j})
            }
        }
    }
    sort.Slice(coords, func(i, j int) bool {
        if coords[i][0] == coords[j][0] {
            return coords[i][1] > coords[j][1]
        }
        return coords[i][0] < coords[j][0]
    })

    return coords
}



   
import (
    "sort"
)

// Given an array of non-negative integers, return a copy of the given array after sorting,
// you will sort the given array in ascending order if the sum( first index value, last index value) is odd,
// or sort it in descending order if the sum( first index value, last index value) is even.
// 
// Note:
// * don't change the given array.
// 
// Examples:
// * SortArray([]) => []
// * SortArray([5]) => [5]
// * SortArray([2, 4, 3, 0, 1, 5]) => [0, 1, 2, 3, 4, 5]
// * SortArray([2, 4, 3, 0, 1, 5, 6]) => [6, 5, 4, 3, 2, 1, 0]


func SortArray(array []int) []int {

    arr := make([]int, len(array))
    copy(arr, array)
    if len(arr) == 0 {
        return arr
    }
    if (arr[0]+arr[len(arr)-1])%2 == 0 {
        sort.Slice(arr, func(i, j int) bool {
            return arr[i] > arr[j]
        })
    } else {
        sort.Slice(arr, func(i, j int) bool {
            return arr[i] < arr[j]
        })
    }
    return arr
}



   
import (
    "strings"
)

// Create a function Encrypt that takes a string as an argument and
// returns a string Encrypted with the alphabet being rotated.
// The alphabet should be rotated in a manner such that the letters
// shift down by two multiplied to two places.
// For example:
// Encrypt('hi') returns 'lm'
// Encrypt('asdfghjkl') returns 'ewhjklnop'
// Encrypt('gf') returns 'kj'
// Encrypt('et') returns 'ix'


func Encrypt(s string) string {

    d := "abcdefghijklmnopqrstuvwxyz"
    out := make([]rune, 0, len(s))
    for _, c := range s {
        pos := strings.IndexRune(d, c)
        if pos != -1 {
            out = append(out, []rune(d)[(pos+2*2)%26])
        } else {
            out = append(out, c)
        }
    }
    return string(out)
}



   
import (
    "math"
    "sort"
)

// You are given a list of integers.
// Write a function NextSmallest() that returns the 2nd smallest element of the list.
// Return nil if there is no such element.
// 
// NextSmallest([1, 2, 3, 4, 5]) == 2
// NextSmallest([5, 1, 4, 3, 2]) == 2
// NextSmallest([]) == nil
// NextSmallest([1, 1]) == nil


func NextSmallest(lst []int) interface{} {

    set := make(map[int]struct{})
    for _, i := range lst {
        set[i] = struct{}{}
    }
    vals := make([]int, 0, len(set))
    for k := range set {
        vals = append(vals, k)
    }
    sort.Slice(vals, func(i, j int) bool {
        return vals[i] < vals[j]
    })
    if len(vals) < 2 {
        return nil
    }
    return vals[1]
}



   
import (
    "regexp"
)

// You'll be given a string of words, and your task is to count the number
// of boredoms. A boredom is a sentence that starts with the word "I".
// Sentences are delimited by '.', '?' or '!'.
// 
// For example:
// >>> IsBored("Hello world")
// 0
// >>> IsBored("The sky is blue. The sun is shining. I love this weather")
// 1


func IsBored(S string) int {

    r, _ := regexp.Compile(`[.?!]\s*`)
    sentences := r.Split(S, -1)
    sum := 0
    for _, s := range sentences {
        if len(s) >= 2 && s[:2] == "I " {
            sum++
        }
    }
    return sum
}



   

// Create a function that takes 3 numbers.
// Returns true if one of the numbers is equal to the sum of the other two, and all numbers are integers.
// Returns false in any other cases.
// 
// Examples
// AnyInt(5, 2, 7) ➞ true
// 
// AnyInt(3, 2, 2) ➞ false
// 
// AnyInt(3, -2, 1) ➞ true
// 
// AnyInt(3.6, -2.2, 2) ➞ false


func AnyInt(x, y, z interface{}) bool {

    xx, ok := x.(int)
    if !ok {
        return false
    }
    yy, ok := y.(int)
    if !ok {
        return false
    }
    zz, ok := z.(int)
    if !ok {
        return false
    }

    if (xx+yy == zz) || (xx+zz == yy) || (yy+zz == xx) {
        return true
    }
    return false
}



   
import (
    "strings"
)

// Write a function that takes a message, and Encodes in such a
// way that it swaps case of all letters, replaces all vowels in
// the message with the letter that appears 2 places ahead of that
// vowel in the english alphabet.
// Assume only letters.
// 
// Examples:
// >>> Encode('test')
// 'TGST'
// >>> Encode('This is a message')
// 'tHKS KS C MGSSCGG'


func Encode(message string) string {

    vowels := "aeiouAEIOU"
    vowels_replace := make(map[rune]rune)
    for _, c := range vowels {
        vowels_replace[c] = c + 2
    }
    result := make([]rune, 0, len(message))
    for _, c := range message {
        if 'a' <= c && c <= 'z' {
            c += 'A' - 'a'
        } else if 'A' <= c && c <= 'Z' {
            c += 'a' - 'A'
        }
        if strings.ContainsRune(vowels, c) {
            result = append(result, vowels_replace[c])
        } else {
            result = append(result, c)
        }
    }
    return string(result)
}



   
import (
    "math"
    "strconv"
)

// You are given a list of integers.
// You need to find the largest prime value and return the sum of its digits.
// 
// Examples:
// For lst = [0,3,2,1,3,5,7,4,5,5,5,2,181,32,4,32,3,2,32,324,4,3] the output should be 10
// For lst = [1,0,1,8,2,4597,2,1,3,40,1,2,1,2,4,2,5,1] the output should be 25
// For lst = [1,3,1,32,5107,34,83278,109,163,23,2323,32,30,1,9,3] the output should be 13
// For lst = [0,724,32,71,99,32,6,0,5,91,83,0,5,6] the output should be 11
// For lst = [0,81,12,3,1,21] the output should be 3
// For lst = [0,8,1,2,1,7] the output should be 7


func Skjkasdkd(lst []int) int {

    isPrime := func(n int) bool {
        for i := 2; i < int(math.Pow(float64(n), 0.5)+1); i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }
    maxx := 0
    i := 0
    for i < len(lst) {
        if lst[i] > maxx && isPrime(lst[i]) {
            maxx = lst[i]
        }
        i++
    }
    sum := 0
    for _, d := range strconv.Itoa(maxx) {
        sum += int(d - '0')
    }
    return sum
}



   
import (
    "strings"
)

// Given a dictionary, return true if all keys are strings in lower
// case or all keys are strings in upper case, else return false.
// The function should return false is the given dictionary is empty.
// Examples:
// CheckDictCase({"a":"apple", "b":"banana"}) should return true.
// CheckDictCase({"a":"apple", "A":"banana", "B":"banana"}) should return false.
// CheckDictCase({"a":"apple", 8:"banana", "a":"apple"}) should return false.
// CheckDictCase({"Name":"John", "Age":"36", "City":"Houston"}) should return false.
// CheckDictCase({"STATE":"NC", "ZIP":"12345" }) should return true.


func CheckDictCase(dict map[interface{}]interface{}) bool {

    if len(dict) == 0 {
        return false
    }
    state := "start"
    key := ""
    ok := false
    for k := range dict {
        if key, ok = k.(string); !ok {
            state = "mixed"
            break
        }
        if state == "start" {
            if key == strings.ToUpper(key) {
                state = "upper"
            } else if key == strings.ToLower(key) {
                state = "lower"
            } else {
                break
            }
        } else if (state == "upper" && key != strings.ToUpper(key)) || (state == "lower" && key != strings.ToLower(key)) {
            state = "mixed"
            break
        } else {
            break
        }
    }
    return state == "upper" || state == "lower"
}



   

// Implement a function that takes an non-negative integer and returns an array of the first n
// integers that are prime numbers and less than n.
// for example:
// CountUpTo(5) => [2,3]
// CountUpTo(11) => [2,3,5,7]
// CountUpTo(0) => []
// CountUpTo(20) => [2,3,5,7,11,13,17,19]
// CountUpTo(1) => []
// CountUpTo(18) => [2,3,5,7,11,13,17]


func CountUpTo(n int) []int {

    primes := make([]int, 0)
    for i := 2; i < n; i++ {
        is_prime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                is_prime = false
                break
            }
        }
        if is_prime {
            primes = append(primes, i)
        }
    }
    return primes
}



   
import (
    "math"
)

// Complete the function that takes two integers and returns
// the product of their unit digits.
// Assume the input is always valid.
// Examples:
// Multiply(148, 412) should return 16.
// Multiply(19, 28) should return 72.
// Multiply(2020, 1851) should return 0.
// Multiply(14,-15) should return 20.


func Multiply(a, b int) int {

    return int(math.Abs(float64(a%10)) * math.Abs(float64(b%10)))
}



   
import (
    "strings"
)

// Given a string s, count the number of uppercase vowels in even indices.
// 
// For example:
// CountUpper('aBCdEf') returns 1
// CountUpper('abcdefg') returns 0
// CountUpper('dBBE') returns 0


func CountUpper(s string) int {

    count := 0
    runes := []rune(s)
    for i := 0; i < len(runes); i += 2 {
        if strings.ContainsRune("AEIOU", runes[i]) {
            count += 1
        }
    }
    return count
}



   
import (
    "math"
    "strconv"
    "strings"
)

// Create a function that takes a value (string) representing a number
// and returns the closest integer to it. If the number is equidistant
// from two integers, round it away from zero.
// 
// Examples
// >>> ClosestInteger("10")
// 10
// >>> ClosestInteger("15.3")
// 15
// 
// Note:
// Rounding away from zero means that if the given number is equidistant
// from two integers, the one you should return is the one that is the
// farthest from zero. For example ClosestInteger("14.5") should
// return 15 and ClosestInteger("-14.5") should return -15.


func ClosestInteger(value string) int {

    if strings.Count(value, ".") == 1 {
        // remove trailing zeros
        for value[len(value)-1] == '0' {
            value = value[:len(value)-1]
        }
    }
    var res float64
    num, _ := strconv.ParseFloat(value, 64)
    if len(value) >= 2 && value[len(value)-2:] == ".5" {
        if num > 0 {
            res = math.Ceil(num)
        } else {
            res = math.Floor(num)
        }
    } else if len(value) > 0 {
        res = math.Round(num)
    } else {
        res = 0
    }

    return int(res)
}



   

// Given a positive integer n, you have to make a pile of n levels of stones.
// The first level has n stones.
// The number of stones in the next level is:
// - the next odd number if n is odd.
// - the next even number if n is even.
// Return the number of stones in each level in a list, where element at index
// i represents the number of stones in the level (i+1).
// 
// Examples:
// >>> MakeAPile(3)
// [3, 5, 7]


func MakeAPile(n int) []int {

    result := make([]int, 0, n)
    for i := 0; i < n; i++ {
        result = append(result, n+2*i)
    }
    return result
}



   
import (
    "strings"
)

// You will be given a string of words separated by commas or spaces. Your task is
// to split the string into words and return an array of the words.
// 
// For example:
// WordsString("Hi, my name is John") == ["Hi", "my", "name", "is", "John"]
// WordsString("One, two, three, four, five, six") == ["One", "two", "three", "four", "five", "six"]


func WordsString(s string) []string {

    s_list := make([]rune, 0)

    for _, c := range s {
        if c == ',' {
            s_list = append(s_list, ' ')
        } else {
            s_list = append(s_list, c)
        }
    }
    return strings.Fields(string(s_list))
}



   

// This function takes two positive numbers x and y and returns the
// biggest even integer number that is in the range [x, y] inclusive. If
// there's no such number, then the function should return -1.
// 
// For example:
// ChooseNum(12, 15) = 14
// ChooseNum(13, 12) = -1


func ChooseNum(x, y int) int {

    if x > y {
        return -1
    }
    if y % 2 == 0 {
        return y
    }
    if x == y {
        return -1
    }
    return y - 1
}



   
import (
    "fmt"
    "math"
)

// You are given two positive integers n and m, and your task is to compute the
// average of the integers from n through m (including n and m).
// Round the answer to the nearest integer and convert that to binary.
// If n is greater than m, return -1.
// Example:
// RoundedAvg(1, 5) => "0b11"
// RoundedAvg(7, 5) => -1
// RoundedAvg(10, 20) => "0b1111"
// RoundedAvg(20, 33) => "0b11010"


func RoundedAvg(n, m int) interface{} {

    if m < n {
        return -1
    }
    summation := 0
    for i := n;i < m+1;i++{
        summation += i
    }
    return fmt.Sprintf("0b%b", int(math.Round(float64(summation)/float64(m - n + 1))))
}



   
import (
    "sort"
    "strconv"
)

// Given a list of positive integers x. return a sorted list of all
// elements that hasn't any even digit.
// 
// Note: Returned list should be sorted in increasing order.
// 
// For example:
// >>> UniqueDigits([15, 33, 1422, 1])
// [1, 15, 33]
// >>> UniqueDigits([152, 323, 1422, 10])
// []


func UniqueDigits(x []int) []int {

    odd_digit_elements := make([]int, 0)
    OUTER:
    for _, i := range x {
        for _, c := range strconv.Itoa(i) {
            if (c - '0') % 2 == 0 {
                continue OUTER
            }
        }
            odd_digit_elements = append(odd_digit_elements, i)
    }
    sort.Slice(odd_digit_elements, func(i, j int) bool {
        return odd_digit_elements[i] < odd_digit_elements[j]
    })
    return odd_digit_elements
}



   
import (
    "sort"
)

// Given an array of integers, sort the integers that are between 1 and 9 inclusive,
// reverse the resulting array, and then replace each digit by its corresponding name from
// "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine".
// 
// For example:
// arr = [2, 1, 1, 4, 5, 8, 2, 3]
// -> sort arr -> [1, 1, 2, 2, 3, 4, 5, 8]
// -> reverse arr -> [8, 5, 4, 3, 2, 2, 1, 1]
// return ["Eight", "Five", "Four", "Three", "Two", "Two", "One", "One"]
// 
// If the array is empty, return an empty array:
// arr = []
// return []
// 
// If the array has any strange number ignore it:
// arr = [1, -1 , 55]
// -> sort arr -> [-1, 1, 55]
// -> reverse arr -> [55, 1, -1]
// return = ['One']


func ByLength(arr []int)[]string {

    dic := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
    new_arr := make([]string, 0)
    for _, item := range arr {
        if v, ok := dic[item]; ok {
            new_arr = append(new_arr, v)
        }
    }
    return new_arr
}



   


// Implement the Function F that takes n as a parameter,
// and returns a list oF size n, such that the value oF the element at index i is the Factorial oF i iF i is even
// or the sum oF numbers From 1 to i otherwise.
// i starts From 1.
// the Factorial oF i is the multiplication oF the numbers From 1 to i (1 * 2 * ... * i).
// Example:
// F(5) == [1, 2, 6, 24, 15]


func F(n int) []int {

    ret := make([]int, 0, 5)
    for i:=1;i<n+1;i++{
        if i%2 == 0 {
            x := 1
            for j:=1;j<i+1;j++{
                x*=j
            }
            ret = append(ret, x)
        }else {
            x := 0
            for j:=1;j<i+1;j++{
                x+=j
            }
            ret = append(ret, x)
        }
    }
    return ret
}



   
import (
    "strconv"
)

// Given a positive integer n, return a tuple that has the number of even and odd
// integer palindromes that fall within the range(1, n), inclusive.
// 
// Example 1:
// 
// Input: 3
// Output: (1, 2)
// Explanation:
// Integer palindrome are 1, 2, 3. one of them is even, and two of them are odd.
// 
// Example 2:
// 
// Input: 12
// Output: (4, 6)
// Explanation:
// Integer palindrome are 1, 2, 3, 4, 5, 6, 7, 8, 9, 11. four of them are even, and 6 of them are odd.
// 
// Note:
// 1. 1 <= n <= 10^3
// 2. returned tuple has the number of even and odd integer palindromes respectively.


func EvenOddPalindrome(n int) [2]int {

    is_palindrome := func (n int) bool {
        s := strconv.Itoa(n)
        for i := 0;i < len(s)>>1;i++ {
            if s[i] != s[len(s)-i-1] {
                return false
            }
        }
        return true
    }

    even_palindrome_count := 0
    odd_palindrome_count := 0

    for i :=1;i<n+1;i++ {
        if i%2 == 1 && is_palindrome(i){
                odd_palindrome_count ++
        } else if i%2 == 0 && is_palindrome(i) {
            even_palindrome_count ++
        }
    }
    return [2]int{even_palindrome_count, odd_palindrome_count}
}



   
import (
    "math"
    "strconv"
)

// Write a function CountNums which takes an array of integers and returns
// the number of elements which has a sum of digits > 0.
// If a number is negative, then its first signed digit will be negative:
// e.g. -123 has signed digits -1, 2, and 3.
// >>> CountNums([]) == 0
// >>> CountNums([-1, 11, -11]) == 1
// >>> CountNums([1, 1, 2]) == 3


func CountNums(arr []int) int {

    digits_sum:= func (n int) int {
        neg := 1
        if n < 0 {
             n, neg = -1 * n, -1 
        }
        r := make([]int,0)
        for _, c := range strconv.Itoa(n) {
            r = append(r, int(c-'0'))
        }
        r[0] *= neg
        sum := 0
        for _, i := range r {
            sum += i
        }
        return sum
    }
    count := 0
    for _, i := range arr {
        x := digits_sum(i)
        if x > 0 {
            count++
        }
    }
    return count
}



   
import (
    "math"
    "sort"
)

// We have an array 'arr' of N integers arr[1], arr[2], ..., arr[N].The
// numbers in the array will be randomly ordered. Your task is to determine if
// it is possible to get an array sorted in non-decreasing order by performing
// the following operation on the given array:
// You are allowed to perform right shift operation any number of times.
// 
// One right shift operation means shifting all elements of the array by one
// position in the right direction. The last element of the array will be moved to
// the starting position in the array i.e. 0th index.
// 
// If it is possible to obtain the sorted array by performing the above operation
// then return true else return false.
// If the given array is empty then return true.
// 
// Note: The given list is guaranteed to have unique elements.
// 
// For Example:
// 
// MoveOneBall([3, 4, 5, 1, 2])==>true
// Explanation: By performin 2 right shift operations, non-decreasing order can
// be achieved for the given array.
// MoveOneBall([3, 5, 4, 1, 2])==>false
// Explanation:It is not possible to get non-decreasing order for the given
// array by performing any number of right shift operations.


func MoveOneBall(arr []int) bool {

    if len(arr)==0 {
      return true
    }
    sorted_array := make([]int, len(arr))
    copy(sorted_array, arr)
    sort.Slice(sorted_array, func(i, j int) bool {
        return sorted_array[i] < sorted_array[j]
    })    
    min_value := math.MaxInt
    min_index := -1
    for i, x := range arr {
        if i < min_value {
            min_index, min_value = i, x
        }
    }
    my_arr := make([]int, len(arr[min_index:]))
    copy(my_arr, arr[min_index:])
    my_arr = append(my_arr, arr[0:min_index]...)
    for i :=0;i<len(arr);i++ {
      if my_arr[i]!=sorted_array[i]{
        return false
      }
    }
    return true
}



   

// In this problem, you will implement a function that takes two lists of numbers,
// and determines whether it is possible to perform an Exchange of elements
// between them to make lst1 a list of only even numbers.
// There is no limit on the number of Exchanged elements between lst1 and lst2.
// If it is possible to Exchange elements between the lst1 and lst2 to make
// all the elements of lst1 to be even, return "YES".
// Otherwise, return "NO".
// For example:
// Exchange([1, 2, 3, 4], [1, 2, 3, 4]) => "YES"
// Exchange([1, 2, 3, 4], [1, 5, 3, 4]) => "NO"
// It is assumed that the input lists will be non-empty.


func Exchange(lst1, lst2 []int) string {

    odd := 0
    even := 0
    for _, i := range lst1 {
        if i%2 == 1 {
            odd++
        }
    }
    for _, i := range lst2 {
        if i%2 == 0 {
            even++
        }
    }
    if even >= odd {
        return "YES"
    }
    return "NO"
}
            



   
import (
    "strings"
)

// Given a string representing a space separated lowercase letters, return a dictionary
// of the letter with the most repetition and containing the corresponding count.
// If several letters have the same occurrence, return all of them.
// 
// Example:
// Histogram('a b c') == {'a': 1, 'b': 1, 'c': 1}
// Histogram('a b b a') == {'a': 2, 'b': 2}
// Histogram('a b c a b') == {'a': 2, 'b': 2}
// Histogram('b b b b a') == {'b': 4}
// Histogram('') == {}


func Histogram(test string) map[rune]int {

    dict1 := make(map[rune]int)
    list1 := strings.Fields(test)
    t := 0
    count := func(lst []string, v string) int {
        cnt := 0
        for _, i := range lst {
            if i == v {
                cnt++
            }
        }
        return cnt
    }
    for _, i := range list1 {
        if c := count(list1, i); c>t && i!="" {
            t=c
        }
    }
    if t>0 {
        for _, i := range list1 {
            if count(list1, i)==t {
                dict1[[]rune(i)[0]]=t
            }
        }
    }
    return dict1
}



   
import (
    "strings"
)

// Task
// We are given two strings s and c, you have to deleted all the characters in s that are equal to any character in c
// then check if the result string is palindrome.
// A string is called palindrome if it reads the same backward as forward.
// You should return a tuple containing the result string and true/false for the check.
// Example
// For s = "abcde", c = "ae", the result should be ('bcd',false)
// For s = "abcdef", c = "b"  the result should be ('acdef',false)
// For s = "abcdedcba", c = "ab", the result should be ('cdedc',true)


func ReverseDelete(s,c string) [2]interface{} {

    rs := make([]rune, 0)
    for _, r := range s {
        if !strings.ContainsRune(c, r) {
            rs = append(rs, r)
        }
    }
    t := true
    for i := 0;i < len(rs)>>1;i++ {
        if rs[i] != rs[len(rs)-i-1] {
            t=false
            break
        }
    }
    return [2]interface{}{string(rs), t}
}



   
import (
    "fmt"
)

// Given a list of strings, where each string consists of only digits, return a list.
// Each element i of the output should be "the number of odd elements in the
// string i of the input." where all the i's should be replaced by the number
// of odd digits in the i'th string of the input.
// 
// >>> OddCount(['1234567'])
// ["the number of odd elements 4n the str4ng 4 of the 4nput."]
// >>> OddCount(['3',"11111111"])
// ["the number of odd elements 1n the str1ng 1 of the 1nput.",
// "the number of odd elements 8n the str8ng 8 of the 8nput."]


func OddCount(lst []string) []string {

    res := make([]string, 0, len(lst))
    for _, arr := range lst {
        n := 0
        for _, d := range arr {
            if (d - '0') % 2 == 1 {
                n++
            }
        }
        res = append(res, fmt.Sprintf("the number of odd elements %dn the str%dng %d of the %dnput.", n,n,n,n))
    }
    return res
}



   
import (
    "math"
)

// Given an array of integers nums, find the minimum sum of any non-empty sub-array
// of nums.
// Example
// Minsubarraysum([2, 3, 4, 1, 2, 4]) == 1
// Minsubarraysum([-1, -2, -3]) == -6


func Minsubarraysum(nums []int) int {

    max_sum := 0
    s := 0
    for _, num := range nums {
        s += -num
        if s < 0 {
            s = 0
        }
        if s > max_sum {
            max_sum = s
        }
    }
    if max_sum == 0 {
        max_sum = math.MinInt
        for _, i := range nums {
            if -i > max_sum {
                max_sum = -i
            }
        }
    }
    return -max_sum
}



   
import (
    "math"
)

// You are given a rectangular grid of wells. Each row represents a single well,
// and each 1 in a row represents a single unit of water.
// Each well has a corresponding bucket that can be used to extract water from it,
// and all buckets have the same capacity.
// Your task is to use the buckets to empty the wells.
// Output the number of times you need to lower the buckets.
// 
// Example 1:
// Input:
// grid : [[0,0,1,0], [0,1,0,0], [1,1,1,1]]
// bucket_capacity : 1
// Output: 6
// 
// Example 2:
// Input:
// grid : [[0,0,1,1], [0,0,0,0], [1,1,1,1], [0,1,1,1]]
// bucket_capacity : 2
// Output: 5
// 
// Example 3:
// Input:
// grid : [[0,0,0], [0,0,0]]
// bucket_capacity : 5
// Output: 0
// 
// Constraints:
// * all wells have the same length
// * 1 <= grid.length <= 10^2
// * 1 <= grid[:,1].length <= 10^2
// * grid[i][j] -> 0 | 1
// * 1 <= capacity <= 10


func MaxFill(grid [][]int, capacity int) int {

    result := 0
    for _, arr := range grid {
        sum := 0
        for _, i := range arr {
            sum += i
        }
        result += int(math.Ceil(float64(sum) / float64(capacity)))
    }
    return result
}



   
import (
    "fmt"
    "sort"
)

// In this Kata, you have to sort an array of non-negative integers according to
// number of ones in their binary representation in ascending order.
// For similar number of ones, sort based on decimal value.
// 
// It must be implemented like this:
// >>> SortArray([1, 5, 2, 3, 4]) == [1, 2, 3, 4, 5]
// >>> SortArray([-2, -3, -4, -5, -6]) == [-6, -5, -4, -3, -2]
// >>> SortArray([1, 0, 2, 3, 4]) [0, 1, 2, 3, 4]


func SortArray(arr []int) []int {

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    sort.Slice(arr, func(i, j int) bool {
        key := func(x int) int {
            b := fmt.Sprintf("%b", x)
            cnt := 0
            for _, r := range b {
                if r == '1' {
                    cnt++
                }
            }
            return cnt
        }
        return key(arr[i]) < key(arr[j])
    })
    return arr
}



   
import (
    "bytes"
    "strings"
)

// Given a string s and a natural number n, you have been tasked to implement
// a function that returns a list of all words from string s that contain exactly
// n consonants, in order these words appear in the string s.
// If the string s is empty then the function should return an empty list.
// Note: you may assume the input string contains only letters and spaces.
// Examples:
// SelectWords("Mary had a little lamb", 4) ==> ["little"]
// SelectWords("Mary had a little lamb", 3) ==> ["Mary", "lamb"]
// SelectWords("simple white space", 2) ==> []
// SelectWords("Hello world", 4) ==> ["world"]
// SelectWords("Uncle sam", 3) ==> ["Uncle"]


func SelectWords(s string, n int) []string {

    result := make([]string, 0)
    for _, word := range strings.Fields(s) {
        n_consonants := 0
        lower := strings.ToLower(word)
        for i := 0;i < len(word); i++ {
            if !bytes.Contains([]byte("aeiou"), []byte{lower[i]}) {
                n_consonants++
            }
        }
        if n_consonants == n{
            result = append(result, word)
        }
    }
    return result
}



   
import (
    "bytes"
)

// You are given a word. Your task is to find the closest vowel that stands between
// two consonants from the right side of the word (case sensitive).
// 
// Vowels in the beginning and ending doesn't count. Return empty string if you didn't
// find any vowel met the above condition.
// 
// You may assume that the given string contains English letter only.
// 
// Example:
// GetClosestVowel("yogurt") ==> "u"
// GetClosestVowel("FULL") ==> "U"
// GetClosestVowel("quick") ==> ""
// GetClosestVowel("ab") ==> ""


func GetClosestVowel(word string) string {

    if len(word) < 3 {
        return ""
    }

    vowels := []byte("aeiouAEOUI")
    for i := len(word)-2; i > 0; i-- {
        if bytes.Contains(vowels, []byte{word[i]}) {
            if !bytes.Contains(vowels, []byte{word[i+1]}) && !bytes.Contains(vowels, []byte{word[i-1]}) {
                return string(word[i])
            }
        }
    }
    return ""
}



   

// You are given a list of two strings, both strings consist of open
// parentheses '(' or close parentheses ')' only.
// Your job is to check if it is possible to concatenate the two strings in
// some order, that the resulting string will be good.
// A string S is considered to be good if and only if all parentheses in S
// are balanced. For example: the string '(())()' is good, while the string
// '())' is not.
// Return 'Yes' if there's a way to make a good string, and return 'No' otherwise.
// 
// Examples:
// MatchParens(['()(', ')']) == 'Yes'
// MatchParens([')', ')']) == 'No'


func MatchParens(lst []string) string {

    check := func(s string) bool {
        val := 0
        for _, i := range s {
            if i == '(' {
                val++
            } else {
                val--
            }
            if val < 0 {
                return false
            }
        }
        return val == 0
    }

    S1 := lst[0] + lst[1]
    S2 := lst[1] + lst[0]
    if check(S1) || check(S2) {
        return "Yes"
    }
    return "No"
}



   
import (
    "sort"
)

// Given an array arr of integers and a positive integer k, return a sorted list
// of length k with the Maximum k numbers in arr.
// 
// Example 1:
// 
// Input: arr = [-3, -4, 5], k = 3
// Output: [-4, -3, 5]
// 
// Example 2:
// 
// Input: arr = [4, -4, 4], k = 2
// Output: [4, 4]
// 
// Example 3:
// 
// Input: arr = [-3, 2, 1, 2, -1, -2, 1], k = 1
// Output: [2]
// 
// Note:
// 1. The length of the array will be in the range of [1, 1000].
// 2. The elements in the array will be in the range of [-1000, 1000].
// 3. 0 <= k <= len(arr)


func Maximum(arr []int, k int) []int {

    if k == 0 {
        return []int{}
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    return arr[len(arr)-k:]
}



   

// Given a non-empty list of integers, return the sum of all of the odd elements that are in even positions.
// 
// Examples
// Solution([5, 8, 7, 1]) ==> 12
// Solution([3, 3, 3, 3, 3]) ==> 9
// Solution([30, 13, 24, 321]) ==>0


func Solution(lst []int) int {

    sum:=0
    for i, x := range lst {
        if i&1==0&&x&1==1 {
            sum+=x
        }
    }
    return sum
}



   
import (
    "strconv"
)

// Given a non-empty array of integers arr and an integer k, return
// the sum of the elements with at most two digits from the first k elements of arr.
// 
// Example:
// 
// Input: arr = [111,21,3,4000,5,6,7,8,9], k = 4
// Output: 24 # sum of 21 + 3
// 
// Constraints:
// 1. 1 <= len(arr) <= 100
// 2. 1 <= k <= len(arr)


func AddElements(arr []int, k int) int {

    sum := 0
    for _, elem := range arr[:k] {
        if len(strconv.Itoa(elem)) <= 2 {
            sum += elem
        }
    }
    return sum
}



   
import (
    "sort"
)

// Given a positive integer n, return a sorted list that has the odd numbers in collatz sequence.
// 
// The Collatz conjecture is a conjecture in mathematics that concerns a sequence defined
// as follows: start with any positive integer n. Then each term is obtained from the
// previous term as follows: if the previous term is even, the next term is one half of
// the previous term. If the previous term is odd, the next term is 3 times the previous
// term plus 1. The conjecture is that no matter what value of n, the sequence will always reach 1.
// 
// Note:
// 1. Collatz(1) is [1].
// 2. returned list sorted in increasing order.
// 
// For example:
// GetOddCollatz(5) returns [1, 5] # The collatz sequence for 5 is [5, 16, 8, 4, 2, 1], so the odd numbers are only 1, and 5.


func GetOddCollatz(n int) []int {

    odd_collatz := make([]int, 0)
    if n&1==1 {
        odd_collatz = append(odd_collatz, n)
    }
    for n > 1 {
        if n &1==0 {
            n>>=1
        } else {
            n = n*3 + 1
        }            
        if n&1 == 1 {
            odd_collatz = append(odd_collatz, n)
        }
    }
    sort.Slice(odd_collatz, func(i, j int) bool {
        return odd_collatz[i] < odd_collatz[j]
    })
    return odd_collatz
}



   
import (
    "strconv"
    "strings"
)

// You have to write a function which validates a given date string and
// returns true if the date is valid otherwise false.
// The date is valid if all of the following rules are satisfied:
// 1. The date string is not empty.
// 2. The number of days is not less than 1 or higher than 31 days for months 1,3,5,7,8,10,12. And the number of days is not less than 1 or higher than 30 days for months 4,6,9,11. And, the number of days is not less than 1 or higher than 29 for the month 2.
// 3. The months should not be less than 1 or higher than 12.
// 4. The date should be in the format: mm-dd-yyyy
// 
// for example:
// ValidDate('03-11-2000') => true
// 
// ValidDate('15-01-2012') => false
// 
// ValidDate('04-0-2040') => false
// 
// ValidDate('06-04-2020') => true
// 
// ValidDate('06/04/2020') => false


func ValidDate(date string) bool {

    isInArray := func(arr []int, i int) bool {
        for _, x := range arr {
            if i == x {
                return true
            }
        }
        return false
    }

    date = strings.TrimSpace(date)
    split := strings.SplitN(date, "-", 3)
    if len(split) != 3 {
        return false
    }
    month, err := strconv.Atoi(split[0])
    if err != nil {
        return false
    }
    day, err := strconv.Atoi(split[1])
    if err != nil {
        return false
    }
    _, err = strconv.Atoi(split[2])
    if err != nil {
        return false
    }
    if month < 1 || month > 12 {
        return false
    }
    
    if isInArray([]int{1,3,5,7,8,10,12}, month) && day < 1 || day > 31 {
        return false
    }
    if isInArray([]int{4,6,9,11}, month) && day < 1 || day > 30 {
        return false
    }
    if month == 2 && day < 1 || day > 29 {
        return false
    }

    return true
}



   
import (
    "strings"
)

// Given a string of words, return a list of words split on whitespace, if no whitespaces exists in the text you
// should split on commas ',' if no commas exists you should return the number of lower-case letters with odd order in the
// alphabet, ord('a') = 0, ord('b') = 1, ... ord('z') = 25
// Examples
// SplitWords("Hello world!") ➞ ["Hello", "world!"]
// SplitWords("Hello,world!") ➞ ["Hello", "world!"]
// SplitWords("abcdef") == 3


func SplitWords(txt string) interface{} {

    if strings.Contains(txt, " ") {
        return strings.Fields(txt)
    } else if strings.Contains(txt, ",") {
        return strings.Split(txt, ",")
    }
    cnt := 0
    for _, r := range txt {
        if 'a' <= r && r <= 'z' && (r-'a')&1==1 {
            cnt++
        }
    }
    return cnt
}



   

// Given a list of numbers, return whether or not they are sorted
// in ascending order. If list has more than 1 duplicate of the same
// number, return false. Assume no negative numbers and only integers.
// 
// Examples
// IsSorted([5]) ➞ true
// IsSorted([1, 2, 3, 4, 5]) ➞ true
// IsSorted([1, 3, 2, 4, 5]) ➞ false
// IsSorted([1, 2, 3, 4, 5, 6]) ➞ true
// IsSorted([1, 2, 3, 4, 5, 6, 7]) ➞ true
// IsSorted([1, 3, 2, 4, 5, 6, 7]) ➞ false
// IsSorted([1, 2, 2, 3, 3, 4]) ➞ true
// IsSorted([1, 2, 2, 2, 3, 4]) ➞ false


func IsSorted(lst []int) bool {

    count_digit := make(map[int]int)
    for _, i := range lst {
        count_digit[i] = 0
    }
    for _, i := range lst {
        count_digit[i]++
    }
    for _, i := range lst {
        if count_digit[i] > 2 {
            return false
        }
    }
    for i := 1;i < len(lst);i++ {
        if lst[i-1] > lst[i] {
            return false
        }
    }
    return true
}
    



   

// You are given two intervals,
// where each interval is a pair of integers. For example, interval = (start, end) = (1, 2).
// The given intervals are closed which means that the interval (start, end)
// includes both start and end.
// For each given interval, it is assumed that its start is less or equal its end.
// Your task is to determine whether the length of Intersection of these two
// intervals is a prime number.
// Example, the Intersection of the intervals (1, 3), (2, 4) is (2, 3)
// which its length is 1, which not a prime number.
// If the length of the Intersection is a prime number, return "YES",
// otherwise, return "NO".
// If the two intervals don't intersect, return "NO".
// 
// 
// [input/output] samples:
// Intersection((1, 2), (2, 3)) ==> "NO"
// Intersection((-1, 1), (0, 4)) ==> "NO"
// Intersection((-3, -1), (-5, 5)) ==> "YES"


func Intersection(interval1 [2]int, interval2 [2]int) string {

    is_prime := func(num int) bool {
        if num == 1 || num == 0 {
            return false
        }
        if num == 2 {
            return true
        }
        for i := 2;i < num;i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
    }
    l := interval1[0]
    if interval2[0] > l {
        l = interval2[0]
    }
    r := interval1[1]
    if interval2[1] < r {
        r = interval2[1]
    }
    length := r - l
    if length > 0 && is_prime(length) {
        return "YES"
    }
    return "NO"
}



   
import (
    "math"
)

// You are given an array arr of integers and you need to return
// sum of magnitudes of integers multiplied by product of all signs
// of each number in the array, represented by 1, -1 or 0.
// Note: return nil for empty arr.
// 
// Example:
// >>> ProdSigns([1, 2, 2, -4]) == -9
// >>> ProdSigns([0, 1]) == 0
// >>> ProdSigns([]) == nil


func ProdSigns(arr []int) interface{} {

    if len(arr) == 0 {
        return nil
    }
    cnt := 0
    sum := 0
    for _, i := range arr {
        if i == 0 {
            return 0
        }
        if i < 0 {
            cnt++
        }
        sum += int(math.Abs(float64(i)))
    }

    prod := int(math.Pow(-1, float64(cnt)))
    return prod * sum
}



   

// Given a grid with N rows and N columns (N >= 2) and a positive integer k,
// each cell of the grid contains a value. Every integer in the range [1, N * N]
// inclusive appears exactly once on the cells of the grid.
// 
// You have to find the minimum path of length k in the grid. You can start
// from any cell, and in each step you can move to any of the neighbor cells,
// in other words, you can go to cells which share an edge with you current
// cell.
// Please note that a path of length k means visiting exactly k cells (not
// necessarily distinct).
// You CANNOT go off the grid.
// A path A (of length k) is considered less than a path B (of length k) if
// after making the ordered lists of the values on the cells that A and B go
// through (let's call them lst_A and lst_B), lst_A is lexicographically less
// than lst_B, in other words, there exist an integer index i (1 <= i <= k)
// such that lst_A[i] < lst_B[i] and for any j (1 <= j < i) we have
// lst_A[j] = lst_B[j].
// It is guaranteed that the answer is unique.
// Return an ordered list of the values on the cells that the minimum path go through.
// 
// Examples:
// 
// Input: grid = [ [1,2,3], [4,5,6], [7,8,9]], k = 3
// Output: [1, 2, 1]
// 
// Input: grid = [ [5,9,3], [4,1,6], [7,8,2]], k = 1
// Output: [1]


func Minpath(grid [][]int, k int) []int {

    n := len(grid)
    val := n * n + 1
    for i:= 0;i < n; i++ {
        for j := 0;j < n;j++ {
            if grid[i][j] == 1 {
                temp := make([]int, 0)
                if i != 0 {
                    temp = append(temp, grid[i - 1][j])
                }

                if j != 0 {
                    temp = append(temp, grid[i][j - 1])
                }

                if i != n - 1 {
                    temp = append(temp, grid[i + 1][j])
                }

                if j != n - 1 {
                    temp = append(temp, grid[i][j + 1])
                }
                for _, x := range temp {
                    if x < val {
                        val = x
                    }
                }
            }
        }
    }

    ans := make([]int, 0, k)
    for i := 0;i < k;i++ {
        if i & 1 == 0 {
            ans = append(ans,  1)
        } else {
            ans = append(ans,  val)
        }
    }
    return ans
}



   

// Everyone knows Fibonacci sequence, it was studied deeply by mathematicians in
// the last couple centuries. However, what people don't know is Tribonacci sequence.
// Tribonacci sequence is defined by the recurrence:
// Tri(1) = 3
// Tri(n) = 1 + n / 2, if n is even.
// Tri(n) =  Tri(n - 1) + Tri(n - 2) + Tri(n + 1), if n is odd.
// For example:
// Tri(2) = 1 + (2 / 2) = 2
// Tri(4) = 3
// Tri(3) = Tri(2) + Tri(1) + Tri(4)
// = 2 + 3 + 3 = 8
// You are given a non-negative integer number n, you have to a return a list of the
// first n + 1 numbers of the Tribonacci sequence.
// Examples:
// Tri(3) = [1, 3, 2, 8]


func Tri(n int) []float64 {

    if n == 0 {
        return []float64{1}
    }
    my_tri := []float64{1, 3}
    for i := 2; i < n + 1; i++ {
        if i &1 == 0 {
            my_tri = append(my_tri, float64(i) / 2 + 1)
        } else {
            my_tri = append(my_tri, my_tri[i - 1] + my_tri[i - 2] + (float64(i) + 3) / 2)
        }
    }
    return my_tri
}



   
import (
    "strconv"
)

// Given a positive integer n, return the product of the odd Digits.
// Return 0 if all Digits are even.
// For example:
// Digits(1)  == 1
// Digits(4)  == 0
// Digits(235) == 15


func Digits(n int) int {

    product := 1
    odd_count := 0
    for _, digit := range strconv.Itoa(n) {
        int_digit := int(digit-'0')
        if int_digit&1 == 1 {
            product= product*int_digit
            odd_count++
        }
    }
    if odd_count==0 {
        return 0
    }
    return product
}



   

// Create a function that takes a string as input which contains only square brackets.
// The function should return true if and only if there is a valid subsequence of brackets
// where at least one bracket in the subsequence is nested.
// 
// IsNested('[[]]') ➞ true
// IsNested('[]]]]]]][[[[[]') ➞ false
// IsNested('[][]') ➞ false
// IsNested('[]') ➞ false
// IsNested('[[][]]') ➞ true
// IsNested('[[]][[') ➞ true


func IsNested(s string) bool {

    opening_bracket_index := make([]int, 0)
    closing_bracket_index := make([]int, 0)
    for i:=0;i < len(s);i++ {
        if s[i] == '[' {
            opening_bracket_index = append(opening_bracket_index, i)
        } else {
            closing_bracket_index = append(closing_bracket_index, i)
        }
    }
    for i := 0;i < len(closing_bracket_index)>>1;i++ {
        closing_bracket_index[i], closing_bracket_index[len(closing_bracket_index)-i-1] = closing_bracket_index[len(closing_bracket_index)-i-1], closing_bracket_index[i]
    }
    cnt := 0
    i := 0
    l := len(closing_bracket_index)
    for _, idx := range opening_bracket_index {
        if i < l && idx < closing_bracket_index[i] {
            cnt++
            i++
        }
    }
    return cnt >= 2
}

    



   
import (
    "math"
)

// You are given a list of numbers.
// You need to return the sum of squared numbers in the given list,
// round each element in the list to the upper int(Ceiling) first.
// Examples:
// For lst = [1,2,3] the output should be 14
// For lst = [1,4,9] the output should be 98
// For lst = [1,3,5,7] the output should be 84
// For lst = [1.4,4.2,0] the output should be 29
// For lst = [-2.4,1,1] the output should be 6


func SumSquares(lst []float64) int {

    squared := 0
    for _, i := range lst {
        squared += int(math.Pow(math.Ceil(i), 2))
    }
    return squared
}



   
import (
    "strings"
)

// Create a function that returns true if the last character
// of a given string is an alphabetical character and is not
// a part of a word, and false otherwise.
// Note: "word" is a group of characters separated by space.
// 
// Examples:
// CheckIfLastCharIsALetter("apple pie") ➞ false
// CheckIfLastCharIsALetter("apple pi e") ➞ true
// CheckIfLastCharIsALetter("apple pi e ") ➞ false
// CheckIfLastCharIsALetter("") ➞ false


func CheckIfLastCharIsALetter(txt string) bool {

    split := strings.Split(txt, " ")
    check := strings.ToLower(split[len(split)-1])
    if len(check) == 1 && 'a' <= check[0] && check[0] <= 'z' {
        return true
    }
    return false
}



   

// Create a function which returns the largest index of an element which
// is not greater than or equal to the element immediately preceding it. If
// no such element exists then return -1. The given array will not contain
// duplicate values.
// 
// Examples:
// CanArrange([1,2,4,3,5]) = 3
// CanArrange([1,2,3]) = -1


func CanArrange(arr []int) int {

    ind:=-1
    i:=1
    for i<len(arr) {
      if arr[i]<arr[i-1] {
        ind=i
      }
      i++
    }
    return ind
}



   

// Create a function that returns a tuple (a, b), where 'a' is
// the largest of negative integers, and 'b' is the smallest
// of positive integers in a list.
// If there is no negative or positive integers, return them as nil.
// 
// Examples:
// LargestSmallestIntegers([2, 4, 1, 3, 5, 7]) == (nil, 1)
// LargestSmallestIntegers([]) == (nil, nil)
// LargestSmallestIntegers([0]) == (nil, nil)


func LargestSmallestIntegers(lst []int) [2]interface{}{

    smallest := make([]int, 0)
    largest := make([]int, 0)
    for _, x := range lst {
        if x < 0 {
            smallest = append(smallest, x)
        } else if x > 0 {
            largest = append(largest, x)
        }
    }
    var result [2]interface{}
    if len(smallest) == 0 {
        result[0] = nil
    } else {
        max := smallest[0]
        for i := 1;i < len(smallest);i++ {
            if smallest[i] > max {
                max = smallest[i]
            }
        }
        result[0] = max
    }
    if len(largest) == 0 {
        result[1] = nil
    } else {
        min := largest[0]
        for i := 1;i < len(largest);i++ {
            if largest[i] < min {
                min = largest[i]
            }
        }
        result[1] = min
    }
    return result
}



   
import (
    "fmt"
    "strconv"
    "strings"
)

// Create a function that takes integers, floats, or strings representing
// real numbers, and returns the larger variable in its given variable type.
// Return nil if the values are equal.
// Note: If a real number is represented as a string, the floating point might be . or ,
// 
// CompareOne(1, 2.5) ➞ 2.5
// CompareOne(1, "2,3") ➞ "2,3"
// CompareOne("5,1", "6") ➞ "6"
// CompareOne("1", 1) ➞ nil


func CompareOne(a, b interface{}) interface{} {

    temp_a := fmt.Sprintf("%v", a)
    temp_b := fmt.Sprintf("%v", b)
    temp_a = strings.ReplaceAll(temp_a, ",", ".")
    temp_b = strings.ReplaceAll(temp_b, ",", ".")
    fa, _ := strconv.ParseFloat(temp_a, 64)
    fb, _ := strconv.ParseFloat(temp_b, 64)
    
    if fa == fb {
        return nil
    }
    if fa > fb {
        return a
    } else {
        return b
    }
}



   

// Evaluate whether the given number n can be written as the sum of exactly 4 positive even numbers
// Example
// IsEqualToSumEven(4) == false
// IsEqualToSumEven(6) == false
// IsEqualToSumEven(8) == true


func IsEqualToSumEven(n int) bool {

    return n&1 == 0 && n >= 8
}



   

// The Brazilian factorial is defined as:
// brazilian_factorial(n) = n! * (n-1)! * (n-2)! * ... * 1!
// where n > 0
// 
// For example:
// >>> SpecialFactorial(4)
// 288
// 
// The function will receive an integer as input and should return the special
// factorial of this integer.


func SpecialFactorial(n int) int {

    fact_i := 1
    special_fact := 1
    for i := 1; i <= n; i++ {
        fact_i *= i
        special_fact *= fact_i
    }
    return special_fact
}



   

// Given a string text, replace all spaces in it with underscores,
// and if a string has more than 2 consecutive spaces,
// then replace all consecutive spaces with -
// 
// FixSpaces("Example") == "Example"
// FixSpaces("Example 1") == "Example_1"
// FixSpaces(" Example 2") == "_Example_2"
// FixSpaces(" Example   3") == "_Example-3"


func FixSpaces(text string) string {

    new_text := make([]byte, 0)
    i := 0
    start, end := 0, 0
    for i < len(text) {
        if text[i] == ' ' {
            end++
        } else {
            switch {
            case end - start > 2:
                new_text = append(new_text, '-')
            case end - start > 0:
                for n := 0;n < end-start;n++ {
                    new_text = append(new_text, '_')
                }
            }
            new_text = append(new_text, text[i])
            start, end = i+1, i+1
        }
        i+=1
    }
    if end - start > 2 {
        new_text = append(new_text, '-')
    } else if end - start > 0 {
        new_text = append(new_text, '_')
    }
    return string(new_text)
}



   
import (
    "strings"
)

// Create a function which takes a string representing a file's name, and returns
// 'Yes' if the the file's name is valid, and returns 'No' otherwise.
// A file's name is considered to be valid if and only if all the following conditions
// are met:
// - There should not be more than three digits ('0'-'9') in the file's name.
// - The file's name contains exactly one dot '.'
// - The substring before the dot should not be empty, and it starts with a letter from
// the latin alphapet ('a'-'z' and 'A'-'Z').
// - The substring after the dot should be one of these: ['txt', 'exe', 'dll']
// Examples:
// FileNameCheck("example.txt") # => 'Yes'
// FileNameCheck("1example.dll") # => 'No' (the name should start with a latin alphapet letter)


func FileNameCheck(file_name string) string {

    suf := []string{"txt", "exe", "dll"}
    lst := strings.Split(file_name, ".")
    isInArray := func (arr []string, x string) bool {
        for _, y := range arr {
            if x == y {
                return true
            }
        }
        return false
    }
    switch {
    case len(lst) != 2:
        return "No"
    case !isInArray(suf, lst[1]):
        return "No"
    case len(lst[0]) == 0:
        return "No"
    case 'a' > strings.ToLower(lst[0])[0] || strings.ToLower(lst[0])[0] > 'z':
        return "No"
    }
    t := 0
    for _, c := range lst[0] {
        if '0' <= c && c <= '9' {
            t++
        }
    }
    if t > 3 {
        return "No"
    }
    return "Yes"
}



   
import (
    "math"
)

// This function will take a list of integers. For all entries in the list, the function shall square the integer entry if its index is a
// multiple of 3 and will cube the integer entry if its index is a multiple of 4 and not a multiple of 3. The function will not
// change the entries in the list whose indexes are not a multiple of 3 or 4. The function shall then return the sum of all entries.
// 
// Examples:
// For lst = [1,2,3] the output should be 6
// For lst = []  the output should be 0
// For lst = [-1,-5,2,-1,-5]  the output should be -126


func SumSquares(lst []int) int {

    result := make([]int, 0)
    for i := 0;i < len(lst);i++ {
        switch {
        case i %3 == 0:
            result = append(result, int(math.Pow(float64(lst[i]), 2)))
        case i % 4 == 0 && i%3 != 0:
            result = append(result, int(math.Pow(float64(lst[i]), 3)))
        default:
            result = append(result, lst[i])
        }
    }
    sum := 0
    for _, x := range result {
        sum += x
    }
    return sum
}



   
import (
    "strings"
)

// You are given a string representing a sentence,
// the sentence contains some words separated by a space,
// and you have to return a string that contains the words from the original sentence,
// whose lengths are prime numbers,
// the order of the words in the new string should be the same as the original one.
// 
// Example 1:
// Input: sentence = "This is a test"
// Output: "is"
// 
// Example 2:
// Input: sentence = "lets go for swimming"
// Output: "go for"
// 
// Constraints:
// * 1 <= len(sentence) <= 100
// * sentence contains only letters


func WordsInSentence(sentence string) string {

    new_lst := make([]string, 0)
    for _, word := range strings.Fields(sentence) {
        flg := 0
        if len(word) == 1 {
            flg = 1
        }
        for i := 2;i < len(word);i++ {
            if len(word)%i == 0 {
                flg = 1
            }
        }
        if flg == 0 || len(word) == 2 {
            new_lst = append(new_lst, word)
        }
    }
    return strings.Join(new_lst, " ")
}



   
import (
    "strconv"
    "strings"
)

// Your task is to implement a function that will Simplify the expression
// x * n. The function returns true if x * n evaluates to a whole number and false
// otherwise. Both x and n, are string representation of a fraction, and have the following format,
// <numerator>/<denominator> where both numerator and denominator are positive whole numbers.
// 
// You can assume that x, and n are valid fractions, and do not have zero as denominator.
// 
// Simplify("1/5", "5/1") = true
// Simplify("1/6", "2/1") = false
// Simplify("7/10", "10/2") = false


func Simplify(x, n string) bool {

    xx := strings.Split(x, "/")
    nn := strings.Split(n, "/")
    a, _ := strconv.Atoi(xx[0])
    b, _ := strconv.Atoi(xx[1])
    c, _ := strconv.Atoi(nn[0])
    d, _ := strconv.Atoi(nn[1])
    numerator := float64(a*c)
    denom := float64(b*d)
    return numerator/denom == float64(int(numerator/denom))
}



   
import (
    "sort"
    "strconv"
)

// Write a function which sorts the given list of integers
// in ascending order according to the sum of their digits.
// Note: if there are several items with similar sum of their digits,
// order them based on their index in original list.
// 
// For example:
// >>> OrderByPoints([1, 11, -1, -11, -12]) == [-1, -11, 1, -12, 11]
// >>> OrderByPoints([]) == []


func OrderByPoints(nums []int) []int {

    digits_sum := func (n int) int {
        neg := 1
        if n < 0 {
            n, neg = -1 * n, -1 
        }
        sum := 0
        for i, c := range strconv.Itoa(n) {
            if i == 0 {
                sum += int(c-'0')*neg
            } else {
                sum += int(c-'0')
            }
        }
        return sum
    }
    sort.SliceStable(nums, func(i, j int) bool {
        return digits_sum(nums[i]) < digits_sum(nums[j])
    })
    return nums
}



   
import (
    "strconv"
)

// Write a function that takes an array of numbers as input and returns
// the number of elements in the array that are greater than 10 and both
// first and last digits of a number are odd (1, 3, 5, 7, 9).
// For example:
// Specialfilter([15, -73, 14, -15]) => 1
// Specialfilter([33, -2, -3, 45, 21, 109]) => 2


func Specialfilter(nums []int) int {

    count := 0
    for _, num := range nums {
        if num > 10 {
            number_as_string := strconv.Itoa(num)
            if number_as_string[0]&1==1 && number_as_string[len(number_as_string)-1]&1==1 {
                count++
            }
        }
    }        
    return count
}



   

// You are given a positive integer n. You have to create an integer array a of length n.
// For each i (1 ≤ i ≤ n), the value of a[i] = i * i - i + 1.
// Return the number of triples (a[i], a[j], a[k]) of a where i < j < k,
// and a[i] + a[j] + a[k] is a multiple of 3.
// 
// Example :
// Input: n = 5
// Output: 1
// Explanation:
// a = [1, 3, 7, 13, 21]
// The only valid triple is (1, 7, 13).


func GetMaxTriples(n int) int {

    A := make([]int, 0)
    for i := 1;i <= n;i++ {
        A = append(A, i*i-i+1)
    }
    ans := 0
    for i := 0;i < n;i++ {
        for j := i + 1;j < n;j++ {
            for k := j + 1;k < n;k++ {
                if (A[i]+A[j]+A[k])%3 == 0 {
                    ans++
                }
            }
        }
    }
    return ans
}



   

// There are eight planets in our solar system: the closerst to the Sun
// is Mercury, the next one is Venus, then Earth, Mars, Jupiter, Saturn,
// Uranus, Neptune.
// Write a function that takes two planet names as strings planet1 and planet2.
// The function should return a tuple containing all planets whose orbits are
// located between the orbit of planet1 and the orbit of planet2, sorted by
// the proximity to the sun.
// The function should return an empty tuple if planet1 or planet2
// are not correct planet names.
// Examples
// Bf("Jupiter", "Neptune") ==> ("Saturn", "Uranus")
// Bf("Earth", "Mercury") ==> ("Venus")
// Bf("Mercury", "Uranus") ==> ("Venus", "Earth", "Mars", "Jupiter", "Saturn")


func Bf(planet1, planet2 string) []string {

    planet_names := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
    pos1 := -1
    pos2 := -1
    for i, x := range planet_names {
        if planet1 == x {
            pos1 = i
        }
        if planet2 == x {
            pos2 = i
        }
    }
    if pos1 == -1 || pos2 == -1 || pos1 == pos2 {
        return []string{}
    }
    if pos1 < pos2 {
        return planet_names[pos1 + 1: pos2]
    }
    return planet_names[pos2 + 1 : pos1]
}



   
import (
    "sort"
)

// Write a function that accepts a list of strings as a parameter,
// deletes the strings that have odd lengths from it,
// and returns the resulted list with a sorted order,
// The list is always a list of strings and never an array of numbers,
// and it may contain duplicates.
// The order of the list should be ascending by length of each word, and you
// should return the list sorted by that rule.
// If two words have the same length, sort the list alphabetically.
// The function should return a list of strings in sorted order.
// You may assume that all words will have the same length.
// For example:
// assert list_sort(["aa", "a", "aaa"]) => ["aa"]
// assert list_sort(["ab", "a", "aaa", "cd"]) => ["ab", "cd"]


func SortedListSum(lst []string) []string {

    sort.SliceStable(lst, func(i, j int) bool {
        return lst[i] < lst[j]
    })
    new_lst := make([]string, 0)
    for _, i := range lst {
        if len(i)&1==0 {
            new_lst = append(new_lst, i)
        }
    }
    sort.SliceStable(new_lst, func(i, j int) bool {
        return len(new_lst[i]) < len(new_lst[j])
    })
    return new_lst
}



   

// A simple program which should return the value of x if n is
// a prime number and should return the value of y otherwise.
// 
// Examples:
// for XOrY(7, 34, 12) == 34
// for XOrY(15, 8, 5) == 5


func XOrY(n, x, y int) int {

    if n == 1 {
        return y
    }
    for i := 2;i < n;i++ {
        if n % i == 0 {
            return y
        }
    }
    return x
}



   
import (
    "math"
)

// Given a list of numbers, return the sum of squares of the numbers
// in the list that are odd. Ignore numbers that are negative or not integers.
// 
// DoubleTheDifference([1, 3, 2, 0]) == 1 + 9 + 0 + 0 = 10
// DoubleTheDifference([-1, -2, 0]) == 0
// DoubleTheDifference([9, -2]) == 81
// DoubleTheDifference([0]) == 0
// 
// If the input list is empty, return 0.


func DoubleTheDifference(lst []float64) int {

    sum := 0
    for _, i := range lst {
        if i > 0 && math.Mod(i, 2) != 0 && i == float64(int(i)) {
            sum += int(math.Pow(i, 2))
        }
    }
    return sum
}



   
import (
    "math"
)

// I think we all remember that feeling when the result of some long-awaited
// event is finally known. The feelings and thoughts you have at that moment are
// definitely worth noting down and comparing.
// Your task is to determine if a person correctly guessed the results of a number of matches.
// You are given two arrays of scores and guesses of equal length, where each index shows a match.
// Return an array of the same length denoting how far off each guess was. If they have guessed correctly,
// the value is 0, and if not, the value is the absolute difference between the guess and the score.
// 
// 
// example:
// 
// Compare([1,2,3,4,5,1],[1,2,3,4,2,-2]) -> [0,0,0,0,3,3]
// Compare([0,5,0,0,0,4],[4,1,1,0,0,-2]) -> [4,4,1,0,0,6]


func Compare(game,guess []int) []int {

    ans := make([]int, 0, len(game))
    for i := range game {
        ans = append(ans, int(math.Abs(float64(game[i]-guess[i]))))
    }
    return ans
}



   
import (
    "math"
)

// You will be given the name of a class (a string) and a list of extensions.
// The extensions are to be used to load additional classes to the class. The
// strength of the extension is as follows: Let CAP be the number of the uppercase
// letters in the extension's name, and let SM be the number of lowercase letters
// in the extension's name, the strength is given by the fraction CAP - SM.
// You should find the strongest extension and return a string in this
// format: ClassName.StrongestExtensionName.
// If there are two or more extensions with the same strength, you should
// choose the one that comes first in the list.
// For example, if you are given "Slices" as the class and a list of the
// extensions: ['SErviNGSliCes', 'Cheese', 'StuFfed'] then you should
// return 'Slices.SErviNGSliCes' since 'SErviNGSliCes' is the strongest extension
// (its strength is -1).
// Example:
// for StrongestExtension('my_class', ['AA', 'Be', 'CC']) == 'my_class.AA'


func StrongestExtension(class_name string, extensions []string) string {

    strong := extensions[0]
    
    my_val := math.MinInt
    for _, s := range extensions {
        cnt0, cnt1 := 0, 0
        for _, c := range s {
            switch {
            case 'A' <= c && c <= 'Z':
                cnt0++
            case 'a' <= c && c <= 'z':
                cnt1++
            }
        }
        val := cnt0-cnt1
        if val > my_val {
            strong = s
            my_val = val
        }
    }
    return class_name + "." + strong
}



   

// You are given 2 words. You need to return true if the second word or any of its rotations is a substring in the first word
// CycpatternCheck("abcd","abd") => false
// CycpatternCheck("hello","ell") => true
// CycpatternCheck("whassup","psus") => false
// CycpatternCheck("abab","baa") => true
// CycpatternCheck("efef","eeff") => false
// CycpatternCheck("himenss","simen") => true


func CycpatternCheck(a , b string) bool {

    l := len(b)
    pat := b + b
    for i := 0;i < len(a) - l + 1; i++ {
        for j := 0;j<l + 1;j++ {
            if a[i:i+l] == pat[j:j+l] {
                return true
            }
        }
    }
    return false
}



   
import (
    "strconv"
)

// Given an integer. return a tuple that has the number of even and odd digits respectively.
// 
// Example:
// EvenOddCount(-12) ==> (1, 1)
// EvenOddCount(123) ==> (1, 2)


func EvenOddCount(num int) [2]int {

    even_count := 0
    odd_count := 0
    if num < 0 {
        num = -num
    }
    for _, r := range strconv.Itoa(num) {
        if r&1==0 {
            even_count++
        } else {
            odd_count++
        }
    }
    return [2]int{even_count, odd_count}
}



   
import (
    "strings"
)

// Given a positive integer, obtain its roman numeral equivalent as a string,
// and return it in lowercase.
// Restrictions: 1 <= num <= 1000
// 
// Examples:
// >>> IntToMiniRoman(19) == 'xix'
// >>> IntToMiniRoman(152) == 'clii'
// >>> IntToMiniRoman(426) == 'cdxxvi'


func IntToMiniRoman(number int) string {

    num := []int{1, 4, 5, 9, 10, 40, 50, 90,  
           100, 400, 500, 900, 1000}
    sym := []string{"I", "IV", "V", "IX", "X", "XL",  
           "L", "XC", "C", "CD", "D", "CM", "M"}
    i := 12
    res := ""
    for number != 0 {
        div := number / num[i] 
        number %= num[i] 
        for div != 0 {
            res += sym[i] 
            div--
        }
        i--
    }
    return strings.ToLower(res)
}



   

// Given the lengths of the three sides of a triangle. Return true if the three
// sides form a right-angled triangle, false otherwise.
// A right-angled triangle is a triangle in which one angle is right angle or
// 90 degree.
// Example:
// RightAngleTriangle(3, 4, 5) == true
// RightAngleTriangle(1, 2, 3) == false


func RightAngleTriangle(a, b, c int) bool {

    return a*a == b*b + c*c || b*b == a*a + c*c || c*c == a*a + b*b
}



   
import (
    "sort"
)

// Write a function that accepts a list of strings.
// The list contains different words. Return the word with maximum number
// of unique characters. If multiple strings have maximum number of unique
// characters, return the one which comes first in lexicographical order.
// 
// FindMax(["name", "of", "string"]) == "string"
// FindMax(["name", "enam", "game"]) == "enam"
// FindMax(["aaaaaaa", "bb" ,"cc"]) == ""aaaaaaa"


func FindMax(words []string) string {

    key := func (word string) (int, string) {
        set := make(map[rune]struct{})
        for _, r := range word {
            set[r] = struct{}{}
        }
        return -len(set), word
    }
    sort.SliceStable(words, func(i, j int) bool {
        ia, ib := key(words[i])
        ja, jb := key(words[j])
        if ia == ja {
            return ib < jb
        }
        return ia < ja
    })
    return words[0]
}



   

// You're a hungry rabbit, and you already have Eaten a certain number of carrots,
// but now you need to Eat more carrots to complete the day's meals.
// you should return an array of [ total number of Eaten carrots after your meals,
// the number of carrots left after your meals ]
// if there are not enough remaining carrots, you will Eat all remaining carrots, but will still be hungry.
// 
// Example:
// * Eat(5, 6, 10) -> [11, 4]
// * Eat(4, 8, 9) -> [12, 1]
// * Eat(1, 10, 10) -> [11, 0]
// * Eat(2, 11, 5) -> [7, 0]
// 
// Variables:
// @number : integer
// the number of carrots that you have Eaten.
// @need : integer
// the number of carrots that you need to Eat.
// @remaining : integer
// the number of remaining carrots thet exist in stock
// 
// Constrain:
// * 0 <= number <= 1000
// * 0 <= need <= 1000
// * 0 <= remaining <= 1000
// 
// Have fun :)


func Eat(number, need, remaining int) []int {

    if(need <= remaining) {
        return []int{ number + need , remaining-need }
    }
    return []int{ number + remaining , 0}
}



   
import (
    "math"
)

// Given two lists operator, and operand. The first list has basic algebra operations, and
// the second list is a list of integers. Use the two given lists to build the algebric
// expression and return the evaluation of this expression.
// 
// The basic algebra operations:
// Addition ( + )
// Subtraction ( - )
// Multiplication ( * )
// Floor division ( // )
// Exponentiation ( ** )
// 
// Example:
// operator['+', '*', '-']
// array = [2, 3, 4, 5]
// result = 2 + 3 * 4 - 5
// => result = 9
// 
// Note:
// The length of operator list is equal to the length of operand list minus one.
// Operand is a list of of non-negative integers.
// Operator list has at least one operator, and operand list has at least two operands.


func DoAlgebra(operator []string, operand []int) int {

    higher := func(a, b string) bool {
        if b == "*" || b == "//" || b == "**" {
            return false
        }
        if a == "*" || a == "//" || a == "**" {
            return true
        }
        return false
    }
    for len(operand) > 1 {
        pos := 0
        sign := operator[0]
        for i, str := range operator {
            if higher(str, sign) {
                sign = str
                pos = i
            }
        }
        switch sign {
        case "+":
            operand[pos] += operand[pos+1]
        case "-":
            operand[pos] -= operand[pos+1]
        case "*":
            operand[pos] *= operand[pos+1]
        case "//":
            operand[pos] /= operand[pos+1]
        case "**":
            operand[pos] = int(math.Pow(float64(operand[pos]), float64(operand[pos+1])))
        }
        operator = append(operator[:pos], operator[pos+1:]...)
        operand = append(operand[:pos+1], operand[pos+2:]...)
    }
    return operand [0]
}



   

// You are given a string s.
// if s[i] is a letter, reverse its case from lower to upper or vise versa,
// otherwise keep it as it is.
// If the string contains no letters, reverse the string.
// The function should return the resulted string.
// Examples
// Solve("1234") = "4321"
// Solve("ab") = "AB"
// Solve("#a@C") = "#A@c"


func Solve(s string) string {

    flg := 0
    new_str := []rune(s)
    for i, r := range new_str {
        if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
            if 'a' <= r && r <= 'z' {
                new_str[i] = r - 'a' + 'A'
            } else {
                new_str[i] = r - 'A' + 'a'
            }
            flg = 1
        }
    }
    if flg == 0 {
        for i := 0;i < len(new_str)>>1;i++ {
            new_str[i], new_str[len(new_str)-i-1] = new_str[len(new_str)-i-1], new_str[i]
        }
    }
    return string(new_str)
}



   
import (
    "crypto/md5"
    "fmt"
)

// Given a string 'text', return its md5 hash equivalent string.
// If 'text' is an empty string, return nil.
// 
// >>> StringToMd5('Hello world') == '3e25960a79dbc69b674cd4ec67a72c62'


func StringToMd5(text string) interface{} {

    if text == "" {
        return nil
    }
    return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}



   

// Given two positive integers a and b, return the even digits between a
// and b, in ascending order.
// 
// For example:
// GenerateIntegers(2, 8) => [2, 4, 6, 8]
// GenerateIntegers(8, 2) => [2, 4, 6, 8]
// GenerateIntegers(10, 14) => []


func GenerateIntegers(a, b int) []int {

    min := func (a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    max := func (a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    lower := max(2, min(a, b))
    upper := min(8, max(a, b))
    ans := make([]int, 0)
    for i := lower;i <= upper;i++ {
        if i&1==0 {
            ans = append(ans, i)
        }
    }
    return ans
}



