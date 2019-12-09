class School {
  type DB = Map[Int, Seq[String]]

  private var data = Map[Int, Seq[String]]() // TODO is there a way to use the DB type here?

  def add(name: String, g: Int) = {
    data = data + (g -> (grade(g) :+ name))
  }

  def db: DB = data

  def grade(g: Int): Seq[String] = db getOrElse (g, Nil)

  def sorted: DB = db.toSeq.sortBy(_._1).toMap.mapValues(_.sortBy(s => s))
}
