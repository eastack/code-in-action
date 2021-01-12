package ch02

fun getWarmth(color: Color) = when(color) {
    Color.RED, Color.ORANGE, Color.YELLOW -> "warm"
    Color.GREEN, Color.BLACK, Color.WHITE -> "neutral"
    Color.BLUE, Color.INDIGO, Color.VIOLET -> "cold"
}

fun main() {
    println("${Color.RED} warmth is ${getWarmth(Color.RED)}")
}