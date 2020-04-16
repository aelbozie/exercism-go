scalaVersion := "2.12.10"
val gitCommitCountTask = taskKey[String]("Prints commit count of the current branch")
libraryDependencies += "org.scalatest" %% "scalatest" % "3.0.1" % "test"

gitCommitCountTask := {
  val branch = scala.sys.process.Process("git symbolic-ref -q HEAD").lines.head.replace("refs/heads/","")
  val commitCount = scala.sys.process.Process(s"git rev-list --count $branch").lines.head
  println(s"total number of commits on [$branch]: $commitCount")
  commitCount
}
