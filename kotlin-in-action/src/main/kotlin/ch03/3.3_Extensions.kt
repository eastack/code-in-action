package ch03

fun String.lastChar(): Char = get(length - 1)

fun main() {
    println("Kotlin".lastChar())
}