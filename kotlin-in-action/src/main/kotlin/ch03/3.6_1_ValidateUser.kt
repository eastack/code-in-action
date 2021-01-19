package ch03

import java.lang.IllegalArgumentException

class User(val id: Int, val name: String, val address: String)

fun saveUser(user: User)  {
    if (user.name.isEmpty()) {
        throw IllegalArgumentException(
            "Can't save user ${user.id}: empty Name"
        )
    }

    if (user.address.isEmpty()) {
        throw IllegalArgumentException(
            "Can't save user ${user.id}: empty Address"
        )
    }

    // Save user to the database
}

fun saveUserWithLocalFunction(user: User) {
    fun validate(
        value: String,
        fieldName: String
    ) {
        if (value.isEmpty()) {
            throw IllegalArgumentException(
                "Can't save user ${user.id}: empty $fieldName"
            )
        }
    }

    validate(user.name, "Name")
    validate(user.address, "Address")
}

fun User.validateBeforeSave() {
    fun validate(value: String, fieldName: String) {
        if (value.isEmpty()) {
            throw IllegalArgumentException(
                "Can't save user $id: empty $fieldName"
            )
        }

        validate(name, "Name")
        validate(address, "Address")
    }
}

fun saveUserWithExtensionFunction(user: User) {
    user.validateBeforeSave()
}

fun main() {
    saveUser(User(1, "", ""))
    saveUserWithLocalFunction(User(1, "", ""))
    saveUserWithExtensionFunction(User(1, "", ""))
}