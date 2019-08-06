
object Bob {
  def response(statement: String): String = {
    statement.trim match {
      case s if s.isEmpty =>
        "Fine. Be that way!"
      case s if s.toUpperCase == s && hasLetters(s) && s.endsWith("?") =>
        "Calm down, I know what I'm doing!"
      case s if s.endsWith("?") =>
        "Sure."
      case s if s.toUpperCase == s && hasLetters(s) =>
        "Whoa, chill out!"
      case _ =>
        "Whatever."
    }
  }
  def hasLetters(statement: String): Boolean =
    statement.toLowerCase.exists(('a' to 'z').contains(_))
}
