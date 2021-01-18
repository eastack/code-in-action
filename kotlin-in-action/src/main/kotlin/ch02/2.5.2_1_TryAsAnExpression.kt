package ch02

import java.io.BufferedReader
import java.io.StringReader
import java.lang.NumberFormatException

fun printNumber(reader: BufferedReader) {
    val number = try {
        Integer.parseInt(reader.readLine())
    } catch (e: NumberFormatException) {
        return
    }
}

fun printNumberOrNull(reader: BufferedReader) {
    val number = try {
        Integer.parseInt(reader.readLine())
    } catch (e: NumberFormatException) {
        null
    }

    println(number)
}

fun main() {
    val reader = BufferedReader(StringReader("not a number"))
    printNumber(reader)
    printNumberOrNull(reader)
}