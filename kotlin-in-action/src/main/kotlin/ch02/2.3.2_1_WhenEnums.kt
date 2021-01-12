package ch02

fun getMnemonic(color: Color) =
    when (color) {
        Color.RED -> "Richard"
        Color.ORANGE -> "Of"
        Color.YELLOW -> "York"
        Color.GREEN -> "Gave"
        Color.BLUE -> "Battle"
        Color.INDIGO -> "In"
        Color.VIOLET -> "Vain"
        Color.BLACK, Color.WHITE -> ""
    }

fun main() {
    println(getMnemonic(Color.RED))
    println(getMnemonic(Color.BLACK))
}