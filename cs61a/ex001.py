# https://pythontutor.com/render.html#mode=display
def horse(mask):
  horse = mask
  def mask(horse):
    return horse
  return horse(mask)

mask = lambda horse: horse(2)
horse(mask)
