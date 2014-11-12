lines = $stdin.read.split("\n")

num = lines[0].to_i

dates = {}
(1..num).each do |i|
  dates[lines[i]] = true
end

stu = lines[num+1].to_i
result = 0
(num+2..num+1+stu).each do |i|
  result += 1 if dates[lines[i]]
end

puts result
