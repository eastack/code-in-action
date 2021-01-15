package tmp

import kotlinx.coroutines.*


fun main() = runBlocking {
    val job = launch {
        repeat(1000) { i ->
            println("I'm sleeping $i ...")
            delay(500L)
        }
    }
    println("a")
    delay(1300L) // just quit after delay
    job.cancelAndJoin()
    println("b")
}
