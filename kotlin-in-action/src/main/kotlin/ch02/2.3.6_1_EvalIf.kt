package ch02

fun main() {
    fun eval(e: Expr): Int =
        if (e is Num) {
            e.value
        } else if (e is Sum) {
            eval(e.right) + eval(e.left)
        } else {
            throw IllegalArgumentException("Unknown expression")
        }

    println(eval(Sum(Sum(Num(1), Num(2)), Num(4))))
}




