package tmp.workflow

interface Element {
}

class Arc {
    var pair: Pair<Node, Node> = Node.NONE to Node.NONE
    var exec: () -> Unit = {}
}

class Origin {
    var node = Node.NONE
}

class Terminus {
    var node = Node.NONE
}

fun arc(inti: Arc.() -> Unit): Arc {
    val arc = Arc()
    arc.inti()
    return arc
}

enum class Node {
    NONE,
    DRAFT,
    MU_YUAN_AUDIT,
    BANK_AUDIT,
}

fun main() {
    val arcs = setOf(
        arc {
            pair = Node.DRAFT to Node.MU_YUAN_AUDIT
            exec = { println("$pair") }
        },
        arc {
            pair = Node.MU_YUAN_AUDIT to Node.BANK_AUDIT
            exec = { println("$pair") }
        }
    )


    fun exec(pair: Pair<Node, Node>) {
        arcs.filter { it.pair == pair }.first().exec()
    }

    exec(Node.MU_YUAN_AUDIT to Node.BANK_AUDIT)
}


