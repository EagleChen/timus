garden, cord = gets.split(" ").map(&:to_f)

half = garden / 2.0
hypotenuse = half * (2 ** 0.5)
if cord >= hypotenuse
  puts "%.3f" % (garden * garden)
else
  if cord < half
    puts "%.3f" % (Math::PI * cord * cord)
  else
    angle = Math.acos(half / cord) * 2
    result = (Math::PI - 2 * angle) * cord * cord + 4 * ((cord * cord - half * half) ** 0.5) * half
    puts "%.3f" % result
  end
end
