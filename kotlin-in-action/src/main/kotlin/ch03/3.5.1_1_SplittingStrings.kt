package ch03

fun main() {
    println("12.345-6.A".split("[.\\-]".toRegex()))
    println("12.345-6.A".split(".", "-"))
}