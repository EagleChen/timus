num, marks = gets.chomp.split(" ")

num = num.to_i
times = marks.length

result = 1
while num > 0
  result *= num
  num -= times
end
puts result
