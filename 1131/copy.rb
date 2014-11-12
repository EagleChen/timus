n, k = gets.split(" ").map(&:to_f)

if n == 1
  puts 0
  exit
end

basic = Math.log2(n).ceil
before = Math.log2(k).ceil
if (basic <= before)
  puts basic
  exit
end
after = ((n - 2 ** before) / k).ceil

puts before + after
