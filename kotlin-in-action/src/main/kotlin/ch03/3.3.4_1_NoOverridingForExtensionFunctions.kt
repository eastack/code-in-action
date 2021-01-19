package ch03

open class View() {
    open fun click() = println("View clicked")
}

open class Button: View() {
    override fun click() = println("Button clicked")
}

fun View.showOff() = println("I'm a view!")
fun Button.showOff() = println("I'm a button!")

fun main() {
    val view: View = Button()
    view.click()

    /**
     * 扩展函数实际是一个接收调用者的静态函数
     * 与多态在运行时确定不同，其在编译时确定
     * 因此无法进行重写。
     */
    view.showOff()
}
