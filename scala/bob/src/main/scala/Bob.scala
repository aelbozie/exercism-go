object Bob {
  def response(statement: String): String = {
    statement.trim match {
      case s if s.isEmpty =>
        "Fine. Be that way!"
      case s if isShoutingQuestion(s) =>
        "Calm down, I know what I'm doing!"
      case s if isQuestion(s) =>
        "Sure."
      case s if isShouting(s) =>
        "Whoa, chill out!"
      case _ =>
        "Whatever."
    }
  }

  def isShoutingQuestion(statement: String): Boolean =
    statement.toUpperCase == statement &&
      hasLetters(statement) &&
      isQuestion(statement)

  def isShouting(statement: String): Boolean =
    statement.toUpperCase == statement &&
      hasLetters(statement)

  def isQuestion(statement: String): Boolean =
    statement.endsWith("?")

  def hasLetters(statement: String): Boolean =
    statement.toLowerCase.exists(('a' to 'z').contains(_))
}
