package ch02

import ch02.Color.*

fun getSaturation(color: Color) = when(color) {
    RED, GREEN ,BLUE -> "high"
    BLACK, WHITE -> "none"
    ORANGE, YELLOW ,INDIGO, VIOLET, -> "low"
}

fun main() {
    println("$RED saturation is ${getSaturation(RED)}")
}