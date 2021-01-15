package ch02

import java.lang.IllegalArgumentException

fun eval(e: Expr): Int =
    when(e) {
        is Num -> e.value
        is Sum -> eval(e.left) + eval(e.right)
        else -> throw IllegalArgumentException("Unknown expression")
    }

fun main() {
    println(eval(Sum(Num(1), Num(2))))
}