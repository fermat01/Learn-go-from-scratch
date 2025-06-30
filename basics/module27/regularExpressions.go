package main

import (
    "fmt"
    "regexp"
)

func main() {
    // --- BASIC LEVEL ---
    // 1. Match a literal string
    match, _ := regexp.MatchString("hello", "hello world")
    fmt.Println("Basic: Match 'hello' in 'hello world':", match) // true

    // 2. Match any digit (using character class)
    reDigit := regexp.MustCompile(`\d`)
    hasDigit := reDigit.MatchString("abc123")
    fmt.Println("Basic: Has digit in 'abc123':", hasDigit) // true

    // --- INTERMEDIATE LEVEL ---
    // 3. Match one or more digits (quantifier)
    reDigits := regexp.MustCompile(`\d+`)
    digits := reDigits.FindString("abc123def456")
    fmt.Println("Intermediate: First sequence of digits in 'abc123def456':", digits) // "123"

    // 4. Match either 'apple' or 'orange' (alternation)
    reFruit := regexp.MustCompile(`apple|orange`)
    fruits := reFruit.FindAllString("apple orange banana", -1)
    fmt.Println("Intermediate: Match fruits 'apple' or 'orange':", fruits) // [apple orange]

    // 5. Match email-like pattern (character classes and quantifiers)
    reEmail := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    email := reEmail.MatchString("user@example.com")
    fmt.Println("Intermediate: Is 'user@example.com' an email?:", email) // true

    // --- ADVANCED LEVEL ---
    // 6. Match and extract groups (grouping and capturing)
    reDate := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
    dateMatch := reDate.FindStringSubmatch("2025-06-29")
    if len(dateMatch) > 0 {
        fmt.Printf("Advanced: Date groups (year, month, day): %s, %s, %s\n", dateMatch[1], dateMatch[2], dateMatch[3]) // "2025", "06", "29"
    }

    // 7. Lookahead assertion (positive and negative)
    // Go regexp package does not support lookahead/lookbehind, but here's a workaround example for lookahead-like logic using capture groups.
    // Note: This is not a true lookahead, but demonstrates intent.
    // reLookahead := regexp.MustCompile(`\d+(?=px)`) // Go does NOT support lookahead, but this example is for clarification.
    // Instead, in Go, you might use: `(\d+)px` and extract group 1.
    reWorkaround := regexp.MustCompile(`(\d+)px`)
    lookaheadMatch := reWorkaround.FindStringSubmatch("Width: 300px")
    if len(lookaheadMatch) > 0 {
        fmt.Println("Advanced: Value before 'px' (workaround):", lookaheadMatch[1]) // "300"
    }

    // 8. Named groups (Go does not support named groups natively, but here's how you might document intent)
    // reNamedGroup := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`) // NOT supported in Go
    // For Go, use numbered groups as in example 6.
    fmt.Println("Note: Go does not support named groups or lookahead/lookbehind assertions natively.")
}
