package ch03

fun main() {
    val strings = listOf("first", "second", "fourteenth")
    println(strings.last())

    val numbers = setOf(1, 14, 2)
    println(numbers.maxOrNull())
}