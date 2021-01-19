package ch03

fun main(args: Array<String>) {
    val list = listOf("args: ", *args)
    println(list)
}