object SpaceAge {
  
  def onEarth(seconds:Double, divisor: Double = 31557600): Double =
    seconds / divisor

  def onMercury(seconds: Double, divisor: Double = 0.2408467): Double =
    onEarth(seconds) / divisor

  def onVenus(seconds:Double, divisor: Double = 0.61519726): Double =
    onEarth(seconds) / divisor

  def onMars(seconds:Double, divisor: Double = 1.880815): Double =
    onEarth(seconds) / divisor

  def onJupiter(seconds:Double, divisor: Double = 11.862615): Double =
    onEarth(seconds) / divisor

  def onSaturn(seconds: Double, divisor: Double = 29.447498): Double =
    onEarth(seconds) / divisor

  def onUranus(seconds: Double, divisor: Double = 84.016846): Double =
    onEarth(seconds) / divisor

  def onNeptune(seconds: Double, divisor: Double = 164.79132): Double =
    onEarth(seconds) / divisor
} 
